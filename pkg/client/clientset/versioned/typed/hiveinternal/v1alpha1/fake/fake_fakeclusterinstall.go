// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/openshift/hive/apis/hiveinternal/v1alpha1"
	hiveinternalv1alpha1 "github.com/openshift/hive/pkg/client/applyconfiguration/hiveinternal/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFakeClusterInstalls implements FakeClusterInstallInterface
type FakeFakeClusterInstalls struct {
	Fake *FakeHiveinternalV1alpha1
	ns   string
}

var fakeclusterinstallsResource = v1alpha1.SchemeGroupVersion.WithResource("fakeclusterinstalls")

var fakeclusterinstallsKind = v1alpha1.SchemeGroupVersion.WithKind("FakeClusterInstall")

// Get takes name of the fakeClusterInstall, and returns the corresponding fakeClusterInstall object, and an error if there is any.
func (c *FakeFakeClusterInstalls) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FakeClusterInstall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(fakeclusterinstallsResource, c.ns, name), &v1alpha1.FakeClusterInstall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FakeClusterInstall), err
}

// List takes label and field selectors, and returns the list of FakeClusterInstalls that match those selectors.
func (c *FakeFakeClusterInstalls) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FakeClusterInstallList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(fakeclusterinstallsResource, fakeclusterinstallsKind, c.ns, opts), &v1alpha1.FakeClusterInstallList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FakeClusterInstallList{ListMeta: obj.(*v1alpha1.FakeClusterInstallList).ListMeta}
	for _, item := range obj.(*v1alpha1.FakeClusterInstallList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fakeClusterInstalls.
func (c *FakeFakeClusterInstalls) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(fakeclusterinstallsResource, c.ns, opts))

}

// Create takes the representation of a fakeClusterInstall and creates it.  Returns the server's representation of the fakeClusterInstall, and an error, if there is any.
func (c *FakeFakeClusterInstalls) Create(ctx context.Context, fakeClusterInstall *v1alpha1.FakeClusterInstall, opts v1.CreateOptions) (result *v1alpha1.FakeClusterInstall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(fakeclusterinstallsResource, c.ns, fakeClusterInstall), &v1alpha1.FakeClusterInstall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FakeClusterInstall), err
}

// Update takes the representation of a fakeClusterInstall and updates it. Returns the server's representation of the fakeClusterInstall, and an error, if there is any.
func (c *FakeFakeClusterInstalls) Update(ctx context.Context, fakeClusterInstall *v1alpha1.FakeClusterInstall, opts v1.UpdateOptions) (result *v1alpha1.FakeClusterInstall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(fakeclusterinstallsResource, c.ns, fakeClusterInstall), &v1alpha1.FakeClusterInstall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FakeClusterInstall), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFakeClusterInstalls) UpdateStatus(ctx context.Context, fakeClusterInstall *v1alpha1.FakeClusterInstall, opts v1.UpdateOptions) (*v1alpha1.FakeClusterInstall, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(fakeclusterinstallsResource, "status", c.ns, fakeClusterInstall), &v1alpha1.FakeClusterInstall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FakeClusterInstall), err
}

// Delete takes name of the fakeClusterInstall and deletes it. Returns an error if one occurs.
func (c *FakeFakeClusterInstalls) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(fakeclusterinstallsResource, c.ns, name, opts), &v1alpha1.FakeClusterInstall{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFakeClusterInstalls) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(fakeclusterinstallsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FakeClusterInstallList{})
	return err
}

// Patch applies the patch and returns the patched fakeClusterInstall.
func (c *FakeFakeClusterInstalls) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FakeClusterInstall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fakeclusterinstallsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FakeClusterInstall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FakeClusterInstall), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied fakeClusterInstall.
func (c *FakeFakeClusterInstalls) Apply(ctx context.Context, fakeClusterInstall *hiveinternalv1alpha1.FakeClusterInstallApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FakeClusterInstall, err error) {
	if fakeClusterInstall == nil {
		return nil, fmt.Errorf("fakeClusterInstall provided to Apply must not be nil")
	}
	data, err := json.Marshal(fakeClusterInstall)
	if err != nil {
		return nil, err
	}
	name := fakeClusterInstall.Name
	if name == nil {
		return nil, fmt.Errorf("fakeClusterInstall.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fakeclusterinstallsResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha1.FakeClusterInstall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FakeClusterInstall), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeFakeClusterInstalls) ApplyStatus(ctx context.Context, fakeClusterInstall *hiveinternalv1alpha1.FakeClusterInstallApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FakeClusterInstall, err error) {
	if fakeClusterInstall == nil {
		return nil, fmt.Errorf("fakeClusterInstall provided to Apply must not be nil")
	}
	data, err := json.Marshal(fakeClusterInstall)
	if err != nil {
		return nil, err
	}
	name := fakeClusterInstall.Name
	if name == nil {
		return nil, fmt.Errorf("fakeClusterInstall.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fakeclusterinstallsResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1alpha1.FakeClusterInstall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FakeClusterInstall), err
}