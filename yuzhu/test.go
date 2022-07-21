package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	// val := 3
	res := removeDuplicates(nums)
	fmt.Println(res)
}

func removeDuplicates(nums []int) int {
	index := 0
	for index+2 < len(nums) {
		if nums[index] == nums[index+2] {
			nums = append(nums[:index], nums[index+1:]...)
			index--
		}
		index++
	}
	fmt.Println(nums)
	return len(nums)
}
