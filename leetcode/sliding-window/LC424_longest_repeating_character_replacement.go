package main

import "fmt"

/*
*
s = "AABABBA" k = 1
replace first B with A to get AAAABBA (4 As)
replace third A with B to get AABBBBA (4 Bs)

So, given a string of upper case characters, replace k (or less? Yes, "at most k times") characters
to get the longest sequence of repeating upper case chars possible.
Return the length of the sequence.

Edge cases: k = 0, len(s) = 1, len(s) = 2
ABABA
f(0, _) = 1
f(1, k) = 1+[c | c in f(0, k) if s[1] == s[0]] ++ 1+[c | c in f(0,k-1) if s[1] != s[0] and k > 0] ++ 1
f(2, k) = 1+ [c | c in f(1, k) if s[1] == s[0]] ++ 1+[c | c in f(1,k-1) if s[2] != s[1] and k > 0] ++ 1

f(i, k) = max repeating char at i with k or less replacements
f(i+1, k) = 1 + f(i, k) if s[i+1] == max repeating char else 1 + f(i, k-1)
*/
func main() {
	s := "AABABBA"
	k := 1
	fmt.Printf("for string: %v and k: %v, max seq count: %v\n", s, k, characterReplacement(s, k))
}

func characterReplacement(s string, k int) int {
	counts := []int{}
	for i := len(s) - 1; i >= 0; i-- {
		counts = append(counts, f(s, i, k))
		fmt.Printf("counts after i = %v: %v\n", i, counts)
	}
	max := 0
	for _, n := range counts {
		if n > max {
			max = n
		}
	}
	return max
}

func f(s string, i int, k int) int {
	fmt.Printf("s: %v i: %v k: %v\n", s, i, k)
	if i == 0 {
		return 1
	}
	if s[i] == s[i-1] {
		return 1 + f(s, i-1, k)
	}
	if k > 0 {
		changeThis := 1 + f(s[:i]+string(s[i-1])+string([]rune(s)[i+1:]), i-1, k-1)
		changeNext := 1 + f(s[:i-1]+string(s[i])+string([]rune(s)[i:]), i-1, k-1)
		if changeThis > changeNext {
			return changeThis
		}
		return changeNext
	}
	return 1

}
