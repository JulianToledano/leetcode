package main

import (
	"reflect"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	if !reflect.DeepEqual(romanToInt("I"), 1) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(romanToInt("II"), 2) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(romanToInt("III"), 3) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(romanToInt("IV"), 4) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(romanToInt("VIII"), 8) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(romanToInt("IX"), 9) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(romanToInt("LVIII"), 58) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(romanToInt("DCXXI"), 621) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(romanToInt("MDCCCLXXXIV"), 1884) {
		t.Error("ERROR. Expected another result")
	}
	if !reflect.DeepEqual(romanToInt("MMCCCXCIX"), 2399) {
		t.Error("ERROR. Expected another result")
	}
}
