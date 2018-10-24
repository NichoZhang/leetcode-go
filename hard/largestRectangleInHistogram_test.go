package main

import (
	"testing"
)

func Test1(t *testing.T) {
	heights := []int{2, 1, 5, 6, 2, 3}
	areas := largestRectangleArea(heights)
	if areas != 10 {
		t.Fatal("wrong answers!")
	}
}

func Test2(t *testing.T) {
	heights := []int{2, 0, 1, 3, 2, 2, 3, 1}
	areas := largestRectangleArea(heights)
	if areas != 8 {
		t.Fatal("wrong answers!")
	}
}

func Test3(t *testing.T) {
	heights := []int{2, 0, 1, 3, 5, 1, 3, 1}
	areas := largestRectangleArea(heights)
	if areas != 6 {
		t.Fatal("wrong answers!")
	}
}

func Test4(t *testing.T) {
	heights := []int{2, 0, 1, 3, 7, 1, 3, 1}
	areas := largestRectangleArea(heights)
	if areas != 7 {
		t.Fatal("wrong answers!")
	}
}

func Test5(t *testing.T) {
	heights := []int{1, 1}
	areas := largestRectangleArea(heights)
	if areas != 2 {
		t.Fatal("wrong answers!")
	}
}

func Test6(t *testing.T) {
	heights := []int{1, 1, 1}
	areas := largestRectangleArea(heights)
	if areas != 3 {
		t.Fatal("wrong answers!")
	}
}
