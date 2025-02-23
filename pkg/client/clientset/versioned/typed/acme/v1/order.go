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

	v1 "github.com/cert-manager/cert-manager/pkg/apis/acme/v1"
	scheme "github.com/cert-manager/cert-manager/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OrdersGetter has a method to return a OrderInterface.
// A group's client should implement this interface.
type OrdersGetter interface {
	Orders(namespace string) OrderInterface
}

// OrderInterface has methods to work with Order resources.
type OrderInterface interface {
	Create(ctx context.Context, order *v1.Order, opts metav1.CreateOptions) (*v1.Order, error)
	Update(ctx context.Context, order *v1.Order, opts metav1.UpdateOptions) (*v1.Order, error)
	UpdateStatus(ctx context.Context, order *v1.Order, opts metav1.UpdateOptions) (*v1.Order, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Order, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.OrderList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Order, err error)
	OrderExpansion
}

// orders implements OrderInterface
type orders struct {
	client  rest.Interface
	cluster string
	ns      string
}

// newOrders returns a Orders
func newOrders(c *AcmeV1Client, namespace string) *orders {
	return &orders{
		client:  c.RESTClient(),
		cluster: c.cluster,
		ns:      namespace,
	}
}

// Get takes name of the order, and returns the corresponding order object, and an error if there is any.
func (c *orders) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Order, err error) {
	result = &v1.Order{}
	err = c.client.Get().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("orders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Orders that match those selectors.
func (c *orders) List(ctx context.Context, opts metav1.ListOptions) (result *v1.OrderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.OrderList{}
	err = c.client.Get().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("orders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested orders.
func (c *orders) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("orders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a order and creates it.  Returns the server's representation of the order, and an error, if there is any.
func (c *orders) Create(ctx context.Context, order *v1.Order, opts metav1.CreateOptions) (result *v1.Order, err error) {
	result = &v1.Order{}
	err = c.client.Post().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("orders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(order).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a order and updates it. Returns the server's representation of the order, and an error, if there is any.
func (c *orders) Update(ctx context.Context, order *v1.Order, opts metav1.UpdateOptions) (result *v1.Order, err error) {
	result = &v1.Order{}
	err = c.client.Put().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("orders").
		Name(order.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(order).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *orders) UpdateStatus(ctx context.Context, order *v1.Order, opts metav1.UpdateOptions) (result *v1.Order, err error) {
	result = &v1.Order{}
	err = c.client.Put().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("orders").
		Name(order.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(order).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the order and deletes it. Returns an error if one occurs.
func (c *orders) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("orders").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *orders) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("orders").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched order.
func (c *orders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Order, err error) {
	result = &v1.Order{}
	err = c.client.Patch(pt).
		Cluster(c.cluster).
		Namespace(c.ns).
		Resource("orders").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
