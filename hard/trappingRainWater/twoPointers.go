package main

func main() {
	print(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	ans := 0
	left := new(int)
	*left = 0
	right := new(int)
	*right = len(height) - 1
	left_max := 0
	right_max := 0
	for *left < *right {
		if height[*left] < height[*right] {
			if height[*left] >= left_max {
				left_max = height[*left]
			} else {
				ans += left_max - height[*left]
			}
			*left++
		} else {
			if height[*right] >= right_max {
				right_max = height[*right]
			} else {
				ans += right_max - height[*right]
			}
			*right--
		}
	}

	return ans
}
