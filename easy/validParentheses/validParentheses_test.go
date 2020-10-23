package main

import (
	"reflect"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	if !reflect.DeepEqual(isValid("()"), true) {
		t.Error("ERROR. Expected another result")
	}

	if !reflect.DeepEqual(isValid("()[]{}"), true) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(isValid("(]"), false) {
		t.Error("ERROR. Expected another result")
	}

	if !reflect.DeepEqual(isValid("([)]"), false) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(isValid("{[]}"), true) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(isValid("(){}}{"), false) {
		t.Error("ERROR. Expected another result")
	}

}
