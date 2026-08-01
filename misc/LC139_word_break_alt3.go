package main

import (
	"fmt"
	"regexp"
)

/*
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

*/

func main2342j1jijoj() {
	var s = "catsandog"
	var wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	/*

			   cats -> andog
			   dog -> ""
			   sand ->og
			   and->og
			   cat->sandog

					var s = "applepenapple"
					var wordDict = []string{"apple", "pen"}

				apple->penapple
				pen->appleapple

					var s = "leetcode"
					var wordDict = []string{"leet", "code"}

				leet -> code
				code -> ""

					var s = "leetcodeend"
					var wordDict = []string{"leet", "code", "end"}

				leet->codeend
				code->end
				end->""

						var s = "leetcodend"
						var wordDict = []string{"leet", "code", "end"}
				leet->codend
				code->nd
				end->""
		1. map (m1) of first letter of every word in wordDict to corresponding word
		l -> leet
		c -> code
		e -> end

		2. beginning with first letter (l1) of s, find all words in wordDict beginning with l1 v1 = (m1[l1])
		3. find all words in v1 that match the beginning of s
		4. map each match to the substring (s1) of s that remains when the match is removed from the beginning
		5. repeat 2-4 for s1

	*/
	fmt.Print(wordBreak(s, wordDict))
}

/*
	var s = "leetcode"
	var wordDict = []string{"leet", "code"}
*/

func wordBreak(s string, wordDict []string) bool {
	var m1 = map[string][]string{}
	for _, word := range wordDict {
		var u = word[0:1]
		if v, ok := m1[u]; ok {
			m1[u] = append(v, word)
		} else {
			m1[u] = []string{word}
		}
	}
	s2 := s[0:1]
	if v, ok := m1[s2]; !ok {
		return false
	} else {
		for _, s3 := range v {
			compile := regexp.MustCompile(s3)
			find := compile.Find([]byte(s))
			if find == nil {
				return false
			}

			m1[string(find)] = []string{s[len(find):]}
		}
	}
	return false
}
