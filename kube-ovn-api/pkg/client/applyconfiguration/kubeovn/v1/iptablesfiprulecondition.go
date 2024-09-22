// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	kubeovnv1 "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IptablesFIPRuleConditionApplyConfiguration represents a declarative configuration of the IptablesFIPRuleCondition type for use
// with apply.
type IptablesFIPRuleConditionApplyConfiguration struct {
	Type               *kubeovnv1.ConditionType `json:"type,omitempty"`
	Status             *corev1.ConditionStatus  `json:"status,omitempty"`
	Reason             *string                  `json:"reason,omitempty"`
	Message            *string                  `json:"message,omitempty"`
	LastUpdateTime     *metav1.Time             `json:"lastUpdateTime,omitempty"`
	LastTransitionTime *metav1.Time             `json:"lastTransitionTime,omitempty"`
}

// IptablesFIPRuleConditionApplyConfiguration constructs a declarative configuration of the IptablesFIPRuleCondition type for use with
// apply.
func IptablesFIPRuleCondition() *IptablesFIPRuleConditionApplyConfiguration {
	return &IptablesFIPRuleConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *IptablesFIPRuleConditionApplyConfiguration) WithType(value kubeovnv1.ConditionType) *IptablesFIPRuleConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *IptablesFIPRuleConditionApplyConfiguration) WithStatus(value corev1.ConditionStatus) *IptablesFIPRuleConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *IptablesFIPRuleConditionApplyConfiguration) WithReason(value string) *IptablesFIPRuleConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *IptablesFIPRuleConditionApplyConfiguration) WithMessage(value string) *IptablesFIPRuleConditionApplyConfiguration {
	b.Message = &value
	return b
}

// WithLastUpdateTime sets the LastUpdateTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastUpdateTime field is set to the value of the last call.
func (b *IptablesFIPRuleConditionApplyConfiguration) WithLastUpdateTime(value metav1.Time) *IptablesFIPRuleConditionApplyConfiguration {
	b.LastUpdateTime = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *IptablesFIPRuleConditionApplyConfiguration) WithLastTransitionTime(value metav1.Time) *IptablesFIPRuleConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}
