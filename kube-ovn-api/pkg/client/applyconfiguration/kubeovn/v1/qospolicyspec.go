// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	kubeovnv1 "github.com/DiMalovanyy/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
)

// QoSPolicySpecApplyConfiguration represents a declarative configuration of the QoSPolicySpec type for use
// with apply.
type QoSPolicySpecApplyConfiguration struct {
	BandwidthLimitRules *kubeovnv1.QoSPolicyBandwidthLimitRules `json:"bandwidthLimitRules,omitempty"`
	Shared              *bool                                   `json:"shared,omitempty"`
	BindingType         *kubeovnv1.QoSPolicyBindingType         `json:"bindingType,omitempty"`
}

// QoSPolicySpecApplyConfiguration constructs a declarative configuration of the QoSPolicySpec type for use with
// apply.
func QoSPolicySpec() *QoSPolicySpecApplyConfiguration {
	return &QoSPolicySpecApplyConfiguration{}
}

// WithBandwidthLimitRules sets the BandwidthLimitRules field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BandwidthLimitRules field is set to the value of the last call.
func (b *QoSPolicySpecApplyConfiguration) WithBandwidthLimitRules(value kubeovnv1.QoSPolicyBandwidthLimitRules) *QoSPolicySpecApplyConfiguration {
	b.BandwidthLimitRules = &value
	return b
}

// WithShared sets the Shared field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Shared field is set to the value of the last call.
func (b *QoSPolicySpecApplyConfiguration) WithShared(value bool) *QoSPolicySpecApplyConfiguration {
	b.Shared = &value
	return b
}

// WithBindingType sets the BindingType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BindingType field is set to the value of the last call.
func (b *QoSPolicySpecApplyConfiguration) WithBindingType(value kubeovnv1.QoSPolicyBindingType) *QoSPolicySpecApplyConfiguration {
	b.BindingType = &value
	return b
}