package main

import "testing"

func TestDominantIndex1(t *testing.T) {
	t.Log(DominantIndex([]int{3, 6, 1, 0}))
}

func TestDominantIndex2(t *testing.T) {
	t.Log(DominantIndex([]int{1, 2, 3, 4}))
}
