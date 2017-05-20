//go:generate genny -in=optional/optional.go -out=optional/testOptional.go gen "Type=Test"
package main

import (
	"fmt"
	"optional/optional"
)

type Test struct{}

func main() {
	test := optional.NotNilTest(nil)
	fmt.Printf("tests value is %v\n", test.GetValue())
	fmt.Printf("test is instanciated? %v\n", test.WasInitialised())
}
