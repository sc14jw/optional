package optional

import "github.com/cheekybits/genny/generic"

// Type provides placeholder for an TypeOptional's type. Altered through code generation.
type Type generic.Type

// TypeOptional provides information for an TypeOptional type such as whether or not a given value was initialised given a some constraits.
type TypeOptional struct {
	value       *Type
	initialised bool
}

// WasInitialised returns whether or not a given TypeOptional contains a value that was initialised given the TypeOptional's constraits.
func (opt *TypeOptional) WasInitialised() bool {
	return opt.initialised
}

// GetValue returns the value of given TypeOptional as a new value.
func (opt *TypeOptional) GetValue() *Type {
	return opt.value
}

// NotNilType returns an TypeOptional stating the given value should not be nil.
func NotNilType(value *Type) *TypeOptional {
	if value == nil {
		return &TypeOptional{value: nil, initialised: false}
	}

	return &TypeOptional{value: value, initialised: true}
}

// WithDefaultTypeValue alter a created TypeOptional to have a default value to be set should it not be initialised
func WithDefaultTypeValue(opt *TypeOptional, def *Type) *TypeOptional {
	if !opt.WasInitialised() {
		opt.value = def
		return opt
	}

	return opt
}
