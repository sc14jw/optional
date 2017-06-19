# Optional [![Build Status](https://travis-ci.org/sc14jw/optional.svg?branch=master)](https://travis-ci.org/sc14jw/optional) [![Coverage Status](https://coveralls.io/repos/github/sc14jw/optional/badge.svg?branch=master)](https://coveralls.io/github/sc14jw/optional?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/sc14jw/optional)](https://goreportcard.com/report/github.com/sc14jw/optional)
A small library for quickly dealing with nil pointers in go based heavily from the Guava package for java (https://github.com/google/guava/wiki/UsingAndAvoidingNullExplained).

Quickly complete a nil check for the following code:

```go
  type testStruct struct {}

  opt, err := optional.NotNil(nil)

  fmt.Println(err)
  // "Given value was nil"
```

As well as being able to quickly check for nils it is also possible to query an Optional to see whether or not:

```go
  opt, _ := optional.NotNil(nil)
  opt.WasInitialized() // false

  opt = optional.NotNil(interface{}(testStruct{}))
  opt.WasInitialized() // true

  // it is also possible to query the Optional's value:
  opt.GetValue() // testStruct{}
```

Through an Optional it is also possible to set a default value for nil values using the "Nillable" strategy:

```go
  opt := optional.Nillable(nil) // note the lack of an error here as values are allowed to be nil
  opt.WasInitialized() // false

  opt.GetValue() // nil

  opt.WithDefaultValue(testStruct{})
  opt.GetValue() // testStruct{}
  ```

  The "WithDefaultValue" method will attempt to preserve run-time thread safety through reflection panicking, should the incorrect type be passed as a parameter.
