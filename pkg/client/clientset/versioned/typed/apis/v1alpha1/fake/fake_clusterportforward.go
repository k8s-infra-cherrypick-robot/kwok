/*
Copyright 2023 The Kubernetes Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "sigs.k8s.io/kwok/pkg/apis/v1alpha1"
)

// FakeClusterPortForwards implements ClusterPortForwardInterface
type FakeClusterPortForwards struct {
	Fake *FakeKwokV1alpha1
}

var clusterportforwardsResource = v1alpha1.SchemeGroupVersion.WithResource("clusterportforwards")

var clusterportforwardsKind = v1alpha1.SchemeGroupVersion.WithKind("ClusterPortForward")

// Get takes name of the clusterPortForward, and returns the corresponding clusterPortForward object, and an error if there is any.
func (c *FakeClusterPortForwards) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterPortForward, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterportforwardsResource, name), &v1alpha1.ClusterPortForward{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPortForward), err
}

// List takes label and field selectors, and returns the list of ClusterPortForwards that match those selectors.
func (c *FakeClusterPortForwards) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterPortForwardList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterportforwardsResource, clusterportforwardsKind, opts), &v1alpha1.ClusterPortForwardList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterPortForwardList{ListMeta: obj.(*v1alpha1.ClusterPortForwardList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterPortForwardList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterPortForwards.
func (c *FakeClusterPortForwards) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterportforwardsResource, opts))
}

// Create takes the representation of a clusterPortForward and creates it.  Returns the server's representation of the clusterPortForward, and an error, if there is any.
func (c *FakeClusterPortForwards) Create(ctx context.Context, clusterPortForward *v1alpha1.ClusterPortForward, opts v1.CreateOptions) (result *v1alpha1.ClusterPortForward, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterportforwardsResource, clusterPortForward), &v1alpha1.ClusterPortForward{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPortForward), err
}

// Update takes the representation of a clusterPortForward and updates it. Returns the server's representation of the clusterPortForward, and an error, if there is any.
func (c *FakeClusterPortForwards) Update(ctx context.Context, clusterPortForward *v1alpha1.ClusterPortForward, opts v1.UpdateOptions) (result *v1alpha1.ClusterPortForward, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterportforwardsResource, clusterPortForward), &v1alpha1.ClusterPortForward{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPortForward), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterPortForwards) UpdateStatus(ctx context.Context, clusterPortForward *v1alpha1.ClusterPortForward, opts v1.UpdateOptions) (*v1alpha1.ClusterPortForward, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterportforwardsResource, "status", clusterPortForward), &v1alpha1.ClusterPortForward{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPortForward), err
}

// Delete takes name of the clusterPortForward and deletes it. Returns an error if one occurs.
func (c *FakeClusterPortForwards) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterportforwardsResource, name, opts), &v1alpha1.ClusterPortForward{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterPortForwards) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterportforwardsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterPortForwardList{})
	return err
}

// Patch applies the patch and returns the patched clusterPortForward.
func (c *FakeClusterPortForwards) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterPortForward, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterportforwardsResource, name, pt, data, subresources...), &v1alpha1.ClusterPortForward{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPortForward), err
}
