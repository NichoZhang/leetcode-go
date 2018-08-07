package main

/**
* https://leetcode.com/problems/plus-one/description/
* Example 1:
* Input: [1,2,3]
* Output: [1,2,4]
*
* Example 2:
* Input: [4,3,2,1]
* Output: [4,3,2,2]
*
* 109 / 109 test cases passed.
* beats 100%
* Runtime:0ms
* https://leetcode.com/submissions/detail/163578911/
**/

func PlusOne(digits []int) []int {
	sum := 0
	one := 1
	len := len(digits)
	res := make([]int, len)

	for i := len - 1; i >= 0; i-- {
		sum = one + digits[i]
		one = sum / 10
		res[i] = sum % 10
	}

	if one > 0 {
		res = append([]int{one}, res...)
	}

	return res
}
