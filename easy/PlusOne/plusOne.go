package main

import (
	"math"
	"strconv"
)

func main() {
	plusOne([]int{6, 1, 4, 5, 3, 9, 0, 1, 9, 5, 1, 8, 6, 7, 0, 5, 5, 4, 3})
}

func plusOne(digits []int) []int {
	sum := float64(0)
	for i, v := range digits {
		sum += float64(v) * math.Pow(float64(10), float64(len(digits)-1-i))
	}
	sum += 1
	r := strconv.Itoa(int(sum))
	result := []int{}
	for i := 0; i < len(r); i++ {
		result = append(result, int(r[i]-'0'))
	}
	return result
}
