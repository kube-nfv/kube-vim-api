// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	kubeovnv1 "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
	applyconfigurationkubeovnv1 "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/client/applyconfiguration/kubeovn/v1"
	scheme "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// IptablesFIPRulesGetter has a method to return a IptablesFIPRuleInterface.
// A group's client should implement this interface.
type IptablesFIPRulesGetter interface {
	IptablesFIPRules() IptablesFIPRuleInterface
}

// IptablesFIPRuleInterface has methods to work with IptablesFIPRule resources.
type IptablesFIPRuleInterface interface {
	Create(ctx context.Context, iptablesFIPRule *kubeovnv1.IptablesFIPRule, opts metav1.CreateOptions) (*kubeovnv1.IptablesFIPRule, error)
	Update(ctx context.Context, iptablesFIPRule *kubeovnv1.IptablesFIPRule, opts metav1.UpdateOptions) (*kubeovnv1.IptablesFIPRule, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, iptablesFIPRule *kubeovnv1.IptablesFIPRule, opts metav1.UpdateOptions) (*kubeovnv1.IptablesFIPRule, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubeovnv1.IptablesFIPRule, error)
	List(ctx context.Context, opts metav1.ListOptions) (*kubeovnv1.IptablesFIPRuleList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *kubeovnv1.IptablesFIPRule, err error)
	Apply(ctx context.Context, iptablesFIPRule *applyconfigurationkubeovnv1.IptablesFIPRuleApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.IptablesFIPRule, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, iptablesFIPRule *applyconfigurationkubeovnv1.IptablesFIPRuleApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.IptablesFIPRule, err error)
	IptablesFIPRuleExpansion
}

// iptablesFIPRules implements IptablesFIPRuleInterface
type iptablesFIPRules struct {
	*gentype.ClientWithListAndApply[*kubeovnv1.IptablesFIPRule, *kubeovnv1.IptablesFIPRuleList, *applyconfigurationkubeovnv1.IptablesFIPRuleApplyConfiguration]
}

// newIptablesFIPRules returns a IptablesFIPRules
func newIptablesFIPRules(c *KubeovnV1Client) *iptablesFIPRules {
	return &iptablesFIPRules{
		gentype.NewClientWithListAndApply[*kubeovnv1.IptablesFIPRule, *kubeovnv1.IptablesFIPRuleList, *applyconfigurationkubeovnv1.IptablesFIPRuleApplyConfiguration](
			"iptables-fip-rules",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *kubeovnv1.IptablesFIPRule { return &kubeovnv1.IptablesFIPRule{} },
			func() *kubeovnv1.IptablesFIPRuleList { return &kubeovnv1.IptablesFIPRuleList{} }),
	}
}
