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

	v1alpha1 "kubeform.dev/provider-datadog-api/apis/apikey/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApiKeys implements ApiKeyInterface
type FakeApiKeys struct {
	Fake *FakeApikeyV1alpha1
	ns   string
}

var apikeysResource = schema.GroupVersionResource{Group: "apikey.datadog.kubeform.com", Version: "v1alpha1", Resource: "apikeys"}

var apikeysKind = schema.GroupVersionKind{Group: "apikey.datadog.kubeform.com", Version: "v1alpha1", Kind: "ApiKey"}

// Get takes name of the apiKey, and returns the corresponding apiKey object, and an error if there is any.
func (c *FakeApiKeys) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApiKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apikeysResource, c.ns, name), &v1alpha1.ApiKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiKey), err
}

// List takes label and field selectors, and returns the list of ApiKeys that match those selectors.
func (c *FakeApiKeys) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApiKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apikeysResource, apikeysKind, c.ns, opts), &v1alpha1.ApiKeyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiKeyList{ListMeta: obj.(*v1alpha1.ApiKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiKeys.
func (c *FakeApiKeys) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apikeysResource, c.ns, opts))

}

// Create takes the representation of a apiKey and creates it.  Returns the server's representation of the apiKey, and an error, if there is any.
func (c *FakeApiKeys) Create(ctx context.Context, apiKey *v1alpha1.ApiKey, opts v1.CreateOptions) (result *v1alpha1.ApiKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apikeysResource, c.ns, apiKey), &v1alpha1.ApiKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiKey), err
}

// Update takes the representation of a apiKey and updates it. Returns the server's representation of the apiKey, and an error, if there is any.
func (c *FakeApiKeys) Update(ctx context.Context, apiKey *v1alpha1.ApiKey, opts v1.UpdateOptions) (result *v1alpha1.ApiKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apikeysResource, c.ns, apiKey), &v1alpha1.ApiKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiKeys) UpdateStatus(ctx context.Context, apiKey *v1alpha1.ApiKey, opts v1.UpdateOptions) (*v1alpha1.ApiKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apikeysResource, "status", c.ns, apiKey), &v1alpha1.ApiKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiKey), err
}

// Delete takes name of the apiKey and deletes it. Returns an error if one occurs.
func (c *FakeApiKeys) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apikeysResource, c.ns, name), &v1alpha1.ApiKey{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiKeys) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apikeysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiKeyList{})
	return err
}

// Patch applies the patch and returns the patched apiKey.
func (c *FakeApiKeys) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apikeysResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiKey), err
}
