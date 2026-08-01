package main

func lengthOfLongestSubstring(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		seen := make(map[byte]bool)
		count := 0
		for j := i; j < len(s); j++ {
			if seen[s[j]] {
				break
			}
			seen[s[j]] = true
			count++
		}
		if count > max {
			max = count
		}
	}
	return max
}

func lengthOfLongestSubstringFast(s string) int {
	seen := make(map[byte]int)
	max := 0
	for left, right := 0, 0; right < len(s); right++ {
		if idx, ok := seen[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		seen[s[right]] = right
		if window := right - left + 1; window > max {
			max = window
		}
	}
	return max
}
