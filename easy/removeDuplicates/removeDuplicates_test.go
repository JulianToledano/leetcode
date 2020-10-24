package main

import (
	"reflect"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	if !reflect.DeepEqual(removeDuplicates([]int{1, 1, 2}), 2) {
		t.Error("ERROR. Expected another result")
	}
}
