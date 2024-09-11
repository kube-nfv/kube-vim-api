// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// OvnDnatRuleStatusApplyConfiguration represents a declarative configuration of the OvnDnatRuleStatus type for use
// with apply.
type OvnDnatRuleStatusApplyConfiguration struct {
	Vpc          *string                                  `json:"vpc,omitempty"`
	V4Eip        *string                                  `json:"v4Eip,omitempty"`
	V6Eip        *string                                  `json:"v6Eip,omitempty"`
	ExternalPort *string                                  `json:"externalPort,omitempty"`
	V4Ip         *string                                  `json:"v4Ip,omitempty"`
	V6Ip         *string                                  `json:"v6Ip,omitempty"`
	InternalPort *string                                  `json:"internalPort,omitempty"`
	Protocol     *string                                  `json:"protocol,omitempty"`
	IPName       *string                                  `json:"ipName,omitempty"`
	Ready        *bool                                    `json:"ready,omitempty"`
	Conditions   []OvnDnatRuleConditionApplyConfiguration `json:"conditions,omitempty"`
}

// OvnDnatRuleStatusApplyConfiguration constructs a declarative configuration of the OvnDnatRuleStatus type for use with
// apply.
func OvnDnatRuleStatus() *OvnDnatRuleStatusApplyConfiguration {
	return &OvnDnatRuleStatusApplyConfiguration{}
}

// WithVpc sets the Vpc field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Vpc field is set to the value of the last call.
func (b *OvnDnatRuleStatusApplyConfiguration) WithVpc(value string) *OvnDnatRuleStatusApplyConfiguration {
	b.Vpc = &value
	return b
}

// WithV4Eip sets the V4Eip field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the V4Eip field is set to the value of the last call.
func (b *OvnDnatRuleStatusApplyConfiguration) WithV4Eip(value string) *OvnDnatRuleStatusApplyConfiguration {
	b.V4Eip = &value
	return b
}

// WithV6Eip sets the V6Eip field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the V6Eip field is set to the value of the last call.
func (b *OvnDnatRuleStatusApplyConfiguration) WithV6Eip(value string) *OvnDnatRuleStatusApplyConfiguration {
	b.V6Eip = &value
	return b
}

// WithExternalPort sets the ExternalPort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExternalPort field is set to the value of the last call.
func (b *OvnDnatRuleStatusApplyConfiguration) WithExternalPort(value string) *OvnDnatRuleStatusApplyConfiguration {
	b.ExternalPort = &value
	return b
}

// WithV4Ip sets the V4Ip field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the V4Ip field is set to the value of the last call.
func (b *OvnDnatRuleStatusApplyConfiguration) WithV4Ip(value string) *OvnDnatRuleStatusApplyConfiguration {
	b.V4Ip = &value
	return b
}

// WithV6Ip sets the V6Ip field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the V6Ip field is set to the value of the last call.
func (b *OvnDnatRuleStatusApplyConfiguration) WithV6Ip(value string) *OvnDnatRuleStatusApplyConfiguration {
	b.V6Ip = &value
	return b
}

// WithInternalPort sets the InternalPort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InternalPort field is set to the value of the last call.
func (b *OvnDnatRuleStatusApplyConfiguration) WithInternalPort(value string) *OvnDnatRuleStatusApplyConfiguration {
	b.InternalPort = &value
	return b
}

// WithProtocol sets the Protocol field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Protocol field is set to the value of the last call.
func (b *OvnDnatRuleStatusApplyConfiguration) WithProtocol(value string) *OvnDnatRuleStatusApplyConfiguration {
	b.Protocol = &value
	return b
}

// WithIPName sets the IPName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IPName field is set to the value of the last call.
func (b *OvnDnatRuleStatusApplyConfiguration) WithIPName(value string) *OvnDnatRuleStatusApplyConfiguration {
	b.IPName = &value
	return b
}

// WithReady sets the Ready field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Ready field is set to the value of the last call.
func (b *OvnDnatRuleStatusApplyConfiguration) WithReady(value bool) *OvnDnatRuleStatusApplyConfiguration {
	b.Ready = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *OvnDnatRuleStatusApplyConfiguration) WithConditions(values ...*OvnDnatRuleConditionApplyConfiguration) *OvnDnatRuleStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}