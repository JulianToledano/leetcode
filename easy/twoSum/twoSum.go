package main

import "fmt"

func main() {
	x := []int{2, 7, 11, 15}
	fmt.Println(twoSum(x, 9))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		complement := target - v
		c, ok := m[complement]
		if ok {
			return []int{c, i}
		}
		m[v] = i
	}
	return []int{0}
}
