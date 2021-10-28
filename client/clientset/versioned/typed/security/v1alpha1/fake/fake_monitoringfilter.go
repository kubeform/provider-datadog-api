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

	v1alpha1 "kubeform.dev/provider-datadog-api/apis/security/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMonitoringFilters implements MonitoringFilterInterface
type FakeMonitoringFilters struct {
	Fake *FakeSecurityV1alpha1
	ns   string
}

var monitoringfiltersResource = schema.GroupVersionResource{Group: "security.datadog.kubeform.com", Version: "v1alpha1", Resource: "monitoringfilters"}

var monitoringfiltersKind = schema.GroupVersionKind{Group: "security.datadog.kubeform.com", Version: "v1alpha1", Kind: "MonitoringFilter"}

// Get takes name of the monitoringFilter, and returns the corresponding monitoringFilter object, and an error if there is any.
func (c *FakeMonitoringFilters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MonitoringFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(monitoringfiltersResource, c.ns, name), &v1alpha1.MonitoringFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitoringFilter), err
}

// List takes label and field selectors, and returns the list of MonitoringFilters that match those selectors.
func (c *FakeMonitoringFilters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MonitoringFilterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(monitoringfiltersResource, monitoringfiltersKind, c.ns, opts), &v1alpha1.MonitoringFilterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MonitoringFilterList{ListMeta: obj.(*v1alpha1.MonitoringFilterList).ListMeta}
	for _, item := range obj.(*v1alpha1.MonitoringFilterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested monitoringFilters.
func (c *FakeMonitoringFilters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(monitoringfiltersResource, c.ns, opts))

}

// Create takes the representation of a monitoringFilter and creates it.  Returns the server's representation of the monitoringFilter, and an error, if there is any.
func (c *FakeMonitoringFilters) Create(ctx context.Context, monitoringFilter *v1alpha1.MonitoringFilter, opts v1.CreateOptions) (result *v1alpha1.MonitoringFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(monitoringfiltersResource, c.ns, monitoringFilter), &v1alpha1.MonitoringFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitoringFilter), err
}

// Update takes the representation of a monitoringFilter and updates it. Returns the server's representation of the monitoringFilter, and an error, if there is any.
func (c *FakeMonitoringFilters) Update(ctx context.Context, monitoringFilter *v1alpha1.MonitoringFilter, opts v1.UpdateOptions) (result *v1alpha1.MonitoringFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(monitoringfiltersResource, c.ns, monitoringFilter), &v1alpha1.MonitoringFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitoringFilter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMonitoringFilters) UpdateStatus(ctx context.Context, monitoringFilter *v1alpha1.MonitoringFilter, opts v1.UpdateOptions) (*v1alpha1.MonitoringFilter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(monitoringfiltersResource, "status", c.ns, monitoringFilter), &v1alpha1.MonitoringFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitoringFilter), err
}

// Delete takes name of the monitoringFilter and deletes it. Returns an error if one occurs.
func (c *FakeMonitoringFilters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(monitoringfiltersResource, c.ns, name), &v1alpha1.MonitoringFilter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMonitoringFilters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(monitoringfiltersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MonitoringFilterList{})
	return err
}

// Patch applies the patch and returns the patched monitoringFilter.
func (c *FakeMonitoringFilters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MonitoringFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(monitoringfiltersResource, c.ns, name, pt, data, subresources...), &v1alpha1.MonitoringFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitoringFilter), err
}