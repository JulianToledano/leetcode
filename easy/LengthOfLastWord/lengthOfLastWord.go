package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLastWord("this is my last work  "))
}

func lengthOfLastWord(s string) int {
	sum := 0
	pos := len(s) - 1
	for i := pos; i >= 0; i-- {
		if s[i] != ' ' {
			pos = i
			break
		}
	}
	for i := pos; i >= 0; i-- {
		if s[i] == ' ' {
			return sum
		}
		sum += 1
	}
	return sum
}
