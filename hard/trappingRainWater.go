package main

// return 4ms, beats 100%
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	ans := 0
	lens := len(height)
	maxLeft := make([]int, lens)
	maxRight := make([]int, lens)
	maxLeft[0] = height[0]
	for i := 1; i < lens; i++ {
		maxLeft[i] = max(height[i], maxLeft[i-1])
	}

	maxRight[lens-1] = height[lens-1]
	for i := lens - 2; i > 0; i-- {
		maxRight[i] = max(height[i], maxRight[i+1])
	}
	for i := 1; i < lens-1; i++ {
		ans += min(maxLeft[i], maxRight[i]) - height[i]
	}

	return ans
}

// runtime 8 ms, beats 28.26%
func twoPointTrap(height []int) int {
	left := 0
	right := len(height) - 1
	ans := 0
	maxLeft, maxRight := 0, 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				ans += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				ans += maxRight - height[right]
			}
			right--
		}
	}

	return ans
}

// runtime 76ms, beats 10.87%
func bruteForceTrap(height []int) int {
	ans := 0
	lens := len(height)

	for i := 0; i < lens-1; i++ {
		maxLeft, maxRight := 0, 0
		for j := i; j >= 0; j-- {
			maxLeft = max(maxLeft, height[j])
		}
		for j := i; j < lens; j++ {
			maxRight = max(maxRight, height[j])
		}
		ans += min(maxLeft, maxRight) - height[i]
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
