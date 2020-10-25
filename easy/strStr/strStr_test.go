package main

import (
	"reflect"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	if !reflect.DeepEqual(strStr("hello", "ll"), 2) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(strStr("aaaaa", "bba"), -1) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(strStr("", ""), 0) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(strStr("mississippi", "issip"), 4) {
		t.Error("ERROR. Expected another result")
	}
}
