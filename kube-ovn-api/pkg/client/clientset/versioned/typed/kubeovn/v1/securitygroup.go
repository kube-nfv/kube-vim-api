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

// SecurityGroupsGetter has a method to return a SecurityGroupInterface.
// A group's client should implement this interface.
type SecurityGroupsGetter interface {
	SecurityGroups() SecurityGroupInterface
}

// SecurityGroupInterface has methods to work with SecurityGroup resources.
type SecurityGroupInterface interface {
	Create(ctx context.Context, securityGroup *kubeovnv1.SecurityGroup, opts metav1.CreateOptions) (*kubeovnv1.SecurityGroup, error)
	Update(ctx context.Context, securityGroup *kubeovnv1.SecurityGroup, opts metav1.UpdateOptions) (*kubeovnv1.SecurityGroup, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, securityGroup *kubeovnv1.SecurityGroup, opts metav1.UpdateOptions) (*kubeovnv1.SecurityGroup, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubeovnv1.SecurityGroup, error)
	List(ctx context.Context, opts metav1.ListOptions) (*kubeovnv1.SecurityGroupList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *kubeovnv1.SecurityGroup, err error)
	Apply(ctx context.Context, securityGroup *applyconfigurationkubeovnv1.SecurityGroupApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.SecurityGroup, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, securityGroup *applyconfigurationkubeovnv1.SecurityGroupApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.SecurityGroup, err error)
	SecurityGroupExpansion
}

// securityGroups implements SecurityGroupInterface
type securityGroups struct {
	*gentype.ClientWithListAndApply[*kubeovnv1.SecurityGroup, *kubeovnv1.SecurityGroupList, *applyconfigurationkubeovnv1.SecurityGroupApplyConfiguration]
}

// newSecurityGroups returns a SecurityGroups
func newSecurityGroups(c *KubeovnV1Client) *securityGroups {
	return &securityGroups{
		gentype.NewClientWithListAndApply[*kubeovnv1.SecurityGroup, *kubeovnv1.SecurityGroupList, *applyconfigurationkubeovnv1.SecurityGroupApplyConfiguration](
			"security-groups",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *kubeovnv1.SecurityGroup { return &kubeovnv1.SecurityGroup{} },
			func() *kubeovnv1.SecurityGroupList { return &kubeovnv1.SecurityGroupList{} }),
	}
}
