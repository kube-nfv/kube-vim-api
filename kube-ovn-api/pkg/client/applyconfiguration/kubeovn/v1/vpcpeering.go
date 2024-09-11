// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// VpcPeeringApplyConfiguration represents a declarative configuration of the VpcPeering type for use
// with apply.
type VpcPeeringApplyConfiguration struct {
	RemoteVpc      *string `json:"remoteVpc,omitempty"`
	LocalConnectIP *string `json:"localConnectIP,omitempty"`
}

// VpcPeeringApplyConfiguration constructs a declarative configuration of the VpcPeering type for use with
// apply.
func VpcPeering() *VpcPeeringApplyConfiguration {
	return &VpcPeeringApplyConfiguration{}
}

// WithRemoteVpc sets the RemoteVpc field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RemoteVpc field is set to the value of the last call.
func (b *VpcPeeringApplyConfiguration) WithRemoteVpc(value string) *VpcPeeringApplyConfiguration {
	b.RemoteVpc = &value
	return b
}

// WithLocalConnectIP sets the LocalConnectIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LocalConnectIP field is set to the value of the last call.
func (b *VpcPeeringApplyConfiguration) WithLocalConnectIP(value string) *VpcPeeringApplyConfiguration {
	b.LocalConnectIP = &value
	return b
}