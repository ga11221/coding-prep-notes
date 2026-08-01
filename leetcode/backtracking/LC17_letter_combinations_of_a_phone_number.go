package main

import (
	"fmt"
)

/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

*/

func mainaisodjfiaoihoi() {
	var s = "323"
	fmt.Print(letterCombinations(s))
}

func letterCombinations(digits string) []string {
	// for each letter mapping to digits[i], append each letter corresponding to digits[i+1]
	var dialPad = map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	var combinations = dialPad[string(digits[0])]
	return _letterCombinations(digits[1:], combinations, dialPad)
}

func _letterCombinations(digits string, combinations []string, dialPad map[string][]string) []string {
	if digits == "" {
		return combinations
	}
	if _, ok := dialPad[string(digits[0])]; !ok {
		return _letterCombinations(digits[1:], combinations, dialPad)
	}
	var combos = []string{}
	for _, combination := range combinations {
		for _, letter := range dialPad[string(digits[0])] {
			combos = append(combos, combination+letter)
		}
	}
	return _letterCombinations(digits[1:], combos, dialPad)
}
