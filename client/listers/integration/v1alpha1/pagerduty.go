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

// PagerdutyLister helps list Pagerduties.
// All objects returned here must be treated as read-only.
type PagerdutyLister interface {
	// List lists all Pagerduties in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Pagerduty, err error)
	// Pagerduties returns an object that can list and get Pagerduties.
	Pagerduties(namespace string) PagerdutyNamespaceLister
	PagerdutyListerExpansion
}

// pagerdutyLister implements the PagerdutyLister interface.
type pagerdutyLister struct {
	indexer cache.Indexer
}

// NewPagerdutyLister returns a new PagerdutyLister.
func NewPagerdutyLister(indexer cache.Indexer) PagerdutyLister {
	return &pagerdutyLister{indexer: indexer}
}

// List lists all Pagerduties in the indexer.
func (s *pagerdutyLister) List(selector labels.Selector) (ret []*v1alpha1.Pagerduty, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Pagerduty))
	})
	return ret, err
}

// Pagerduties returns an object that can list and get Pagerduties.
func (s *pagerdutyLister) Pagerduties(namespace string) PagerdutyNamespaceLister {
	return pagerdutyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PagerdutyNamespaceLister helps list and get Pagerduties.
// All objects returned here must be treated as read-only.
type PagerdutyNamespaceLister interface {
	// List lists all Pagerduties in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Pagerduty, err error)
	// Get retrieves the Pagerduty from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Pagerduty, error)
	PagerdutyNamespaceListerExpansion
}

// pagerdutyNamespaceLister implements the PagerdutyNamespaceLister
// interface.
type pagerdutyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Pagerduties in the indexer for a given namespace.
func (s pagerdutyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Pagerduty, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Pagerduty))
	})
	return ret, err
}

// Get retrieves the Pagerduty from the indexer for a given namespace and name.
func (s pagerdutyNamespaceLister) Get(name string) (*v1alpha1.Pagerduty, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("pagerduty"), name)
	}
	return obj.(*v1alpha1.Pagerduty), nil
}
