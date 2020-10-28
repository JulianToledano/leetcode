package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	if !reflect.DeepEqual(numJewelsInStones("aA", "aAAbbbb"), 3) {
		t.Error("ERROR. Expected another result")
	}

	if !reflect.DeepEqual(numJewelsInStones("z", "ZZZ"), 0) {
		t.Error("ERROR. Expected another result")
	}

}
