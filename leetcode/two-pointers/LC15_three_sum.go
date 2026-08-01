package main

import (
	"fmt"
	"slices"
)

/*
15. 3Sum (Medium)

Given an integer array nums, return all triplets [nums[i], nums[j], nums[k]]
such that i != j, i != k, j != k, and nums[i] + nums[j] + nums[k] == 0.
No duplicate triplets.

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]

Approach: sort + fix one element + two-pointer on rest — O(n^2)
- sort: O(n log n)
- fix nums[i], two-pointer on remaining: O(n) per i → O(n^2)
- skip duplicates to avoid duplicate triplets

Notes:
- After finding a match (j+k == target), continue moving both pointers — don't break.
  Multiple pairs can sum to the same target for a fixed i.
- Dedup via map[3]int works but is O(k) space. More efficient: skip duplicates in
  outer loop (if nums[i] == nums[i-1], continue). Inner pointer dedup is handled
  naturally by moving j++ and k-- past duplicates.
- Return type [][3]int is fixed-size — Go idiom for known-length tuples. LeetCode
  expects [][]int (variable-length slices).
*/

func main() {
	fmt.Println(threeSum([]int{-4, -1, -1, 0, 1, 2})) //[-1,-1,2], [-1,0,1]
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
}

func threeSum(nums []int) [][3]int {
	slices.Sort(nums)
	triplets := map[[3]int]bool{}
	for i := 0; i < len(nums); i++ {
		first := 0 - nums[i]
		for j, k := i+1, len(nums)-1; j < k; {
			if nums[j]+nums[k] == first {
				triplets[[3]int{nums[i], nums[j], nums[k]}] = true
				j++
				k--
			} else if nums[j]+nums[k] > first {
				k--
			} else {
				j++
			}
		}
	}
	all_triplets := [][3]int{}
	for k := range triplets {
		all_triplets = append(all_triplets, k)
	}
	return all_triplets
}
