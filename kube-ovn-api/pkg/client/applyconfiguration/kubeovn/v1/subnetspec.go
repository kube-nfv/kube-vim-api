// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// SubnetSpecApplyConfiguration represents a declarative configuration of the SubnetSpec type for use
// with apply.
type SubnetSpecApplyConfiguration struct {
	Default                *bool                                     `json:"default,omitempty"`
	Vpc                    *string                                   `json:"vpc,omitempty"`
	Protocol               *string                                   `json:"protocol,omitempty"`
	Namespaces             []string                                  `json:"namespaces,omitempty"`
	CIDRBlock              *string                                   `json:"cidrBlock,omitempty"`
	Gateway                *string                                   `json:"gateway,omitempty"`
	ExcludeIps             []string                                  `json:"excludeIps,omitempty"`
	Provider               *string                                   `json:"provider,omitempty"`
	GatewayType            *string                                   `json:"gatewayType,omitempty"`
	GatewayNode            *string                                   `json:"gatewayNode,omitempty"`
	NatOutgoing            *bool                                     `json:"natOutgoing,omitempty"`
	ExternalEgressGateway  *string                                   `json:"externalEgressGateway,omitempty"`
	PolicyRoutingPriority  *uint32                                   `json:"policyRoutingPriority,omitempty"`
	PolicyRoutingTableID   *uint32                                   `json:"policyRoutingTableID,omitempty"`
	Mtu                    *uint32                                   `json:"mtu,omitempty"`
	Private                *bool                                     `json:"private,omitempty"`
	AllowSubnets           []string                                  `json:"allowSubnets,omitempty"`
	Vlan                   *string                                   `json:"vlan,omitempty"`
	Vips                   []string                                  `json:"vips,omitempty"`
	LogicalGateway         *bool                                     `json:"logicalGateway,omitempty"`
	DisableGatewayCheck    *bool                                     `json:"disableGatewayCheck,omitempty"`
	DisableInterConnection *bool                                     `json:"disableInterConnection,omitempty"`
	EnableDHCP             *bool                                     `json:"enableDHCP,omitempty"`
	DHCPv4Options          *string                                   `json:"dhcpV4Options,omitempty"`
	DHCPv6Options          *string                                   `json:"dhcpV6Options,omitempty"`
	EnableIPv6RA           *bool                                     `json:"enableIPv6RA,omitempty"`
	IPv6RAConfigs          *string                                   `json:"ipv6RAConfigs,omitempty"`
	Acls                   []ACLApplyConfiguration                   `json:"acls,omitempty"`
	AllowEWTraffic         *bool                                     `json:"allowEWTraffic,omitempty"`
	NatOutgoingPolicyRules []NatOutgoingPolicyRuleApplyConfiguration `json:"natOutgoingPolicyRules,omitempty"`
	U2OInterconnectionIP   *string                                   `json:"u2oInterconnectionIP,omitempty"`
	U2OInterconnection     *bool                                     `json:"u2oInterconnection,omitempty"`
	EnableLb               *bool                                     `json:"enableLb,omitempty"`
	EnableEcmp             *bool                                     `json:"enableEcmp,omitempty"`
	EnableMulticastSnoop   *bool                                     `json:"enableMulticastSnoop,omitempty"`
	RouteTable             *string                                   `json:"routeTable,omitempty"`
}

// SubnetSpecApplyConfiguration constructs a declarative configuration of the SubnetSpec type for use with
// apply.
func SubnetSpec() *SubnetSpecApplyConfiguration {
	return &SubnetSpecApplyConfiguration{}
}

// WithDefault sets the Default field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Default field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithDefault(value bool) *SubnetSpecApplyConfiguration {
	b.Default = &value
	return b
}

// WithVpc sets the Vpc field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Vpc field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithVpc(value string) *SubnetSpecApplyConfiguration {
	b.Vpc = &value
	return b
}

