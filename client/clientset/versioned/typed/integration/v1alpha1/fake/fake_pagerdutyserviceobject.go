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

// FakePagerdutyServiceObjects implements PagerdutyServiceObjectInterface
type FakePagerdutyServiceObjects struct {
	Fake *FakeIntegrationV1alpha1
	ns   string
}

var pagerdutyserviceobjectsResource = schema.GroupVersionResource{Group: "integration.datadog.kubeform.com", Version: "v1alpha1", Resource: "pagerdutyserviceobjects"}

var pagerdutyserviceobjectsKind = schema.GroupVersionKind{Group: "integration.datadog.kubeform.com", Version: "v1alpha1", Kind: "PagerdutyServiceObject"}

// Get takes name of the pagerdutyServiceObject, and returns the corresponding pagerdutyServiceObject object, and an error if there is any.
func (c *FakePagerdutyServiceObjects) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PagerdutyServiceObject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pagerdutyserviceobjectsResource, c.ns, name), &v1alpha1.PagerdutyServiceObject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PagerdutyServiceObject), err
}

// List takes label and field selectors, and returns the list of PagerdutyServiceObjects that match those selectors.
func (c *FakePagerdutyServiceObjects) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PagerdutyServiceObjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pagerdutyserviceobjectsResource, pagerdutyserviceobjectsKind, c.ns, opts), &v1alpha1.PagerdutyServiceObjectList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PagerdutyServiceObjectList{ListMeta: obj.(*v1alpha1.PagerdutyServiceObjectList).ListMeta}
	for _, item := range obj.(*v1alpha1.PagerdutyServiceObjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pagerdutyServiceObjects.
func (c *FakePagerdutyServiceObjects) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pagerdutyserviceobjectsResource, c.ns, opts))

}

// Create takes the representation of a pagerdutyServiceObject and creates it.  Returns the server's representation of the pagerdutyServiceObject, and an error, if there is any.
func (c *FakePagerdutyServiceObjects) Create(ctx context.Context, pagerdutyServiceObject *v1alpha1.PagerdutyServiceObject, opts v1.CreateOptions) (result *v1alpha1.PagerdutyServiceObject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pagerdutyserviceobjectsResource, c.ns, pagerdutyServiceObject), &v1alpha1.PagerdutyServiceObject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PagerdutyServiceObject), err
}

// Update takes the representation of a pagerdutyServiceObject and updates it. Returns the server's representation of the pagerdutyServiceObject, and an error, if there is any.
func (c *FakePagerdutyServiceObjects) Update(ctx context.Context, pagerdutyServiceObject *v1alpha1.PagerdutyServiceObject, opts v1.UpdateOptions) (result *v1alpha1.PagerdutyServiceObject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pagerdutyserviceobjectsResource, c.ns, pagerdutyServiceObject), &v1alpha1.PagerdutyServiceObject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PagerdutyServiceObject), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePagerdutyServiceObjects) UpdateStatus(ctx context.Context, pagerdutyServiceObject *v1alpha1.PagerdutyServiceObject, opts v1.UpdateOptions) (*v1alpha1.PagerdutyServiceObject, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pagerdutyserviceobjectsResource, "status", c.ns, pagerdutyServiceObject), &v1alpha1.PagerdutyServiceObject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PagerdutyServiceObject), err
}

// Delete takes name of the pagerdutyServiceObject and deletes it. Returns an error if one occurs.
func (c *FakePagerdutyServiceObjects) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pagerdutyserviceobjectsResource, c.ns, name), &v1alpha1.PagerdutyServiceObject{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePagerdutyServiceObjects) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pagerdutyserviceobjectsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PagerdutyServiceObjectList{})
	return err
}

// Patch applies the patch and returns the patched pagerdutyServiceObject.
func (c *FakePagerdutyServiceObjects) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PagerdutyServiceObject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pagerdutyserviceobjectsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PagerdutyServiceObject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PagerdutyServiceObject), err
}