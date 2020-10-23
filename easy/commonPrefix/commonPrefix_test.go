package main

import (
	"reflect"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	if !reflect.DeepEqual(longestCommonPrefix([]string{"flower", "flow", "flight"}), "fl") {
		t.Error("ERROR. Expected another result")
	}

	if !reflect.DeepEqual(longestCommonPrefix([]string{"dog", "racecar", "car"}), "") {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(longestCommonPrefix([]string{"ab", "a"}), "a") {
		t.Error("ERROR. Expected another result")
	}

}
