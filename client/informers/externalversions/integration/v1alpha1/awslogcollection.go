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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	integrationv1alpha1 "kubeform.dev/provider-datadog-api/apis/integration/v1alpha1"
	versioned "kubeform.dev/provider-datadog-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-datadog-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-datadog-api/client/listers/integration/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AwsLogCollectionInformer provides access to a shared informer and lister for
// AwsLogCollections.
type AwsLogCollectionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AwsLogCollectionLister
}

type awsLogCollectionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAwsLogCollectionInformer constructs a new informer for AwsLogCollection type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAwsLogCollectionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAwsLogCollectionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAwsLogCollectionInformer constructs a new informer for AwsLogCollection type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAwsLogCollectionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IntegrationV1alpha1().AwsLogCollections(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IntegrationV1alpha1().AwsLogCollections(namespace).Watch(context.TODO(), options)
			},
		},
		&integrationv1alpha1.AwsLogCollection{},
		resyncPeriod,
		indexers,
	)
}

func (f *awsLogCollectionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAwsLogCollectionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *awsLogCollectionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&integrationv1alpha1.AwsLogCollection{}, f.defaultInformer)
}

func (f *awsLogCollectionInformer) Lister() v1alpha1.AwsLogCollectionLister {
	return v1alpha1.NewAwsLogCollectionLister(f.Informer().GetIndexer())
}