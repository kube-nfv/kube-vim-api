// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// IptablesFIPRuleStatusApplyConfiguration represents a declarative configuration of the IptablesFIPRuleStatus type for use
// with apply.
type IptablesFIPRuleStatusApplyConfiguration struct {
	Ready      *bool                                        `json:"ready,omitempty"`
	V4ip       *string                                      `json:"v4ip,omitempty"`
	V6ip       *string                                      `json:"v6ip,omitempty"`
	NatGwDp    *string                                      `json:"natGwDp,omitempty"`
	Redo       *string                                      `json:"redo,omitempty"`
	InternalIP *string                                      `json:"internalIp,omitempty"`
	Conditions []IptablesFIPRuleConditionApplyConfiguration `json:"conditions,omitempty"`
}

// IptablesFIPRuleStatusApplyConfiguration constructs a declarative configuration of the IptablesFIPRuleStatus type for use with
// apply.
func IptablesFIPRuleStatus() *IptablesFIPRuleStatusApplyConfiguration {
	return &IptablesFIPRuleStatusApplyConfiguration{}
}

// WithReady sets the Ready field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Ready field is set to the value of the last call.
func (b *IptablesFIPRuleStatusApplyConfiguration) WithReady(value bool) *IptablesFIPRuleStatusApplyConfiguration {
	b.Ready = &value
	return b
}

// WithV4ip sets the V4ip field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the V4ip field is set to the value of the last call.
func (b *IptablesFIPRuleStatusApplyConfiguration) WithV4ip(value string) *IptablesFIPRuleStatusApplyConfiguration {
	b.V4ip = &value
	return b
}

// WithV6ip sets the V6ip field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the V6ip field is set to the value of the last call.
func (b *IptablesFIPRuleStatusApplyConfiguration) WithV6ip(value string) *IptablesFIPRuleStatusApplyConfiguration {
	b.V6ip = &value
	return b
}

// WithNatGwDp sets the NatGwDp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NatGwDp field is set to the value of the last call.
func (b *IptablesFIPRuleStatusApplyConfiguration) WithNatGwDp(value string) *IptablesFIPRuleStatusApplyConfiguration {
	b.NatGwDp = &value
	return b
}

// WithRedo sets the Redo field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Redo field is set to the value of the last call.
func (b *IptablesFIPRuleStatusApplyConfiguration) WithRedo(value string) *IptablesFIPRuleStatusApplyConfiguration {
	b.Redo = &value
	return b
}

// WithInternalIP sets the InternalIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InternalIP field is set to the value of the last call.
func (b *IptablesFIPRuleStatusApplyConfiguration) WithInternalIP(value string) *IptablesFIPRuleStatusApplyConfiguration {
	b.InternalIP = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *IptablesFIPRuleStatusApplyConfiguration) WithConditions(values ...*IptablesFIPRuleConditionApplyConfiguration) *IptablesFIPRuleStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}
