// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// VpcDNSSpecApplyConfiguration represents a declarative configuration of the VpcDNSSpec type for use
// with apply.
type VpcDNSSpecApplyConfiguration struct {
	Replicas *int32  `json:"replicas,omitempty"`
	Vpc      *string `json:"vpc,omitempty"`
	Subnet   *string `json:"subnet,omitempty"`
}

// VpcDNSSpecApplyConfiguration constructs a declarative configuration of the VpcDNSSpec type for use with
// apply.
func VpcDNSSpec() *VpcDNSSpecApplyConfiguration {
	return &VpcDNSSpecApplyConfiguration{}
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *VpcDNSSpecApplyConfiguration) WithReplicas(value int32) *VpcDNSSpecApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithVpc sets the Vpc field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Vpc field is set to the value of the last call.
func (b *VpcDNSSpecApplyConfiguration) WithVpc(value string) *VpcDNSSpecApplyConfiguration {
	b.Vpc = &value
	return b
}

// WithSubnet sets the Subnet field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Subnet field is set to the value of the last call.
func (b *VpcDNSSpecApplyConfiguration) WithSubnet(value string) *VpcDNSSpecApplyConfiguration {
	b.Subnet = &value
	return b
}
