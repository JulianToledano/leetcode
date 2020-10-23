package main

import (
	"reflect"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	if !reflect.DeepEqual(isPalindrome(123), false) {
		t.Error("ERROR. Expected another result")
	}

	if !reflect.DeepEqual(isPalindrome(-123), false) {
		t.Error("ERROR. Expected another result")
	}

	if !reflect.DeepEqual(isPalindrome(121), true) {
		t.Error("ERROR. Expected another result")
	}

	if !reflect.DeepEqual(isPalindrome(10), false) {
		t.Error("ERROR. Expected another result")
	}

	if !reflect.DeepEqual(isPalindrome(0), true) {
		t.Error("ERROR. Expected another result")
	}

}
