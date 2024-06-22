package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var value int64 = 5
	var p1 = &value
	// A pointer of type unsafe.Pointer can override the type system of Go.
	// This is unquestionably fast but it can also be dangerous if used incorrectly
	// or carelessly. Additionally, it gives developers more control over data.
	var p2 = (*int32)(unsafe.Pointer(p1))

	fmt.Println("*p1: ", *p1)
	fmt.Println("*p2: ", *p2)
	*p1 = 5434123412312431212
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)
	*p1 = 54341234
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)
}
