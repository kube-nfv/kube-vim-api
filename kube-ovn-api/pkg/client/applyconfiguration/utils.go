// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1 "github.com/DiMalovanyy/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
	internal "github.com/DiMalovanyy/kube-vim-api/kube-ovn-api/pkg/client/applyconfiguration/internal"
	kubeovnv1 "github.com/DiMalovanyy/kube-vim-api/kube-ovn-api/pkg/client/applyconfiguration/kubeovn/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=kubeovn.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("ACL"):
		return &kubeovnv1.ACLApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Condition"):
		return &kubeovnv1.ConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomInterface"):
		return &kubeovnv1.CustomInterfaceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IP"):
		return &kubeovnv1.IPApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IPPool"):
		return &kubeovnv1.IPPoolApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IPPoolCondition"):
		return &kubeovnv1.IPPoolConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IPPoolSpec"):
		return &kubeovnv1.IPPoolSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IPPoolStatus"):
		return &kubeovnv1.IPPoolStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IPSpec"):
		return &kubeovnv1.IPSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IptablesDnatRule"):
		return &kubeovnv1.IptablesDnatRuleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IptablesDnatRuleCondition"):
		return &kubeovnv1.IptablesDnatRuleConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IptablesDnatRuleSpec"):
		return &kubeovnv1.IptablesDnatRuleSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IptablesDnatRuleStatus"):
		return &kubeovnv1.IptablesDnatRuleStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IptablesEIP"):
		return &kubeovnv1.IptablesEIPApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IptablesEIPCondition"):
		return &kubeovnv1.IptablesEIPConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IptablesEipSpec"):
		return &kubeovnv1.IptablesEipSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IptablesEipStatus"):
		return &kubeovnv1.IptablesEipStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IptablesFIPRule"):
		return &kubeovnv1.IptablesFIPRuleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IptablesFIPRuleCondition"):
		return &kubeovnv1.IptablesFIPRuleConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IptablesFIPRuleSpec"):
		return &kubeovnv1.IptablesFIPRuleSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IptablesFIPRuleStatus"):
		return &kubeovnv1.IptablesFIPRuleStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IptablesSnatRule"):
		return &kubeovnv1.IptablesSnatRuleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IptablesSnatRuleCondition"):
		return &kubeovnv1.IptablesSnatRuleConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IptablesSnatRuleSpec"):
		return &kubeovnv1.IptablesSnatRuleSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IptablesSnatRuleStatus"):
		return &kubeovnv1.IptablesSnatRuleStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NatOutGoingPolicyMatch"):
		return &kubeovnv1.NatOutGoingPolicyMatchApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NatOutgoingPolicyRule"):
		return &kubeovnv1.NatOutgoingPolicyRuleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NatOutgoingPolicyRuleStatus"):
		return &kubeovnv1.NatOutgoingPolicyRuleStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OvnDnatRule"):
		return &kubeovnv1.OvnDnatRuleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OvnDnatRuleCondition"):
		return &kubeovnv1.OvnDnatRuleConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OvnDnatRuleSpec"):
		return &kubeovnv1.OvnDnatRuleSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OvnDnatRuleStatus"):
		return &kubeovnv1.OvnDnatRuleStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OvnEip"):
		return &kubeovnv1.OvnEipApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OvnEipCondition"):
		return &kubeovnv1.OvnEipConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OvnEipSpec"):
		return &kubeovnv1.OvnEipSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OvnEipStatus"):
		return &kubeovnv1.OvnEipStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OvnFip"):
		return &kubeovnv1.OvnFipApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OvnFipCondition"):
		return &kubeovnv1.OvnFipConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OvnFipSpec"):
		return &kubeovnv1.OvnFipSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OvnFipStatus"):
		return &kubeovnv1.OvnFipStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OvnSnatRule"):
		return &kubeovnv1.OvnSnatRuleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OvnSnatRuleCondition"):
		return &kubeovnv1.OvnSnatRuleConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OvnSnatRuleSpec"):
		return &kubeovnv1.OvnSnatRuleSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OvnSnatRuleStatus"):
		return &kubeovnv1.OvnSnatRuleStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PolicyRoute"):
		return &kubeovnv1.PolicyRouteApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ProviderNetwork"):
		return &kubeovnv1.ProviderNetworkApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ProviderNetworkCondition"):
		return &kubeovnv1.ProviderNetworkConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ProviderNetworkSpec"):
		return &kubeovnv1.ProviderNetworkSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ProviderNetworkStatus"):
		return &kubeovnv1.ProviderNetworkStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("QoSPolicy"):
		return &kubeovnv1.QoSPolicyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("QoSPolicyBandwidthLimitRule"):
		return &kubeovnv1.QoSPolicyBandwidthLimitRuleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("QoSPolicyCondition"):
		return &kubeovnv1.QoSPolicyConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("QoSPolicySpec"):
		return &kubeovnv1.QoSPolicySpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("QoSPolicyStatus"):
		return &kubeovnv1.QoSPolicyStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SecurityGroup"):
		return &kubeovnv1.SecurityGroupApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SecurityGroupSpec"):
		return &kubeovnv1.SecurityGroupSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SecurityGroupStatus"):
		return &kubeovnv1.SecurityGroupStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SgRule"):
		return &kubeovnv1.SgRuleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SlrPort"):
		return &kubeovnv1.SlrPortApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("StaticRoute"):
		return &kubeovnv1.StaticRouteApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Subnet"):
		return &kubeovnv1.SubnetApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SubnetCondition"):
		return &kubeovnv1.SubnetConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SubnetSpec"):
		return &kubeovnv1.SubnetSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SubnetStatus"):
		return &kubeovnv1.SubnetStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SwitchLBRule"):
		return &kubeovnv1.SwitchLBRuleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SwitchLBRuleCondition"):
		return &kubeovnv1.SwitchLBRuleConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SwitchLBRuleSpec"):
		return &kubeovnv1.SwitchLBRuleSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SwitchLBRuleStatus"):
		return &kubeovnv1.SwitchLBRuleStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Vip"):
		return &kubeovnv1.VipApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VipCondition"):
		return &kubeovnv1.VipConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VipSpec"):
		return &kubeovnv1.VipSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VipStatus"):
		return &kubeovnv1.VipStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Vlan"):
		return &kubeovnv1.VlanApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VlanCondition"):
		return &kubeovnv1.VlanConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VlanSpec"):
		return &kubeovnv1.VlanSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VlanStatus"):
		return &kubeovnv1.VlanStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Vpc"):
		return &kubeovnv1.VpcApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VpcBgpSpeaker"):
		return &kubeovnv1.VpcBgpSpeakerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VpcCondition"):
		return &kubeovnv1.VpcConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VpcDns"):
		return &kubeovnv1.VpcDnsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VpcDNSCondition"):
		return &kubeovnv1.VpcDNSConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VpcDNSSpec"):
		return &kubeovnv1.VpcDNSSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VpcDNSStatus"):
		return &kubeovnv1.VpcDNSStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VpcNatGateway"):
		return &kubeovnv1.VpcNatGatewayApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VpcNatSpec"):
		return &kubeovnv1.VpcNatSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VpcNatStatus"):
		return &kubeovnv1.VpcNatStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VpcPeering"):
		return &kubeovnv1.VpcPeeringApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VpcSpec"):
		return &kubeovnv1.VpcSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VpcStatus"):
		return &kubeovnv1.VpcStatusApplyConfiguration{}

	}
	return nil
}

func NewTypeConverter(scheme *runtime.Scheme) *testing.TypeConverter {
	return &testing.TypeConverter{Scheme: scheme, TypeResolver: internal.Parser()}
}