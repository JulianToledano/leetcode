package main

import (
	"reflect"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	if !reflect.DeepEqual(searchInsert([]int{1, 3, 5, 6}, 5), 2) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(searchInsert([]int{1, 3, 5, 6}, 2), 1) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(searchInsert([]int{1, 3, 5, 6}, 7), 4) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(searchInsert([]int{1, 3, 5, 6}, 0), 0) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(searchInsert([]int{1}, 0), 0) {
		t.Error("ERROR. Expected another result")
	}

}
