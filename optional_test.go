package optional

import "testing"

const initializedError = "Optional was Initialized. Initialized = %v."
const notInitializedError = "Optional was not Initialized. Initialized = %v."
const getValueError = "GetValue returned %v not expected %v."
const wrongError = "Error was %v not expected."
const errorExpected = "No Error was returned. Expected %v."
const panicExpected = "The Optional did not panic."

const helloWorld = "Hello World!"

var test interface{} = interface{}(testStruct{v: 12, n: "test"})
var anotherTest interface{} = interface{}(testStruct{v: 13, n: "anotherTest"})
var wrongTypeTest interface{} = interface{}(wrongTypeStruct{})
var nilTestStruct testStruct
var nilTest interface{}

type testStruct struct {
	v int
	n string
}

type wrongTypeStruct struct{}

func TestOptionalCreation(t *testing.T) {

	var err error
	opt, err := NotNil(&test)

	if !opt.WasInitialized() {
		t.Errorf(notInitializedError, opt.WasInitialized())
	} else if opt.GetValue() != test {
		t.Errorf(getValueError, opt.GetValue(), &test)
	} else if err != nil {
		t.Errorf(wrongError, err)
	}

	opt = Nillable(&test)

	if !opt.WasInitialized() {
		t.Errorf(notInitializedError, opt.WasInitialized())
	} else if opt.GetValue() != test {
		t.Errorf(getValueError, opt.GetValue(), &test)
	}

}

func TestOptionalNilCreation(t *testing.T) {
	var err error
	opt, err := NotNil(&nilTest)

	if opt.WasInitialized() {
		t.Errorf(initializedError, opt.WasInitialized())
	} else if opt.GetValue() != nilTest {
		t.Errorf(getValueError, opt.GetValue(), &nilTest)
	} else if err == nil {
		t.Errorf(errorExpected, NILVALUE)
	}

	opt = Nillable(&nilTest)

	if opt.WasInitialized() {
		t.Errorf(initializedError, opt.WasInitialized())
	} else if opt.GetValue() != nilTest {
		t.Errorf(getValueError, opt.GetValue(), &nilTest)
	}
}

func TestDefaultValue(t *testing.T) {
	opt := Nillable(&test).WithDefaultTypeValue(&anotherTest)

	if !opt.WasInitialized() {
		t.Errorf(notInitializedError, opt.WasInitialized())
	} else if opt.GetValue() != test {
		t.Errorf(getValueError, opt.GetValue(), test)
	}

	opt = Nillable(&nilTest).WithDefaultTypeValue(&anotherTest)

	if opt.WasInitialized() {
		t.Errorf(initializedError, opt.WasInitialized())
	} else if opt.GetValue() != anotherTest {
		t.Errorf(getValueError, opt.GetValue(), anotherTest)
	}
}

func TestDefaultValueWrongType(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Errorf(panicExpected)
		}
	}()

	opt := Nillable(&test).WithDefaultTypeValue(&wrongTypeTest)
	opt.GetValue()
}
