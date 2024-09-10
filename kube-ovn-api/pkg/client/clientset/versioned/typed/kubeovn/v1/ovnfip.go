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

// OvnFipsGetter has a method to return a OvnFipInterface.
// A group's client should implement this interface.
type OvnFipsGetter interface {
	OvnFips() OvnFipInterface
}

// OvnFipInterface has methods to work with OvnFip resources.
type OvnFipInterface interface {
	Create(ctx context.Context, ovnFip *kubeovnv1.OvnFip, opts metav1.CreateOptions) (*kubeovnv1.OvnFip, error)
	Update(ctx context.Context, ovnFip *kubeovnv1.OvnFip, opts metav1.UpdateOptions) (*kubeovnv1.OvnFip, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, ovnFip *kubeovnv1.OvnFip, opts metav1.UpdateOptions) (*kubeovnv1.OvnFip, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubeovnv1.OvnFip, error)
	List(ctx context.Context, opts metav1.ListOptions) (*kubeovnv1.OvnFipList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *kubeovnv1.OvnFip, err error)
	Apply(ctx context.Context, ovnFip *applyconfigurationkubeovnv1.OvnFipApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.OvnFip, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, ovnFip *applyconfigurationkubeovnv1.OvnFipApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.OvnFip, err error)
	OvnFipExpansion
}

// ovnFips implements OvnFipInterface
type ovnFips struct {
	*gentype.ClientWithListAndApply[*kubeovnv1.OvnFip, *kubeovnv1.OvnFipList, *applyconfigurationkubeovnv1.OvnFipApplyConfiguration]
}

// newOvnFips returns a OvnFips
func newOvnFips(c *KubeovnV1Client) *ovnFips {
	return &ovnFips{
		gentype.NewClientWithListAndApply[*kubeovnv1.OvnFip, *kubeovnv1.OvnFipList, *applyconfigurationkubeovnv1.OvnFipApplyConfiguration](
			"ovn-fips",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *kubeovnv1.OvnFip { return &kubeovnv1.OvnFip{} },
			func() *kubeovnv1.OvnFipList { return &kubeovnv1.OvnFipList{} }),
	}
}
