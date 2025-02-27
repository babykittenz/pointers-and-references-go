package main

import "fmt"

func main() {
	// Declare an integer variable
	var num int = 42

	// Declare a pointer to an integer and assign the memory address of 'num' to it
	var ptr *int = &num

	// Print the value of 'num' and the memory address stored in 'ptr'
	fmt.Printf("Value of 'num': %d\n", num)
	fmt.Printf("Memory address of 'num': %p\n", &num)

	// Print the value of the memory address pointed by 'ptr'
	fmt.Printf("Value pointed by 'ptr': %d\n", *ptr)

	// Change the value of the memory address pointed by 'ptr'
	*ptr = 100

	// The value of 'num' is now changed because 'ptr' points to 'num'
	fmt.Printf("New value of 'num': %d\n", num)
}
