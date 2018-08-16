package main

/**
* https://leetcode.com/problems/two-sum/description/
* Example:
* Given nums = [2, 7, 11, 15], target = 9,
* return [0, 1]
*
* 29 / 29 test cases passed.
* beats 100%
* Runtime:4ms
* https://leetcode.com/submissions/detail/169801698/
**/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	len := len(nums)
	for i := 0; i < len; i++ {
		ret := target - nums[i]
		if j, ok := m[ret]; ok {
			return []int{j, i}
		}
		m[nums[i]] = i
	}
	return nil
}
