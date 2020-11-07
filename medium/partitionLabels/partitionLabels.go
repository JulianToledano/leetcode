package main

import "fmt"

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}

func partitionLabels(S string) []int {
	r := []int{}
	charLastPos := map[rune]int{}
	for i, v := range S {
		charLastPos[v] = i
	}
	anchor := 0
	max := 0
	for i, v := range S {
		if max < charLastPos[v] {
			max = charLastPos[v]
		}
		if i == max {
			r = append(r, i-anchor+1)
			anchor = i + 1
		}
	}
	return r
}
