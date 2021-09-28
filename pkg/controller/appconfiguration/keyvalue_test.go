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
	"testing"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/crossplane-runtime/pkg/test"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"github.com/stone-payments/appconfig-go-sdk/appconfig/keyvalues"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/provider-azure/apis/appconfiguration/v1alpha1"
	azure "github.com/crossplane/provider-azure/pkg/clients"
	"github.com/crossplane/provider-azure/pkg/clients/appconfiguration/fake"
)

var (
	tags             = map[string]string{"created_by": "crossplane"}
	contentType      = azure.ToStringPtr("text/plain")
	value            = "the-secret-value"
	unexpectedObject resource.Managed
	locked           = azure.ToBoolPtr(true)
	label            = "the-label"
	key              = "the-key"
)

var (
	errorBoom = errors.New("boom")
)

type keyvaluesResourceModifier func(*v1alpha1.KeyValue)

func withConditions(c ...xpv1.Condition) keyvaluesResourceModifier {
	return func(r *v1alpha1.KeyValue) { r.Status.ConditionedStatus.Conditions = c }
}

func withContentType(c *string) keyvaluesResourceModifier {
	return func(r *v1alpha1.KeyValue) { r.Spec.ForProvider.ContentType = c }
}

func instance(rm ...keyvaluesResourceModifier) *v1alpha1.KeyValue {
	cr := &v1alpha1.KeyValue{
		Spec: v1alpha1.KeyValueSpec{
			ForProvider: v1alpha1.KeyValueParameters{
				Value:  value,
				Label:  label,
				Key:    key,
				Tags:   tags,
				Locked: locked,
			},
		},
	}

	for _, m := range rm {
		m(cr)
	}

	return cr
}

