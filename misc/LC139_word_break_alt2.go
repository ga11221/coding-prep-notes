//go:build ignore

package main

import (
	"testing"
)

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "[100,4,200,1,3,2]",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "[0,3,7,2,5,8,4,6,0,1]",
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
		{
			name: "[1,0,1,2]",
			nums: []int{1, 0, 1, 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreak(tt.nums)
			if got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
