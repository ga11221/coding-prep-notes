package main

// @TODO

import "math"

/*
*
	[1,2]

	f(0) = nums[0]
	f(1) = nums[1]+f(0)

	[-1,2]

	f(0) = nums[0]
	f(1) = max(nums[1]+f(0), nums[1])

	[-1,2, 3]

	f(0) = nums[0]
	f(1) = max(nums[1]+f(0), nums[1]) = max(2+(-1), 2) = max(1,2) = 2
	f(2) = max(nums[2]+f(1), nums[2]) = max(3+2, 3) = max(5,3) = 5

	f(i, j) =
		max(nums[i], nums[i+1:j], nums[i:j-1])
		this does examine every subarray (some more than once) - and it's a 2d state space
		think about collapsing into 1D

	
	base case = dp[0] = nums[0]

	Given an integer array nums, find the subarray with the largest sum, and return its sum.

	nums = [-2,1,-3,4,-1,2,1,-5,4] - sum (0, end) = 1 sum(1, end) = 3 sum
	output -> 6

	nums = [2,1,-3,4,-1,2,1,-5,4] - sum (0, end) = 5 sum(0, 6) = 6 sum
	output -> 6

	nums = [5,4,-1,7,8]
	output -> 23

subarray-	must be contiguous elems
two pointers to nums that start at 0
n^2 - increment i each time j reaches end - stop when i reaches end
dp[1][3]: sum (1...4)=2
tabulate sum at each index i in dp[i][j]
return maximum from dp

***should have had dp table diagrammed first***
edge cases - [], singleton

tail recursion approach - still n^2?

O(n) approach? kadane's algo - single pass O(1) space complexity - current max and global max
divide-conquer approach?
*/
func main() {}
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// [-2,-1]
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
		// [0][0] = -2
		// [1][1] = -1
		dp[i][i] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			//[0][1] = [0][0] + nums[1] = -2+-1 = -3

			dp[i][j] = dp[i][j-1] + nums[j]
		}
	}
	/**
	_____________
	| -2  |  -3 |
	_____________
	|  0  | 1   |
	*/
	max := math.MinInt32
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}
	return max
}
