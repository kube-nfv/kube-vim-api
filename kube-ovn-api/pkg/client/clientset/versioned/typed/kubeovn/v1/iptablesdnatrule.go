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

// IptablesDnatRulesGetter has a method to return a IptablesDnatRuleInterface.
// A group's client should implement this interface.
type IptablesDnatRulesGetter interface {
	IptablesDnatRules() IptablesDnatRuleInterface
}

// IptablesDnatRuleInterface has methods to work with IptablesDnatRule resources.
type IptablesDnatRuleInterface interface {
	Create(ctx context.Context, iptablesDnatRule *kubeovnv1.IptablesDnatRule, opts metav1.CreateOptions) (*kubeovnv1.IptablesDnatRule, error)
	Update(ctx context.Context, iptablesDnatRule *kubeovnv1.IptablesDnatRule, opts metav1.UpdateOptions) (*kubeovnv1.IptablesDnatRule, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, iptablesDnatRule *kubeovnv1.IptablesDnatRule, opts metav1.UpdateOptions) (*kubeovnv1.IptablesDnatRule, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubeovnv1.IptablesDnatRule, error)
	List(ctx context.Context, opts metav1.ListOptions) (*kubeovnv1.IptablesDnatRuleList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *kubeovnv1.IptablesDnatRule, err error)
	Apply(ctx context.Context, iptablesDnatRule *applyconfigurationkubeovnv1.IptablesDnatRuleApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.IptablesDnatRule, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, iptablesDnatRule *applyconfigurationkubeovnv1.IptablesDnatRuleApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.IptablesDnatRule, err error)
	IptablesDnatRuleExpansion
}

// iptablesDnatRules implements IptablesDnatRuleInterface
type iptablesDnatRules struct {
	*gentype.ClientWithListAndApply[*kubeovnv1.IptablesDnatRule, *kubeovnv1.IptablesDnatRuleList, *applyconfigurationkubeovnv1.IptablesDnatRuleApplyConfiguration]
}

// newIptablesDnatRules returns a IptablesDnatRules
func newIptablesDnatRules(c *KubeovnV1Client) *iptablesDnatRules {
	return &iptablesDnatRules{
		gentype.NewClientWithListAndApply[*kubeovnv1.IptablesDnatRule, *kubeovnv1.IptablesDnatRuleList, *applyconfigurationkubeovnv1.IptablesDnatRuleApplyConfiguration](
			"iptables-dnat-rules",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *kubeovnv1.IptablesDnatRule { return &kubeovnv1.IptablesDnatRule{} },
			func() *kubeovnv1.IptablesDnatRuleList { return &kubeovnv1.IptablesDnatRuleList{} }),
	}
}
