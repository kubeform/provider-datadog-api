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

	v1alpha1 "kubeform.dev/provider-datadog-api/apis/service/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLevelObjectives implements LevelObjectiveInterface
type FakeLevelObjectives struct {
	Fake *FakeServiceV1alpha1
	ns   string
}

var levelobjectivesResource = schema.GroupVersionResource{Group: "service.datadog.kubeform.com", Version: "v1alpha1", Resource: "levelobjectives"}

var levelobjectivesKind = schema.GroupVersionKind{Group: "service.datadog.kubeform.com", Version: "v1alpha1", Kind: "LevelObjective"}

// Get takes name of the levelObjective, and returns the corresponding levelObjective object, and an error if there is any.
func (c *FakeLevelObjectives) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LevelObjective, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(levelobjectivesResource, c.ns, name), &v1alpha1.LevelObjective{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LevelObjective), err
}

// List takes label and field selectors, and returns the list of LevelObjectives that match those selectors.
func (c *FakeLevelObjectives) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LevelObjectiveList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(levelobjectivesResource, levelobjectivesKind, c.ns, opts), &v1alpha1.LevelObjectiveList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LevelObjectiveList{ListMeta: obj.(*v1alpha1.LevelObjectiveList).ListMeta}
	for _, item := range obj.(*v1alpha1.LevelObjectiveList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested levelObjectives.
func (c *FakeLevelObjectives) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(levelobjectivesResource, c.ns, opts))

}

// Create takes the representation of a levelObjective and creates it.  Returns the server's representation of the levelObjective, and an error, if there is any.
func (c *FakeLevelObjectives) Create(ctx context.Context, levelObjective *v1alpha1.LevelObjective, opts v1.CreateOptions) (result *v1alpha1.LevelObjective, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(levelobjectivesResource, c.ns, levelObjective), &v1alpha1.LevelObjective{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LevelObjective), err
}

// Update takes the representation of a levelObjective and updates it. Returns the server's representation of the levelObjective, and an error, if there is any.
func (c *FakeLevelObjectives) Update(ctx context.Context, levelObjective *v1alpha1.LevelObjective, opts v1.UpdateOptions) (result *v1alpha1.LevelObjective, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(levelobjectivesResource, c.ns, levelObjective), &v1alpha1.LevelObjective{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LevelObjective), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLevelObjectives) UpdateStatus(ctx context.Context, levelObjective *v1alpha1.LevelObjective, opts v1.UpdateOptions) (*v1alpha1.LevelObjective, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(levelobjectivesResource, "status", c.ns, levelObjective), &v1alpha1.LevelObjective{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LevelObjective), err
}

// Delete takes name of the levelObjective and deletes it. Returns an error if one occurs.
func (c *FakeLevelObjectives) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(levelobjectivesResource, c.ns, name), &v1alpha1.LevelObjective{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLevelObjectives) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(levelobjectivesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LevelObjectiveList{})
	return err
}

// Patch applies the patch and returns the patched levelObjective.
func (c *FakeLevelObjectives) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LevelObjective, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(levelobjectivesResource, c.ns, name, pt, data, subresources...), &v1alpha1.LevelObjective{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LevelObjective), err
}
