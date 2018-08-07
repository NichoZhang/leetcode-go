package main

/**
* https://leetcode.com/problems/pascals-triangle/description/
* Example:
* Input: 5
* Output:
* [
*      [1],
*     [1,1],
*    [1,2,1],
*   [1,3,3,1],
*  [1,4,6,4,1]
* ]
*
* 15 / 15 test cases passed.
* beats 100%
* Runtime:0ms
* https://leetcode.com/submissions/detail/163374898/
**/

func PascalsTriangle(numRows int) [][]int {
	matrix := make([][]int, numRows)
	if numRows == 0 {
		return matrix
	}
	matrix[0] = []int{1}
	for i := 1; i < numRows; i++ {
		matrix[i] = make([]int, i+1)
		matrix[i][0] = 1
		matrix[i][i] = 1
		for j := 1; j < i; j++ {
			matrix[i][j] = matrix[i-1][j-1] + matrix[i-1][j]
		}
	}

	return matrix
}
