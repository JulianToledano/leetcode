package main

import (
	"reflect"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	if !reflect.DeepEqual(removeElement([]int{1, 1, 2}, 2), 2) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(removeElement([]int{3, 2, 2, 3}, 3), 2) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2), 5) {
		t.Error("ERROR. Expected another result")
	}
}
