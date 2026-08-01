package main

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "nums = [1,3,5], target = 3",
			nums:   []int{1, 3, 5},
			target: 3,
			want:   1,
		},
		{
			name:   "nums = [1,3,5,6,10], target = 6",
			nums:   []int{1, 3, 5, 6, 10},
			target: 6,
			want:   3,
		},
		{
			name:   "nums = [1,3,5,6], target = 5",
			nums:   []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			name:   "nums = [1,3,5,6], target = 2",
			nums:   []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			name:   "nums = [1,3,5,6], target = 7",
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchInsert(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
