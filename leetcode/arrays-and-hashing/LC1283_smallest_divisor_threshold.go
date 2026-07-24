package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 5, 9}
	fmt.Println(smallestDivisor(nums, 6)) // 5

	nums2 := []int{44, 22, 33, 11, 1}
	fmt.Println(smallestDivisor(nums2, 5)) // 44
}

/*
Given an array of integers nums and an integer threshold, we will choose a
positive integer divisor, divide all the array by it, and sum the division's
result. Find the smallest divisor such that the result mentioned above is
less than or equal to threshold.

Each result of the division is rounded to the nearest integer greater than
or equal to that element. (For example: 7/3 = 3 and 10/2 = 5).


Binary search on divisor value in range [1, max(nums)].
- If sum(ceil(nums[i]/divisor)) > threshold → divisor too small, search right
- If sum(ceil(nums[i]/divisor)) <= threshold → divisor works, try smaller, search left
Answer is the smallest divisor where sum <= threshold.
*/

func smallestDivisor(nums []int, threshold int) int {
	left, right := 1, 0
	for _, n := range nums {
		if n > right {
			right = n
		}
	}

	for left < right {
		mid := left + (right-left)/2
		if divSum(nums, mid) > threshold {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func divSum(nums []int, divisor int) int {
	sum := 0
	for _, n := range nums {
		sum += int(math.Ceil(float64(n) / float64(divisor)))
	}
	return sum
}
