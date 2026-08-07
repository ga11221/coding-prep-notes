package main

/*
LC1759 Count Number of Homogenous Substrings (String, Math)

A string is homogenous if all the characters of the string are the same.

For example, "aaa" is homogenous, while "aab" is not.

You are given a string s, consisting of lowercase letters. Return the number
of homogenous substrings of s. Since the answer may be too large, return it
modulo 10^9 + 7.

A substring is a contiguous sequence of characters within the string.

ccc
ccc cc cc c c c

cccc
cccc ccc ccc cc cc cc
|s| = 4 s4 = 1 s3 = 2 s2 = 3 s1 = 4
Input: s = "abbcccaa"
Output: 13
Explanation: The homogenous substrings are listed as follows:
"a" appears 3 times, "bb" appears 3 times, "ccc" appears 6 times,
"aa" appears 1 time. Total = 3 + 3 + 6 + 1 = 13.

Input: s = "xy"
Output: 2
Explanation: The homogenous substrings are "x" and "y".

Input: s = "zzzzz"
s5 = 1 s4 = 2 s3 = 3 s2 = 4 s1 = 5
Output: 15

Constraints:
- 1 <= s.length <= 10^5
- s consists of lowercase English letters.
*/

func countHomogenous(s string) int {
	const mod = 1_000_000_007
	if len(s) == 1 {
		return 1
	}
	total := 0
	i, j := 0, 1
	for i < len(s) {
		for j < len(s) && s[i] == s[j] {
			j++
		}
		n := j - i
		total = (total + (n * (n + 1) / 2)) % mod
		i = j
	}

	return total
}
