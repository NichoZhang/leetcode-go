package main

import "testing"

func TestFindPivotIndex1(t *testing.T) {
	t.Log(FindPivotIndex([]int{1, 7, 3, 6, 5, 6}))
}

func TestFindPivotIndex2(t *testing.T) {
	t.Log(FindPivotIndex([]int{1, 2, 3}))
}

func TestFindPivotIndex3(t *testing.T) {
	t.Log(FindPivotIndex([]int{}))
}