// WithProtocol sets the Protocol field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Protocol field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithProtocol(value string) *SubnetSpecApplyConfiguration {
	b.Protocol = &value
	return b
}

// WithNamespaces adds the given value to the Namespaces field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Namespaces field.
func (b *SubnetSpecApplyConfiguration) WithNamespaces(values ...string) *SubnetSpecApplyConfiguration {
	for i := range values {
		b.Namespaces = append(b.Namespaces, values[i])
	}
	return b
}

// WithCIDRBlock sets the CIDRBlock field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CIDRBlock field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithCIDRBlock(value string) *SubnetSpecApplyConfiguration {
	b.CIDRBlock = &value
	return b
}

// WithGateway sets the Gateway field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Gateway field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithGateway(value string) *SubnetSpecApplyConfiguration {
	b.Gateway = &value
	return b
}

// WithExcludeIps adds the given value to the ExcludeIps field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ExcludeIps field.
func (b *SubnetSpecApplyConfiguration) WithExcludeIps(values ...string) *SubnetSpecApplyConfiguration {
	for i := range values {
		b.ExcludeIps = append(b.ExcludeIps, values[i])
	}
	return b
}

// WithProvider sets the Provider field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Provider field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithProvider(value string) *SubnetSpecApplyConfiguration {
	b.Provider = &value
	return b
}

// WithGatewayType sets the GatewayType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GatewayType field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithGatewayType(value string) *SubnetSpecApplyConfiguration {
	b.GatewayType = &value
	return b
}

// WithGatewayNode sets the GatewayNode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GatewayNode field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithGatewayNode(value string) *SubnetSpecApplyConfiguration {
	b.GatewayNode = &value
	return b
}

// WithNatOutgoing sets the NatOutgoing field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NatOutgoing field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithNatOutgoing(value bool) *SubnetSpecApplyConfiguration {
	b.NatOutgoing = &value
	return b
}

// WithExternalEgressGateway sets the ExternalEgressGateway field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExternalEgressGateway field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithExternalEgressGateway(value string) *SubnetSpecApplyConfiguration {
	b.ExternalEgressGateway = &value
	return b
}

// WithPolicyRoutingPriority sets the PolicyRoutingPriority field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PolicyRoutingPriority field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithPolicyRoutingPriority(value uint32) *SubnetSpecApplyConfiguration {
	b.PolicyRoutingPriority = &value
	return b
}

// WithPolicyRoutingTableID sets the PolicyRoutingTableID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PolicyRoutingTableID field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithPolicyRoutingTableID(value uint32) *SubnetSpecApplyConfiguration {
	b.PolicyRoutingTableID = &value
	return b
}

// WithMtu sets the Mtu field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Mtu field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithMtu(value uint32) *SubnetSpecApplyConfiguration {
	b.Mtu = &value
	return b
}

// WithPrivate sets the Private field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Private field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithPrivate(value bool) *SubnetSpecApplyConfiguration {
	b.Private = &value
	return b
}

// WithAllowSubnets adds the given value to the AllowSubnets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AllowSubnets field.
func (b *SubnetSpecApplyConfiguration) WithAllowSubnets(values ...string) *SubnetSpecApplyConfiguration {
	for i := range values {
		b.AllowSubnets = append(b.AllowSubnets, values[i])
	}
	return b
}

// WithVlan sets the Vlan field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Vlan field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithVlan(value string) *SubnetSpecApplyConfiguration {
	b.Vlan = &value
	return b
}

// WithVips adds the given value to the Vips field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Vips field.
func (b *SubnetSpecApplyConfiguration) WithVips(values ...string) *SubnetSpecApplyConfiguration {
	for i := range values {
		b.Vips = append(b.Vips, values[i])
	}
	return b
}

// WithLogicalGateway sets the LogicalGateway field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LogicalGateway field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithLogicalGateway(value bool) *SubnetSpecApplyConfiguration {
	b.LogicalGateway = &value
	return b
}

