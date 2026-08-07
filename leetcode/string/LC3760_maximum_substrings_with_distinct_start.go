package main

/*
LC3760 Maximum Substrings With Distinct Start (String)

You are given a string s consisting of lowercase English letters.

Return an integer denoting the maximum number of substrings you can split s
into such that each substring starts with a distinct character (i.e., no two
substrings start with the same character).

A substring is a contiguous, non-empty sequence of characters within the string.

Input: s = "abab"
Output: 2
Explanation: "abab" can be split into "a" and "bab"; each starts with a
distinct character ('a' and 'b'). The answer is 2.

Input: s = "abcd"
Output: 4
Explanation: "abcd" can be split into "a", "b", "c", "d", each with a
distinct starting character.

Input: s = "aaaa"
Output: 1
Explanation: All characters are 'a'; only one substring can start with 'a'.

Constraints:
- 1 <= s.length <= 10^5
- s consists of lowercase English letters.
*/
func maxDistinct(s string) int {
	distinct := map[rune]int{}
	count := 0
	for _, c := range s {
		if _, ok := distinct[c]; !ok {
			count++
			distinct[c] = 1
		}
	}
	return count
}
