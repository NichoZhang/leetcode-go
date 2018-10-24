package main

func largestRectangleArea(heights []int) int {
	area := 0
	lens := len(heights)
	if lens == 0 {
		return area
	}
	if lens == 1 {
		return heights[0]
	}

	for i := 0; i < lens; i++ {
		minLeft := heights[i]
		for j := i; j >= 0; j-- {
			minLeft = min(heights[j], minLeft)
			area = max(minLeft*(i-j+1), area)
		}
	}
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
