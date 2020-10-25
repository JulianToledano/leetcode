package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	lastValPos := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums[i] = nums[lastValPos]
			lastValPos--
		}
	}
	return lastValPos + 1
}
