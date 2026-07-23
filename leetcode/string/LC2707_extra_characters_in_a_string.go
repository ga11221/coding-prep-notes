package main

import (
	"cmp"
	"slices"
	"strings"
)

func main() {}

/*
You are given a 0-indexed string s and a dictionary of words dictionary.
You have to break s into one or more non-overlapping substrings such that each substring is present in dictionary.
There may be some extra characters in s which are not present in any of the substrings.

Return the minimum number of extra characters left over if you break up s optimally.

	Example 1:

	Input: s = "leetscode", dictionary = ["leet", "leets", "code","leetcode"]
	Output: 1
	Explanation: We can break s in two substrings: "leet" from index 0 to 3 and "code" from index 5 to 8.
	There is only 1 unused character (at index 4), so we return 1.

	Example 2:

	Input: s = "sayhelloworld", dictionary = ["hello","world"]
	Output: 3
	Explanation: We can break s in two substrings: "hello" from index 3 to 7 and "world" from index 8 to 12.
	The characters at indices 0, 1, 2 are not used in any substring and thus are considered as extra characters. Hence, we return 3.

	group words in dictionary by 1st letter
	for each word in dictionary where 1st letter == s[0] from longest to shortest:
	  remove word from s
	extra chars, if any, remain in s
*/
func minExtraChar(s string, dictionary []string) int {
	dictMap := map[string][]string{}
	cmpFn := func(a, b string) int {
		return -cmp.Compare(len(a), len(b))
	}
	for _, word := range dictionary {
		firstLetter := string(word[0])
		if words, ok := dictMap[firstLetter]; ok {
			i, _ := slices.BinarySearchFunc(words, word, cmpFn)
			dictMap[firstLetter] = slices.Insert(words, i, word)
		} else {
			dictMap[firstLetter] = []string{word}
		}
	}
	extraLetters := []int{}
	for i := 0; i < len(s); {
		if words, ok := dictMap[string(s[i])]; ok {
			for _, w := range words {
				if strings.HasPrefix(s[i:], w) {
					i = i + len(w)
				}
			}
		} else {
			extraLetters = append(extraLetters, i)
			i++
		}
	}
	return len(extraLetters)
}
