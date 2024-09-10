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

// VlansGetter has a method to return a VlanInterface.
// A group's client should implement this interface.
type VlansGetter interface {
	Vlans() VlanInterface
}

// VlanInterface has methods to work with Vlan resources.
type VlanInterface interface {
	Create(ctx context.Context, vlan *kubeovnv1.Vlan, opts metav1.CreateOptions) (*kubeovnv1.Vlan, error)
	Update(ctx context.Context, vlan *kubeovnv1.Vlan, opts metav1.UpdateOptions) (*kubeovnv1.Vlan, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, vlan *kubeovnv1.Vlan, opts metav1.UpdateOptions) (*kubeovnv1.Vlan, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubeovnv1.Vlan, error)
	List(ctx context.Context, opts metav1.ListOptions) (*kubeovnv1.VlanList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *kubeovnv1.Vlan, err error)
	Apply(ctx context.Context, vlan *applyconfigurationkubeovnv1.VlanApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.Vlan, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, vlan *applyconfigurationkubeovnv1.VlanApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.Vlan, err error)
	VlanExpansion
}

// vlans implements VlanInterface
type vlans struct {
	*gentype.ClientWithListAndApply[*kubeovnv1.Vlan, *kubeovnv1.VlanList, *applyconfigurationkubeovnv1.VlanApplyConfiguration]
}

// newVlans returns a Vlans
func newVlans(c *KubeovnV1Client) *vlans {
	return &vlans{
		gentype.NewClientWithListAndApply[*kubeovnv1.Vlan, *kubeovnv1.VlanList, *applyconfigurationkubeovnv1.VlanApplyConfiguration](
			"vlans",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *kubeovnv1.Vlan { return &kubeovnv1.Vlan{} },
			func() *kubeovnv1.VlanList { return &kubeovnv1.VlanList{} }),
	}
}
