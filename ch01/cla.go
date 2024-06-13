package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: cla <num1> <num2>")
		os.Exit(1)
	}

	min, _ := strconv.ParseFloat(os.Args[1], 64)
	max, _ := strconv.ParseFloat(os.Args[1], 64)

	for i := 2; i < len(os.Args); i++ {
		n, _ := strconv.ParseFloat(os.Args[i], 64)

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}
