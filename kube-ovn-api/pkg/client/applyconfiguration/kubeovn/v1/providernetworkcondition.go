// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ProviderNetworkConditionApplyConfiguration represents a declarative configuration of the ProviderNetworkCondition type for use
// with apply.
type ProviderNetworkConditionApplyConfiguration struct {
	Node *string `json:"node,omitempty"`
}

// ProviderNetworkConditionApplyConfiguration constructs a declarative configuration of the ProviderNetworkCondition type for use with
// apply.
func ProviderNetworkCondition() *ProviderNetworkConditionApplyConfiguration {
	return &ProviderNetworkConditionApplyConfiguration{}
}

// WithNode sets the Node field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Node field is set to the value of the last call.
func (b *ProviderNetworkConditionApplyConfiguration) WithNode(value string) *ProviderNetworkConditionApplyConfiguration {
	b.Node = &value
	return b
}