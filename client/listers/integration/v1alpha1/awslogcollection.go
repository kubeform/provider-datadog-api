/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-datadog-api/apis/integration/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AwsLogCollectionLister helps list AwsLogCollections.
// All objects returned here must be treated as read-only.
type AwsLogCollectionLister interface {
	// List lists all AwsLogCollections in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AwsLogCollection, err error)
	// AwsLogCollections returns an object that can list and get AwsLogCollections.
	AwsLogCollections(namespace string) AwsLogCollectionNamespaceLister
	AwsLogCollectionListerExpansion
}

// awsLogCollectionLister implements the AwsLogCollectionLister interface.
type awsLogCollectionLister struct {
	indexer cache.Indexer
}

// NewAwsLogCollectionLister returns a new AwsLogCollectionLister.
func NewAwsLogCollectionLister(indexer cache.Indexer) AwsLogCollectionLister {
	return &awsLogCollectionLister{indexer: indexer}
}

// List lists all AwsLogCollections in the indexer.
func (s *awsLogCollectionLister) List(selector labels.Selector) (ret []*v1alpha1.AwsLogCollection, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsLogCollection))
	})
	return ret, err
}

// AwsLogCollections returns an object that can list and get AwsLogCollections.
func (s *awsLogCollectionLister) AwsLogCollections(namespace string) AwsLogCollectionNamespaceLister {
	return awsLogCollectionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsLogCollectionNamespaceLister helps list and get AwsLogCollections.
// All objects returned here must be treated as read-only.
type AwsLogCollectionNamespaceLister interface {
	// List lists all AwsLogCollections in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AwsLogCollection, err error)
	// Get retrieves the AwsLogCollection from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AwsLogCollection, error)
	AwsLogCollectionNamespaceListerExpansion
}

// awsLogCollectionNamespaceLister implements the AwsLogCollectionNamespaceLister
// interface.
type awsLogCollectionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsLogCollections in the indexer for a given namespace.
func (s awsLogCollectionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsLogCollection, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsLogCollection))
	})
	return ret, err
}

// Get retrieves the AwsLogCollection from the indexer for a given namespace and name.
func (s awsLogCollectionNamespaceLister) Get(name string) (*v1alpha1.AwsLogCollection, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awslogcollection"), name)
	}
	return obj.(*v1alpha1.AwsLogCollection), nil
}
