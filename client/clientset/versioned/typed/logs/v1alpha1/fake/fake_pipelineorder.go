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

	v1alpha1 "kubeform.dev/provider-datadog-api/apis/logs/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePipelineOrders implements PipelineOrderInterface
type FakePipelineOrders struct {
	Fake *FakeLogsV1alpha1
	ns   string
}

var pipelineordersResource = schema.GroupVersionResource{Group: "logs.datadog.kubeform.com", Version: "v1alpha1", Resource: "pipelineorders"}

var pipelineordersKind = schema.GroupVersionKind{Group: "logs.datadog.kubeform.com", Version: "v1alpha1", Kind: "PipelineOrder"}

// Get takes name of the pipelineOrder, and returns the corresponding pipelineOrder object, and an error if there is any.
func (c *FakePipelineOrders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PipelineOrder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pipelineordersResource, c.ns, name), &v1alpha1.PipelineOrder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PipelineOrder), err
}

// List takes label and field selectors, and returns the list of PipelineOrders that match those selectors.
func (c *FakePipelineOrders) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PipelineOrderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pipelineordersResource, pipelineordersKind, c.ns, opts), &v1alpha1.PipelineOrderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PipelineOrderList{ListMeta: obj.(*v1alpha1.PipelineOrderList).ListMeta}
	for _, item := range obj.(*v1alpha1.PipelineOrderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pipelineOrders.
func (c *FakePipelineOrders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pipelineordersResource, c.ns, opts))

}

// Create takes the representation of a pipelineOrder and creates it.  Returns the server's representation of the pipelineOrder, and an error, if there is any.
func (c *FakePipelineOrders) Create(ctx context.Context, pipelineOrder *v1alpha1.PipelineOrder, opts v1.CreateOptions) (result *v1alpha1.PipelineOrder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pipelineordersResource, c.ns, pipelineOrder), &v1alpha1.PipelineOrder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PipelineOrder), err
}

// Update takes the representation of a pipelineOrder and updates it. Returns the server's representation of the pipelineOrder, and an error, if there is any.
func (c *FakePipelineOrders) Update(ctx context.Context, pipelineOrder *v1alpha1.PipelineOrder, opts v1.UpdateOptions) (result *v1alpha1.PipelineOrder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pipelineordersResource, c.ns, pipelineOrder), &v1alpha1.PipelineOrder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PipelineOrder), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePipelineOrders) UpdateStatus(ctx context.Context, pipelineOrder *v1alpha1.PipelineOrder, opts v1.UpdateOptions) (*v1alpha1.PipelineOrder, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pipelineordersResource, "status", c.ns, pipelineOrder), &v1alpha1.PipelineOrder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PipelineOrder), err
}

// Delete takes name of the pipelineOrder and deletes it. Returns an error if one occurs.
func (c *FakePipelineOrders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pipelineordersResource, c.ns, name), &v1alpha1.PipelineOrder{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePipelineOrders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pipelineordersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PipelineOrderList{})
	return err
}

// Patch applies the patch and returns the patched pipelineOrder.
func (c *FakePipelineOrders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PipelineOrder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pipelineordersResource, c.ns, name, pt, data, subresources...), &v1alpha1.PipelineOrder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PipelineOrder), err
}
