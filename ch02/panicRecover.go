package main

import "fmt"

func a() {
	fmt.Println("Inside a()")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in a()", r)
		}
	}()
	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited")
	fmt.Println("Exiting a()")
}

func b() {
	fmt.Println("Inside b()")
	panic("Panic in b()!")
	fmt.Println("Exiting b()")
}

func main() {
	a()
	fmt.Println("main() ended!")
}
