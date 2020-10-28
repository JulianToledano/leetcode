package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
}

func defangIPaddr(address string) string {
	var d byte
	d = 46
	lastDot := 0
	result := [][]byte{}
	for i := range address {
		if d == address[i] {
			result = append(result, []byte(address[lastDot:i]))
			lastDot = i + 1
		}
	}
	result = append(result, []byte(address[lastDot:]))
	return string(bytes.Join(result, []byte("[.]")))
}
