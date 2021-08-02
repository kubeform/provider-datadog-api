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

	v1alpha1 "kubeform.dev/provider-datadog-api/apis/integration/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGcps implements GcpInterface
type FakeGcps struct {
	Fake *FakeIntegrationV1alpha1
	ns   string
}

var gcpsResource = schema.GroupVersionResource{Group: "integration.datadog.kubeform.com", Version: "v1alpha1", Resource: "gcps"}

var gcpsKind = schema.GroupVersionKind{Group: "integration.datadog.kubeform.com", Version: "v1alpha1", Kind: "Gcp"}

// Get takes name of the gcp, and returns the corresponding gcp object, and an error if there is any.
func (c *FakeGcps) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Gcp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gcpsResource, c.ns, name), &v1alpha1.Gcp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Gcp), err
}

// List takes label and field selectors, and returns the list of Gcps that match those selectors.
func (c *FakeGcps) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GcpList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gcpsResource, gcpsKind, c.ns, opts), &v1alpha1.GcpList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GcpList{ListMeta: obj.(*v1alpha1.GcpList).ListMeta}
	for _, item := range obj.(*v1alpha1.GcpList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gcps.
func (c *FakeGcps) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gcpsResource, c.ns, opts))

}

// Create takes the representation of a gcp and creates it.  Returns the server's representation of the gcp, and an error, if there is any.
func (c *FakeGcps) Create(ctx context.Context, gcp *v1alpha1.Gcp, opts v1.CreateOptions) (result *v1alpha1.Gcp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gcpsResource, c.ns, gcp), &v1alpha1.Gcp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Gcp), err
}

// Update takes the representation of a gcp and updates it. Returns the server's representation of the gcp, and an error, if there is any.
func (c *FakeGcps) Update(ctx context.Context, gcp *v1alpha1.Gcp, opts v1.UpdateOptions) (result *v1alpha1.Gcp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gcpsResource, c.ns, gcp), &v1alpha1.Gcp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Gcp), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGcps) UpdateStatus(ctx context.Context, gcp *v1alpha1.Gcp, opts v1.UpdateOptions) (*v1alpha1.Gcp, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(gcpsResource, "status", c.ns, gcp), &v1alpha1.Gcp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Gcp), err
}

// Delete takes name of the gcp and deletes it. Returns an error if one occurs.
func (c *FakeGcps) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(gcpsResource, c.ns, name), &v1alpha1.Gcp{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGcps) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gcpsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GcpList{})
	return err
}

// Patch applies the patch and returns the patched gcp.
func (c *FakeGcps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Gcp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gcpsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Gcp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Gcp), err
}