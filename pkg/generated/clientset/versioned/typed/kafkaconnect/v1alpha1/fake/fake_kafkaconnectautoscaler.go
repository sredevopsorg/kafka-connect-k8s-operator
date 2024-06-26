// Copyright 2020 The Amadeus Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/sredevopsdev/kafka-connect-k8s-operator/pkg/apis/kafkaconnect/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKafkaConnectAutoScalers implements KafkaConnectAutoScalerInterface
type FakeKafkaConnectAutoScalers struct {
	Fake *FakeKafkaconnectV1alpha1
	ns   string
}

var kafkaconnectautoscalersResource = schema.GroupVersionResource{Group: "kafkaconnect.operator.io", Version: "v1alpha1", Resource: "kafkaconnectautoscalers"}

var kafkaconnectautoscalersKind = schema.GroupVersionKind{Group: "kafkaconnect.operator.io", Version: "v1alpha1", Kind: "KafkaConnectAutoScaler"}

// Get takes name of the kafkaConnectAutoScaler, and returns the corresponding kafkaConnectAutoScaler object, and an error if there is any.
func (c *FakeKafkaConnectAutoScalers) Get(name string, options v1.GetOptions) (result *v1alpha1.KafkaConnectAutoScaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kafkaconnectautoscalersResource, c.ns, name), &v1alpha1.KafkaConnectAutoScaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KafkaConnectAutoScaler), err
}

// List takes label and field selectors, and returns the list of KafkaConnectAutoScalers that match those selectors.
func (c *FakeKafkaConnectAutoScalers) List(opts v1.ListOptions) (result *v1alpha1.KafkaConnectAutoScalerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kafkaconnectautoscalersResource, kafkaconnectautoscalersKind, c.ns, opts), &v1alpha1.KafkaConnectAutoScalerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KafkaConnectAutoScalerList{ListMeta: obj.(*v1alpha1.KafkaConnectAutoScalerList).ListMeta}
	for _, item := range obj.(*v1alpha1.KafkaConnectAutoScalerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kafkaConnectAutoScalers.
func (c *FakeKafkaConnectAutoScalers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kafkaconnectautoscalersResource, c.ns, opts))

}

// Create takes the representation of a kafkaConnectAutoScaler and creates it.  Returns the server's representation of the kafkaConnectAutoScaler, and an error, if there is any.
func (c *FakeKafkaConnectAutoScalers) Create(kafkaConnectAutoScaler *v1alpha1.KafkaConnectAutoScaler) (result *v1alpha1.KafkaConnectAutoScaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kafkaconnectautoscalersResource, c.ns, kafkaConnectAutoScaler), &v1alpha1.KafkaConnectAutoScaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KafkaConnectAutoScaler), err
}

// Update takes the representation of a kafkaConnectAutoScaler and updates it. Returns the server's representation of the kafkaConnectAutoScaler, and an error, if there is any.
func (c *FakeKafkaConnectAutoScalers) Update(kafkaConnectAutoScaler *v1alpha1.KafkaConnectAutoScaler) (result *v1alpha1.KafkaConnectAutoScaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kafkaconnectautoscalersResource, c.ns, kafkaConnectAutoScaler), &v1alpha1.KafkaConnectAutoScaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KafkaConnectAutoScaler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKafkaConnectAutoScalers) UpdateStatus(kafkaConnectAutoScaler *v1alpha1.KafkaConnectAutoScaler) (*v1alpha1.KafkaConnectAutoScaler, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kafkaconnectautoscalersResource, "status", c.ns, kafkaConnectAutoScaler), &v1alpha1.KafkaConnectAutoScaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KafkaConnectAutoScaler), err
}

// Delete takes name of the kafkaConnectAutoScaler and deletes it. Returns an error if one occurs.
func (c *FakeKafkaConnectAutoScalers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kafkaconnectautoscalersResource, c.ns, name), &v1alpha1.KafkaConnectAutoScaler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKafkaConnectAutoScalers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kafkaconnectautoscalersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.KafkaConnectAutoScalerList{})
	return err
}

// Patch applies the patch and returns the patched kafkaConnectAutoScaler.
func (c *FakeKafkaConnectAutoScalers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KafkaConnectAutoScaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kafkaconnectautoscalersResource, c.ns, name, pt, data, subresources...), &v1alpha1.KafkaConnectAutoScaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KafkaConnectAutoScaler), err
}
