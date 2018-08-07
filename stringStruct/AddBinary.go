package main

/**
* https://leetcode.com/problems/add-binary/description/
* Example 1:
* input: a = "11", b = "1"
* output: "100"
*
* Example 2:
* input: a = "1010", b = "1011"
* output: "10101"
*
* 294 / 294 test cases passed.
* beats 100%
* Runtime:0ms
* https://leetcode.com/submissions/detail/168095093/
**/

func AddBinary(a string, b string) string {
	ai, bi := len(a)-1, len(b)-1
	if len(a) < len(b) {
		ai, bi = bi, ai
		a, b = b, a
	}

	sum := 0
	result := make([]byte, len(a)+1)
	idx := len(result) - 1
	for ai >= 0 || bi >= 0 {
		if ai >= 0 {
			sum += int(a[ai] - '0')
		}

		if bi >= 0 {
			sum += int(b[bi] - '0')
		}

		result[idx] = byte(sum&1) + '0'
		sum = sum >> 1
		idx--
		ai--
		bi--
	}

	if sum > 0 {
		result[0] = '1'
		return string(result)
	}

	return string(result[1:])
}
