package main

import "fmt"

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(heights)) // 10

	heights2 := []int{2, 4}
	fmt.Println(largestRectangleArea(heights2)) // 4
}

/*
Given an array of integers heights representing the histogram's bar height
where the width of each bar is 1, return the area of the largest rectangle
in the histogram.

Monotonic stack approach.
*/

func largestRectangleArea(heights []int) int {
	// TODO
	return 0
}
