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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/sredevopsdev/kafka-connect-k8s-operator/pkg/apis/kafkaconnect/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KafkaConnectLister helps list KafkaConnects.
type KafkaConnectLister interface {
	// List lists all KafkaConnects in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.KafkaConnect, err error)
	// KafkaConnects returns an object that can list and get KafkaConnects.
	KafkaConnects(namespace string) KafkaConnectNamespaceLister
	KafkaConnectListerExpansion
}

// kafkaConnectLister implements the KafkaConnectLister interface.
type kafkaConnectLister struct {
	indexer cache.Indexer
}

// NewKafkaConnectLister returns a new KafkaConnectLister.
func NewKafkaConnectLister(indexer cache.Indexer) KafkaConnectLister {
	return &kafkaConnectLister{indexer: indexer}
}

// List lists all KafkaConnects in the indexer.
func (s *kafkaConnectLister) List(selector labels.Selector) (ret []*v1alpha1.KafkaConnect, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KafkaConnect))
	})
	return ret, err
}

// KafkaConnects returns an object that can list and get KafkaConnects.
func (s *kafkaConnectLister) KafkaConnects(namespace string) KafkaConnectNamespaceLister {
	return kafkaConnectNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KafkaConnectNamespaceLister helps list and get KafkaConnects.
type KafkaConnectNamespaceLister interface {
	// List lists all KafkaConnects in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.KafkaConnect, err error)
	// Get retrieves the KafkaConnect from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.KafkaConnect, error)
	KafkaConnectNamespaceListerExpansion
}

// kafkaConnectNamespaceLister implements the KafkaConnectNamespaceLister
// interface.
type kafkaConnectNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KafkaConnects in the indexer for a given namespace.
func (s kafkaConnectNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.KafkaConnect, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KafkaConnect))
	})
	return ret, err
}

// Get retrieves the KafkaConnect from the indexer for a given namespace and name.
func (s kafkaConnectNamespaceLister) Get(name string) (*v1alpha1.KafkaConnect, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kafkaconnect"), name)
	}
	return obj.(*v1alpha1.KafkaConnect), nil
}
