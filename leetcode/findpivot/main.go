package main

import "fmt"

func main() {
	//nums := []int{1, 2, 3, 4}
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums))
}

func pivotIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println(sum)
	leftsum := 0
	for i := 0; i < len(nums); i++ {
		if leftsum == sum-nums[i]-leftsum {
			return i
		}
		leftsum += nums[i]
	}
	return -1
}
