package main

import (
	"reflect"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	if !reflect.DeepEqual(reverse(123), 321) {
		t.Error("ERROR. Expected another result")
	}

	if !reflect.DeepEqual(reverse(-123), -321) {
		t.Error("ERROR. Expected another result")
	}

	if !reflect.DeepEqual(reverse(120), 21) {
		t.Error("ERROR. Expected another result")
	}

	if !reflect.DeepEqual(reverse(1534236469), 0) {
		t.Error("ERROR. Expected another result")
	}
}
