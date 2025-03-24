package testpackage

import "fmt"

func MyFunction() {
	fmt.Println("Steps 1...4")
}

func myPrivateFunction() {
	fmt.Println("Step 5")
	fmt.Println("End of steps")
}