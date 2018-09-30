package main

import "testing"

func TestFunc1(t *testing.T) {
	s := "abcabcbb"
	ret := lengthOfLongestSubstring(s)
	if ret != 3 {
		t.Fatal("wrong answer")
	}
}

func TestFunc2(t *testing.T) {
	s := "bbbbb"
	ret := lengthOfLongestSubstring(s)
	if ret != 1 {
		t.Fatal("wrong answer")
	}
}

func TestFunc3(t *testing.T) {
	s := "pwwkew"
	ret := lengthOfLongestSubstring(s)
	if ret != 3 {
		t.Fatal("wrong answer")
	}
}

func TestFunc4(t *testing.T) {
	s := "aababcabcd"
	ret := lengthOfLongestSubstring(s)
	if ret != 4 {
		t.Fatal("wrong answer")
	}
}
