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

	io.WriteString(os.Stdout, mystr)
	io.WriteString(os.Stdout, "\n")
}
