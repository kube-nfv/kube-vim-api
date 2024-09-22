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

// IPsGetter has a method to return a IPInterface.
// A group's client should implement this interface.
type IPsGetter interface {
	IPs() IPInterface
}

// IPInterface has methods to work with IP resources.
type IPInterface interface {
	Create(ctx context.Context, iP *kubeovnv1.IP, opts metav1.CreateOptions) (*kubeovnv1.IP, error)
	Update(ctx context.Context, iP *kubeovnv1.IP, opts metav1.UpdateOptions) (*kubeovnv1.IP, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubeovnv1.IP, error)
	List(ctx context.Context, opts metav1.ListOptions) (*kubeovnv1.IPList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *kubeovnv1.IP, err error)
	Apply(ctx context.Context, iP *applyconfigurationkubeovnv1.IPApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.IP, err error)
	IPExpansion
}

// iPs implements IPInterface
type iPs struct {
	*gentype.ClientWithListAndApply[*kubeovnv1.IP, *kubeovnv1.IPList, *applyconfigurationkubeovnv1.IPApplyConfiguration]
}

// newIPs returns a IPs
func newIPs(c *KubeovnV1Client) *iPs {
	return &iPs{
		gentype.NewClientWithListAndApply[*kubeovnv1.IP, *kubeovnv1.IPList, *applyconfigurationkubeovnv1.IPApplyConfiguration](
			"ips",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *kubeovnv1.IP { return &kubeovnv1.IP{} },
			func() *kubeovnv1.IPList { return &kubeovnv1.IPList{} }),
	}
}
