/*
Copyright The Kubernetes Authors.

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

package v1

import (
	"time"

	v1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	scheme "github.com/rook/rook/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CephObjectStoreUsersGetter has a method to return a CephObjectStoreUserInterface.
// A group's client should implement this interface.
type CephObjectStoreUsersGetter interface {
	CephObjectStoreUsers(namespace string) CephObjectStoreUserInterface
}

// CephObjectStoreUserInterface has methods to work with CephObjectStoreUser resources.
type CephObjectStoreUserInterface interface {
	Create(*v1.CephObjectStoreUser) (*v1.CephObjectStoreUser, error)
	Update(*v1.CephObjectStoreUser) (*v1.CephObjectStoreUser, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.CephObjectStoreUser, error)
	List(opts metav1.ListOptions) (*v1.CephObjectStoreUserList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.CephObjectStoreUser, err error)
	CephObjectStoreUserExpansion
}

// cephObjectStoreUsers implements CephObjectStoreUserInterface
type cephObjectStoreUsers struct {
	client rest.Interface
	ns     string
}

// newCephObjectStoreUsers returns a CephObjectStoreUsers
func newCephObjectStoreUsers(c *CephV1Client, namespace string) *cephObjectStoreUsers {
	return &cephObjectStoreUsers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cephObjectStoreUser, and returns the corresponding cephObjectStoreUser object, and an error if there is any.
func (c *cephObjectStoreUsers) Get(name string, options metav1.GetOptions) (result *v1.CephObjectStoreUser, err error) {
	result = &v1.CephObjectStoreUser{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cephobjectstoreusers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CephObjectStoreUsers that match those selectors.
func (c *cephObjectStoreUsers) List(opts metav1.ListOptions) (result *v1.CephObjectStoreUserList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CephObjectStoreUserList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cephobjectstoreusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cephObjectStoreUsers.
func (c *cephObjectStoreUsers) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cephobjectstoreusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cephObjectStoreUser and creates it.  Returns the server's representation of the cephObjectStoreUser, and an error, if there is any.
func (c *cephObjectStoreUsers) Create(cephObjectStoreUser *v1.CephObjectStoreUser) (result *v1.CephObjectStoreUser, err error) {
	result = &v1.CephObjectStoreUser{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cephobjectstoreusers").
		Body(cephObjectStoreUser).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cephObjectStoreUser and updates it. Returns the server's representation of the cephObjectStoreUser, and an error, if there is any.
func (c *cephObjectStoreUsers) Update(cephObjectStoreUser *v1.CephObjectStoreUser) (result *v1.CephObjectStoreUser, err error) {
	result = &v1.CephObjectStoreUser{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cephobjectstoreusers").
		Name(cephObjectStoreUser.Name).
		Body(cephObjectStoreUser).
		Do().
		Into(result)
	return
}

// Delete takes name of the cephObjectStoreUser and deletes it. Returns an error if one occurs.
func (c *cephObjectStoreUsers) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cephobjectstoreusers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cephObjectStoreUsers) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cephobjectstoreusers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cephObjectStoreUser.
func (c *cephObjectStoreUsers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.CephObjectStoreUser, err error) {
	result = &v1.CephObjectStoreUser{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cephobjectstoreusers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}