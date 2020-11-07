package main

import "fmt"

func main() {
	print(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	ans := 0
	size := len(height)
	lm := make([]int, size)
	rm := make([]int, size)
	lm[0] = height[0]
	rm[size-1] = height[size-1]
	for i := 1; i < size; i++ {
		lm[i] = max(height[i], lm[i-1])
	}
	fmt.Println(lm)
	for i := size - 2; i >= 0; i-- {
		rm[i] = max(height[i], rm[i+1])
	}
	fmt.Println(rm)
	for i := 1; i < size-1; i++ {
		ans += min(lm[i], rm[i]) - height[i]
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
