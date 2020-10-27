package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	if !reflect.DeepEqual(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}), 4) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(numIdenticalPairs([]int{1, 1, 1, 1}), 6) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(numIdenticalPairs([]int{1, 2, 3}), 0) {
		t.Error("ERROR. Expected another result")
	}

}
