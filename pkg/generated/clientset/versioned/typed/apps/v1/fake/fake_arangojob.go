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

	appsv1 "github.com/dalet-oss/arangodb-operator/pkg/apis/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeArangoJobs implements ArangoJobInterface
type FakeArangoJobs struct {
	Fake *FakeAppsV1
	ns   string
}

var arangojobsResource = schema.GroupVersionResource{Group: "apps.arangodb.com", Version: "v1", Resource: "arangojobs"}

var arangojobsKind = schema.GroupVersionKind{Group: "apps.arangodb.com", Version: "v1", Kind: "ArangoJob"}

// Get takes name of the arangoJob, and returns the corresponding arangoJob object, and an error if there is any.
func (c *FakeArangoJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *appsv1.ArangoJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(arangojobsResource, c.ns, name), &appsv1.ArangoJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.ArangoJob), err
}

// List takes label and field selectors, and returns the list of ArangoJobs that match those selectors.
func (c *FakeArangoJobs) List(ctx context.Context, opts v1.ListOptions) (result *appsv1.ArangoJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(arangojobsResource, arangojobsKind, c.ns, opts), &appsv1.ArangoJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &appsv1.ArangoJobList{ListMeta: obj.(*appsv1.ArangoJobList).ListMeta}
	for _, item := range obj.(*appsv1.ArangoJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested arangoJobs.
func (c *FakeArangoJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(arangojobsResource, c.ns, opts))

}

// Create takes the representation of a arangoJob and creates it.  Returns the server's representation of the arangoJob, and an error, if there is any.
func (c *FakeArangoJobs) Create(ctx context.Context, arangoJob *appsv1.ArangoJob, opts v1.CreateOptions) (result *appsv1.ArangoJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(arangojobsResource, c.ns, arangoJob), &appsv1.ArangoJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.ArangoJob), err
}

// Update takes the representation of a arangoJob and updates it. Returns the server's representation of the arangoJob, and an error, if there is any.
func (c *FakeArangoJobs) Update(ctx context.Context, arangoJob *appsv1.ArangoJob, opts v1.UpdateOptions) (result *appsv1.ArangoJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(arangojobsResource, c.ns, arangoJob), &appsv1.ArangoJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.ArangoJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeArangoJobs) UpdateStatus(ctx context.Context, arangoJob *appsv1.ArangoJob, opts v1.UpdateOptions) (*appsv1.ArangoJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(arangojobsResource, "status", c.ns, arangoJob), &appsv1.ArangoJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.ArangoJob), err
}

// Delete takes name of the arangoJob and deletes it. Returns an error if one occurs.
func (c *FakeArangoJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(arangojobsResource, c.ns, name), &appsv1.ArangoJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeArangoJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(arangojobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &appsv1.ArangoJobList{})
	return err
}

// Patch applies the patch and returns the patched arangoJob.
func (c *FakeArangoJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *appsv1.ArangoJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(arangojobsResource, c.ns, name, pt, data, subresources...), &appsv1.ArangoJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.ArangoJob), err
}
