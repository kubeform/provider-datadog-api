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

// FakeSlackChannels implements SlackChannelInterface
type FakeSlackChannels struct {
	Fake *FakeIntegrationV1alpha1
	ns   string
}

var slackchannelsResource = schema.GroupVersionResource{Group: "integration.datadog.kubeform.com", Version: "v1alpha1", Resource: "slackchannels"}

var slackchannelsKind = schema.GroupVersionKind{Group: "integration.datadog.kubeform.com", Version: "v1alpha1", Kind: "SlackChannel"}

// Get takes name of the slackChannel, and returns the corresponding slackChannel object, and an error if there is any.
func (c *FakeSlackChannels) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SlackChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(slackchannelsResource, c.ns, name), &v1alpha1.SlackChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SlackChannel), err
}

// List takes label and field selectors, and returns the list of SlackChannels that match those selectors.
func (c *FakeSlackChannels) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SlackChannelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(slackchannelsResource, slackchannelsKind, c.ns, opts), &v1alpha1.SlackChannelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SlackChannelList{ListMeta: obj.(*v1alpha1.SlackChannelList).ListMeta}
	for _, item := range obj.(*v1alpha1.SlackChannelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested slackChannels.
func (c *FakeSlackChannels) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(slackchannelsResource, c.ns, opts))

}

// Create takes the representation of a slackChannel and creates it.  Returns the server's representation of the slackChannel, and an error, if there is any.
func (c *FakeSlackChannels) Create(ctx context.Context, slackChannel *v1alpha1.SlackChannel, opts v1.CreateOptions) (result *v1alpha1.SlackChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(slackchannelsResource, c.ns, slackChannel), &v1alpha1.SlackChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SlackChannel), err
}

// Update takes the representation of a slackChannel and updates it. Returns the server's representation of the slackChannel, and an error, if there is any.
func (c *FakeSlackChannels) Update(ctx context.Context, slackChannel *v1alpha1.SlackChannel, opts v1.UpdateOptions) (result *v1alpha1.SlackChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(slackchannelsResource, c.ns, slackChannel), &v1alpha1.SlackChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SlackChannel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSlackChannels) UpdateStatus(ctx context.Context, slackChannel *v1alpha1.SlackChannel, opts v1.UpdateOptions) (*v1alpha1.SlackChannel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(slackchannelsResource, "status", c.ns, slackChannel), &v1alpha1.SlackChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SlackChannel), err
}

// Delete takes name of the slackChannel and deletes it. Returns an error if one occurs.
func (c *FakeSlackChannels) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(slackchannelsResource, c.ns, name), &v1alpha1.SlackChannel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSlackChannels) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(slackchannelsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SlackChannelList{})
	return err
}

// Patch applies the patch and returns the patched slackChannel.
func (c *FakeSlackChannels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SlackChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(slackchannelsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SlackChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SlackChannel), err
}