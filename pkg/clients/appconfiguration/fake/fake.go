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

package fake

import (
	"github.com/stone-payments/appconfig-go-sdk/appconfig/keyvalues"
)

var _ keyvalues.Client = &MockClient{}

// MockClient is a fake implementation of keyvalues.KeyValue.
type MockClient struct {
	keyvalues.Client

	MockCreateOrUpdateKeyValue func(kvargs keyvalues.CreateOrUpdateKeyValueArgs) (keyvalues.KeyValue, error)
	MockDeleteKeyValue         func(key, label string) error
	MockGetKeyValue            func(key, label string) (keyvalues.KeyValue, error)
}

// CreateOrUpdateKeyValue calls the MockClient's MockCreateOrUpdateKeyValue method.
func (c *MockClient) CreateOrUpdateKeyValue(kvargs keyvalues.CreateOrUpdateKeyValueArgs) (keyvalues.KeyValue, error) {
	return c.MockCreateOrUpdateKeyValue(kvargs)
}

// DeleteKeyValue calls the MockClient's MockDeleteKeyValue method.
func (c *MockClient) DeleteKeyValue(key, label string) error {
	return c.MockDeleteKeyValue(key, label)
}

// GetKeyValue calls the MockClient's MockGetKeyValue method.
func (c *MockClient) GetKeyValue(key, label string) (keyvalues.KeyValue, error) {
	return c.MockGetKeyValue(key, label)
}
