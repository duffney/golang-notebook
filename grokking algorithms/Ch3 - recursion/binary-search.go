package main

import "fmt"

func main() {
	slice := []int{1, 3, 5, 7, 9, 10}
	//slice := []int{2, 1}
	fmt.Println(search(slice, 5))
}

func search(list []int, item int) int {
	// what's the base case?
	low := 0
	high := len(list) - 1
	guess := (low + high) / 2

	if guess == item {
		return guess
	}
	if guess > item {
		high = guess - 1
	}
	if guess < item {
		low = guess + 1
	}
	return search(list, guess)
}
