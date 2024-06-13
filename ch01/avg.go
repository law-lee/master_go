package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	var sum float64
	var count float64
	if len(args) == 1 {
		fmt.Println("The average number = 0")
		return
	}
	for i := 1; i < len(args); i++ {
		n, err := strconv.ParseFloat(args[i], 64)
		if err != nil {
			continue
		}
		sum += n
		count++
	}
	fmt.Println("The average number= ", sum/count)
}
