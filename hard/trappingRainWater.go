package main

// runtime 8 ms, beats 28.26%
func trap(height []int) int {
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
