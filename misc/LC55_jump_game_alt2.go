package main

import (
	"testing"
)

func TestCanJump(t *testing.T) {
	t.Skip()
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "[8,2,4,4,4,9,5,2,5,8,8,0,8,6,9,1,1,6,3,5,1,2,6,6,0,4,8,6,0,3,2,8,7,6,5,1,7,0,3,4,8,3,5,9,0,4,0,1,0,5,9,2,0,7,0,2,1,0,8,2,5,1,2,3,9,7,4,7,0,0,1,8,5,6,7,5,1,9,9,3,5,0,7,5]",
			nums: []int{8, 2, 4, 4, 4, 9, 5, 2, 5, 8, 8, 0, 8, 6, 9, 1, 1, 6, 3, 5, 1, 2, 6, 6, 0, 4, 8, 6, 0, 3, 2, 8, 7, 6, 5, 1, 7, 0, 3, 4, 8, 3, 5, 9, 0, 4, 0, 1, 0, 5, 9, 2, 0, 7, 0, 2, 1, 0, 8, 2, 5, 1, 2, 3, 9, 7, 4, 7, 0, 0, 1, 8, 5, 6, 7, 5, 1, 9, 9, 3, 5, 0, 7, 5},
			want: true,
		},
		{
			name: "[3,4,3,1,0,7,0,3,0,2,0,3]",
			// unit: nonZeroSegment
			// state: pointer -> 0
			// if I can advance pointer one or more nonZeroSegments at a time until
			// pointer >= end of nums
			// f(nonZeroSegmenti) -> nonZeroSegmentj or len(nums) or x where j> i and x > len(nums)
			// build nonZeroSegment table...
			// can jump to nZS1 if i in nZS0 + nums[i] > nZS1(last) or > nZS1(first)
			// and if i in nZS0 > nZS1(last), check next nZS2
			/*
				check every index of each nonZeroSegment to determine if pointer can be advanced to a later nonZeroSegment or the end or beyond



				for each pointer starting at nonZeroSegment[i][j] for i from
			*/

			// from first nonZeroSegment, if I can get to a later nonZeroSegment
			nums: []int{3, 4, 3, 1, 0, 7, 0, 3, 0, 2, 0, 3},
			want: true,
		},
		{
			name: "[5,9,3,2,1,0,2,3,3,1,0,0]",
			//          0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11
			nums: []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0},
			want: true,
		},
		{
			name: "[2,3,1,1,4]",
			//          0, 1, 2, 3, 4
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		}, {
			name: "[3,0,8,2,0,0,1]",
			//          0, 1, 2, 3, 4, 5, 6
			nums: []int{3, 0, 8, 2, 0, 0, 1},
			want: true,
		}, {
			name: "[3,2,1,0,4]",
			//          0, 1, 2, 3, 4
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		}, {
			name: "[3,2,2,0,4]",
			//          0, 1, 2, 3, 4
			nums: []int{3, 2, 2, 0, 4},
			want: true,
		}, {
			name: "[1,0,0,2,0,4]",
			//          0, 1, 2, 3, 4, 5
			nums: []int{1, 0, 0, 2, 0, 4},
			want: false,
		}, {
			name: "[1,100,0,2,0,4]",
			//          0, 1,   2, 3, 4, 5
			nums: []int{1, 100, 0, 2, 0, 4},
			want: true,
		}, {
			name: "[0]",
			nums: []int{0},
			want: true,
		}, {
			name: "[0,1]",
			nums: []int{0, 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canJump(tt.nums)
			if got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
