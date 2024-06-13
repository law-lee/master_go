package main

import (
	"io"
	"os"
)

func main() {
	mystr := ""
	args := os.Args

	if len(args) == 1 {
		mystr = "Please input one argument"
	} else {
		mystr = args[1]
	}

	io.WriteString(os.Stdout, "This is a standard output\n")
	io.WriteString(os.Stderr, mystr)
	io.WriteString(os.Stderr, "\n")
}
