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

// QoSPoliciesGetter has a method to return a QoSPolicyInterface.
// A group's client should implement this interface.
type QoSPoliciesGetter interface {
	QoSPolicies() QoSPolicyInterface
}

// QoSPolicyInterface has methods to work with QoSPolicy resources.
type QoSPolicyInterface interface {
	Create(ctx context.Context, qoSPolicy *kubeovnv1.QoSPolicy, opts metav1.CreateOptions) (*kubeovnv1.QoSPolicy, error)
	Update(ctx context.Context, qoSPolicy *kubeovnv1.QoSPolicy, opts metav1.UpdateOptions) (*kubeovnv1.QoSPolicy, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, qoSPolicy *kubeovnv1.QoSPolicy, opts metav1.UpdateOptions) (*kubeovnv1.QoSPolicy, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubeovnv1.QoSPolicy, error)
	List(ctx context.Context, opts metav1.ListOptions) (*kubeovnv1.QoSPolicyList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *kubeovnv1.QoSPolicy, err error)
	Apply(ctx context.Context, qoSPolicy *applyconfigurationkubeovnv1.QoSPolicyApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.QoSPolicy, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, qoSPolicy *applyconfigurationkubeovnv1.QoSPolicyApplyConfiguration, opts metav1.ApplyOptions) (result *kubeovnv1.QoSPolicy, err error)
	QoSPolicyExpansion
}

// qoSPolicies implements QoSPolicyInterface
type qoSPolicies struct {
	*gentype.ClientWithListAndApply[*kubeovnv1.QoSPolicy, *kubeovnv1.QoSPolicyList, *applyconfigurationkubeovnv1.QoSPolicyApplyConfiguration]
}

// newQoSPolicies returns a QoSPolicies
func newQoSPolicies(c *KubeovnV1Client) *qoSPolicies {
	return &qoSPolicies{
		gentype.NewClientWithListAndApply[*kubeovnv1.QoSPolicy, *kubeovnv1.QoSPolicyList, *applyconfigurationkubeovnv1.QoSPolicyApplyConfiguration](
			"qos-policies",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *kubeovnv1.QoSPolicy { return &kubeovnv1.QoSPolicy{} },
			func() *kubeovnv1.QoSPolicyList { return &kubeovnv1.QoSPolicyList{} }),
	}
}
