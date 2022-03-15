package main

import (
	"reflect"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	if !reflect.DeepEqual(plusOne([]int{1, 2, 3}), []int{1, 2, 4}) {
		t.Error("ERROR. Expected another result")
	}

	if !reflect.DeepEqual(plusOne([]int{9}), []int{1, 0}) {
		t.Error("ERROR. Expected another result")
	}

	if !reflect.DeepEqual(plusOne([]int{9, 9, 9}), []int{1, 0, 0, 0}) {
		t.Error("ERROR. Expected another result")
	}

}
