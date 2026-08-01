package main

// arrays/strings:
// sliding window
// pointer arithmetic
// bit array/buckets
// sort first
// dp
// working backwards
// maps/sets (for strings)
// queues (for strings)
// search (binary = log(n))

// [1,3,5,6] target=2
func searchInsert(nums []int, target int) int {
	// binary search
	last := len(nums) - 1 // 3
	first := 0
	if target > nums[last] { //2>6?
		return last + 1
	}
	if target < nums[first] { //2 < 1?
		return first
	}
	midpt := (last - first) / 2 // 3-0/2 = 1
	if target > nums[midpt] {   // 2 > nums[1] (3)? no
		return bs(nums, target, (midpt)+1, last)
	}
	return bs(nums, target, first, midpt) //bs(nums, 2, 0, 1)
}

// bs(nums, 2, 0, 1)
func bs(nums []int, target, first, last int) int {
	if last-first == 0 {
		if target <= nums[first] {
			return first
		}
		return first + 1
	}
	midpt := ((last - first) / 2) + first // 1-0/2 + 0 = 0
	if target > nums[midpt] {             // 2 > nums[0](1)yes
		return bs(nums, target, midpt+1, last) //bs(nums, 2,1,1)
	}
	return bs(nums, target, first, midpt) //bs(nums,3,3)
}
