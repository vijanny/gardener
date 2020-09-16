/*
SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors

SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeShoots implements ShootInterface
type FakeShoots struct {
	Fake *FakeCoreV1alpha1
	ns   string
}

var shootsResource = schema.GroupVersionResource{Group: "core.gardener.cloud", Version: "v1alpha1", Resource: "shoots"}

var shootsKind = schema.GroupVersionKind{Group: "core.gardener.cloud", Version: "v1alpha1", Kind: "Shoot"}

// Get takes name of the shoot, and returns the corresponding shoot object, and an error if there is any.
func (c *FakeShoots) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Shoot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(shootsResource, c.ns, name), &v1alpha1.Shoot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Shoot), err
}

// List takes label and field selectors, and returns the list of Shoots that match those selectors.
func (c *FakeShoots) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ShootList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(shootsResource, shootsKind, c.ns, opts), &v1alpha1.ShootList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ShootList{ListMeta: obj.(*v1alpha1.ShootList).ListMeta}
	for _, item := range obj.(*v1alpha1.ShootList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested shoots.
func (c *FakeShoots) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(shootsResource, c.ns, opts))

}

// Create takes the representation of a shoot and creates it.  Returns the server's representation of the shoot, and an error, if there is any.
func (c *FakeShoots) Create(ctx context.Context, shoot *v1alpha1.Shoot, opts v1.CreateOptions) (result *v1alpha1.Shoot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(shootsResource, c.ns, shoot), &v1alpha1.Shoot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Shoot), err
}

// Update takes the representation of a shoot and updates it. Returns the server's representation of the shoot, and an error, if there is any.
func (c *FakeShoots) Update(ctx context.Context, shoot *v1alpha1.Shoot, opts v1.UpdateOptions) (result *v1alpha1.Shoot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(shootsResource, c.ns, shoot), &v1alpha1.Shoot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Shoot), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeShoots) UpdateStatus(ctx context.Context, shoot *v1alpha1.Shoot, opts v1.UpdateOptions) (*v1alpha1.Shoot, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(shootsResource, "status", c.ns, shoot), &v1alpha1.Shoot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Shoot), err
}

// Delete takes name of the shoot and deletes it. Returns an error if one occurs.
func (c *FakeShoots) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(shootsResource, c.ns, name), &v1alpha1.Shoot{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeShoots) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(shootsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ShootList{})
	return err
}

// Patch applies the patch and returns the patched shoot.
func (c *FakeShoots) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Shoot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(shootsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Shoot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Shoot), err
}
