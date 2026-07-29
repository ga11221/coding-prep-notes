package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2})) // 1
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2})) // 0
	fmt.Println(findMin([]int{11, 13, 15, 17})) // 11
}

/*
Suppose an array of length n sorted in ascending order is rotated between 1 and n times.
Notice that rotation [0, 1, 2, 4, 5, 6, 7] might become [4, 5, 6, 7, 0, 1, 2]
if it was rotated 4 times.

Given the sorted rotated array nums of unique elements, return the minimum element of this array.

You must write an algorithm that runs in O(log n) time.
*/

func findMin(nums []int) int {
	// TODO
	return -1
}
