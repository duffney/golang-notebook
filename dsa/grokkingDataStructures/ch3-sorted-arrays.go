package main

import (
	"fmt"
	"log"
)

//TODO: unsorted

type sortedArray struct {
	data [5]int
	count int
}

// get better at sorting algos bro
func (a *sortedArray) Insert(value int) {
	if a.count == len(a.data) {
		log.Fatal("array is full")
	}
	for i:= a.count; i > 0; i-- {
		if a.data[i-1] <= value {
			a.data[i] = value
			a.count++
			return
		} else {
			a.data[i] = a.data[i-1]
		}
	}
	a.data[0] = value
	a.count++
}

func main() {
	a := &sortedArray{}
	a.data = [5]int{1,3,0,0,0}
	a.count = 2
	a.Insert(2)
	fmt.Println(a.data)
}
