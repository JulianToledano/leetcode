package main

import (
	"reflect"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	if !reflect.DeepEqual(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}), []int{5, 6}) {
		t.Error("ERROR. Expected another result")
	}

}
