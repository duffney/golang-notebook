package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		firstNum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			secondNum := nums[j]
			if firstNum+secondNum == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
