package optional

import (
	"fmt"
	"reflect"
)

const (
	// INCORRECTTYPE provides error text should an optional be given the wrong type.
	INCORRECTTYPE = "The given type %v cannot be used with an Optional of type %v"
	// NILVALUE provides error text should an optional be given a nil value.
	NILVALUE = "Given value was nil"
)

// Optional provides mechanism for eligantly dealing with Nil pointer errors in your code using reflection for type safety (Note: this will be altered if/when go implements Generics).
// Through the Optional struct it is possible to fail fast should a given value be nil or provide a default value for Nil values.
// From the Optional struct it is also possible to query whether or not a value has been instanciated and get the value itself.
//
// Author: Jack Wainwright - 20-05-2017
type Optional struct {
	v    *interface{}
	t    *reflect.Type
	init bool
}

// WasInitialized returns whether or not a given Optional contains a value that was Initialized given the Optional's constraints.
func (opt *Optional) WasInitialized() bool {
	return opt.init
}

// GetValue returns the value of given Optional as a new value.
func (opt *Optional) GetValue() interface{} {
	return *opt.v
}

// NotNil allows a not nil value to be added to given Optional. Returns the Optional and should the Optional's value be nil a (NILVALUE) error
func (opt *Optional) NotNil() (*Optional, error) {

	if *opt.v == nil || *opt.v == (interface{}(nil)) {
		return opt, fmt.Errorf(NILVALUE)
	}

	opt.init = true
	return opt, nil
}

// WithDefaultTypeValue set a given Optional to use a given default value should the Optional be un-Initialized. Panics if default value is incorrect type for the current Optional. Returns the Optional.
func (opt *Optional) WithDefaultTypeValue(def *interface{}) *Optional {
	typeCheck(*opt.t, def)

	if !opt.WasInitialized() {
		opt.v = def
	}

	return opt
}

// Nillable allows a value either nil or otherwise to be added to the given Optional. Returns the Optional.
func (opt *Optional) Nillable() *Optional {

	fmt.Printf("opt.v = %v\n", *opt.v)

	if *opt.v == nil || *opt.v == (interface{}(reflect.New(*opt.t))) {
		return opt
	}

	opt.init = true
	return opt
}

// From returns an Optional from a given value
func From(v *interface{}) *Optional {
	var t reflect.Type
	if v == nil {
		t = reflect.TypeOf(v).Elem()
	} else {
		t = reflect.TypeOf(v)
	}
	return &Optional{v: v, t: &t, init: false}
}

func typeCheck(optType reflect.Type, value *interface{}) {
	t := reflect.ValueOf(*value)

	if t.Type() != optType {
		panic(fmt.Errorf(INCORRECTTYPE, t, optType))
	}
}