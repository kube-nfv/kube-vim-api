// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	kubeovnv1 "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
)

// QoSPolicyStatusApplyConfiguration represents a declarative configuration of the QoSPolicyStatus type for use
// with apply.
type QoSPolicyStatusApplyConfiguration struct {
	BandwidthLimitRules *kubeovnv1.QoSPolicyBandwidthLimitRules `json:"bandwidthLimitRules,omitempty"`
	Shared              *bool                                   `json:"shared,omitempty"`
	BindingType         *kubeovnv1.QoSPolicyBindingType         `json:"bindingType,omitempty"`
	Conditions          []QoSPolicyConditionApplyConfiguration  `json:"conditions,omitempty"`
}

// QoSPolicyStatusApplyConfiguration constructs a declarative configuration of the QoSPolicyStatus type for use with
// apply.
func QoSPolicyStatus() *QoSPolicyStatusApplyConfiguration {
	return &QoSPolicyStatusApplyConfiguration{}
}

// WithBandwidthLimitRules sets the BandwidthLimitRules field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BandwidthLimitRules field is set to the value of the last call.
func (b *QoSPolicyStatusApplyConfiguration) WithBandwidthLimitRules(value kubeovnv1.QoSPolicyBandwidthLimitRules) *QoSPolicyStatusApplyConfiguration {
	b.BandwidthLimitRules = &value
	return b
}

// WithShared sets the Shared field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Shared field is set to the value of the last call.
func (b *QoSPolicyStatusApplyConfiguration) WithShared(value bool) *QoSPolicyStatusApplyConfiguration {
	b.Shared = &value
	return b
}

// WithBindingType sets the BindingType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BindingType field is set to the value of the last call.
func (b *QoSPolicyStatusApplyConfiguration) WithBindingType(value kubeovnv1.QoSPolicyBindingType) *QoSPolicyStatusApplyConfiguration {
	b.BindingType = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *QoSPolicyStatusApplyConfiguration) WithConditions(values ...*QoSPolicyConditionApplyConfiguration) *QoSPolicyStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}