// WithDisableGatewayCheck sets the DisableGatewayCheck field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisableGatewayCheck field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithDisableGatewayCheck(value bool) *SubnetSpecApplyConfiguration {
	b.DisableGatewayCheck = &value
	return b
}

// WithDisableInterConnection sets the DisableInterConnection field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisableInterConnection field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithDisableInterConnection(value bool) *SubnetSpecApplyConfiguration {
	b.DisableInterConnection = &value
	return b
}

// WithEnableDHCP sets the EnableDHCP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableDHCP field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithEnableDHCP(value bool) *SubnetSpecApplyConfiguration {
	b.EnableDHCP = &value
	return b
}

// WithDHCPv4Options sets the DHCPv4Options field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DHCPv4Options field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithDHCPv4Options(value string) *SubnetSpecApplyConfiguration {
	b.DHCPv4Options = &value
	return b
}

// WithDHCPv6Options sets the DHCPv6Options field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DHCPv6Options field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithDHCPv6Options(value string) *SubnetSpecApplyConfiguration {
	b.DHCPv6Options = &value
	return b
}

// WithEnableIPv6RA sets the EnableIPv6RA field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableIPv6RA field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithEnableIPv6RA(value bool) *SubnetSpecApplyConfiguration {
	b.EnableIPv6RA = &value
	return b
}

// WithIPv6RAConfigs sets the IPv6RAConfigs field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IPv6RAConfigs field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithIPv6RAConfigs(value string) *SubnetSpecApplyConfiguration {
	b.IPv6RAConfigs = &value
	return b
}

// WithAcls adds the given value to the Acls field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Acls field.
func (b *SubnetSpecApplyConfiguration) WithAcls(values ...*ACLApplyConfiguration) *SubnetSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAcls")
		}
		b.Acls = append(b.Acls, *values[i])
	}
	return b
}

// WithAllowEWTraffic sets the AllowEWTraffic field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllowEWTraffic field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithAllowEWTraffic(value bool) *SubnetSpecApplyConfiguration {
	b.AllowEWTraffic = &value
	return b
}

// WithNatOutgoingPolicyRules adds the given value to the NatOutgoingPolicyRules field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NatOutgoingPolicyRules field.
func (b *SubnetSpecApplyConfiguration) WithNatOutgoingPolicyRules(values ...*NatOutgoingPolicyRuleApplyConfiguration) *SubnetSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithNatOutgoingPolicyRules")
		}
		b.NatOutgoingPolicyRules = append(b.NatOutgoingPolicyRules, *values[i])
	}
	return b
}

// WithU2OInterconnectionIP sets the U2OInterconnectionIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the U2OInterconnectionIP field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithU2OInterconnectionIP(value string) *SubnetSpecApplyConfiguration {
	b.U2OInterconnectionIP = &value
	return b
}

// WithU2OInterconnection sets the U2OInterconnection field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the U2OInterconnection field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithU2OInterconnection(value bool) *SubnetSpecApplyConfiguration {
	b.U2OInterconnection = &value
	return b
}

// WithEnableLb sets the EnableLb field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableLb field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithEnableLb(value bool) *SubnetSpecApplyConfiguration {
	b.EnableLb = &value
	return b
}

// WithEnableEcmp sets the EnableEcmp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableEcmp field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithEnableEcmp(value bool) *SubnetSpecApplyConfiguration {
	b.EnableEcmp = &value
	return b
}

// WithEnableMulticastSnoop sets the EnableMulticastSnoop field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableMulticastSnoop field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithEnableMulticastSnoop(value bool) *SubnetSpecApplyConfiguration {
	b.EnableMulticastSnoop = &value
	return b
}

// WithRouteTable sets the RouteTable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RouteTable field is set to the value of the last call.
func (b *SubnetSpecApplyConfiguration) WithRouteTable(value string) *SubnetSpecApplyConfiguration {
	b.RouteTable = &value
	return b
}
