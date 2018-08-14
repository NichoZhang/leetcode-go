package main

/**
* https://leetcode.com/problems/binary-search/description/
* Example 1:
* Input: nums = [-1, 0, 3, 5, 9, 12], target = 9
* Output: 4
*
* Example 2:
* Input: nums = [-1, 0, 3, 5, 9, 12], target = 2
* Output: -1
*
* 46 / 46 test cases passed.
* beats 24.39%
* Runtime:52ms
* https://leetcode.com/submissions/detail/169148865/
**/

func Search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left == len(nums) || nums[left] != target {
		return -1
	}
	return left
}
