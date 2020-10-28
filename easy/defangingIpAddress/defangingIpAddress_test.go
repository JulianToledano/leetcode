package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	if !reflect.DeepEqual(defangIPaddr("1.1.1.1"), "1[.]1[.]1[.]1") {
		t.Error("ERROR. Expected another result")
	}

	if !reflect.DeepEqual(defangIPaddr("12.132.18.145"), "12[.]132[.]18[.]145") {
		t.Error("ERROR. Expected another result")
	}

}
