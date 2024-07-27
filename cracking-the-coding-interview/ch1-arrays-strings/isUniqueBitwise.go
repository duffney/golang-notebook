package main

import (
	"fmt"
)

func main() {
	s := "aabc"
	
	for k,v := range s {
		fmt.Printf("key: %x, value:%x", k, v)
	}
}
