package main

import "fmt"

func main() {
	b := 6

	var b_ptr *int // *int is used to declare variable b_ptr to be a pointer to an int
	b_ptr = &b     // b_ptr is assigned value that is the address of where variable b is stored

	// Shorhand for the above two lines is:
	// b_ptr := &b

	fmt.Printf("value stored at b: %d\n", b)

	fmt.Printf("address of b_ptr: %p\n", b_ptr)
	fmt.Printf("value stored at b_ptr: %d\n", *b_ptr)

}
