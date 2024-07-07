package testpackage

import "fmt"

// Public (exported) functions start with a capital letter
// and are accessible from other packages
func MyFunction() {
	fmt.Println("Hello World")
}

// Private functions begin with a lowercase letter 
// and are not accessible from other packages
func myPrivateFunction() {
	fmt.Println("Hello Private World")
}
