// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	context "context"
	json "encoding/json"
	fmt "fmt"

	v1 "github.com/DiMalovanyy/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
	kubeovnv1 "github.com/DiMalovanyy/kube-vim-api/kube-ovn-api/pkg/client/applyconfiguration/kubeovn/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOvnDnatRules implements OvnDnatRuleInterface
type FakeOvnDnatRules struct {
	Fake *FakeKubeovnV1
}

var ovndnatrulesResource = v1.SchemeGroupVersion.WithResource("ovn-dnat-rules")

var ovndnatrulesKind = v1.SchemeGroupVersion.WithKind("OvnDnatRule")

// Get takes name of the ovnDnatRule, and returns the corresponding ovnDnatRule object, and an error if there is any.
func (c *FakeOvnDnatRules) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.OvnDnatRule, err error) {
	emptyResult := &v1.OvnDnatRule{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(ovndnatrulesResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OvnDnatRule), err
}

// List takes label and field selectors, and returns the list of OvnDnatRules that match those selectors.
func (c *FakeOvnDnatRules) List(ctx context.Context, opts metav1.ListOptions) (result *v1.OvnDnatRuleList, err error) {
	emptyResult := &v1.OvnDnatRuleList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(ovndnatrulesResource, ovndnatrulesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.OvnDnatRuleList{ListMeta: obj.(*v1.OvnDnatRuleList).ListMeta}
	for _, item := range obj.(*v1.OvnDnatRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ovnDnatRules.
func (c *FakeOvnDnatRules) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(ovndnatrulesResource, opts))
}

// Create takes the representation of a ovnDnatRule and creates it.  Returns the server's representation of the ovnDnatRule, and an error, if there is any.
func (c *FakeOvnDnatRules) Create(ctx context.Context, ovnDnatRule *v1.OvnDnatRule, opts metav1.CreateOptions) (result *v1.OvnDnatRule, err error) {
	emptyResult := &v1.OvnDnatRule{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(ovndnatrulesResource, ovnDnatRule, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OvnDnatRule), err
}

// Update takes the representation of a ovnDnatRule and updates it. Returns the server's representation of the ovnDnatRule, and an error, if there is any.
func (c *FakeOvnDnatRules) Update(ctx context.Context, ovnDnatRule *v1.OvnDnatRule, opts metav1.UpdateOptions) (result *v1.OvnDnatRule, err error) {
	emptyResult := &v1.OvnDnatRule{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(ovndnatrulesResource, ovnDnatRule, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OvnDnatRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOvnDnatRules) UpdateStatus(ctx context.Context, ovnDnatRule *v1.OvnDnatRule, opts metav1.UpdateOptions) (result *v1.OvnDnatRule, err error) {
	emptyResult := &v1.OvnDnatRule{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(ovndnatrulesResource, "status", ovnDnatRule, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OvnDnatRule), err
}

// Delete takes name of the ovnDnatRule and deletes it. Returns an error if one occurs.
func (c *FakeOvnDnatRules) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(ovndnatrulesResource, name, opts), &v1.OvnDnatRule{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOvnDnatRules) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(ovndnatrulesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.OvnDnatRuleList{})
	return err
}

// Patch applies the patch and returns the patched ovnDnatRule.
func (c *FakeOvnDnatRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.OvnDnatRule, err error) {
	emptyResult := &v1.OvnDnatRule{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(ovndnatrulesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OvnDnatRule), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied ovnDnatRule.
func (c *FakeOvnDnatRules) Apply(ctx context.Context, ovnDnatRule *kubeovnv1.OvnDnatRuleApplyConfiguration, opts metav1.ApplyOptions) (result *v1.OvnDnatRule, err error) {
	if ovnDnatRule == nil {
		return nil, fmt.Errorf("ovnDnatRule provided to Apply must not be nil")
	}
	data, err := json.Marshal(ovnDnatRule)
	if err != nil {
		return nil, err
	}
	name := ovnDnatRule.Name
	if name == nil {
		return nil, fmt.Errorf("ovnDnatRule.Name must be provided to Apply")
	}
	emptyResult := &v1.OvnDnatRule{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(ovndnatrulesResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OvnDnatRule), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeOvnDnatRules) ApplyStatus(ctx context.Context, ovnDnatRule *kubeovnv1.OvnDnatRuleApplyConfiguration, opts metav1.ApplyOptions) (result *v1.OvnDnatRule, err error) {
	if ovnDnatRule == nil {
		return nil, fmt.Errorf("ovnDnatRule provided to Apply must not be nil")
	}
	data, err := json.Marshal(ovnDnatRule)
	if err != nil {
		return nil, err
	}
	name := ovnDnatRule.Name
	if name == nil {
		return nil, fmt.Errorf("ovnDnatRule.Name must be provided to Apply")
	}
	emptyResult := &v1.OvnDnatRule{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(ovndnatrulesResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OvnDnatRule), err
}