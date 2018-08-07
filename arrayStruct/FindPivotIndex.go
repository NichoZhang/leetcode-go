package main

/**
* https://leetcode.com/problems/find-pivot-index/description/
* Example 1:
* Input: [1, 7, 3, 6, 5, 6]
* Output: 3
*
* Example 2:
* Input: [1, 2, 3]
* Output: -1
*
* 741 / 741 test cases passed.
* beats 100%
* Runtime:28ms
* https://leetcode.com/submissions/detail/163195320/
**/

func FindPivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	var sum int
	for index := range nums {
		sum += nums[index]
	}
	if (sum - nums[0]) == 0 {
		return 0
	}
	ret := -1
	fSum := 0 //front part sum
	l := len(nums)
	for i := 1; i < l; i++ {
		fSum += nums[i-1]
		if (sum - nums[i]) == fSum*2 {
			ret = i
			break
		}
	}
	return ret
}
