package main

/**
* https://leetcode.com/problems/largest-number-at-least-twice-of-others/description/
* Example 1:
* Input: nums = [3, 6, 1, 0]
* Output: 1
*
* Example 2:
* Input: nums = [1, 2, 3, 4]
* Output: -1
*
* 250 / 250 test cases passed.
* beats 100%
* Runtimes:0ms
* https://leetcode.com/submissions/detail/163205857/
**/

func DominantIndex(nums []int) int {
	l := len(nums)
	if l == 1 {
		return 0
	}

	topTwo := []int{0, 0}
	if nums[0] > nums[1] {
		topTwo = []int{0, 1}
	} else {
		topTwo = []int{1, 0}
	}

	for i := 2; i < l; i++ {
		if nums[i] > nums[topTwo[0]] {
			topTwo = []int{i, topTwo[0]}
		} else if nums[i] > nums[topTwo[1]] {
			topTwo = []int{topTwo[0], i}
		}
	}
	if nums[topTwo[1]] == 0 {
		return topTwo[0]
	}
	if (nums[topTwo[0]] / nums[topTwo[1]]) > 1 {
		return topTwo[0]
	}
	return -1
}
