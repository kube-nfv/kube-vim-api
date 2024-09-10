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

// OvnEipsGetter has a method to return a OvnEipInterface.
// A group's client should implement this interface.
type OvnEipsGetter interface {
	OvnEips() OvnEipInterface
}

// OvnEipInterface has methods to work with OvnEip resources.
type OvnEipInterface interface {
	Create(ctx context.Context, ovnEip *kubeovnv1.OvnEip, opts metav1.CreateOptions) (*kubeovnv1.OvnEip, error)
	Update(ctx context.Context, ovnEip *kubeovnv1.OvnEip, opts metav1.UpdateOptions) (*kubeovnv1.OvnEip, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, ovnEip *kubeovnv1.OvnEip, opts metav1.UpdateOptions) (*kubeovnv1.OvnEip, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubeovnv1.OvnEip, error)
	List(ctx context.Context, opts metav1.ListOptions) (*kubeovnv1.OvnEipList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *kubeovnv1.OvnEip, err error)
	Apply(ctx context.Context, ovnEip *applyconfigurationkubeovnv1.OvnEipApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.OvnEip, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, ovnEip *applyconfigurationkubeovnv1.OvnEipApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.OvnEip, err error)
	OvnEipExpansion
}

// ovnEips implements OvnEipInterface
type ovnEips struct {
	*gentype.ClientWithListAndApply[*kubeovnv1.OvnEip, *kubeovnv1.OvnEipList, *applyconfigurationkubeovnv1.OvnEipApplyConfiguration]
}

// newOvnEips returns a OvnEips
func newOvnEips(c *KubeovnV1Client) *ovnEips {
	return &ovnEips{
		gentype.NewClientWithListAndApply[*kubeovnv1.OvnEip, *kubeovnv1.OvnEipList, *applyconfigurationkubeovnv1.OvnEipApplyConfiguration](
			"ovn-eips",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *kubeovnv1.OvnEip { return &kubeovnv1.OvnEip{} },
			func() *kubeovnv1.OvnEipList { return &kubeovnv1.OvnEipList{} }),
	}
}
