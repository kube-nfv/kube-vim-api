// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	kubeovnv1 "github.com/DiMalovanyy/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
	applyconfigurationkubeovnv1 "github.com/DiMalovanyy/kube-vim-api/kube-ovn-api/pkg/client/applyconfiguration/kubeovn/v1"
	scheme "github.com/DiMalovanyy/kube-vim-api/kube-ovn-api/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// IptablesSnatRulesGetter has a method to return a IptablesSnatRuleInterface.
// A group's client should implement this interface.
type IptablesSnatRulesGetter interface {
	IptablesSnatRules() IptablesSnatRuleInterface
}

// IptablesSnatRuleInterface has methods to work with IptablesSnatRule resources.
type IptablesSnatRuleInterface interface {
	Create(ctx context.Context, iptablesSnatRule *kubeovnv1.IptablesSnatRule, opts metav1.CreateOptions) (*kubeovnv1.IptablesSnatRule, error)
	Update(ctx context.Context, iptablesSnatRule *kubeovnv1.IptablesSnatRule, opts metav1.UpdateOptions) (*kubeovnv1.IptablesSnatRule, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, iptablesSnatRule *kubeovnv1.IptablesSnatRule, opts metav1.UpdateOptions) (*kubeovnv1.IptablesSnatRule, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubeovnv1.IptablesSnatRule, error)
	List(ctx context.Context, opts metav1.ListOptions) (*kubeovnv1.IptablesSnatRuleList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *kubeovnv1.IptablesSnatRule, err error)
	Apply(ctx context.Context, iptablesSnatRule *applyconfigurationkubeovnv1.IptablesSnatRuleApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.IptablesSnatRule, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, iptablesSnatRule *applyconfigurationkubeovnv1.IptablesSnatRuleApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.IptablesSnatRule, err error)
	IptablesSnatRuleExpansion
}

// iptablesSnatRules implements IptablesSnatRuleInterface
type iptablesSnatRules struct {
	*gentype.ClientWithListAndApply[*kubeovnv1.IptablesSnatRule, *kubeovnv1.IptablesSnatRuleList, *applyconfigurationkubeovnv1.IptablesSnatRuleApplyConfiguration]
}

// newIptablesSnatRules returns a IptablesSnatRules
func newIptablesSnatRules(c *KubeovnV1Client) *iptablesSnatRules {
	return &iptablesSnatRules{
		gentype.NewClientWithListAndApply[*kubeovnv1.IptablesSnatRule, *kubeovnv1.IptablesSnatRuleList, *applyconfigurationkubeovnv1.IptablesSnatRuleApplyConfiguration](
			"iptables-snat-rules",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *kubeovnv1.IptablesSnatRule { return &kubeovnv1.IptablesSnatRule{} },
			func() *kubeovnv1.IptablesSnatRuleList { return &kubeovnv1.IptablesSnatRuleList{} }),
	}
}