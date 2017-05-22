# Optional
A small library for quickly dealing with nil pointers in go based heavily from the Guava package for java (https://github.com/google/guava/wiki/UsingAndAvoidingNullExplained).

Quickly complete a nil check for the following code:

```go
  type testStruct struct {}

  opt, err := optional.From(nil).NotNil()

  fmt.Println(err)
  // "Given value was nil"
```

As well as being able to quickly check for nils it is also possible to query an Optional to see whether or not:

```go
  opt, _ := optional.From(nil).NotNil()
  opt.WasInitialized() // false

  opt = optional.From(testStruct{}).NotNil()
  opt.WasInitialized() // true

  // it is also possible to query the Optional's value:
  opt.GetValue() // testStruct{}
```

Through an Optional it is also possible to set a default value for nil values using the "Nillable" strategy:

```go
  opt := optional.From(nil).Nillable() // note the lack of an error here as values are allowed to be nil
  opt.WasInitialized() // false

  opt.GetValue() // nil

  opt.WithDefaultValue(testStruct{})
  opt.GetValue() // testStruct{}
  ```

  The "WithDefaultValue" method will attempt to preserve run-time thread safety through reflection panicing should the incorrect type be passed as a parameter.
