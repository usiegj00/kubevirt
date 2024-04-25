/*
Copyright The KubeVirt Authors.

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
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	corev1 "kubevirt.io/api/core/v1"
)

// FakeVirtualMachineInstances implements VirtualMachineInstanceInterface
type FakeVirtualMachineInstances struct {
	Fake *FakeKubevirtV1
	ns   string
}

var virtualmachineinstancesResource = schema.GroupVersionResource{Group: "kubevirt.io", Version: "v1", Resource: "virtualmachineinstances"}

var virtualmachineinstancesKind = schema.GroupVersionKind{Group: "kubevirt.io", Version: "v1", Kind: "VirtualMachineInstance"}

// Get takes name of the virtualMachineInstance, and returns the corresponding virtualMachineInstance object, and an error if there is any.
func (c *FakeVirtualMachineInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *corev1.VirtualMachineInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachineinstancesResource, c.ns, name), &corev1.VirtualMachineInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.VirtualMachineInstance), err
}

// List takes label and field selectors, and returns the list of VirtualMachineInstances that match those selectors.
func (c *FakeVirtualMachineInstances) List(ctx context.Context, opts v1.ListOptions) (result *corev1.VirtualMachineInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachineinstancesResource, virtualmachineinstancesKind, c.ns, opts), &corev1.VirtualMachineInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.VirtualMachineInstanceList{ListMeta: obj.(*corev1.VirtualMachineInstanceList).ListMeta}
	for _, item := range obj.(*corev1.VirtualMachineInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMachineInstances.
func (c *FakeVirtualMachineInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachineinstancesResource, c.ns, opts))

}

// Create takes the representation of a virtualMachineInstance and creates it.  Returns the server's representation of the virtualMachineInstance, and an error, if there is any.
func (c *FakeVirtualMachineInstances) Create(ctx context.Context, virtualMachineInstance *corev1.VirtualMachineInstance, opts v1.CreateOptions) (result *corev1.VirtualMachineInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachineinstancesResource, c.ns, virtualMachineInstance), &corev1.VirtualMachineInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.VirtualMachineInstance), err
}

// Update takes the representation of a virtualMachineInstance and updates it. Returns the server's representation of the virtualMachineInstance, and an error, if there is any.
func (c *FakeVirtualMachineInstances) Update(ctx context.Context, virtualMachineInstance *corev1.VirtualMachineInstance, opts v1.UpdateOptions) (result *corev1.VirtualMachineInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachineinstancesResource, c.ns, virtualMachineInstance), &corev1.VirtualMachineInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.VirtualMachineInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualMachineInstances) UpdateStatus(ctx context.Context, virtualMachineInstance *corev1.VirtualMachineInstance, opts v1.UpdateOptions) (*corev1.VirtualMachineInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmachineinstancesResource, "status", c.ns, virtualMachineInstance), &corev1.VirtualMachineInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.VirtualMachineInstance), err
}

// Delete takes name of the virtualMachineInstance and deletes it. Returns an error if one occurs.
func (c *FakeVirtualMachineInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmachineinstancesResource, c.ns, name), &corev1.VirtualMachineInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualMachineInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualmachineinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.VirtualMachineInstanceList{})
	return err
}

// Patch applies the patch and returns the patched virtualMachineInstance.
func (c *FakeVirtualMachineInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1.VirtualMachineInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachineinstancesResource, c.ns, name, pt, data, subresources...), &corev1.VirtualMachineInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.VirtualMachineInstance), err
}
