// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	kubeovnv1 "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
)

// StaticRouteApplyConfiguration represents a declarative configuration of the StaticRoute type for use
// with apply.
type StaticRouteApplyConfiguration struct {
	Policy     *kubeovnv1.RoutePolicy `json:"policy,omitempty"`
	CIDR       *string                `json:"cidr,omitempty"`
	NextHopIP  *string                `json:"nextHopIP,omitempty"`
	ECMPMode   *string                `json:"ecmpMode,omitempty"`
	BfdID      *string                `json:"bfdId,omitempty"`
	RouteTable *string                `json:"routeTable,omitempty"`
}

// StaticRouteApplyConfiguration constructs a declarative configuration of the StaticRoute type for use with
// apply.
func StaticRoute() *StaticRouteApplyConfiguration {
	return &StaticRouteApplyConfiguration{}
}

// WithPolicy sets the Policy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Policy field is set to the value of the last call.
func (b *StaticRouteApplyConfiguration) WithPolicy(value kubeovnv1.RoutePolicy) *StaticRouteApplyConfiguration {
	b.Policy = &value
	return b
}

// WithCIDR sets the CIDR field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CIDR field is set to the value of the last call.
func (b *StaticRouteApplyConfiguration) WithCIDR(value string) *StaticRouteApplyConfiguration {
	b.CIDR = &value
	return b
}

// WithNextHopIP sets the NextHopIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NextHopIP field is set to the value of the last call.
func (b *StaticRouteApplyConfiguration) WithNextHopIP(value string) *StaticRouteApplyConfiguration {
	b.NextHopIP = &value
	return b
}

// WithECMPMode sets the ECMPMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ECMPMode field is set to the value of the last call.
func (b *StaticRouteApplyConfiguration) WithECMPMode(value string) *StaticRouteApplyConfiguration {
	b.ECMPMode = &value
	return b
}

// WithBfdID sets the BfdID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BfdID field is set to the value of the last call.
func (b *StaticRouteApplyConfiguration) WithBfdID(value string) *StaticRouteApplyConfiguration {
	b.BfdID = &value
	return b
}

// WithRouteTable sets the RouteTable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RouteTable field is set to the value of the last call.
func (b *StaticRouteApplyConfiguration) WithRouteTable(value string) *StaticRouteApplyConfiguration {
	b.RouteTable = &value
	return b
}
