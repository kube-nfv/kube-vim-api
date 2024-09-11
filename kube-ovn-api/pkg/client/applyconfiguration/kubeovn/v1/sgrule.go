// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	kubeovnv1 "github.com/DiMalovanyy/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
)

// SgRuleApplyConfiguration represents a declarative configuration of the SgRule type for use
// with apply.
type SgRuleApplyConfiguration struct {
	IPVersion           *string                 `json:"ipVersion,omitempty"`
	Protocol            *kubeovnv1.SgProtocol   `json:"protocol,omitempty"`
	Priority            *int                    `json:"priority,omitempty"`
	RemoteType          *kubeovnv1.SgRemoteType `json:"remoteType,omitempty"`
	RemoteAddress       *string                 `json:"remoteAddress,omitempty"`
	RemoteSecurityGroup *string                 `json:"remoteSecurityGroup,omitempty"`
	PortRangeMin        *int                    `json:"portRangeMin,omitempty"`
	PortRangeMax        *int                    `json:"portRangeMax,omitempty"`
	Policy              *kubeovnv1.SgPolicy     `json:"policy,omitempty"`
}

// SgRuleApplyConfiguration constructs a declarative configuration of the SgRule type for use with
// apply.
func SgRule() *SgRuleApplyConfiguration {
	return &SgRuleApplyConfiguration{}
}

// WithIPVersion sets the IPVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IPVersion field is set to the value of the last call.
func (b *SgRuleApplyConfiguration) WithIPVersion(value string) *SgRuleApplyConfiguration {
	b.IPVersion = &value
	return b
}

// WithProtocol sets the Protocol field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Protocol field is set to the value of the last call.
func (b *SgRuleApplyConfiguration) WithProtocol(value kubeovnv1.SgProtocol) *SgRuleApplyConfiguration {
	b.Protocol = &value
	return b
}

// WithPriority sets the Priority field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Priority field is set to the value of the last call.
func (b *SgRuleApplyConfiguration) WithPriority(value int) *SgRuleApplyConfiguration {
	b.Priority = &value
	return b
}

// WithRemoteType sets the RemoteType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RemoteType field is set to the value of the last call.
func (b *SgRuleApplyConfiguration) WithRemoteType(value kubeovnv1.SgRemoteType) *SgRuleApplyConfiguration {
	b.RemoteType = &value
	return b
}

// WithRemoteAddress sets the RemoteAddress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RemoteAddress field is set to the value of the last call.
func (b *SgRuleApplyConfiguration) WithRemoteAddress(value string) *SgRuleApplyConfiguration {
	b.RemoteAddress = &value
	return b
}

// WithRemoteSecurityGroup sets the RemoteSecurityGroup field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RemoteSecurityGroup field is set to the value of the last call.
func (b *SgRuleApplyConfiguration) WithRemoteSecurityGroup(value string) *SgRuleApplyConfiguration {
	b.RemoteSecurityGroup = &value
	return b
}

// WithPortRangeMin sets the PortRangeMin field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PortRangeMin field is set to the value of the last call.
func (b *SgRuleApplyConfiguration) WithPortRangeMin(value int) *SgRuleApplyConfiguration {
	b.PortRangeMin = &value
	return b
}

// WithPortRangeMax sets the PortRangeMax field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PortRangeMax field is set to the value of the last call.
func (b *SgRuleApplyConfiguration) WithPortRangeMax(value int) *SgRuleApplyConfiguration {
	b.PortRangeMax = &value
	return b
}

// WithPolicy sets the Policy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Policy field is set to the value of the last call.
func (b *SgRuleApplyConfiguration) WithPolicy(value kubeovnv1.SgPolicy) *SgRuleApplyConfiguration {
	b.Policy = &value
	return b
}