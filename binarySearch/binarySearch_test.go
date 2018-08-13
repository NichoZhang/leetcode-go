package main

import "testing"

func TestSeach1(t *testing.T) {
	ret := Search([]int{-1, 0, 3, 5, 9, 12}, 9)
	if ret == 4 {
		t.Log("Success")
	} else {
		t.Error("Search fail")
	}
}

func TestSeach2(t *testing.T) {
	ret := Search([]int{-1, 0, 3, 5, 9, 12}, 2)
	if ret == -1 {
		t.Log("Success")
	} else {
		t.Error("Search fail")
	}
}
