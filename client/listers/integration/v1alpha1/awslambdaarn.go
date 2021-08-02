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

// AwsLambdaArnLister helps list AwsLambdaArns.
// All objects returned here must be treated as read-only.
type AwsLambdaArnLister interface {
	// List lists all AwsLambdaArns in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AwsLambdaArn, err error)
	// AwsLambdaArns returns an object that can list and get AwsLambdaArns.
	AwsLambdaArns(namespace string) AwsLambdaArnNamespaceLister
	AwsLambdaArnListerExpansion
}

// awsLambdaArnLister implements the AwsLambdaArnLister interface.
type awsLambdaArnLister struct {
	indexer cache.Indexer
}

// NewAwsLambdaArnLister returns a new AwsLambdaArnLister.
func NewAwsLambdaArnLister(indexer cache.Indexer) AwsLambdaArnLister {
	return &awsLambdaArnLister{indexer: indexer}
}

// List lists all AwsLambdaArns in the indexer.
func (s *awsLambdaArnLister) List(selector labels.Selector) (ret []*v1alpha1.AwsLambdaArn, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsLambdaArn))
	})
	return ret, err
}

// AwsLambdaArns returns an object that can list and get AwsLambdaArns.
func (s *awsLambdaArnLister) AwsLambdaArns(namespace string) AwsLambdaArnNamespaceLister {
	return awsLambdaArnNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsLambdaArnNamespaceLister helps list and get AwsLambdaArns.
// All objects returned here must be treated as read-only.
type AwsLambdaArnNamespaceLister interface {
	// List lists all AwsLambdaArns in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AwsLambdaArn, err error)
	// Get retrieves the AwsLambdaArn from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AwsLambdaArn, error)
	AwsLambdaArnNamespaceListerExpansion
}

// awsLambdaArnNamespaceLister implements the AwsLambdaArnNamespaceLister
// interface.
type awsLambdaArnNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsLambdaArns in the indexer for a given namespace.
func (s awsLambdaArnNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsLambdaArn, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsLambdaArn))
	})
	return ret, err
}

// Get retrieves the AwsLambdaArn from the indexer for a given namespace and name.
func (s awsLambdaArnNamespaceLister) Get(name string) (*v1alpha1.AwsLambdaArn, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awslambdaarn"), name)
	}
	return obj.(*v1alpha1.AwsLambdaArn), nil
}