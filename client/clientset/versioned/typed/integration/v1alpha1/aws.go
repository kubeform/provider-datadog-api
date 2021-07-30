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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubeform.dev/provider-datadog-api/apis/integration/v1alpha1"
	scheme "kubeform.dev/provider-datadog-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AwsesGetter has a method to return a AwsInterface.
// A group's client should implement this interface.
type AwsesGetter interface {
	Awses(namespace string) AwsInterface
}

// AwsInterface has methods to work with Aws resources.
type AwsInterface interface {
	Create(ctx context.Context, aws *v1alpha1.Aws, opts v1.CreateOptions) (*v1alpha1.Aws, error)
	Update(ctx context.Context, aws *v1alpha1.Aws, opts v1.UpdateOptions) (*v1alpha1.Aws, error)
	UpdateStatus(ctx context.Context, aws *v1alpha1.Aws, opts v1.UpdateOptions) (*v1alpha1.Aws, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Aws, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AwsList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Aws, err error)
	AwsExpansion
}

// awses implements AwsInterface
type awses struct {
	client rest.Interface
	ns     string
}

// newAwses returns a Awses
func newAwses(c *IntegrationV1alpha1Client, namespace string) *awses {
	return &awses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the aws, and returns the corresponding aws object, and an error if there is any.
func (c *awses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Aws, err error) {
	result = &v1alpha1.Aws{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Awses that match those selectors.
func (c *awses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AwsList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awses.
func (c *awses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a aws and creates it.  Returns the server's representation of the aws, and an error, if there is any.
func (c *awses) Create(ctx context.Context, aws *v1alpha1.Aws, opts v1.CreateOptions) (result *v1alpha1.Aws, err error) {
	result = &v1alpha1.Aws{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aws).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a aws and updates it. Returns the server's representation of the aws, and an error, if there is any.
func (c *awses) Update(ctx context.Context, aws *v1alpha1.Aws, opts v1.UpdateOptions) (result *v1alpha1.Aws, err error) {
	result = &v1alpha1.Aws{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awses").
		Name(aws.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aws).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *awses) UpdateStatus(ctx context.Context, aws *v1alpha1.Aws, opts v1.UpdateOptions) (result *v1alpha1.Aws, err error) {
	result = &v1alpha1.Aws{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awses").
		Name(aws.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aws).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the aws and deletes it. Returns an error if one occurs.
func (c *awses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched aws.
func (c *awses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Aws, err error) {
	result = &v1alpha1.Aws{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
