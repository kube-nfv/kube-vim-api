// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ACLApplyConfiguration represents a declarative configuration of the ACL type for use
// with apply.
type ACLApplyConfiguration struct {
	Direction *string `json:"direction,omitempty"`
	Priority  *int    `json:"priority,omitempty"`
	Match     *string `json:"match,omitempty"`
	Action    *string `json:"action,omitempty"`
}

// ACLApplyConfiguration constructs a declarative configuration of the ACL type for use with
// apply.
func ACL() *ACLApplyConfiguration {
	return &ACLApplyConfiguration{}
}

// WithDirection sets the Direction field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Direction field is set to the value of the last call.
func (b *ACLApplyConfiguration) WithDirection(value string) *ACLApplyConfiguration {
	b.Direction = &value
	return b
}

// WithPriority sets the Priority field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Priority field is set to the value of the last call.
func (b *ACLApplyConfiguration) WithPriority(value int) *ACLApplyConfiguration {
	b.Priority = &value
	return b
}

// WithMatch sets the Match field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Match field is set to the value of the last call.
func (b *ACLApplyConfiguration) WithMatch(value string) *ACLApplyConfiguration {
	b.Match = &value
	return b
}

// WithAction sets the Action field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Action field is set to the value of the last call.
func (b *ACLApplyConfiguration) WithAction(value string) *ACLApplyConfiguration {
	b.Action = &value
	return b
}