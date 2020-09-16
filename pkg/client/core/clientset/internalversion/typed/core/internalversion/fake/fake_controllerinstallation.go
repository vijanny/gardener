/*
SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors

SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	core "github.com/gardener/gardener/pkg/apis/core"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeControllerInstallations implements ControllerInstallationInterface
type FakeControllerInstallations struct {
	Fake *FakeCore
}

var controllerinstallationsResource = schema.GroupVersionResource{Group: "core.gardener.cloud", Version: "", Resource: "controllerinstallations"}

var controllerinstallationsKind = schema.GroupVersionKind{Group: "core.gardener.cloud", Version: "", Kind: "ControllerInstallation"}

// Get takes name of the controllerInstallation, and returns the corresponding controllerInstallation object, and an error if there is any.
func (c *FakeControllerInstallations) Get(ctx context.Context, name string, options v1.GetOptions) (result *core.ControllerInstallation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(controllerinstallationsResource, name), &core.ControllerInstallation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core.ControllerInstallation), err
}

// List takes label and field selectors, and returns the list of ControllerInstallations that match those selectors.
func (c *FakeControllerInstallations) List(ctx context.Context, opts v1.ListOptions) (result *core.ControllerInstallationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(controllerinstallationsResource, controllerinstallationsKind, opts), &core.ControllerInstallationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &core.ControllerInstallationList{ListMeta: obj.(*core.ControllerInstallationList).ListMeta}
	for _, item := range obj.(*core.ControllerInstallationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested controllerInstallations.
func (c *FakeControllerInstallations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(controllerinstallationsResource, opts))
}

// Create takes the representation of a controllerInstallation and creates it.  Returns the server's representation of the controllerInstallation, and an error, if there is any.
func (c *FakeControllerInstallations) Create(ctx context.Context, controllerInstallation *core.ControllerInstallation, opts v1.CreateOptions) (result *core.ControllerInstallation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(controllerinstallationsResource, controllerInstallation), &core.ControllerInstallation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core.ControllerInstallation), err
}

// Update takes the representation of a controllerInstallation and updates it. Returns the server's representation of the controllerInstallation, and an error, if there is any.
func (c *FakeControllerInstallations) Update(ctx context.Context, controllerInstallation *core.ControllerInstallation, opts v1.UpdateOptions) (result *core.ControllerInstallation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(controllerinstallationsResource, controllerInstallation), &core.ControllerInstallation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core.ControllerInstallation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeControllerInstallations) UpdateStatus(ctx context.Context, controllerInstallation *core.ControllerInstallation, opts v1.UpdateOptions) (*core.ControllerInstallation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(controllerinstallationsResource, "status", controllerInstallation), &core.ControllerInstallation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core.ControllerInstallation), err
}

// Delete takes name of the controllerInstallation and deletes it. Returns an error if one occurs.
func (c *FakeControllerInstallations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(controllerinstallationsResource, name), &core.ControllerInstallation{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeControllerInstallations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(controllerinstallationsResource, listOpts)

	_, err := c.Fake.Invokes(action, &core.ControllerInstallationList{})
	return err
}

// Patch applies the patch and returns the patched controllerInstallation.
func (c *FakeControllerInstallations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *core.ControllerInstallation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(controllerinstallationsResource, name, pt, data, subresources...), &core.ControllerInstallation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core.ControllerInstallation), err
}
