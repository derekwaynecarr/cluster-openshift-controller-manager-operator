// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	operator_v1 "github.com/openshift/api/operator/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceCAs implements ServiceCAInterface
type FakeServiceCAs struct {
	Fake *FakeOperatorV1
}

var servicecasResource = schema.GroupVersionResource{Group: "operator.openshift.io", Version: "v1", Resource: "servicecas"}

var servicecasKind = schema.GroupVersionKind{Group: "operator.openshift.io", Version: "v1", Kind: "ServiceCA"}

// Get takes name of the serviceCA, and returns the corresponding serviceCA object, and an error if there is any.
func (c *FakeServiceCAs) Get(name string, options v1.GetOptions) (result *operator_v1.ServiceCA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(servicecasResource, name), &operator_v1.ServiceCA{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operator_v1.ServiceCA), err
}

// List takes label and field selectors, and returns the list of ServiceCAs that match those selectors.
func (c *FakeServiceCAs) List(opts v1.ListOptions) (result *operator_v1.ServiceCAList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(servicecasResource, servicecasKind, opts), &operator_v1.ServiceCAList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operator_v1.ServiceCAList{ListMeta: obj.(*operator_v1.ServiceCAList).ListMeta}
	for _, item := range obj.(*operator_v1.ServiceCAList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceCAs.
func (c *FakeServiceCAs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(servicecasResource, opts))
}

// Create takes the representation of a serviceCA and creates it.  Returns the server's representation of the serviceCA, and an error, if there is any.
func (c *FakeServiceCAs) Create(serviceCA *operator_v1.ServiceCA) (result *operator_v1.ServiceCA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(servicecasResource, serviceCA), &operator_v1.ServiceCA{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operator_v1.ServiceCA), err
}

// Update takes the representation of a serviceCA and updates it. Returns the server's representation of the serviceCA, and an error, if there is any.
func (c *FakeServiceCAs) Update(serviceCA *operator_v1.ServiceCA) (result *operator_v1.ServiceCA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(servicecasResource, serviceCA), &operator_v1.ServiceCA{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operator_v1.ServiceCA), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceCAs) UpdateStatus(serviceCA *operator_v1.ServiceCA) (*operator_v1.ServiceCA, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(servicecasResource, "status", serviceCA), &operator_v1.ServiceCA{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operator_v1.ServiceCA), err
}

// Delete takes name of the serviceCA and deletes it. Returns an error if one occurs.
func (c *FakeServiceCAs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(servicecasResource, name), &operator_v1.ServiceCA{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceCAs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(servicecasResource, listOptions)

	_, err := c.Fake.Invokes(action, &operator_v1.ServiceCAList{})
	return err
}

// Patch applies the patch and returns the patched serviceCA.
func (c *FakeServiceCAs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *operator_v1.ServiceCA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(servicecasResource, name, data, subresources...), &operator_v1.ServiceCA{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operator_v1.ServiceCA), err
}
