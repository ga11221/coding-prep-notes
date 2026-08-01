package main

/*
Given an integer array nums, return true if you can partition the array into two subsets such that
the sum of the elements in both subsets is equal or false otherwise.

Constraints:

1 <= nums.length <= 200
1 <= nums[i] <= 100

*/

func main() {}

/**

if the dp recurrence relation is:
	dp[i][s] = dp[i-1][s] || dp[i-1][s - nums[i]]
And state def is:
	dp[i][s] = true if subset of nums[0..i] sum exactly to s
Base case:
	dp[0][s] = s == nums[0]
	dp[i][0] = true for all i (empty subset)

Work backwards to define recursive func...

    f(idx, target) {
		if target == 0 {
			return true
		}
		if idx == len(nums) {
			return false
		}
		if nums[idx] == target {
			return true
		}
		if nums[idx] < target {
			return f(idx+1, target-nums[idx])
		}
		return f(idx+1, target)
	}

*/
/**
Example 1:

Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].

Example 2:

Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.
*/

/*
partition - P1: nums[0] and P2: nums[1:]
if sum(P1) < sum(P2) - rebalance - move first from P2 to P1 and recompute sums
if sum(P1) now > sum(P2) -

---------------------------------------------------------------------------------------------
let idx be the start of the right partition
idx wouldn't work - can pick elements at any indices for right or left partition
f(idx) = true if sum(nums[0:idx]) == sum(nums[idx:])
f(0) = false for any non-empty or singleton array
f(1) = nums[0] == sum(nums[1:])
---------------------------------------------------------------------------------------------

all the distinct combinations that include the ith elem and have at-most n-1 elems:

	[1], [1,5], [1,5,11], [1,5,5], [1,11]
	[5], [5,11], [5,5], [5,11,5]
	[11]

	2d = [][]int
	[[1,2,3], [1,2,3], [1,2,3]]
	1 2 3
	1 2 3
	3d = [][][]int
	    [   [[1,21]]      , .... ]
	nums = [1,5,11,5]
	s = sum(nums)/2 = 22/2 = 11
	dp[s] = dp[s] || dp[s-nums[i]]
	dp[0] = true
	dp[11] = dp[11] || dp[11-1] = dp[10]
	dp[10] = dp[10] || dp[10-5] = dp[5]
	dp[5] = dp[5] || dp[5-

	dp[0] = [(1, 21)]
	dp[1] = (dp[0][0]+nums[1],dp[0][1]-nums[1]= (6, 16) for all i < 1: ?
	                                                         for every tuple in dp[i]:
															 	tuple[0]+nums[1] == tuple[1]-nums[1]?
	dp[2] = (dp[0][0]+nums[2],dp[0][1]-nums[2]= (12, 10) for all i < 2, dp[i][0]+nums[2] == dp[i][1]-nums[2]?
	        (dp[1][0]+nums[2],dp[1][1]-num[2] = (17, 5)
	dp[2] = [(12,10), (17,5)]
	dp[3] = 		...
	where does answer accumulate in dp?

	f([1,5,11,5]) -> P1: [[1]] ++ f([5,11,5])
	              -> P2: [[5,11,5]]

	f(0)          -> [1] ++ ...
	f([5,11,5]) -> [[1], [1,5]] ++ f([11,5])
	f(1)          -> [1, 6] ++ ...
	f([11,5]) -> [[1],	[1,5],	[1,11],	[1,5,11]] ++ f([5])
	f(2)	  -> [1, 6, 12, 17] ++ ...
	f([5]) 	  ->  P1:  [[1], 	   [1,5], 	[1,11],  [1,5,11], [1,5,5]]
			      P2:  [[5,11,5],  [5,11],  [5,5],   [5],      [11]]
	f(3) 	-> ...
*/
func canPartition(nums []int) bool {

}
