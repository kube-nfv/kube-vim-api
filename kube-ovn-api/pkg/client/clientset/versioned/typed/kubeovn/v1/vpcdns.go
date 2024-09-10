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

// VpcDnsesGetter has a method to return a VpcDnsInterface.
// A group's client should implement this interface.
type VpcDnsesGetter interface {
	VpcDnses() VpcDnsInterface
}

// VpcDnsInterface has methods to work with VpcDns resources.
type VpcDnsInterface interface {
	Create(ctx context.Context, vpcDns *kubeovnv1.VpcDns, opts metav1.CreateOptions) (*kubeovnv1.VpcDns, error)
	Update(ctx context.Context, vpcDns *kubeovnv1.VpcDns, opts metav1.UpdateOptions) (*kubeovnv1.VpcDns, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, vpcDns *kubeovnv1.VpcDns, opts metav1.UpdateOptions) (*kubeovnv1.VpcDns, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubeovnv1.VpcDns, error)
	List(ctx context.Context, opts metav1.ListOptions) (*kubeovnv1.VpcDnsList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *kubeovnv1.VpcDns, err error)
	Apply(ctx context.Context, vpcDns *applyconfigurationkubeovnv1.VpcDnsApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.VpcDns, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, vpcDns *applyconfigurationkubeovnv1.VpcDnsApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.VpcDns, err error)
	VpcDnsExpansion
}

// vpcDnses implements VpcDnsInterface
type vpcDnses struct {
	*gentype.ClientWithListAndApply[*kubeovnv1.VpcDns, *kubeovnv1.VpcDnsList, *applyconfigurationkubeovnv1.VpcDnsApplyConfiguration]
}

// newVpcDnses returns a VpcDnses
func newVpcDnses(c *KubeovnV1Client) *vpcDnses {
	return &vpcDnses{
		gentype.NewClientWithListAndApply[*kubeovnv1.VpcDns, *kubeovnv1.VpcDnsList, *applyconfigurationkubeovnv1.VpcDnsApplyConfiguration](
			"vpc-dnses",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *kubeovnv1.VpcDns { return &kubeovnv1.VpcDns{} },
			func() *kubeovnv1.VpcDnsList { return &kubeovnv1.VpcDnsList{} }),
	}
}
