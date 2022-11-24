// Select sort algorithm
package main

import "fmt"

func main() {
	arr := []int{10, 33, 5, 7, 9}
	fmt.Println(selectionSort(arr))
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func selectionSort(arr []int) []int {
	var newArr []int
	copyArr := arr

	for i := 0; i < len(arr); i++ {
		smallest := findSmallest(copyArr)
		newArr = append(newArr, copyArr[smallest])
		copyArr = removeIndex(copyArr, smallest)
	}
	return newArr
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
