package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 20}
	fmt.Println(max(list))
}

func max(list []int) int {
	if len(list) == 2 {
		if list[0] > list[1] {
			return list[0]
		}
		return list[1]
	}
	if list[0] > max(list[1:]) {
		return list[0]
	} else {
		return max(list[1:])
	}
}
