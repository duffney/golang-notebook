// Quicksort
package main

import "fmt"

func main() {
	list := []int{33, 10, 15, 45, 65, 16, 5}

	fmt.Println(quicksort(list))
}

func quicksort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	if len(list) == 2 {
		left, right := 0, len(list)-1
		if list[0] > list[1] {
			list[left], list[right] = list[right], list[left]
		}
		return list
	}

	pivot := list[0]
	var less []int
	var greater []int

	for i := range list {
		if list[i] < pivot {
			less = append(less, list[i])
		}
		if list[i] > pivot {
			greater = append(greater, list[i])
		}
	}

	return append(append(quicksort(less), pivot), quicksort(greater)...)
}
