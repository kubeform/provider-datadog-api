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
	v1alpha1 "kubeform.dev/provider-datadog-api/apis/logs/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ArchiveLister helps list Archives.
// All objects returned here must be treated as read-only.
type ArchiveLister interface {
	// List lists all Archives in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Archive, err error)
	// Archives returns an object that can list and get Archives.
	Archives(namespace string) ArchiveNamespaceLister
	ArchiveListerExpansion
}

// archiveLister implements the ArchiveLister interface.
type archiveLister struct {
	indexer cache.Indexer
}

// NewArchiveLister returns a new ArchiveLister.
func NewArchiveLister(indexer cache.Indexer) ArchiveLister {
	return &archiveLister{indexer: indexer}
}

// List lists all Archives in the indexer.
func (s *archiveLister) List(selector labels.Selector) (ret []*v1alpha1.Archive, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Archive))
	})
	return ret, err
}

// Archives returns an object that can list and get Archives.
func (s *archiveLister) Archives(namespace string) ArchiveNamespaceLister {
	return archiveNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ArchiveNamespaceLister helps list and get Archives.
// All objects returned here must be treated as read-only.
type ArchiveNamespaceLister interface {
	// List lists all Archives in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Archive, err error)
	// Get retrieves the Archive from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Archive, error)
	ArchiveNamespaceListerExpansion
}

// archiveNamespaceLister implements the ArchiveNamespaceLister
// interface.
type archiveNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Archives in the indexer for a given namespace.
func (s archiveNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Archive, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Archive))
	})
	return ret, err
}

// Get retrieves the Archive from the indexer for a given namespace and name.
func (s archiveNamespaceLister) Get(name string) (*v1alpha1.Archive, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("archive"), name)
	}
	return obj.(*v1alpha1.Archive), nil
}
