/*
Copyright The cert-manager Authors.

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
	"context"
	"time"

	v1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	scheme "github.com/cert-manager/cert-manager/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IssuersGetter has a method to return a IssuerInterface.
// A group's client should implement this interface.
type IssuersGetter interface {
	Issuers(namespace string) IssuerInterface
}

// IssuerInterface has methods to work with Issuer resources.
type IssuerInterface interface {
	Create(ctx context.Context, issuer *v1.Issuer, opts metav1.CreateOptions) (*v1.Issuer, error)
	Update(ctx context.Context, issuer *v1.Issuer, opts metav1.UpdateOptions) (*v1.Issuer, error)
	UpdateStatus(ctx context.Context, issuer *v1.Issuer, opts metav1.UpdateOptions) (*v1.Issuer, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Issuer, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.IssuerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Issuer, err error)
	IssuerExpansion
}

// issuers implements IssuerInterface
type issuers struct {
	client  rest.Interface
	cluster string
	ns      string
}

// newIssuers returns a Issuers
func newIssuers(c *CertmanagerV1Client, namespace string) *issuers {
	return &issuers{
		client:  c.RESTClient(),
		cluster: c.cluster,
		ns:      namespace,
	}
}

// Get takes name of the issuer, and returns the corresponding issuer object, and an error if there is any.
func (c *issuers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Issuer, err error) {
	result = &v1.Issuer{}
	err = c.client.Get().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("issuers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Issuers that match those selectors.
func (c *issuers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.IssuerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.IssuerList{}
	err = c.client.Get().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("issuers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested issuers.
func (c *issuers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("issuers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a issuer and creates it.  Returns the server's representation of the issuer, and an error, if there is any.
func (c *issuers) Create(ctx context.Context, issuer *v1.Issuer, opts metav1.CreateOptions) (result *v1.Issuer, err error) {
	result = &v1.Issuer{}
	err = c.client.Post().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("issuers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(issuer).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a issuer and updates it. Returns the server's representation of the issuer, and an error, if there is any.
func (c *issuers) Update(ctx context.Context, issuer *v1.Issuer, opts metav1.UpdateOptions) (result *v1.Issuer, err error) {
	result = &v1.Issuer{}
	err = c.client.Put().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("issuers").
		Name(issuer.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(issuer).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *issuers) UpdateStatus(ctx context.Context, issuer *v1.Issuer, opts metav1.UpdateOptions) (result *v1.Issuer, err error) {
	result = &v1.Issuer{}
	err = c.client.Put().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("issuers").
		Name(issuer.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(issuer).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the issuer and deletes it. Returns an error if one occurs.
func (c *issuers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("issuers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *issuers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("issuers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched issuer.
func (c *issuers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Issuer, err error) {
	result = &v1.Issuer{}
	err = c.client.Patch(pt).
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("issuers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
