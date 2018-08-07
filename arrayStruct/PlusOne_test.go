package main

import "testing"

func TestPlusOne1(t *testing.T) {
	t.Log(PlusOne([]int{4, 3, 2, 1}))
}

func TestPlusOne2(t *testing.T) {
	t.Log(PlusOne([]int{0, 0, 0, 0}))
}

func TestPlusOne3(t *testing.T) {
	t.Log(PlusOne([]int{9, 9, 9}))
}
