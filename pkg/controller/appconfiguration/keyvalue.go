/*
Copyright 2021 The Crossplane Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package appconfiguration

import (
	"context"
	"strings"
	"time"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"github.com/stone-payments/appconfig-go-sdk/appconfig/keyvalues"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	"github.com/crossplane/provider-azure/apis/appconfiguration/v1alpha1"
	azure "github.com/crossplane/provider-azure/pkg/clients"
	appclient "github.com/crossplane/provider-azure/pkg/clients/appconfiguration"
)

const (
	errConnectFailed         = "cannot connect to Azure API"
	errGetFailed             = "cannot get KeyValue"
	errCreateFailed          = "cannot create KeyValue"
	errUpdateFailed          = "cannot update KeyValue"
	errKubeUpdateFailed      = "cannot update custom resource in kubernetes"
	errDeleteFailed          = "cannot delete KeyValue"
	errCannotCreateNewClient = "cannot create NewClient"
	errNotKeyValue           = "unexpected resource, it must be KeyValue"
)

// SetupAppConfiguration adds a controller that reconciles KeyValue resources.
func SetupAppConfiguration(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter, poll time.Duration) error {
	name := managed.ControllerName(v1alpha1.KeyValueKind)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(controller.Options{
			RateLimiter: ratelimiter.NewDefaultManagedRateLimiter(rl),
		}).
		For(&v1alpha1.KeyValue{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(v1alpha1.KeyValueGroupVersionKind),
			managed.WithExternalConnecter(&connector{kube: mgr.GetClient()}),
			managed.WithPollInterval(poll),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

type connector struct {
	kube client.Client
}

func (c connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	content, _, err := azure.GetAuthInfo(ctx, c.kube, mg)
	if err != nil {
		return nil, errors.Wrap(err, errConnectFailed)
	}

	cr, ok := mg.(*v1alpha1.KeyValue)
	if !ok {
		return &external{}, errors.New(errNotKeyValue)
	}

	clientArgs := keyvalues.NewClientAzureADArgs{
		ClientID:         content["clientId"],
		ClientSecret:     content["clientSecret"],
		TenantID:         content["tenantId"],
		AADEndpoint:      content["activeDirectoryEndpointUrl"],
		ResourceEndpoint: cr.Spec.ForProvider.Endpoint,
	}

	cl, err := keyvalues.NewClientAzureAD(clientArgs)
	if err != nil {
		return nil, errors.Wrap(err, errCannotCreateNewClient)
	}
	return &external{kube: c.kube, client: cl}, nil
}

type external struct {
	kube   client.Client
	client keyvalues.Client
}

func (c *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*v1alpha1.KeyValue)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotKeyValue)
	}

	key := cr.Spec.ForProvider.Key
	label := cr.Spec.ForProvider.Label

	kv, err := c.client.GetKeyValue(key, label)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return managed.ExternalObservation{}, nil
		}
		return managed.ExternalObservation{}, errors.Wrap(err, errGetFailed)
	}

	// Import/Sync
	lateInit := false
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	appclient.LateInitialize(&cr.Spec.ForProvider, kv)
	if !cmp.Equal(currentSpec, &cr.Spec.ForProvider) {
		lateInit = true
	}

	cr.Status.SetConditions(xpv1.Available())
	appclient.GenerateObservation(&cr.Status.AtProvider, kv)

	upToDate, err := appclient.IsUpToDate(cr.Spec.ForProvider, &kv)
	if err != nil {
		return managed.ExternalObservation{}, err
	}

	return managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		ResourceLateInitialized: lateInit,
	}, nil
}

func (c *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*v1alpha1.KeyValue)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errNotKeyValue)
	}

	return managed.ExternalCreation{}, errors.Wrap(c.SetKeyValue(cr.Spec.ForProvider), errCreateFailed)
}

func (c *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*v1alpha1.KeyValue)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errNotKeyValue)
	}
	return managed.ExternalUpdate{}, errors.Wrap(c.SetKeyValue(cr.Spec.ForProvider), errUpdateFailed)
}

func (c *external) Delete(ctx context.Context, mg resource.Managed) error {
	cr, ok := mg.(*v1alpha1.KeyValue)
	if !ok {
		return errors.New(errNotKeyValue)
	}

	err := c.client.DeleteKeyValue(cr.Spec.ForProvider.Key, cr.Spec.ForProvider.Label)
	if err != nil {
		return errors.Wrap(err, errDeleteFailed)
	}

	return nil
}

func (c *external) SetKeyValue(params v1alpha1.KeyValueParameters) error {
	instance := keyvalues.CreateOrUpdateKeyValueArgs{
		Key:   params.Key,
		Value: params.Value,
		Label: params.Label,
	}
	if params.Tags != nil {
		instance.Tags = *params.Tags
	}
	if params.ContentType != nil {
		instance.ContentType = *params.ContentType
	}

	_, err := c.client.CreateOrUpdateKeyValue(instance)
	return err
}
