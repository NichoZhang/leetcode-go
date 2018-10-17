package main

import "testing"

func TestTrap1(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	trapRain := trap(height)
	if trapRain != 6 {
		t.Fatal("wrong result number")
	}
}

func TestTrap2(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 3, 1}
	trapRain := trap(height)
	if trapRain != 8 {
		t.Fatal("wrong result number")
	}
}

func TestTrap3(t *testing.T) {
	height := []int{0, 0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	trapRain := trap(height)
	if trapRain != 6 {
		t.Fatal("wrong result number")
	}
}

func TestTrap4(t *testing.T) {
	height := []int{7, 0, 4}
	trapRain := trap(height)
	if trapRain != 4 {
		t.Fatal("wrong result number")
	}
}
