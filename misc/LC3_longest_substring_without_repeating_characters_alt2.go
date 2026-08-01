package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"a", 1},
		{"au", 2},
		{"dvdf", 3},
	}
	for _, tt := range tests {
		got := lengthOfLongestSubstring(tt.input)
		if got != tt.want {
			t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestLengthOfLongestSubstringFast(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"a", 1},
		{"au", 2},
		{"dvdf", 3},
	}
	for _, tt := range tests {
		got := lengthOfLongestSubstringFast(tt.input)
		if got != tt.want {
			t.Errorf("lengthOfLongestSubstringFast(%q) = %d, want %d", tt.input, got, tt.want)
		}
	}
}
