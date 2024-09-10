// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	kubeovnv1 "github.com/DiMalovanyy/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OvnDnatRuleConditionApplyConfiguration represents a declarative configuration of the OvnDnatRuleCondition type for use
// with apply.
type OvnDnatRuleConditionApplyConfiguration struct {
	Type               *kubeovnv1.ConditionType `json:"type,omitempty"`
	Status             *corev1.ConditionStatus  `json:"status,omitempty"`
	Reason             *string                  `json:"reason,omitempty"`
	Message            *string                  `json:"message,omitempty"`
	LastUpdateTime     *metav1.Time             `json:"lastUpdateTime,omitempty"`
	LastTransitionTime *metav1.Time             `json:"lastTransitionTime,omitempty"`
}

// OvnDnatRuleConditionApplyConfiguration constructs a declarative configuration of the OvnDnatRuleCondition type for use with
// apply.
func OvnDnatRuleCondition() *OvnDnatRuleConditionApplyConfiguration {
	return &OvnDnatRuleConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *OvnDnatRuleConditionApplyConfiguration) WithType(value kubeovnv1.ConditionType) *OvnDnatRuleConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *OvnDnatRuleConditionApplyConfiguration) WithStatus(value corev1.ConditionStatus) *OvnDnatRuleConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *OvnDnatRuleConditionApplyConfiguration) WithReason(value string) *OvnDnatRuleConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *OvnDnatRuleConditionApplyConfiguration) WithMessage(value string) *OvnDnatRuleConditionApplyConfiguration {
	b.Message = &value
	return b
}

// WithLastUpdateTime sets the LastUpdateTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastUpdateTime field is set to the value of the last call.
func (b *OvnDnatRuleConditionApplyConfiguration) WithLastUpdateTime(value metav1.Time) *OvnDnatRuleConditionApplyConfiguration {
	b.LastUpdateTime = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *OvnDnatRuleConditionApplyConfiguration) WithLastTransitionTime(value metav1.Time) *OvnDnatRuleConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}
