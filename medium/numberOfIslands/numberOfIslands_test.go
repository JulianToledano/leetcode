package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	if !reflect.DeepEqual(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}), 1) {
		t.Error("ERROR. Expected another result")
	}

	if !reflect.DeepEqual(numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}), 3) {
		t.Error("ERROR. Expected another result")
	}

}
