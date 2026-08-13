package main

import "fmt"

/*
LC187 Repeated DNA Sequences (Hash Table, Sliding Window, Bit Manipulation)

The DNA sequence is composed of a series of nucleotides abbreviated as 'A',
'C', 'G', and 'T'.

For example, "ACGAATTCCG" is a DNA sequence.

When studying DNA, it is useful to identify repeated sequences within the DNA.

Given a string s that represents a DNA sequence, return all the 10-letter-long
sequences (substrings) that occur more than once in a DNA molecule. You may
return the answer in any order.

Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
Output: ["AAAAACCCCC","CCCCCAAAAA"]
        ["AAAAACCCCC","CCCCCAAAAA"]

Input: s = "AAAAAAAAAAAAA"
	   s = "AAAAAAAAAAAAA"
Output: ["AAAAAAAAAA"]
        ["AAAAAAAAAA"]

Constraints:
- 1 <= s.length <= 10^5
- s[i] is either 'A', 'C', 'G', or 'T'.
*/

/*
multiset of all 10-letter sequences/substrings in s with count > 1

if len(s) = 11, there exist 2 10-letter substrings, one starting at 0 and the second at 1
if len(s) = 12, there exist 3 10-letter substrings, starting at 0, 1, and 2

	n-10+1 substrings of length 10 in s

recast:

	map of 10letter chars
*/
func main() {
	//s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	//s := "AAAAAAAAAAAAA"
	s := "AAAAAAAAAA"
	fmt.Print(findRepeatedDnaSequences(s))
}

func findRepeatedDnaSequences(s string) []string {
	bag := map[string]int{}
	for i := range s {
		if i+10 > len(s) {
			break
		}
		bag[s[i:i+10]]++
	}
	repeats := []string{}
	for k, v := range bag {
		if v > 1 {
			repeats = append(repeats, k)
		}
	}
	return repeats
}
