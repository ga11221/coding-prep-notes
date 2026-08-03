package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))       // 1
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2})) // 0
	fmt.Println(findMin([]int{11, 13, 15, 17}))      // 11
}

/*
Suppose an array of length n sorted in ascending order is rotated between 1 and n times.
Notice that rotation [0, 1, 2, 4, 5, 6, 7] might become [4, 5, 6, 7, 0, 1, 2]
                                                         0  1  2  3  4  5  6
															left = 0, mid = 3, right = 6
															left = 3, mid = 4, right = 6
															left = 3, mid = 3, right = 4
															left = 3, mid = 3, right = 4
														[3, 4, 5, 1, 2]
                                                         0  1  2  3  4
															left = 0, mid = 2, right = 4
															left = 2, mid = 3, right = 4
															left = 2, mid = 2, right = 3
														[0, 1, 2, 4, 5, 6 ]
                                                         0  1  2  3  4  5
															left = 0, mid = 2, right = 5
															left = 0, mid = 1, right = 2
															left = 0, mid = 0, right = 1

if it was rotated 4 times.

Given the sorted rotated array nums of unique elements, return the minimum element of this array.

You must write an algorithm that runs in O(log n) time.
*/

func findMin(nums []int) int {
	mid := len(nums) / 2
	return _findMin(nums, 0, mid, len(nums)-1)
}
func _findMin(nums []int, left, mid, right int) int {
	if left == mid {
		if nums[left] < nums[right] {
			return nums[left]
		}
		return nums[right]
	}
	if nums[mid] > nums[right] {
		return _findMin(nums, mid, ((right-mid)/2)+mid, right)
	}
	return _findMin(nums, left, ((mid-left)/2)+left, mid)
}
