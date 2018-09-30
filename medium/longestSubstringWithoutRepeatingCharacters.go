package main

// 首先遍历查找组成最长子串
// 然后定位到一个重复出现的，从重复位置再找
// 将找到的不同的子串的长度进行对比
func lengthOfLongestSubstring(s string) int {
	m := map[rune]int{}
	i, maxLen := 0, 0
	for j, rs := range s {
		if offset, ok := m[rs]; ok {
			i = max(i, offset)
		}
		m[rs] = j + 1
		maxLen = max(maxLen, j-i+1)
	}

	return maxLen
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 4ms implement
// func lengthOfLongestSubstring(s string) int {
//     var chs [256]int
//     maxLen := 0
//     l, r := 0, 0
//     for r < len(s) {
//         ch := s[r]
//         if chs[ch] < 1 {
//             chs[ch]++
//             r++
//         } else {
//             maxLen = Max(maxLen, r-l)
//             chs[s[l]]--
//             l++
//         }
//     }
//     maxLen = Max(maxLen, r-l)
//     return maxLen
// }
