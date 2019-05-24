/*
Copyright 2017 The Kubernetes Authors.

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

package fake

import (
	elasticsearchoperator_v1 "github.com/upmc-enterprises/elasticsearch-operator/pkg/apis/elasticsearchoperator/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeElasticsearchClusters implements ElasticsearchClusterInterface
type FakeElasticsearchClusters struct {
	Fake *FakeEnterprisesV1
	ns   string
}

var elasticsearchclustersResource = schema.GroupVersionResource{Group: "enterprises.upmc.com", Version: "v1", Resource: "elasticsearchclusters"}

var elasticsearchclustersKind = schema.GroupVersionKind{Group: "enterprises.upmc.com", Version: "v1", Kind: "ElasticsearchCluster"}

// Get takes name of the elasticsearchCluster, and returns the corresponding elasticsearchCluster object, and an error if there is any.
func (c *FakeElasticsearchClusters) Get(name string, options v1.GetOptions) (result *elasticsearchoperator_v1.ElasticsearchCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(elasticsearchclustersResource, c.ns, name), &elasticsearchoperator_v1.ElasticsearchCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*elasticsearchoperator_v1.ElasticsearchCluster), err
}

// List takes label and field selectors, and returns the list of ElasticsearchClusters that match those selectors.
func (c *FakeElasticsearchClusters) List(opts v1.ListOptions) (result *elasticsearchoperator_v1.ElasticsearchClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(elasticsearchclustersResource, elasticsearchclustersKind, c.ns, opts), &elasticsearchoperator_v1.ElasticsearchClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &elasticsearchoperator_v1.ElasticsearchClusterList{}
	for _, item := range obj.(*elasticsearchoperator_v1.ElasticsearchClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested elasticsearchClusters.
func (c *FakeElasticsearchClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(elasticsearchclustersResource, c.ns, opts))

}

// Create takes the representation of a elasticsearchCluster and creates it.  Returns the server's representation of the elasticsearchCluster, and an error, if there is any.
func (c *FakeElasticsearchClusters) Create(elasticsearchCluster *elasticsearchoperator_v1.ElasticsearchCluster) (result *elasticsearchoperator_v1.ElasticsearchCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(elasticsearchclustersResource, c.ns, elasticsearchCluster), &elasticsearchoperator_v1.ElasticsearchCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*elasticsearchoperator_v1.ElasticsearchCluster), err
}

// Update takes the representation of a elasticsearchCluster and updates it. Returns the server's representation of the elasticsearchCluster, and an error, if there is any.
func (c *FakeElasticsearchClusters) Update(elasticsearchCluster *elasticsearchoperator_v1.ElasticsearchCluster) (result *elasticsearchoperator_v1.ElasticsearchCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(elasticsearchclustersResource, c.ns, elasticsearchCluster), &elasticsearchoperator_v1.ElasticsearchCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*elasticsearchoperator_v1.ElasticsearchCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeElasticsearchClusters) UpdateStatus(elasticsearchCluster *elasticsearchoperator_v1.ElasticsearchCluster) (*elasticsearchoperator_v1.ElasticsearchCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(elasticsearchclustersResource, "status", c.ns, elasticsearchCluster), &elasticsearchoperator_v1.ElasticsearchCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*elasticsearchoperator_v1.ElasticsearchCluster), err
}

// Delete takes name of the elasticsearchCluster and deletes it. Returns an error if one occurs.
func (c *FakeElasticsearchClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(elasticsearchclustersResource, c.ns, name), &elasticsearchoperator_v1.ElasticsearchCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeElasticsearchClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(elasticsearchclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &elasticsearchoperator_v1.ElasticsearchClusterList{})
	return err
}

// Patch applies the patch and returns the patched elasticsearchCluster.
func (c *FakeElasticsearchClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *elasticsearchoperator_v1.ElasticsearchCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(elasticsearchclustersResource, c.ns, name, data, subresources...), &elasticsearchoperator_v1.ElasticsearchCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*elasticsearchoperator_v1.ElasticsearchCluster), err
}