// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// VipStatusApplyConfiguration represents a declarative configuration of the VipStatus type for use
// with apply.
type VipStatusApplyConfiguration struct {
	Conditions []VipConditionApplyConfiguration `json:"conditions,omitempty"`
	Ready      *bool                            `json:"ready,omitempty"`
	Type       *string                          `json:"type,omitempty"`
	V4ip       *string                          `json:"v4ip,omitempty"`
	V6ip       *string                          `json:"v6ip,omitempty"`
	Mac        *string                          `json:"mac,omitempty"`
	Pv4ip      *string                          `json:"pv4ip,omitempty"`
	Pv6ip      *string                          `json:"pv6ip,omitempty"`
	Pmac       *string                          `json:"pmac,omitempty"`
}

// VipStatusApplyConfiguration constructs a declarative configuration of the VipStatus type for use with
// apply.
func VipStatus() *VipStatusApplyConfiguration {
	return &VipStatusApplyConfiguration{}
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *VipStatusApplyConfiguration) WithConditions(values ...*VipConditionApplyConfiguration) *VipStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithReady sets the Ready field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Ready field is set to the value of the last call.
func (b *VipStatusApplyConfiguration) WithReady(value bool) *VipStatusApplyConfiguration {
	b.Ready = &value
	return b
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *VipStatusApplyConfiguration) WithType(value string) *VipStatusApplyConfiguration {
	b.Type = &value
	return b
}

// WithV4ip sets the V4ip field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the V4ip field is set to the value of the last call.
func (b *VipStatusApplyConfiguration) WithV4ip(value string) *VipStatusApplyConfiguration {
	b.V4ip = &value
	return b
}

// WithV6ip sets the V6ip field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the V6ip field is set to the value of the last call.
func (b *VipStatusApplyConfiguration) WithV6ip(value string) *VipStatusApplyConfiguration {
	b.V6ip = &value
	return b
}

// WithMac sets the Mac field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Mac field is set to the value of the last call.
func (b *VipStatusApplyConfiguration) WithMac(value string) *VipStatusApplyConfiguration {
	b.Mac = &value
	return b
}

// WithPv4ip sets the Pv4ip field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Pv4ip field is set to the value of the last call.
func (b *VipStatusApplyConfiguration) WithPv4ip(value string) *VipStatusApplyConfiguration {
	b.Pv4ip = &value
	return b
}

// WithPv6ip sets the Pv6ip field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Pv6ip field is set to the value of the last call.
func (b *VipStatusApplyConfiguration) WithPv6ip(value string) *VipStatusApplyConfiguration {
	b.Pv6ip = &value
	return b
}

// WithPmac sets the Pmac field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Pmac field is set to the value of the last call.
func (b *VipStatusApplyConfiguration) WithPmac(value string) *VipStatusApplyConfiguration {
	b.Pmac = &value
	return b
}