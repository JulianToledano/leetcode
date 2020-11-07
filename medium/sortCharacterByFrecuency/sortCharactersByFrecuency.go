package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(frequencySort("julialn"))
}

type keyValue struct {
	Key rune
	V   int
}

type keyValueList []keyValue

func (kv keyValueList) Len() int           { return len(kv) }
func (kv keyValueList) Less(i, j int) bool { return kv[i].V < kv[j].V }
func (kv keyValueList) Swap(i, j int)      { kv[i], kv[j] = kv[j], kv[i] }

func frequencySort(s string) string {
	// Create map
	// Loop array and sum to map values
	cf := map[rune]int{}
	for _, v := range s {
		if _, ok := cf[v]; !ok {
			cf[v] = 1
		} else {
			cf[v]++
		}
	}
	// Sort map by values
	kv := make(keyValueList, len(cf))
	i := 0
	for k, v := range cf {
		kv[i] = keyValue{k, v}
		i++
	}
	sort.Sort(sort.Reverse(kv))
	// Create new string with stringBuilder
	var sb strings.Builder
	for i := range kv {
		for j := 0; j < kv[i].V; j++ {
			sb.WriteRune(kv[i].Key)
		}
	}
	return sb.String()
}