func TestObserve(t *testing.T) {
	type args struct {
		cr   resource.Managed
		kv   keyvalues.Client
		kube client.Client
	}
	type want struct {
		cr  resource.Managed
		o   managed.ExternalObservation
		err error
	}

	cases := map[string]struct {
		args
		want
	}{
		"ResourceIsNotKeyValue": {
			args: args{
				cr: unexpectedObject,
			},
			want: want{
				o:   managed.ExternalObservation{},
				err: errors.New(errNotKeyValue),
			},
		},
		"Successful": {
			args: args{
				cr: instance(),
				kv: &fake.MockClient{
					MockGetKeyValue: func(key string, label string) (keyvalues.KeyValue, error) {
						return keyvalues.KeyValue{
							Value:  &value,
							Label:  &label,
							Key:    &key,
							Tags:   &tags,
							Locked: locked,
						}, nil
					},
				},
			},
			want: want{
				cr: instance(
					withConditions(xpv1.Available()),
				),
				o:   managed.ExternalObservation{ResourceExists: true, ResourceUpToDate: true},
				err: nil,
			},
		},
		"GetNotFound": {
			args: args{
				cr: instance(),
				kv: &fake.MockClient{
					MockGetKeyValue: func(key string, label string) (keyvalues.KeyValue, error) {
						return keyvalues.KeyValue{}, errors.New("404")
					},
				},
			},
			want: want{
				cr:  instance(),
				o:   managed.ExternalObservation{},
				err: nil,
			},
		},
		"GetInternalError": {
			args: args{
				cr: instance(),
				kv: &fake.MockClient{
					MockGetKeyValue: func(key string, label string) (keyvalues.KeyValue, error) {
						return keyvalues.KeyValue{}, errorBoom
					},
				},
			},
			want: want{
				cr:  instance(),
				o:   managed.ExternalObservation{},
				err: errors.Wrap(errorBoom, errGetFailed),
			},
		},
		"LateInitialized": {
			args: args{
				cr: instance(),
				kube: &test.MockClient{
					MockUpdate: test.NewMockUpdateFn(nil),
				},
				kv: &fake.MockClient{
					MockGetKeyValue: func(key string, label string) (keyvalues.KeyValue, error) {
						return keyvalues.KeyValue{
							Value:       &value,
							Label:       &label,
							Key:         &key,
							Tags:        &tags,
							Locked:      locked,
							ContentType: contentType,
						}, nil
					},
				},
			},
			want: want{
				cr: instance(
					withConditions(xpv1.Available()),
					withContentType(contentType),
				),
				o: managed.ExternalObservation{
					ResourceExists:          true,
					ResourceLateInitialized: true,
					ResourceUpToDate:        true,
				},
				err: nil,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			e := external{
				kube:   tc.kube,
				client: tc.args.kv,
			}
			o, err := e.Observe(context.Background(), tc.args.cr)
			if diff := cmp.Diff(tc.want.cr, tc.args.cr); diff != "" {
				t.Errorf("Observe(...): -want, +got\n%s", diff)
			}
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("Observe(...): -want, +got\n%s", diff)
			}
			if diff := cmp.Diff(tc.want.o, o); diff != "" {
				t.Errorf("Observe(...): -want, +got\n%s", diff)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	type args struct {
		cr resource.Managed
		kv keyvalues.Client
	}
	type want struct {
		cr  resource.Managed
		o   managed.ExternalCreation
		err error
	}

	cases := map[string]struct {
		args
		want
	}{
		"ResourceIsNotKeyValue": {
			args: args{
				cr: unexpectedObject,
			},
			want: want{
				o:   managed.ExternalCreation{},
				err: errors.New(errNotKeyValue),
			},
		},
		"Successful": {
			args: args{
				cr: instance(
					withContentType(contentType),
				),
				kv: &fake.MockClient{
					MockCreateOrUpdateKeyValue: func(kvargs keyvalues.CreateOrUpdateKeyValueArgs) (keyvalues.KeyValue, error) {
						return keyvalues.KeyValue{}, nil
					},
				},
			},
			want: want{
				o: managed.ExternalCreation{},
				cr: instance(
					withContentType(contentType),
				),
			},
		},
		"Failed": {
			args: args{
				cr: instance(),
				kv: &fake.MockClient{
					MockCreateOrUpdateKeyValue: func(kvargs keyvalues.CreateOrUpdateKeyValueArgs) (keyvalues.KeyValue, error) {
						return keyvalues.KeyValue{}, errorBoom
					},
				},
			},
			want: want{
				cr:  instance(),
				o:   managed.ExternalCreation{},
				err: errors.Wrap(errorBoom, errCreateFailed),
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			e := external{client: tc.args.kv}

			c, err := e.Create(context.Background(), tc.args.cr)
			if diff := cmp.Diff(tc.want.cr, tc.args.cr); diff != "" {
				t.Errorf("Create(...): -want, +got\n%s", diff)
			}
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("Create(...): -want, +got\n%s", diff)
			}
			if diff := cmp.Diff(tc.want.o, c); diff != "" {
				t.Errorf("Create(...): -want, +got\n%s", diff)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	type args struct {
		cr resource.Managed
		kv keyvalues.Client
	}
	type want struct {
		cr  resource.Managed
		o   managed.ExternalUpdate
		err error
	}

	cases := map[string]struct {
		args
		want
	}{
		"ResourceIsNotKeyValue": {
			args: args{
				cr: unexpectedObject,
			},
			want: want{
				o:   managed.ExternalUpdate{},
				err: errors.New(errNotKeyValue),
			},
		},
		"Successful": {
			args: args{
				cr: instance(),
				kv: &fake.MockClient{
					MockCreateOrUpdateKeyValue: func(kvargs keyvalues.CreateOrUpdateKeyValueArgs) (keyvalues.KeyValue, error) {
						return keyvalues.KeyValue{}, nil
					},
				},
			},
			want: want{
				o:  managed.ExternalUpdate{},
				cr: instance(),
			},
		},
		"Failed": {
			args: args{
				cr: instance(),
				kv: &fake.MockClient{
					MockCreateOrUpdateKeyValue: func(kvargs keyvalues.CreateOrUpdateKeyValueArgs) (keyvalues.KeyValue, error) {
						return keyvalues.KeyValue{}, errorBoom
					},
				},
			},
			want: want{
				cr:  instance(),
				o:   managed.ExternalUpdate{},
				err: errors.Wrap(errorBoom, errUpdateFailed),
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			e := external{client: tc.args.kv}

			c, err := e.Update(context.Background(), tc.args.cr)
			if diff := cmp.Diff(tc.want.cr, tc.args.cr); diff != "" {
				t.Errorf("Update(...): -want, +got\n%s", diff)
			}
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("Update(...): -want, +got\n%s", diff)
			}
			if diff := cmp.Diff(tc.want.o, c); diff != "" {
				t.Errorf("Update(...): -want, +got\n%s", diff)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		cr resource.Managed
		kv keyvalues.Client
	}
	type want struct {
		cr  resource.Managed
		err error
	}

	cases := map[string]struct {
		args
		want
	}{
		"ResourceIsNotKeyValue": {
			args: args{
				cr: unexpectedObject,
			},
			want: want{
				err: errors.New(errNotKeyValue),
			},
		},
		"Successful": {
			args: args{
				cr: instance(),
				kv: &fake.MockClient{
					MockDeleteKeyValue: func(key, label string) error {
						return nil
					},
				},
			},
			want: want{
				cr:  instance(),
				err: nil,
			},
		},
		"Failed": {
			args: args{
				cr: instance(),
				kv: &fake.MockClient{
					MockDeleteKeyValue: func(key, label string) error {
						return errorBoom
					},
				},
			},
			want: want{
				cr:  instance(),
				err: errors.Wrap(errorBoom, errDeleteFailed),
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			e := external{client: tc.args.kv}

			err := e.Delete(context.Background(), tc.args.cr)
			if diff := cmp.Diff(tc.want.cr, tc.args.cr); diff != "" {
				t.Errorf("Delete(...): -want, +got\n%s", diff)
			}
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("Delete(...): -want, +got\n%s", diff)
			}
		})
	}
}
