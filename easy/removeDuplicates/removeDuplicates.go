package main

func main() {
	removeDuplicates([]int{1, 1, 2, 2, 2, 2, 3, 4, 5, 5, 6, 6, 6, 6, 7, 8, 9, 9, 9})

}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	differentNumbers := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[differentNumbers] {
			differentNumbers++
			nums[differentNumbers] = nums[i]
		}
	}
	return differentNumbers + 1
}
