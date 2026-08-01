package main

import "math"

/*
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.

bitmap (account for negatives by shifting - ie 10^-9 placed in slot 0 and 0 would be placed in slot 10^9)
1. populate bitmap from nums
2. max = 0
3. from i -> 0..len(bitmap) find each gap
4. for each gap: max = math.max(max, start of gap - start of non-zero)
Use pointer to int to allow for nils in bitmap - makes finding gaps easier

-2 < n < 2
[0,1,2,3,4]
*/
/* [100,4,200,1,3,2]
[true,true,true,true,false,...,false,....,true]

*** need to account for duplicates **** look at last test case
*** should have asked about duplicates before starting
@TODO
[]bool -> [][]bool to account for duplicates

*/
func longestConsecutive(nums []int) int {
	// add 1 + 2
	offset := int(math.Pow(10, 9))
	size := 2 * int(math.Pow(10, 9))
	bitmap := make([]bool, size) // all elems should init to false
	for _, v := range nums {     // o(n)
		bitmap[offset+v] = true
	}
	nonGapStart := 0
	maxm := 0
	gap := false
	for i, v := range bitmap { // i =0, v = 1 t ->o(n) space->o(n)
		if v {
			if gap {
				nonGapStart = i
				gap = false
			}
		} else {
			if !gap {
				gap = true
				nonGapWidth := i - nonGapStart
				if nonGapWidth > maxm {
					maxm = nonGapWidth
				}
			}
		}
	}
	return maxm
}
