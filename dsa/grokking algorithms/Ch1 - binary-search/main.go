// binary search
package main

import "fmt"

func main() {
	list := [5]int{1, 3, 5, 7, 9}
	fmt.Println(binary_search(list, 5))
}

func binary_search(list [5]int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= len(list) {
		mid := (low + high) / 2
		guess := list[mid]
		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0
}
