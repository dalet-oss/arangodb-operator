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

package fake

import (
	"context"

	replicationv1 "github.com/dalet-oss/arangodb-operator/pkg/apis/replication/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeArangoDeploymentReplications implements ArangoDeploymentReplicationInterface
type FakeArangoDeploymentReplications struct {
	Fake *FakeReplicationV1
	ns   string
}

var arangodeploymentreplicationsResource = schema.GroupVersionResource{Group: "replication.database.arangodb.com", Version: "v1", Resource: "arangodeploymentreplications"}

var arangodeploymentreplicationsKind = schema.GroupVersionKind{Group: "replication.database.arangodb.com", Version: "v1", Kind: "ArangoDeploymentReplication"}

// Get takes name of the arangoDeploymentReplication, and returns the corresponding arangoDeploymentReplication object, and an error if there is any.
func (c *FakeArangoDeploymentReplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *replicationv1.ArangoDeploymentReplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(arangodeploymentreplicationsResource, c.ns, name), &replicationv1.ArangoDeploymentReplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*replicationv1.ArangoDeploymentReplication), err
}

// List takes label and field selectors, and returns the list of ArangoDeploymentReplications that match those selectors.
func (c *FakeArangoDeploymentReplications) List(ctx context.Context, opts v1.ListOptions) (result *replicationv1.ArangoDeploymentReplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(arangodeploymentreplicationsResource, arangodeploymentreplicationsKind, c.ns, opts), &replicationv1.ArangoDeploymentReplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &replicationv1.ArangoDeploymentReplicationList{ListMeta: obj.(*replicationv1.ArangoDeploymentReplicationList).ListMeta}
	for _, item := range obj.(*replicationv1.ArangoDeploymentReplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested arangoDeploymentReplications.
func (c *FakeArangoDeploymentReplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(arangodeploymentreplicationsResource, c.ns, opts))

}

// Create takes the representation of a arangoDeploymentReplication and creates it.  Returns the server's representation of the arangoDeploymentReplication, and an error, if there is any.
func (c *FakeArangoDeploymentReplications) Create(ctx context.Context, arangoDeploymentReplication *replicationv1.ArangoDeploymentReplication, opts v1.CreateOptions) (result *replicationv1.ArangoDeploymentReplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(arangodeploymentreplicationsResource, c.ns, arangoDeploymentReplication), &replicationv1.ArangoDeploymentReplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*replicationv1.ArangoDeploymentReplication), err
}

// Update takes the representation of a arangoDeploymentReplication and updates it. Returns the server's representation of the arangoDeploymentReplication, and an error, if there is any.
func (c *FakeArangoDeploymentReplications) Update(ctx context.Context, arangoDeploymentReplication *replicationv1.ArangoDeploymentReplication, opts v1.UpdateOptions) (result *replicationv1.ArangoDeploymentReplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(arangodeploymentreplicationsResource, c.ns, arangoDeploymentReplication), &replicationv1.ArangoDeploymentReplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*replicationv1.ArangoDeploymentReplication), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeArangoDeploymentReplications) UpdateStatus(ctx context.Context, arangoDeploymentReplication *replicationv1.ArangoDeploymentReplication, opts v1.UpdateOptions) (*replicationv1.ArangoDeploymentReplication, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(arangodeploymentreplicationsResource, "status", c.ns, arangoDeploymentReplication), &replicationv1.ArangoDeploymentReplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*replicationv1.ArangoDeploymentReplication), err
}

// Delete takes name of the arangoDeploymentReplication and deletes it. Returns an error if one occurs.
func (c *FakeArangoDeploymentReplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(arangodeploymentreplicationsResource, c.ns, name), &replicationv1.ArangoDeploymentReplication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeArangoDeploymentReplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(arangodeploymentreplicationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &replicationv1.ArangoDeploymentReplicationList{})
	return err
}

// Patch applies the patch and returns the patched arangoDeploymentReplication.
func (c *FakeArangoDeploymentReplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *replicationv1.ArangoDeploymentReplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(arangodeploymentreplicationsResource, c.ns, name, pt, data, subresources...), &replicationv1.ArangoDeploymentReplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*replicationv1.ArangoDeploymentReplication), err
}
