//
// DISCLAIMER
//
// Copyright 2023 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// Code generated by client-gen. DO NOT EDIT.

package v2alpha1

import (
	"context"
	"time"

	v2alpha1 "github.com/arangodb/kube-arangodb/pkg/apis/replication/v2alpha1"
	scheme "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ArangoDeploymentReplicationsGetter has a method to return a ArangoDeploymentReplicationInterface.
// A group's client should implement this interface.
type ArangoDeploymentReplicationsGetter interface {
	ArangoDeploymentReplications(namespace string) ArangoDeploymentReplicationInterface
}

// ArangoDeploymentReplicationInterface has methods to work with ArangoDeploymentReplication resources.
type ArangoDeploymentReplicationInterface interface {
	Create(ctx context.Context, arangoDeploymentReplication *v2alpha1.ArangoDeploymentReplication, opts v1.CreateOptions) (*v2alpha1.ArangoDeploymentReplication, error)
	Update(ctx context.Context, arangoDeploymentReplication *v2alpha1.ArangoDeploymentReplication, opts v1.UpdateOptions) (*v2alpha1.ArangoDeploymentReplication, error)
	UpdateStatus(ctx context.Context, arangoDeploymentReplication *v2alpha1.ArangoDeploymentReplication, opts v1.UpdateOptions) (*v2alpha1.ArangoDeploymentReplication, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2alpha1.ArangoDeploymentReplication, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2alpha1.ArangoDeploymentReplicationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.ArangoDeploymentReplication, err error)
	ArangoDeploymentReplicationExpansion
}

// arangoDeploymentReplications implements ArangoDeploymentReplicationInterface
type arangoDeploymentReplications struct {
	client rest.Interface
	ns     string
}

// newArangoDeploymentReplications returns a ArangoDeploymentReplications
func newArangoDeploymentReplications(c *ReplicationV2alpha1Client, namespace string) *arangoDeploymentReplications {
	return &arangoDeploymentReplications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the arangoDeploymentReplication, and returns the corresponding arangoDeploymentReplication object, and an error if there is any.
func (c *arangoDeploymentReplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.ArangoDeploymentReplication, err error) {
	result = &v2alpha1.ArangoDeploymentReplication{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("arangodeploymentreplications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ArangoDeploymentReplications that match those selectors.
func (c *arangoDeploymentReplications) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.ArangoDeploymentReplicationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2alpha1.ArangoDeploymentReplicationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("arangodeploymentreplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested arangoDeploymentReplications.
func (c *arangoDeploymentReplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("arangodeploymentreplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a arangoDeploymentReplication and creates it.  Returns the server's representation of the arangoDeploymentReplication, and an error, if there is any.
func (c *arangoDeploymentReplications) Create(ctx context.Context, arangoDeploymentReplication *v2alpha1.ArangoDeploymentReplication, opts v1.CreateOptions) (result *v2alpha1.ArangoDeploymentReplication, err error) {
	result = &v2alpha1.ArangoDeploymentReplication{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("arangodeploymentreplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(arangoDeploymentReplication).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a arangoDeploymentReplication and updates it. Returns the server's representation of the arangoDeploymentReplication, and an error, if there is any.
func (c *arangoDeploymentReplications) Update(ctx context.Context, arangoDeploymentReplication *v2alpha1.ArangoDeploymentReplication, opts v1.UpdateOptions) (result *v2alpha1.ArangoDeploymentReplication, err error) {
	result = &v2alpha1.ArangoDeploymentReplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("arangodeploymentreplications").
		Name(arangoDeploymentReplication.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(arangoDeploymentReplication).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *arangoDeploymentReplications) UpdateStatus(ctx context.Context, arangoDeploymentReplication *v2alpha1.ArangoDeploymentReplication, opts v1.UpdateOptions) (result *v2alpha1.ArangoDeploymentReplication, err error) {
	result = &v2alpha1.ArangoDeploymentReplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("arangodeploymentreplications").
		Name(arangoDeploymentReplication.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(arangoDeploymentReplication).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the arangoDeploymentReplication and deletes it. Returns an error if one occurs.
func (c *arangoDeploymentReplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("arangodeploymentreplications").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *arangoDeploymentReplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("arangodeploymentreplications").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched arangoDeploymentReplication.
func (c *arangoDeploymentReplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.ArangoDeploymentReplication, err error) {
	result = &v2alpha1.ArangoDeploymentReplication{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("arangodeploymentreplications").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
