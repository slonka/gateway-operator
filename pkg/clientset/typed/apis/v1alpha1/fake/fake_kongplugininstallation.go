/*
Copyright 2022 Kong Inc.

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

	v1alpha1 "github.com/kong/gateway-operator/api/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKongPluginInstallations implements KongPluginInstallationInterface
type FakeKongPluginInstallations struct {
	Fake *FakeApisV1alpha1
	ns   string
}

var kongplugininstallationsResource = v1alpha1.SchemeGroupVersion.WithResource("kongplugininstallations")

var kongplugininstallationsKind = v1alpha1.SchemeGroupVersion.WithKind("KongPluginInstallation")

// Get takes name of the kongPluginInstallation, and returns the corresponding kongPluginInstallation object, and an error if there is any.
func (c *FakeKongPluginInstallations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KongPluginInstallation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kongplugininstallationsResource, c.ns, name), &v1alpha1.KongPluginInstallation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KongPluginInstallation), err
}

// List takes label and field selectors, and returns the list of KongPluginInstallations that match those selectors.
func (c *FakeKongPluginInstallations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KongPluginInstallationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kongplugininstallationsResource, kongplugininstallationsKind, c.ns, opts), &v1alpha1.KongPluginInstallationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KongPluginInstallationList{ListMeta: obj.(*v1alpha1.KongPluginInstallationList).ListMeta}
	for _, item := range obj.(*v1alpha1.KongPluginInstallationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kongPluginInstallations.
func (c *FakeKongPluginInstallations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kongplugininstallationsResource, c.ns, opts))

}

// Create takes the representation of a kongPluginInstallation and creates it.  Returns the server's representation of the kongPluginInstallation, and an error, if there is any.
func (c *FakeKongPluginInstallations) Create(ctx context.Context, kongPluginInstallation *v1alpha1.KongPluginInstallation, opts v1.CreateOptions) (result *v1alpha1.KongPluginInstallation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kongplugininstallationsResource, c.ns, kongPluginInstallation), &v1alpha1.KongPluginInstallation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KongPluginInstallation), err
}

// Update takes the representation of a kongPluginInstallation and updates it. Returns the server's representation of the kongPluginInstallation, and an error, if there is any.
func (c *FakeKongPluginInstallations) Update(ctx context.Context, kongPluginInstallation *v1alpha1.KongPluginInstallation, opts v1.UpdateOptions) (result *v1alpha1.KongPluginInstallation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kongplugininstallationsResource, c.ns, kongPluginInstallation), &v1alpha1.KongPluginInstallation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KongPluginInstallation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKongPluginInstallations) UpdateStatus(ctx context.Context, kongPluginInstallation *v1alpha1.KongPluginInstallation, opts v1.UpdateOptions) (*v1alpha1.KongPluginInstallation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kongplugininstallationsResource, "status", c.ns, kongPluginInstallation), &v1alpha1.KongPluginInstallation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KongPluginInstallation), err
}

// Delete takes name of the kongPluginInstallation and deletes it. Returns an error if one occurs.
func (c *FakeKongPluginInstallations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(kongplugininstallationsResource, c.ns, name, opts), &v1alpha1.KongPluginInstallation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKongPluginInstallations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kongplugininstallationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KongPluginInstallationList{})
	return err
}

// Patch applies the patch and returns the patched kongPluginInstallation.
func (c *FakeKongPluginInstallations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KongPluginInstallation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kongplugininstallationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.KongPluginInstallation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KongPluginInstallation), err
}
