package main

func main() {
	x := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	print(x)
}

func findDisappearedNumbers(nums []int) []int {
	numbers := map[int]int{}
	for _, v := range nums {
		if _, ok := numbers[v]; !ok {
			numbers[v] = 1
		} else {
			numbers[v]++
		}
	}
	r := []int{}
	for i := 1; i < len(nums); i++ {
		if _, ok := numbers[i]; !ok {
			r = append(r, i)
		}
	}
	return r
}
