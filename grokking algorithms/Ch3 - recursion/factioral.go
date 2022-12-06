package main

import (
	"fmt"
)

func main() {
	fmt.Println(fact(3))
}

func fact(x int) int {
	if x == 1 { // base case
		return 1
	} else {
		return x * fact(x-1) // recursive case
	}
}
