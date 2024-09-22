// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	kubeovnv1 "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OvnSnatRuleConditionApplyConfiguration represents a declarative configuration of the OvnSnatRuleCondition type for use
// with apply.
type OvnSnatRuleConditionApplyConfiguration struct {
	Type               *kubeovnv1.ConditionType `json:"type,omitempty"`
	Status             *corev1.ConditionStatus  `json:"status,omitempty"`
	Reason             *string                  `json:"reason,omitempty"`
	Message            *string                  `json:"message,omitempty"`
	LastUpdateTime     *metav1.Time             `json:"lastUpdateTime,omitempty"`
	LastTransitionTime *metav1.Time             `json:"lastTransitionTime,omitempty"`
}

// OvnSnatRuleConditionApplyConfiguration constructs a declarative configuration of the OvnSnatRuleCondition type for use with
// apply.
func OvnSnatRuleCondition() *OvnSnatRuleConditionApplyConfiguration {
	return &OvnSnatRuleConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *OvnSnatRuleConditionApplyConfiguration) WithType(value kubeovnv1.ConditionType) *OvnSnatRuleConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *OvnSnatRuleConditionApplyConfiguration) WithStatus(value corev1.ConditionStatus) *OvnSnatRuleConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *OvnSnatRuleConditionApplyConfiguration) WithReason(value string) *OvnSnatRuleConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *OvnSnatRuleConditionApplyConfiguration) WithMessage(value string) *OvnSnatRuleConditionApplyConfiguration {
	b.Message = &value
	return b
}

// WithLastUpdateTime sets the LastUpdateTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastUpdateTime field is set to the value of the last call.
func (b *OvnSnatRuleConditionApplyConfiguration) WithLastUpdateTime(value metav1.Time) *OvnSnatRuleConditionApplyConfiguration {
	b.LastUpdateTime = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *OvnSnatRuleConditionApplyConfiguration) WithLastTransitionTime(value metav1.Time) *OvnSnatRuleConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}
