package main

import "testing"

func TestTwoSum1(t *testing.T) {
	t.Log(twoSum([]int{3, 2, 4}, 6))
}

func TestTwoSum2(t *testing.T) {
	t.Log(twoSum([]int{2, 7, 11, 15}, 9))
}
