package main

/**
* https://leetcode.com/problems/climbing-stairs/description/
* Example 1:
* Input: 2
* Output: 2
* Explanation: There are two ways to climb to the top.
* 1. 1 step + 1 step
* 2. 2 steps
*
* Example 2:
* Input: 3
* Output: 3
* Explanation: There are three ways to climb to the top.
* 1. 1 step + 1 step + 1 step
* 2. 1 step + 2 step
* 3. 2 step + 1 step
*
* 45 / 45 test cases passed.
* beats 100%
* Runtime:0ms
* https://leetcode.com/submissions/detail/171960488/
**/

func climbStairs(n int) int {
	ret := map[int]int{
		0: 0,
		1: 1,
		2: 2,
	}
	if n <= 2 {
		return ret[n]
	}

	fibN := 1
	fibNMinusOne := 1
	fibNMinusTwo := 2
	for i := 3; i <= n; i++ {
		fibN = fibNMinusOne + fibNMinusTwo

		fibNMinusTwo = fibNMinusOne
		fibNMinusOne = fibN
	}

	return fibN
}
