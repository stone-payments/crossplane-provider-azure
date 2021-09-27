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
	"github.com/google/go-cmp/cmp"
	"github.com/mitchellh/copystructure"
	"github.com/pkg/errors"
	"github.com/stone-payments/appconfig-go-sdk/appconfig/keyvalues"

	"github.com/crossplane/provider-azure/apis/appconfiguration/v1alpha1"
)

const (
	errCheckUpToDate = "unable to determine if external resource is up to date"
)

// LateInitialize fills the cr values that user did not fill with their
// corresponding value in the Azure, if there is any.
func LateInitialize(cr *v1alpha1.KeyValueParameters, p keyvalues.KeyValue) {
	if cr.ContentType == nil && p.ContentType != nil {
		cr.ContentType = p.ContentType
	}

	if cr.Locked == nil && p.Locked != nil {
		cr.Locked = p.Locked
	}

	if cr.Tags == nil && p.Tags != nil {
		cr.Tags = p.Tags
	}
}

// IsUpToDate checks whether KeyValue spec is up to date with remote resource.
func IsUpToDate(in v1alpha1.KeyValueParameters, observed *keyvalues.KeyValue) (bool, error) {
	clone, err := copystructure.Copy(observed)
	if err != nil {
		return true, errors.Wrap(err, errCheckUpToDate)
	}
	external, ok := clone.(*keyvalues.KeyValue)
	if !ok {
		return true, errors.New(errCheckUpToDate)
	}

	desired := overrideParameters(in, *external)

	return cmp.Equal(desired, *observed), nil
}

func overrideParameters(params v1alpha1.KeyValueParameters, desired keyvalues.KeyValue) keyvalues.KeyValue {
	desired.Value = &params.Value
	desired.Label = &params.Label
	desired.Key = &params.Key

	if params.ContentType != nil {
		desired.ContentType = params.ContentType
	}

	if params.Tags != nil {
		desired.Tags = &map[string]string{}
		for key, value := range *params.Tags {
			(*desired.Tags)[key] = value
		}
	}

	if params.Locked != nil {
		desired.Locked = params.Locked
	}

	return desired
}

// GenerateObservation used to Status
func GenerateObservation(cr *v1alpha1.KeyValueObservation, kv keyvalues.KeyValue) {

	if cr.Etag == "" && kv.Etag != nil {
		cr.Etag = *kv.Etag
	}

	if cr.LastModified == "" && kv.LastModified != nil {
		cr.LastModified = *kv.LastModified
	}
}
