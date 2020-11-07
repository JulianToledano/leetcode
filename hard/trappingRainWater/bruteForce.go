package main

func main() {
	print([]int{4, 2, 0, 3, 2, 5})
}

func trap(height []int) int {
	ans := 0
	for i := 0; i < len(height); i++ {
		left_max := 0
		right_max := 0
		for j := i; j >= 0; j-- {
			if left_max < height[j] {
				left_max = height[j]
			}
		}
		for j := i; j < len(height); j++ {
			if right_max < height[j] {
				right_max = height[j]
			}
		}
		min := 0
		if left_max < right_max {
			min = left_max
		} else {
			min = right_max
		}
		ans += min - height[i]
	}
	return ans
}
