package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(reverse(-123))
}
func reverse(x int) int {
	changeSing := 1
	if x < 0 {
		changeSing = -1
		x = x * changeSing
	}
	s1 := strconv.Itoa(x)
	var strNum string
	for i := len(s1) - 1; i >= 0; i-- {
		strNum += string(s1[i])
	}
	r, _ := strconv.Atoi(strNum)
	returnValue := r * changeSing
	if returnValue < math.MaxInt32 && returnValue > math.MinInt32 {
		return returnValue
	}
	return 0
}
