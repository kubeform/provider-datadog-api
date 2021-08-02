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

package fake

import (
	"context"

	v1alpha1 "kubeform.dev/provider-datadog-api/apis/metric/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTagConfigurations implements TagConfigurationInterface
type FakeTagConfigurations struct {
	Fake *FakeMetricV1alpha1
	ns   string
}

var tagconfigurationsResource = schema.GroupVersionResource{Group: "metric.datadog.kubeform.com", Version: "v1alpha1", Resource: "tagconfigurations"}

var tagconfigurationsKind = schema.GroupVersionKind{Group: "metric.datadog.kubeform.com", Version: "v1alpha1", Kind: "TagConfiguration"}

// Get takes name of the tagConfiguration, and returns the corresponding tagConfiguration object, and an error if there is any.
func (c *FakeTagConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TagConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tagconfigurationsResource, c.ns, name), &v1alpha1.TagConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TagConfiguration), err
}

// List takes label and field selectors, and returns the list of TagConfigurations that match those selectors.
func (c *FakeTagConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TagConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tagconfigurationsResource, tagconfigurationsKind, c.ns, opts), &v1alpha1.TagConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TagConfigurationList{ListMeta: obj.(*v1alpha1.TagConfigurationList).ListMeta}
	for _, item := range obj.(*v1alpha1.TagConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tagConfigurations.
func (c *FakeTagConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tagconfigurationsResource, c.ns, opts))

}

// Create takes the representation of a tagConfiguration and creates it.  Returns the server's representation of the tagConfiguration, and an error, if there is any.
func (c *FakeTagConfigurations) Create(ctx context.Context, tagConfiguration *v1alpha1.TagConfiguration, opts v1.CreateOptions) (result *v1alpha1.TagConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tagconfigurationsResource, c.ns, tagConfiguration), &v1alpha1.TagConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TagConfiguration), err
}

// Update takes the representation of a tagConfiguration and updates it. Returns the server's representation of the tagConfiguration, and an error, if there is any.
func (c *FakeTagConfigurations) Update(ctx context.Context, tagConfiguration *v1alpha1.TagConfiguration, opts v1.UpdateOptions) (result *v1alpha1.TagConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tagconfigurationsResource, c.ns, tagConfiguration), &v1alpha1.TagConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TagConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTagConfigurations) UpdateStatus(ctx context.Context, tagConfiguration *v1alpha1.TagConfiguration, opts v1.UpdateOptions) (*v1alpha1.TagConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tagconfigurationsResource, "status", c.ns, tagConfiguration), &v1alpha1.TagConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TagConfiguration), err
}

// Delete takes name of the tagConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeTagConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tagconfigurationsResource, c.ns, name), &v1alpha1.TagConfiguration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTagConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tagconfigurationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TagConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched tagConfiguration.
func (c *FakeTagConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TagConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tagconfigurationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.TagConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TagConfiguration), err
}