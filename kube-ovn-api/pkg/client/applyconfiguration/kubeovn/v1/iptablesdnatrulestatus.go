// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// IptablesDnatRuleStatusApplyConfiguration represents a declarative configuration of the IptablesDnatRuleStatus type for use
// with apply.
type IptablesDnatRuleStatusApplyConfiguration struct {
	Ready        *bool                                         `json:"ready,omitempty"`
	V4ip         *string                                       `json:"v4ip,omitempty"`
	V6ip         *string                                       `json:"v6ip,omitempty"`
	NatGwDp      *string                                       `json:"natGwDp,omitempty"`
	Redo         *string                                       `json:"redo,omitempty"`
	Protocol     *string                                       `json:"protocol,omitempty"`
	InternalIP   *string                                       `json:"internalIp,omitempty"`
	InternalPort *string                                       `json:"internalPort,omitempty"`
	ExternalPort *string                                       `json:"externalPort,omitempty"`
	Conditions   []IptablesDnatRuleConditionApplyConfiguration `json:"conditions,omitempty"`
}

// IptablesDnatRuleStatusApplyConfiguration constructs a declarative configuration of the IptablesDnatRuleStatus type for use with
// apply.
func IptablesDnatRuleStatus() *IptablesDnatRuleStatusApplyConfiguration {
	return &IptablesDnatRuleStatusApplyConfiguration{}
}

// WithReady sets the Ready field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Ready field is set to the value of the last call.
func (b *IptablesDnatRuleStatusApplyConfiguration) WithReady(value bool) *IptablesDnatRuleStatusApplyConfiguration {
	b.Ready = &value
	return b
}

// WithV4ip sets the V4ip field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the V4ip field is set to the value of the last call.
func (b *IptablesDnatRuleStatusApplyConfiguration) WithV4ip(value string) *IptablesDnatRuleStatusApplyConfiguration {
	b.V4ip = &value
	return b
}

// WithV6ip sets the V6ip field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the V6ip field is set to the value of the last call.
func (b *IptablesDnatRuleStatusApplyConfiguration) WithV6ip(value string) *IptablesDnatRuleStatusApplyConfiguration {
	b.V6ip = &value
	return b
}

// WithNatGwDp sets the NatGwDp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NatGwDp field is set to the value of the last call.
func (b *IptablesDnatRuleStatusApplyConfiguration) WithNatGwDp(value string) *IptablesDnatRuleStatusApplyConfiguration {
	b.NatGwDp = &value
	return b
}

// WithRedo sets the Redo field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Redo field is set to the value of the last call.
func (b *IptablesDnatRuleStatusApplyConfiguration) WithRedo(value string) *IptablesDnatRuleStatusApplyConfiguration {
	b.Redo = &value
	return b
}

// WithProtocol sets the Protocol field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Protocol field is set to the value of the last call.
func (b *IptablesDnatRuleStatusApplyConfiguration) WithProtocol(value string) *IptablesDnatRuleStatusApplyConfiguration {
	b.Protocol = &value
	return b
}

// WithInternalIP sets the InternalIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InternalIP field is set to the value of the last call.
func (b *IptablesDnatRuleStatusApplyConfiguration) WithInternalIP(value string) *IptablesDnatRuleStatusApplyConfiguration {
	b.InternalIP = &value
	return b
}

// WithInternalPort sets the InternalPort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InternalPort field is set to the value of the last call.
func (b *IptablesDnatRuleStatusApplyConfiguration) WithInternalPort(value string) *IptablesDnatRuleStatusApplyConfiguration {
	b.InternalPort = &value
	return b
}

// WithExternalPort sets the ExternalPort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExternalPort field is set to the value of the last call.
func (b *IptablesDnatRuleStatusApplyConfiguration) WithExternalPort(value string) *IptablesDnatRuleStatusApplyConfiguration {
	b.ExternalPort = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *IptablesDnatRuleStatusApplyConfiguration) WithConditions(values ...*IptablesDnatRuleConditionApplyConfiguration) *IptablesDnatRuleStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}
