/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Example 2:
Input: strs = [""]
Output: [[""]]

Example 3:
Input: strs = ["a"]
Output: [["a"]]

Constraints:
- 1 <= strs.length <= 10^4
- 0 <= strs[i].length <= 100
- strs[i] consists of lowercase English letters.
*/

package main

import "fmt"

func main() {

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Printf("for strings: %v, anagram groups: %v", strs, groupAnagrams(strs))
}

/*
0. So I'm given a list of strings, and for all strings in str, group together the ones that are anagrams of each other
1. sequences of chars where order doesn't matter and chars can be duplicated within each sequence

with anagram_groups = []
while strs not empty:

	 with str = strs[0] and group = [strs[0]]
		 for j = i+1 in strs and strs[j] not in group:
			with str2 = strs[j]:
				S(str[len(str)], str2) = group ++ str2 and strs = strs \ str2 if str2 == ""
				S(str[idx], str2) = S(str[idx+1:], str2 \ {str[i]}) if str[i] in str2
		 anagram_groups ++ group


		 @todo - using freq maps:
		 ------------------------

with anagram_fm_groups = []
with anagram_groups = []
let freq_maps = {i, {char -> count | for all str in strs}}
while freq_maps not empty:

	 with fm = freq_maps[0] and group = {freq_maps[0]}
		 for j = 1 in freq_maps:
			with fm2 = freq_maps[j]:
				compare entrySets for fm/fm2
				if identical -> add to group and remove fm2 from freq_maps
		 anagram_fm_groups ++ group

populate anagram_groups from anagram_fm_groups
*/
type FM struct {
	freq_map map[string]int
	index    int
}

func groupAnagrams(strs []string) [][]string {
	// build freq_maps
	freq_maps := []FM{}
	for i, s := range strs {
		freq_map := map[string]int{}
		for _, c := range s {
			if _, ok := freq_map[string(c)]; !ok {
				freq_map[string(c)] = 1
			} else {
				freq_map[string(c)]++
			}
		}
		freq_maps = append(freq_maps, FM{
			freq_map,
			i,
		})
	}
	anagram_fm_groups := [][]FM{}
	grouped := map[int]bool{}
	for i := 0; i < len(freq_maps); i++ {
		if _, ok := grouped[i]; !ok {
			fm := freq_maps[i].freq_map
			group := []FM{freq_maps[i]}
			grouped[i] = true
		next:
			for j := i + 1; j < len(freq_maps); j++ {
				if _, ok := grouped[j]; !ok {
					fm2 := freq_maps[j].freq_map
					if len(fm) == len(fm2) {
						for k, v := range fm {
							if v2, ok := fm2[k]; ok {
								if v2 != v {
									continue next
								}
							} else {
								continue next
							}
						}
						group = append(group, freq_maps[j])
						grouped[j] = true
					}
				}
			}
			anagram_fm_groups = append(anagram_fm_groups, group)
		}
	}
	anagram_groups := [][]string{}
	for _, group := range anagram_fm_groups {
		anagram_group := []string{}
		for _, fm := range group {
			anagram_group = append(anagram_group, strs[fm.index])
		}
		anagram_groups = append(anagram_groups, anagram_group)
	}
	return anagram_groups
}
