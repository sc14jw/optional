package optional

import "testing"

const initializedError = "Optional was Initialized. Initialized = %v."
const notInitializedError = "Optional was not Initialized. Initialized = %v."
const getValueError = "GetValue returned %v not expected %v."
const wrongError = "Error was %v not expected."
const errorExpected = "No Error was returned. Expected %v."
const panicExpected = "The Optional did not panic."
const errorText = "This is an error."
const expectedErrorText = "Expected 'This is an error.' But got %v."

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
	opt, err := NotNil(test)

	if !opt.WasInitialized() {
		t.Errorf(notInitializedError, opt.WasInitialized())
	} else if opt.GetValue() != test {
		t.Errorf(getValueError, opt.GetValue(), test)
	} else if err != nil {
		t.Errorf(wrongError, err)
	}

	opt = Nillable(test)

	if !opt.WasInitialized() {
		t.Errorf(notInitializedError, opt.WasInitialized())
	} else if opt.GetValue() != test {
		t.Errorf(getValueError, opt.GetValue(), test)
	}

	var secondErr error
	optSecond, secondErr := NotNilWithMessage(test, errorText)
	if !optSecond.WasInitialized() {
		t.Errorf(notInitializedError, optSecond.WasInitialized())
	} else if optSecond.GetValue() != test {
		t.Errorf(getValueError, optSecond.GetValue(), test)
	} else if secondErr != nil {
		t.Errorf(wrongError, err)
	}

}

func TestOptionalNilCreation(t *testing.T) {
	var err error
	opt, err := NotNil(nilTest)

	if opt.WasInitialized() {
		t.Errorf(initializedError, opt.WasInitialized())
	} else if opt.GetValue() != nilTest {
		t.Errorf(getValueError, opt.GetValue(), nilTest)
	} else if err == nil {
		t.Errorf(errorExpected, NILVALUE)
	}

	var nilPointer *interface{}
	opt, err = NotNil(nilPointer)
	if opt.WasInitialized() {
		t.Errorf(initializedError, opt.WasInitialized())
	} else if err == nil {
		t.Errorf(errorExpected, NILVALUE)
	}

	opt = Nillable(nilTest)

	if opt.WasInitialized() {
		t.Errorf(initializedError, opt.WasInitialized())
	} else if opt.GetValue() != nilTest {
		t.Errorf(getValueError, opt.GetValue(), nilTest)
	}

	opt = Nillable(nilPointer)
	if opt.WasInitialized() {
		t.Errorf(initializedError, opt.WasInitialized())
	}

	secondOpt, secondErr := NotNilWithMessage(nilTest, errorText)
	if secondOpt.WasInitialized() {
		t.Errorf(initializedError, secondOpt.WasInitialized())
	} else if secondOpt.GetValue() != nilTest {
		t.Errorf(getValueError, secondOpt.GetValue(), nilTest)
	} else if secondErr == nil {
		t.Errorf(errorExpected, NILVALUE)
	} else if secondErr.Error() != errorText {
		t.Errorf(expectedErrorText, secondErr)
	}

	opt, err = NotNilWithMessage(nilPointer, errorText)
	if opt.WasInitialized() {
		t.Errorf(initializedError, opt.WasInitialized())
	} else if err == nil {
		t.Errorf(errorExpected, NILVALUE)
	} else if secondErr.Error() != errorText {
		t.Errorf(expectedErrorText, secondErr)
	}
}

func TestDefaultValue(t *testing.T) {
	opt := Nillable(test).WithDefaultTypeValue(anotherTest)

	if !opt.WasInitialized() {
		t.Errorf(notInitializedError, opt.WasInitialized())
	} else if opt.GetValue() != test {
		t.Errorf(getValueError, opt.GetValue(), test)
	}

	opt = Nillable(nilTest).WithDefaultTypeValue(anotherTest)

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

	opt := Nillable(test).WithDefaultTypeValue(wrongTypeTest)
	opt.GetValue()
}
