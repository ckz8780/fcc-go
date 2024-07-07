package testpackage

import "fmt"

// Public (exported) functions start with a capital letter
// and are accessible from other packages
func MyFunction(step int) {
	if step == 1 {
		fmt.Println("Step 1")
	} else if step == 2 {
		fmt.Println("Step 2")
	} else if step == 3 {
		fmt.Println("Step 3")
	} else if step == 4 {
		fmt.Println("Step 4")
	} else {
		fmt.Println("Step not supported")
	}
}

// Private functions begin with a lowercase letter 
// and are not accessible from other packages
func myPrivateFunction() {
	fmt.Println("Hello Private World")
}
