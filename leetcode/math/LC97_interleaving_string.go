package main

import "math"

/*
interleavingString
Given strings s1, s2, and s3, find whether s3 is formed by an interleaving of s1 and s2.

An interleaving of two strings s and t is a configuration where s and t are divided into n and m substrings respectively,
such that:

s = s1 + s2 + ... + sn
t = t1 + t2 + ... + tm
|n - m| <= 1
The interleaving is s1 + t1 + s2 + t2 + s3 + t3 + ... or t1 + s1 + t2 + s2 + t3 + s3 + ...
Note: a + b is the concatenation of strings a and b.

if true, first n chars of s3 must match first n chars of s1 and/or s2

		next o chars must match first o chars of s2, if first n chars matched first n of s1 or vice versa

	 1. find largest matching common prefix from [i = 0.. j = n] between s3 and s1/s2
	 2. if matched s1, find largest substring [n+1, k] between s3 and s2
	    if none matched, goto 1. with i = 0, j = n - 1
*/
func main() {
	var s1 = "aabcc"
	var s2 = "dbbca"
	var s3 = "aadbbcbcac"
	println(isInterleave(s1, s2, s3))
}

func isInterleave(s1 string, s2 string, s3 string) bool {

	matchingString, substringLength := findLargestSubstring(s1, s2, s3, 0, int(math.Max(math.Max(float64(len(s1)), float64(len(s2))), float64(len(s3)))))

	return false
}

func findLargestSubstring(s1 string, s2 string, s3 string, i int, n int) (string, int) {
	j := i
	for ; j < n; j++ {
		if s1[j:j+1] != s3[j:j+1] {
			break
		}
	}
	k := i
	for ; k < n; k++ {
		if s1[k:k+1] != s3[k:k+1] {
			break
		}
	}
	if j > k {
		return "s1", j
	}
	return "s2", k
}
