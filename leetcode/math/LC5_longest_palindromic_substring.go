package main

/*

longestPalindromicSubstring

Given a string s, return the longest palindromic substring in s.

Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"


Constraints:

1 <= s.length <= 1000
s consist of only digits and English letters.

*/

func main() {
	//var s = "babad"
	//var s = "cbbd"
	//var s = "bc"
	var s = "cccc"
	println(longestPalindrome(s))
}

type PalindromeCenter struct {
	Left, Right int
}

func longestPalindrome(s string) string {
	// find all the palindromic centers: where s[i] == s[i+1] or s[i-1] == s[i+1]
	// for all palindromic centers (i, j):
	//  while i > 0 and j < len(s) and s[i] == s[j] continue, else if j-i > currentMax, set longestPalindrome = s[i:j]
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	var palindromeCenters = []PalindromeCenter{}
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			palindromeCenters = append(palindromeCenters, PalindromeCenter{i - 1, i})
		}
		if i+1 < len(s) && s[i-1] == s[i+1] {
			palindromeCenters = append(palindromeCenters, PalindromeCenter{i - 1, i + 1})
		}
	}
	var maxPalindromeLength = 0
	var maxPalindrome = ""
	for _, center := range palindromeCenters {
		for i, j := center.Left, center.Right; i >= 0 && j < len(s); i, j = i-1, j+1 {
			if s[i] != s[j] {
				break
			}
			var length = j - i + 1
			if length > maxPalindromeLength {
				maxPalindromeLength = length
				maxPalindrome = s[i : j+1]
			}
		}
	}
	if maxPalindrome == "" {
		return s[0:1]
	}
	return maxPalindrome
}
