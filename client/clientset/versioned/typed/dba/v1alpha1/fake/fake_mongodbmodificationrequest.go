/*
Copyright The KubeDB Authors.

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
	v1alpha1 "kubedb.dev/apimachinery/apis/dba/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMongoDBModificationRequests implements MongoDBModificationRequestInterface
type FakeMongoDBModificationRequests struct {
	Fake *FakeDbaV1alpha1
	ns   string
}

var mongodbmodificationrequestsResource = schema.GroupVersionResource{Group: "dba.kubedb.com", Version: "v1alpha1", Resource: "mongodbmodificationrequests"}

var mongodbmodificationrequestsKind = schema.GroupVersionKind{Group: "dba.kubedb.com", Version: "v1alpha1", Kind: "MongoDBModificationRequest"}

// Get takes name of the mongoDBModificationRequest, and returns the corresponding mongoDBModificationRequest object, and an error if there is any.
func (c *FakeMongoDBModificationRequests) Get(name string, options v1.GetOptions) (result *v1alpha1.MongoDBModificationRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mongodbmodificationrequestsResource, c.ns, name), &v1alpha1.MongoDBModificationRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MongoDBModificationRequest), err
}

// List takes label and field selectors, and returns the list of MongoDBModificationRequests that match those selectors.
func (c *FakeMongoDBModificationRequests) List(opts v1.ListOptions) (result *v1alpha1.MongoDBModificationRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mongodbmodificationrequestsResource, mongodbmodificationrequestsKind, c.ns, opts), &v1alpha1.MongoDBModificationRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MongoDBModificationRequestList{ListMeta: obj.(*v1alpha1.MongoDBModificationRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.MongoDBModificationRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mongoDBModificationRequests.
func (c *FakeMongoDBModificationRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mongodbmodificationrequestsResource, c.ns, opts))

}

// Create takes the representation of a mongoDBModificationRequest and creates it.  Returns the server's representation of the mongoDBModificationRequest, and an error, if there is any.
func (c *FakeMongoDBModificationRequests) Create(mongoDBModificationRequest *v1alpha1.MongoDBModificationRequest) (result *v1alpha1.MongoDBModificationRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mongodbmodificationrequestsResource, c.ns, mongoDBModificationRequest), &v1alpha1.MongoDBModificationRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MongoDBModificationRequest), err
}

// Update takes the representation of a mongoDBModificationRequest and updates it. Returns the server's representation of the mongoDBModificationRequest, and an error, if there is any.
func (c *FakeMongoDBModificationRequests) Update(mongoDBModificationRequest *v1alpha1.MongoDBModificationRequest) (result *v1alpha1.MongoDBModificationRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mongodbmodificationrequestsResource, c.ns, mongoDBModificationRequest), &v1alpha1.MongoDBModificationRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MongoDBModificationRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMongoDBModificationRequests) UpdateStatus(mongoDBModificationRequest *v1alpha1.MongoDBModificationRequest) (*v1alpha1.MongoDBModificationRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mongodbmodificationrequestsResource, "status", c.ns, mongoDBModificationRequest), &v1alpha1.MongoDBModificationRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MongoDBModificationRequest), err
}

// Delete takes name of the mongoDBModificationRequest and deletes it. Returns an error if one occurs.
func (c *FakeMongoDBModificationRequests) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mongodbmodificationrequestsResource, c.ns, name), &v1alpha1.MongoDBModificationRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMongoDBModificationRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mongodbmodificationrequestsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MongoDBModificationRequestList{})
	return err
}

// Patch applies the patch and returns the patched mongoDBModificationRequest.
func (c *FakeMongoDBModificationRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MongoDBModificationRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mongodbmodificationrequestsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MongoDBModificationRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MongoDBModificationRequest), err
}
