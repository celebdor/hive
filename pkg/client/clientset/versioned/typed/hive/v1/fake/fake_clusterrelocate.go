// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openshift/hive/apis/hive/v1"
	hivev1 "github.com/openshift/hive/pkg/client/applyconfiguration/hive/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterRelocates implements ClusterRelocateInterface
type FakeClusterRelocates struct {
	Fake *FakeHiveV1
}

var clusterrelocatesResource = v1.SchemeGroupVersion.WithResource("clusterrelocates")

var clusterrelocatesKind = v1.SchemeGroupVersion.WithKind("ClusterRelocate")

// Get takes name of the clusterRelocate, and returns the corresponding clusterRelocate object, and an error if there is any.
func (c *FakeClusterRelocates) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClusterRelocate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterrelocatesResource, name), &v1.ClusterRelocate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterRelocate), err
}

// List takes label and field selectors, and returns the list of ClusterRelocates that match those selectors.
func (c *FakeClusterRelocates) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterRelocateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterrelocatesResource, clusterrelocatesKind, opts), &v1.ClusterRelocateList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ClusterRelocateList{ListMeta: obj.(*v1.ClusterRelocateList).ListMeta}
	for _, item := range obj.(*v1.ClusterRelocateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterRelocates.
func (c *FakeClusterRelocates) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterrelocatesResource, opts))
}

// Create takes the representation of a clusterRelocate and creates it.  Returns the server's representation of the clusterRelocate, and an error, if there is any.
func (c *FakeClusterRelocates) Create(ctx context.Context, clusterRelocate *v1.ClusterRelocate, opts metav1.CreateOptions) (result *v1.ClusterRelocate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterrelocatesResource, clusterRelocate), &v1.ClusterRelocate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterRelocate), err
}

// Update takes the representation of a clusterRelocate and updates it. Returns the server's representation of the clusterRelocate, and an error, if there is any.
func (c *FakeClusterRelocates) Update(ctx context.Context, clusterRelocate *v1.ClusterRelocate, opts metav1.UpdateOptions) (result *v1.ClusterRelocate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterrelocatesResource, clusterRelocate), &v1.ClusterRelocate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterRelocate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterRelocates) UpdateStatus(ctx context.Context, clusterRelocate *v1.ClusterRelocate, opts metav1.UpdateOptions) (*v1.ClusterRelocate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterrelocatesResource, "status", clusterRelocate), &v1.ClusterRelocate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterRelocate), err
}

// Delete takes name of the clusterRelocate and deletes it. Returns an error if one occurs.
func (c *FakeClusterRelocates) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterrelocatesResource, name, opts), &v1.ClusterRelocate{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterRelocates) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterrelocatesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ClusterRelocateList{})
	return err
}

// Patch applies the patch and returns the patched clusterRelocate.
func (c *FakeClusterRelocates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterRelocate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterrelocatesResource, name, pt, data, subresources...), &v1.ClusterRelocate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterRelocate), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied clusterRelocate.
func (c *FakeClusterRelocates) Apply(ctx context.Context, clusterRelocate *hivev1.ClusterRelocateApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ClusterRelocate, err error) {
	if clusterRelocate == nil {
		return nil, fmt.Errorf("clusterRelocate provided to Apply must not be nil")
	}
	data, err := json.Marshal(clusterRelocate)
	if err != nil {
		return nil, err
	}
	name := clusterRelocate.Name
	if name == nil {
		return nil, fmt.Errorf("clusterRelocate.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterrelocatesResource, *name, types.ApplyPatchType, data), &v1.ClusterRelocate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterRelocate), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeClusterRelocates) ApplyStatus(ctx context.Context, clusterRelocate *hivev1.ClusterRelocateApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ClusterRelocate, err error) {
	if clusterRelocate == nil {
		return nil, fmt.Errorf("clusterRelocate provided to Apply must not be nil")
	}
	data, err := json.Marshal(clusterRelocate)
	if err != nil {
		return nil, err
	}
	name := clusterRelocate.Name
	if name == nil {
		return nil, fmt.Errorf("clusterRelocate.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterrelocatesResource, *name, types.ApplyPatchType, data, "status"), &v1.ClusterRelocate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterRelocate), err
}