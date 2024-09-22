// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	context "context"
	json "encoding/json"
	fmt "fmt"

	v1 "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
	kubeovnv1 "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/client/applyconfiguration/kubeovn/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIptablesEIPs implements IptablesEIPInterface
type FakeIptablesEIPs struct {
	Fake *FakeKubeovnV1
}

var iptableseipsResource = v1.SchemeGroupVersion.WithResource("iptables-eips")

var iptableseipsKind = v1.SchemeGroupVersion.WithKind("IptablesEIP")

// Get takes name of the iptablesEIP, and returns the corresponding iptablesEIP object, and an error if there is any.
func (c *FakeIptablesEIPs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.IptablesEIP, err error) {
	emptyResult := &v1.IptablesEIP{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(iptableseipsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.IptablesEIP), err
}

// List takes label and field selectors, and returns the list of IptablesEIPs that match those selectors.
func (c *FakeIptablesEIPs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.IptablesEIPList, err error) {
	emptyResult := &v1.IptablesEIPList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(iptableseipsResource, iptableseipsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.IptablesEIPList{ListMeta: obj.(*v1.IptablesEIPList).ListMeta}
	for _, item := range obj.(*v1.IptablesEIPList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iptablesEIPs.
func (c *FakeIptablesEIPs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(iptableseipsResource, opts))
}

// Create takes the representation of a iptablesEIP and creates it.  Returns the server's representation of the iptablesEIP, and an error, if there is any.
func (c *FakeIptablesEIPs) Create(ctx context.Context, iptablesEIP *v1.IptablesEIP, opts metav1.CreateOptions) (result *v1.IptablesEIP, err error) {
	emptyResult := &v1.IptablesEIP{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(iptableseipsResource, iptablesEIP, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.IptablesEIP), err
}

// Update takes the representation of a iptablesEIP and updates it. Returns the server's representation of the iptablesEIP, and an error, if there is any.
func (c *FakeIptablesEIPs) Update(ctx context.Context, iptablesEIP *v1.IptablesEIP, opts metav1.UpdateOptions) (result *v1.IptablesEIP, err error) {
	emptyResult := &v1.IptablesEIP{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(iptableseipsResource, iptablesEIP, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.IptablesEIP), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIptablesEIPs) UpdateStatus(ctx context.Context, iptablesEIP *v1.IptablesEIP, opts metav1.UpdateOptions) (result *v1.IptablesEIP, err error) {
	emptyResult := &v1.IptablesEIP{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(iptableseipsResource, "status", iptablesEIP, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.IptablesEIP), err
}

// Delete takes name of the iptablesEIP and deletes it. Returns an error if one occurs.
func (c *FakeIptablesEIPs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(iptableseipsResource, name, opts), &v1.IptablesEIP{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIptablesEIPs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(iptableseipsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.IptablesEIPList{})
	return err
}

// Patch applies the patch and returns the patched iptablesEIP.
func (c *FakeIptablesEIPs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.IptablesEIP, err error) {
	emptyResult := &v1.IptablesEIP{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(iptableseipsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.IptablesEIP), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied iptablesEIP.
func (c *FakeIptablesEIPs) Apply(ctx context.Context, iptablesEIP *kubeovnv1.IptablesEIPApplyConfiguration, opts metav1.ApplyOptions) (result *v1.IptablesEIP, err error) {
	if iptablesEIP == nil {
		return nil, fmt.Errorf("iptablesEIP provided to Apply must not be nil")
	}
	data, err := json.Marshal(iptablesEIP)
	if err != nil {
		return nil, err
	}
	name := iptablesEIP.Name
	if name == nil {
		return nil, fmt.Errorf("iptablesEIP.Name must be provided to Apply")
	}
	emptyResult := &v1.IptablesEIP{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(iptableseipsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.IptablesEIP), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeIptablesEIPs) ApplyStatus(ctx context.Context, iptablesEIP *kubeovnv1.IptablesEIPApplyConfiguration, opts metav1.ApplyOptions) (result *v1.IptablesEIP, err error) {
	if iptablesEIP == nil {
		return nil, fmt.Errorf("iptablesEIP provided to Apply must not be nil")
	}
	data, err := json.Marshal(iptablesEIP)
	if err != nil {
		return nil, err
	}
	name := iptablesEIP.Name
	if name == nil {
		return nil, fmt.Errorf("iptablesEIP.Name must be provided to Apply")
	}
	emptyResult := &v1.IptablesEIP{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(iptableseipsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.IptablesEIP), err
}
