package main

import (
	"strconv"
)

func main() {
	isPalindrome(0)
}

// func isPalindrome(x int) bool {
// 	if x == 0 {
// 		return true
// 	}
// 	if x < 0 || x%10 == 0 {
// 		return false
// 	}
// 	reverted := 0
// 	for x > reverted {
// 		reverted = reverted*10 + x%10
// 		x /= 10
// 	}
// 	return reverted == x || reverted/10 == x
// }

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}
	s1 := strconv.Itoa(x)
	// 12321 -> 5 | 4 | 5 - 4 -1
	for i := len(s1) - 1; i >= 0; i-- {
		if s1[i] != s1[len(s1)-i-1] {
			return false
		}
	}
	return true
}
