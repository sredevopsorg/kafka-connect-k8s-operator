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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/sredevopsdev/kafka-connect-k8s-operator/pkg/apis/kafkaconnect/v1alpha1"
	scheme "github.com/sredevopsdev/kafka-connect-k8s-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KafkaConnectsGetter has a method to return a KafkaConnectInterface.
// A group's client should implement this interface.
type KafkaConnectsGetter interface {
	KafkaConnects(namespace string) KafkaConnectInterface
}

// KafkaConnectInterface has methods to work with KafkaConnect resources.
type KafkaConnectInterface interface {
	Create(*v1alpha1.KafkaConnect) (*v1alpha1.KafkaConnect, error)
	Update(*v1alpha1.KafkaConnect) (*v1alpha1.KafkaConnect, error)
	UpdateStatus(*v1alpha1.KafkaConnect) (*v1alpha1.KafkaConnect, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.KafkaConnect, error)
	List(opts v1.ListOptions) (*v1alpha1.KafkaConnectList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KafkaConnect, err error)
	KafkaConnectExpansion
}

// kafkaConnects implements KafkaConnectInterface
type kafkaConnects struct {
	client rest.Interface
	ns     string
}

// newKafkaConnects returns a KafkaConnects
func newKafkaConnects(c *KafkaconnectV1alpha1Client, namespace string) *kafkaConnects {
	return &kafkaConnects{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kafkaConnect, and returns the corresponding kafkaConnect object, and an error if there is any.
func (c *kafkaConnects) Get(name string, options v1.GetOptions) (result *v1alpha1.KafkaConnect, err error) {
	result = &v1alpha1.KafkaConnect{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kafkaconnects").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KafkaConnects that match those selectors.
func (c *kafkaConnects) List(opts v1.ListOptions) (result *v1alpha1.KafkaConnectList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KafkaConnectList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kafkaconnects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kafkaConnects.
func (c *kafkaConnects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kafkaconnects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kafkaConnect and creates it.  Returns the server's representation of the kafkaConnect, and an error, if there is any.
func (c *kafkaConnects) Create(kafkaConnect *v1alpha1.KafkaConnect) (result *v1alpha1.KafkaConnect, err error) {
	result = &v1alpha1.KafkaConnect{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kafkaconnects").
		Body(kafkaConnect).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kafkaConnect and updates it. Returns the server's representation of the kafkaConnect, and an error, if there is any.
func (c *kafkaConnects) Update(kafkaConnect *v1alpha1.KafkaConnect) (result *v1alpha1.KafkaConnect, err error) {
	result = &v1alpha1.KafkaConnect{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kafkaconnects").
		Name(kafkaConnect.Name).
		Body(kafkaConnect).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kafkaConnects) UpdateStatus(kafkaConnect *v1alpha1.KafkaConnect) (result *v1alpha1.KafkaConnect, err error) {
	result = &v1alpha1.KafkaConnect{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kafkaconnects").
		Name(kafkaConnect.Name).
		SubResource("status").
		Body(kafkaConnect).
		Do().
		Into(result)
	return
}

// Delete takes name of the kafkaConnect and deletes it. Returns an error if one occurs.
func (c *kafkaConnects) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kafkaconnects").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kafkaConnects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kafkaconnects").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kafkaConnect.
func (c *kafkaConnects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KafkaConnect, err error) {
	result = &v1alpha1.KafkaConnect{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kafkaconnects").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
