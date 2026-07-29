package main

import "fmt"

func main() {
	fmt.Println(takeCharacters("aabaaaacaabc", 2))             // 8
	fmt.Println(takeCharacters("a", 1))                        // -1
	fmt.Println(takeCharacters("abc", 1))                      // 2
	fmt.Println(takeCharacters("abbbbbbbbbbbbbbbbbbbbbbc", 1)) // 3
}

/*
		aabaaaacaabc
			  ^
	   i,j=midpt
	   freq_map(chars)
	   if chars[j+1] expands window to include a char already at k in freq_map, don't advance j
	   if chars[i-1] expands window to include a char already at k in freq_map, don't advance i

*/

/*
You are given a string s consisting of the characters 'a', 'b', and 'c' and a non-negative integer k.
Each minute, you may take either the leftmost character of s, or the rightmost character of s.

Return the minimum number of minutes needed for you to take at least k of each character,
or return -1 if it is not possible to take k of each character.
*/

func takeCharacters(s string, k int) int {
	// TODO
	return -1
}
