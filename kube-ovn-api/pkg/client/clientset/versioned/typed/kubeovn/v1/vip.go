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

// VipsGetter has a method to return a VipInterface.
// A group's client should implement this interface.
type VipsGetter interface {
	Vips() VipInterface
}

// VipInterface has methods to work with Vip resources.
type VipInterface interface {
	Create(ctx context.Context, vip *kubeovnv1.Vip, opts metav1.CreateOptions) (*kubeovnv1.Vip, error)
	Update(ctx context.Context, vip *kubeovnv1.Vip, opts metav1.UpdateOptions) (*kubeovnv1.Vip, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, vip *kubeovnv1.Vip, opts metav1.UpdateOptions) (*kubeovnv1.Vip, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubeovnv1.Vip, error)
	List(ctx context.Context, opts metav1.ListOptions) (*kubeovnv1.VipList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *kubeovnv1.Vip, err error)
	Apply(ctx context.Context, vip *applyconfigurationkubeovnv1.VipApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.Vip, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, vip *applyconfigurationkubeovnv1.VipApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.Vip, err error)
	VipExpansion
}

// vips implements VipInterface
type vips struct {
	*gentype.ClientWithListAndApply[*kubeovnv1.Vip, *kubeovnv1.VipList, *applyconfigurationkubeovnv1.VipApplyConfiguration]
}

// newVips returns a Vips
func newVips(c *KubeovnV1Client) *vips {
	return &vips{
		gentype.NewClientWithListAndApply[*kubeovnv1.Vip, *kubeovnv1.VipList, *applyconfigurationkubeovnv1.VipApplyConfiguration](
			"vips",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *kubeovnv1.Vip { return &kubeovnv1.Vip{} },
			func() *kubeovnv1.VipList { return &kubeovnv1.VipList{} }),
	}
}
