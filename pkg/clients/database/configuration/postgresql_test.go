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

package configuration

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/google/go-cmp/cmp"

	"github.com/crossplane/provider-azure/apis/database/v1beta1"
	azuredbv1beta1 "github.com/crossplane/provider-azure/apis/database/v1beta1"
)

const (
	testValue1 = "testValue1"
	testValue2 = "testValue2"
)

type sqlServerConfigurationParametersModifier func(*azuredbv1beta1.SQLServerConfigurationParameters)

func sqlServerConfigurationParametersWithValue(v *string) sqlServerConfigurationParametersModifier {
	return func(cm *azuredbv1beta1.SQLServerConfigurationParameters) {
		cm.Value = v
	}
}

func sqlServerConfigurationParameters(sm ...sqlServerConfigurationParametersModifier) *azuredbv1beta1.SQLServerConfigurationParameters {
	cm := &azuredbv1beta1.SQLServerConfigurationParameters{}
	for _, m := range sm {
		m(cm)
	}
	return cm
}

type configurationModifier func(configuration *postgresql.Configuration)

func configurationWithValue(v *string) configurationModifier {
	return func(configuration *postgresql.Configuration) {
		configuration.Value = v
	}
}

func configuration(cm ...configurationModifier) *postgresql.Configuration {
	c := &postgresql.Configuration{
		ConfigurationProperties: &postgresql.ConfigurationProperties{},
	}
	for _, m := range cm {
		m(c)
	}
	return c
}

func TestIsPostgreSQLConfigurationUpToDate(t *testing.T) {
	val1, val2 := testValue1, testValue2
	type args struct {
		p  v1beta1.SQLServerConfigurationParameters
		in postgresql.Configuration
	}
	tests := map[string]struct {
		args args
		want bool
	}{
		"UpToDate": {
			args: args{
				p: *sqlServerConfigurationParameters(
					sqlServerConfigurationParametersWithValue(&val1)),
				in: *configuration(configurationWithValue(&val1)),
			},
			want: true,
		},
		"NeedsUpdate": {
			args: args{
				p: *sqlServerConfigurationParameters(
					sqlServerConfigurationParametersWithValue(&val1)),
				in: *configuration(configurationWithValue(&val2)),
			},
			want: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := IsPostgreSQLConfigurationUpToDate(tt.args.p, tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("IsPostgreSQLConfigurationUpToDate(...): -want, +got\n%s", diff)
			}
		})
	}
}
