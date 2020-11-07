package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	if !reflect.DeepEqual(partitionLabels("ababcbacadefegdehijhklij"), []int{9, 7, 8}) {
		t.Error("ERROR. Expected another result")
	}

}
