package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var prefix string
	// samePrefix = true
	isPrefix := true
	word := strs[0]
	for i := range word {
		for _, v := range strs[1:] {
			if len(v) <= i || v[i] != word[i] {
				isPrefix = false
				break
			}
		}
		if isPrefix {
			prefix += string(word[i])
		} else {
			break
		}
	}
	return prefix
}
