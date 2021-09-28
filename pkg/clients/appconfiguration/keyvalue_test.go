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
	"testing"

	"github.com/crossplane/crossplane-runtime/pkg/test"
	"github.com/google/go-cmp/cmp"
	"github.com/stone-payments/appconfig-go-sdk/appconfig/keyvalues"

	"github.com/crossplane/provider-azure/apis/appconfiguration/v1alpha1"
	azure "github.com/crossplane/provider-azure/pkg/clients"
)

var (
	tags         = map[string]string{"created_by": "crossplane"}
	contentType  = azure.ToStringPtr("text/plain")
	locked       = azure.ToBoolPtr(true)
	value        = "the-secret-value"
	etag         = "d00c90f7-971e-4a72-8bf9-de4c6ff95670"
	lastModified = "03/21/3012"
	label        = "the-label"
	key          = "the-key"
)

func keyvalue() *keyvalues.KeyValue {
	return &keyvalues.KeyValue{
		LastModified: &lastModified,
		Etag:         &etag,
		Key:          &key,
		Value:        &value,
		Label:        &label,
		Locked:       locked,
		Tags:         &tags,
	}
}

func params(isLateInitialize bool) *v1alpha1.KeyValueParameters {
	params := &v1alpha1.KeyValueParameters{
		Key:    key,
		Value:  value,
		Label:  label,
		Locked: locked,
		Tags:   tags,
	}
	if isLateInitialize {
		params.Key = ""
		params.Value = ""
		params.Label = ""
		params.Locked = nil
		params.Tags = nil
	}
	return params
}

func TestGenerateObservation(t *testing.T) {
	type args struct {
		obser    v1alpha1.KeyValueObservation
		external keyvalues.KeyValue
	}
	cases := map[string]struct {
		args
		want v1alpha1.KeyValueObservation
	}{
		"FullConversion": {
			args: args{
				obser: v1alpha1.KeyValueObservation{},
				external: keyvalues.KeyValue{
					Etag:         &etag,
					LastModified: &lastModified,
				},
			},
			want: v1alpha1.KeyValueObservation{
				Etag:         etag,
				LastModified: lastModified,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			GenerateObservation(&tc.args.obser, tc.args.external)
			if diff := cmp.Diff(tc.args.obser, tc.want); diff != "" {
				t.Errorf("GenerateObservation(...): -want, +got:\n%s", diff)
			}
		})
	}
}

func TestLateInitialize(t *testing.T) {
	type args struct {
		az   keyvalues.KeyValue
		spec *v1alpha1.KeyValueParameters
	}
	cases := map[string]struct {
		args
		want *v1alpha1.KeyValueParameters
	}{
		"Must use template fields in initialization": {
			args: args{
				spec: &v1alpha1.KeyValueParameters{
					Tags:        tags,
					ContentType: contentType,
					Locked:      locked,
				},
				az: keyvalues.KeyValue{
					Tags:        &map[string]string{"created_by": "crossplane"},
					ContentType: azure.ToStringPtr("application/json"),
					Locked:      locked,
				},
			},
			want: &v1alpha1.KeyValueParameters{
				Tags:        tags,
				ContentType: contentType,
				Locked:      locked,
			},
		},
		"Must initialize template spec field in initialization": {
			args: args{
				spec: &v1alpha1.KeyValueParameters{},
				az: keyvalues.KeyValue{
					ContentType: contentType,
					Tags:        &tags,
					Locked:      locked,
				},
			},
			want: &v1alpha1.KeyValueParameters{
				ContentType: contentType,
				Tags:        tags,
				Locked:      locked,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			LateInitialize(tc.args.spec, tc.args.az)
			if diff := cmp.Diff(tc.args.spec, tc.want); diff != "" {
				t.Errorf("LateInitialize(...): -want, +got:\n%s", diff)
			}
		})
	}
}

func TestIsUpToDate(t *testing.T) {
	type args struct {
		observed *keyvalues.KeyValue
		params   *v1alpha1.KeyValueParameters
	}
	type out struct {
		upToDate bool
		err      error
	}
	cases := map[string]struct {
		reason string
		args
		out
	}{
		"DefinitionUpToDate": {
			reason: "Must return true if the definition is up to date",
			args: args{
				params:   params(false),
				observed: keyvalue(),
			},
			out: out{
				upToDate: true,
				err:      nil,
			},
		},
		"DefinitionNotUpToDate": {
			reason: "Must return false if the definition is not up to date",
			args: args{
				params:   params(false),
				observed: &keyvalues.KeyValue{},
			},
			out: out{
				upToDate: false,
				err:      nil,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			upToDate, err := IsUpToDate(*tc.args.params, tc.args.observed)
			if diff := cmp.Diff(tc.out.upToDate, upToDate); diff != "" {
				t.Errorf("IsUpToDate(...): -want, +got:\n%s", diff)
			}
			if diff := cmp.Diff(tc.out.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("IsUpToDate(...): -want error, +got error:\n%s", diff)
			}
		})
	}
}

func TestOverrideParameters(t *testing.T) {
	type args struct {
		spec    v1alpha1.KeyValueParameters
		desired keyvalues.KeyValue
	}
	cases := map[string]struct {
		reason string
		args
		want keyvalues.KeyValue
	}{
		"GenerateDesiredSuccessful": {
			reason: "Must generate a desired successfully",
			args: args{
				spec: v1alpha1.KeyValueParameters{
					ContentType: contentType,
					Value:       value,
					Key:         key,
					Label:       label,
					Tags:        tags,
					Locked:      locked,
				},
				desired: keyvalues.KeyValue{},
			},
			want: keyvalues.KeyValue{},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			overrideParameters(tc.args.spec, tc.args.desired)
			if diff := cmp.Diff(tc.args.desired, tc.want); diff != "" {
				t.Errorf("overrideParameters(...): -want, +got:\n%s", diff)
			}
		})
	}
}
