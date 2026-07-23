package main

import "math"

func main() {
	println(singleNumber([]int{2, 2, 1}))
	println(singleNumber([]int{4, 1, 2, 1, 2}))
	println(singleNumber([]int{1}))
}

func singleNumber(nums []int) int {
	// let s=nums[len(nums):] be the set of all distinct integers seen at the ith iteration
	// for every int x in nums, add to s (ie append to nums) if not in s
	// if x in s, replace it with sentinel
	// the unique number is that element in s that is not the sentinel

	length := len(nums)
	var x int
	dupe := false
	for i := 0; i < length; i++ {
		x = nums[i]

		for j := length; j < len(nums); j++ {
			if x == nums[j] {
				nums[j] = math.MaxInt
				dupe = true
				break
			}
		}
		if !dupe {
			nums = append(nums, x)
		}
		dupe = false
	}
	for i := length; i < len(nums); i++ {
		if nums[i] != math.MaxInt {
			return nums[i]
		}
	}
	return -1
}
