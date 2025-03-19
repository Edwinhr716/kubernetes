/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

// EndpointHintsApplyConfiguration represents a declarative configuration of the EndpointHints type for use
// with apply.
type EndpointHintsApplyConfiguration struct {
	ForZones []ForZoneApplyConfiguration `json:"forZones,omitempty"`
	ForNodes []ForNodeApplyConfiguration `json:"forNodes,omitempty"`
}

// EndpointHintsApplyConfiguration constructs a declarative configuration of the EndpointHints type for use with
// apply.
func EndpointHints() *EndpointHintsApplyConfiguration {
	return &EndpointHintsApplyConfiguration{}
}

// WithForZones adds the given value to the ForZones field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ForZones field.
func (b *EndpointHintsApplyConfiguration) WithForZones(values ...*ForZoneApplyConfiguration) *EndpointHintsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithForZones")
		}
		b.ForZones = append(b.ForZones, *values[i])
	}
	return b
}

// WithForNodes adds the given value to the ForNodes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ForNodes field.
func (b *EndpointHintsApplyConfiguration) WithForNodes(values ...*ForNodeApplyConfiguration) *EndpointHintsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithForNodes")
		}
		b.ForNodes = append(b.ForNodes, *values[i])
	}
	return b
}
