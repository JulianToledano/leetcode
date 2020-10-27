package main

import "fmt"

func main() {
	x := numIdenticalPairs([]int{1, 2, 3, 1, 1, 3})
	fmt.Println(x)
}

func numIdenticalPairs(nums []int) int {
	goodPairs := map[int][]int{0: []int{0}}
	for i := range nums {
		if _, ok := goodPairs[nums[i]]; !ok {
			goodPairs[nums[i]] = []int{0, 1}
		} else {
			goodPairs[nums[i]][1]++
			goodPairs[nums[i]][0] += goodPairs[nums[i]][1] - 1
			goodPairs[0][0] += goodPairs[nums[i]][1] - 1
		}
	}
	return goodPairs[0][0]
}
