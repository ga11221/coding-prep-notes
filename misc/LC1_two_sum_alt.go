package main

import "fmt"

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/

func main1() {
	//var nums = []int{2, 7, 11, 15}
	//var target = 9

	//var nums = []int{3, 2, 4}
	//var target = 6

	var nums = []int{3, 3}
	var target = 6

	fmt.Print(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	/*
	   	// split nums into 2 maps (k => integer, v => index in nums):
	      // 1. positive integers - pos
	      // 2. negative integers - neg

	*/
	var pos = make(map[int][]int)

	for idx, num := range nums {
		if _, ok := pos[num]; ok {
			pos[num] = append(pos[num], idx)
		} else {
			pos[num] = []int{idx}

		}
	}
	for k, v := range pos {
		var diff = target - k
		if _, ok := pos[diff]; ok {
			if k == diff && len(v) > 1 {
				return []int{v[0], v[1]}
			}
			if _, ok := pos[diff]; ok && pos[diff][0] != v[0] {
				return []int{v[0], pos[diff][0]}
			}
		}
	}
	return nil
}
