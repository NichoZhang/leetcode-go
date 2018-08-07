package main

import "testing"

func TestAddBinary1(t *testing.T) {
	t.Log(AddBinary("1111", "111"))
}

func TestAddBinary2(t *testing.T) {
	t.Log(AddBinary("11", "1"))
}
