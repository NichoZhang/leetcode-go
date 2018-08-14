package main

import "testing"

func TestSqrt1(t *testing.T) {
	ret := Sqrt(4)
	if ret == 2 {
		t.Log("Success")
	} else {
		t.Error("sqrt failed")
	}
}

func TestSqrt2(t *testing.T) {
	ret := Sqrt(8)
	if ret == 2 {
		t.Log("Success")
	} else {
		t.Error("sqrt failed")
	}
}
