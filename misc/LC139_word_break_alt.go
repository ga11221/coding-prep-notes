//go:build ignore

package main

// @TODO - this is after 45 mins

import (
	"sort"
	"strings"
)

/*
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.
*/

/*
map first letter of every word in wordDict to all words in wordDict beginnning with that letter
eg1 wordDict = [leet, code]
 l -> [leet]
 c -> [code]

eg2 
s = catsandog
find largest word(s) starting with c matching s[0:i] - cats
repeat, finding match for s[i+1:j] - and
repeat for "og" - no match - back to s[i+1] - find next largest beginnning with a
@todo: is there a failure function that can save on backtracking ???

wordDict = ["cats","dog","sand","and","cat"]
  c -> [cats, cat] 
  d -> [dog]
  s -> [sand]
  a -> [and]
  ---- or ----
  c -> [nil, nil, nil, [cat], [cats]] // bitmapped by len of each word
  d -> [nil, nil, nil, [dog]]
  ...
   


How large is dictionary?
wordDict sorted? I can sort it in O(log n) if not


examples:

"catsandog", wordDict = ["cats","dog","and","cat"]
"cats" - "cat" = "s" (find words in wordDict starting with a "s")

"catsa ndog"

"ca tsa"

sorted= ["and", "cat", "cats", "dog", "sand"]
rev_sorted= ["sand", "dog", "cats", "cat", "and"]
false
in reverse sorted find first occurrence of word starting with s[0]
find largest string in wordDict whose first letter = s[0] and is contained in s
how to find largest match in wordDict? sort in reverse and lookup
cats largest match
and next largest
og - doesn't match
backtrack to "and" and drop letter "d" - now looking for "an" - no match - drop until empty string -  backtrack to "cats" - drop "s" - now looking "cat" return false only when first substring is empty

if you start at largest match, no need to traverse wordDict further
"cats" "an" "dog"

"catsanddog", wordDict = ["cats","dog","sand","and","cat"]

true

sort dictionary regardless
*/

/*
	                   catsandog
					      |

1.cats, andog   2. cat, sandog      3. ca,tsandog    4. c, atsandog  5. "",catsandog

	  |		            |                    |             |
	andog         	sandog             	atsandog		  catsandog
*/
func wordBreak(s string, wordDict []string) bool {
	sort.Sort(sort.Reverse(sort.StringSlice(wordDict)))
	// find largest prefix
	largestPrefix := findLargestPrefix(s, wordDict)
	if largestPrefix == "" {
		return false
	}
	for i := len(largestPrefix); i >= 0; i-- {
		if !findNext(s[len(largestPrefix):], wordDict) == true {
			return false
		}
	}
	return true
}

// "catsandog", wordDict = ["cats","dog","sand","and","cat"]
// sorted= ["and", "cat", "cats", "dog", "sand"]
// rev_sorted= ["sand", "dog", "cats", "cat", "and"]
func findLargestPrefix(s string, wordDict []string) string {
	for _, word := range wordDict {
		if strings.HasPrefix(s, word) {
			return word // returns first match - largest prefix in s
		}
	}
	return ""
}

// "catsandog", wordDict = ["cats","dog","sand","and","cat"]
// sorted= ["and", "cat", "cats", "dog", "sand"]
// rev_sorted= ["sand", "dog", "cats", "cat", "and"]
func findNext(substring string, wordDict []string) string {
	for _, word := range wordDict { // have largest prefix, so second largest might not be directly adjacent, start search from largest prefix index in wordDict to first word in word dict whose first letter does not match first letter of largestPrefix
		if substring[0] == word[0] && strings.EqualFold(substring, word) {
			return word
		}
	}
	return ""
}
