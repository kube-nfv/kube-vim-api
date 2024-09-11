// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ProviderNetworkStatusApplyConfiguration represents a declarative configuration of the ProviderNetworkStatus type for use
// with apply.
type ProviderNetworkStatusApplyConfiguration struct {
	Ready         *bool                                        `json:"ready,omitempty"`
	ReadyNodes    []string                                     `json:"readyNodes,omitempty"`
	NotReadyNodes []string                                     `json:"notReadyNodes,omitempty"`
	Vlans         []string                                     `json:"vlans,omitempty"`
	Conditions    []ProviderNetworkConditionApplyConfiguration `json:"conditions,omitempty"`
}

// ProviderNetworkStatusApplyConfiguration constructs a declarative configuration of the ProviderNetworkStatus type for use with
// apply.
func ProviderNetworkStatus() *ProviderNetworkStatusApplyConfiguration {
	return &ProviderNetworkStatusApplyConfiguration{}
}

// WithReady sets the Ready field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Ready field is set to the value of the last call.
func (b *ProviderNetworkStatusApplyConfiguration) WithReady(value bool) *ProviderNetworkStatusApplyConfiguration {
	b.Ready = &value
	return b
}

// WithReadyNodes adds the given value to the ReadyNodes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ReadyNodes field.
func (b *ProviderNetworkStatusApplyConfiguration) WithReadyNodes(values ...string) *ProviderNetworkStatusApplyConfiguration {
	for i := range values {
		b.ReadyNodes = append(b.ReadyNodes, values[i])
	}
	return b
}

// WithNotReadyNodes adds the given value to the NotReadyNodes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NotReadyNodes field.
func (b *ProviderNetworkStatusApplyConfiguration) WithNotReadyNodes(values ...string) *ProviderNetworkStatusApplyConfiguration {
	for i := range values {
		b.NotReadyNodes = append(b.NotReadyNodes, values[i])
	}
	return b
}

// WithVlans adds the given value to the Vlans field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Vlans field.
func (b *ProviderNetworkStatusApplyConfiguration) WithVlans(values ...string) *ProviderNetworkStatusApplyConfiguration {
	for i := range values {
		b.Vlans = append(b.Vlans, values[i])
	}
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *ProviderNetworkStatusApplyConfiguration) WithConditions(values ...*ProviderNetworkConditionApplyConfiguration) *ProviderNetworkStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}