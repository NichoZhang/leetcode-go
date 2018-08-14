package main

/**
* https://leetcode.com/problems/sqrtx/description/
* Example 1:
* Input: 4
* Output: 2
*
* Example 2:
* Input: 8
* Output: 2
*
* 1017 / 1017 test cases passed.
* beats 100%
* Runtime:4ms
* https://leetcode.com/submissions/detail/169157048/
**/

func Sqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left := 1
	right := x
	var mid int
	for left < right {
		mid = left + (right-left)/2
		if mid*mid <= x && x < (mid+1)*(mid+1) {
			return mid
		} else if mid*mid < x {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}
