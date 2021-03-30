/*
Copyright Sparebanken Vest

Based on the Kubernetes controller example at
https://github.com/kubernetes/sample-controller

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

	v2beta1 "github.com/SparebankenVest/azure-key-vault-to-kubernetes/pkg/k8s/apis/azurekeyvault/v2beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAzureKeyVaultSecrets implements AzureKeyVaultSecretInterface
type FakeAzureKeyVaultSecrets struct {
	Fake *FakeSpvV2beta1
	ns   string
}

var azurekeyvaultsecretsResource = schema.GroupVersionResource{Group: "spv.no", Version: "v2beta1", Resource: "azurekeyvaultsecrets"}

var azurekeyvaultsecretsKind = schema.GroupVersionKind{Group: "spv.no", Version: "v2beta1", Kind: "AzureKeyVaultSecret"}

// Get takes name of the azureKeyVaultSecret, and returns the corresponding azureKeyVaultSecret object, and an error if there is any.
func (c *FakeAzureKeyVaultSecrets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2beta1.AzureKeyVaultSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(azurekeyvaultsecretsResource, c.ns, name), &v2beta1.AzureKeyVaultSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta1.AzureKeyVaultSecret), err
}

// List takes label and field selectors, and returns the list of AzureKeyVaultSecrets that match those selectors.
func (c *FakeAzureKeyVaultSecrets) List(ctx context.Context, opts v1.ListOptions) (result *v2beta1.AzureKeyVaultSecretList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(azurekeyvaultsecretsResource, azurekeyvaultsecretsKind, c.ns, opts), &v2beta1.AzureKeyVaultSecretList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2beta1.AzureKeyVaultSecretList{ListMeta: obj.(*v2beta1.AzureKeyVaultSecretList).ListMeta}
	for _, item := range obj.(*v2beta1.AzureKeyVaultSecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azureKeyVaultSecrets.
func (c *FakeAzureKeyVaultSecrets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(azurekeyvaultsecretsResource, c.ns, opts))

}

// Create takes the representation of a azureKeyVaultSecret and creates it.  Returns the server's representation of the azureKeyVaultSecret, and an error, if there is any.
func (c *FakeAzureKeyVaultSecrets) Create(ctx context.Context, azureKeyVaultSecret *v2beta1.AzureKeyVaultSecret, opts v1.CreateOptions) (result *v2beta1.AzureKeyVaultSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(azurekeyvaultsecretsResource, c.ns, azureKeyVaultSecret), &v2beta1.AzureKeyVaultSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta1.AzureKeyVaultSecret), err
}

// Update takes the representation of a azureKeyVaultSecret and updates it. Returns the server's representation of the azureKeyVaultSecret, and an error, if there is any.
func (c *FakeAzureKeyVaultSecrets) Update(ctx context.Context, azureKeyVaultSecret *v2beta1.AzureKeyVaultSecret, opts v1.UpdateOptions) (result *v2beta1.AzureKeyVaultSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(azurekeyvaultsecretsResource, c.ns, azureKeyVaultSecret), &v2beta1.AzureKeyVaultSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta1.AzureKeyVaultSecret), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzureKeyVaultSecrets) UpdateStatus(ctx context.Context, azureKeyVaultSecret *v2beta1.AzureKeyVaultSecret, opts v1.UpdateOptions) (*v2beta1.AzureKeyVaultSecret, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(azurekeyvaultsecretsResource, "status", c.ns, azureKeyVaultSecret), &v2beta1.AzureKeyVaultSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta1.AzureKeyVaultSecret), err
}

// Delete takes name of the azureKeyVaultSecret and deletes it. Returns an error if one occurs.
func (c *FakeAzureKeyVaultSecrets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(azurekeyvaultsecretsResource, c.ns, name), &v2beta1.AzureKeyVaultSecret{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzureKeyVaultSecrets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(azurekeyvaultsecretsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v2beta1.AzureKeyVaultSecretList{})
	return err
}

// Patch applies the patch and returns the patched azureKeyVaultSecret.
func (c *FakeAzureKeyVaultSecrets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta1.AzureKeyVaultSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(azurekeyvaultsecretsResource, c.ns, name, pt, data, subresources...), &v2beta1.AzureKeyVaultSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta1.AzureKeyVaultSecret), err
}
