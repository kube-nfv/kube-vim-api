// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	kubeovnv1 "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
)

// PolicyRouteApplyConfiguration represents a declarative configuration of the PolicyRoute type for use
// with apply.
type PolicyRouteApplyConfiguration struct {
	Priority  *int                         `json:"priority,omitempty"`
	Match     *string                      `json:"match,omitempty"`
	Action    *kubeovnv1.PolicyRouteAction `json:"action,omitempty"`
	NextHopIP *string                      `json:"nextHopIP,omitempty"`
}

// PolicyRouteApplyConfiguration constructs a declarative configuration of the PolicyRoute type for use with
// apply.
func PolicyRoute() *PolicyRouteApplyConfiguration {
	return &PolicyRouteApplyConfiguration{}
}

// WithPriority sets the Priority field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Priority field is set to the value of the last call.
func (b *PolicyRouteApplyConfiguration) WithPriority(value int) *PolicyRouteApplyConfiguration {
	b.Priority = &value
	return b
}

// WithMatch sets the Match field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Match field is set to the value of the last call.
func (b *PolicyRouteApplyConfiguration) WithMatch(value string) *PolicyRouteApplyConfiguration {
	b.Match = &value
	return b
}

// WithAction sets the Action field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Action field is set to the value of the last call.
func (b *PolicyRouteApplyConfiguration) WithAction(value kubeovnv1.PolicyRouteAction) *PolicyRouteApplyConfiguration {
	b.Action = &value
	return b
}

// WithNextHopIP sets the NextHopIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NextHopIP field is set to the value of the last call.
func (b *PolicyRouteApplyConfiguration) WithNextHopIP(value string) *PolicyRouteApplyConfiguration {
	b.NextHopIP = &value
	return b
}
