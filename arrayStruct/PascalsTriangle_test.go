package main

import "testing"

func TestPascalsTriangle1(t *testing.T) {
	t.Log(PascalsTriangle(5))
}

func TestPascalsTriangle2(t *testing.T) {
	t.Log(PascalsTriangle(0))
}
