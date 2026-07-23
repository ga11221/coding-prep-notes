package main

func main() {}

/*
You are given a 0-indexed array of integers nums, and an integer target.

Return the length of the longest subsequence of nums that sums up to target.
If no such subsequence exists, return -1.

A subsequence is an array that can be derived from another array by deleting
some or no elements without changing the order of the remaining elements.

		Example 1:

		Input: nums = [1,2,3,4,5], target = 9
		Output: 3
		Explanation: There are 3 subsequences with a sum equal to 9: [4,5], [1,3,5], and [2,3,4].
		The longest subsequences are [1,3,5], and [2,3,4]. Hence, the answer is 3.

		Example 2:

		Input: nums = [4,1,3,2,1,5], target = 7
		Output: 4
		Explanation: There are 5 subsequences with a sum equal to 7: [4,3], [4,1,2], [4,2,1], [1,1,5], and [1,3,2,1].
		The longest subsequence is [1,3,2,1]. Hence, the answer is 4.

		Example 3:

		Input: nums = [1,1,5,4,5], target = 3
		Output: -1
		Explanation: It can be shown that nums has no subsequence that sums up to 3.

		----------------------------------------------------------------------------
		longest Sequence totaling target and starting at nums[i] =
		    1 + longest sequence totaling target - nums[i] and starting at nums[i+1]
			            -------------OR--------------
			longest sequence totaling target and starting at nums[i+1] // had left out the skip branch/choice

	    Need **skip option and ***guard for when elems exceed target

		f(i, remainingTarget) = max(
			f(i+1, remainingTarget), // **skip nums[i]
			1 + f(i+1, remaining - nums[i]) if nums[i] <= remaining // ***guard
		)
		and base case(s): f(n, 0) = 0, f(n, r>0) = -inf (impossible)

		recurrence relation says:
			"longest subsequence starting at i summing to target is the max of either:
				a. longest subsequence starting at i+1 summing to target, or
				b. 1 + longest subsequence starting at i+1 summing to target - nums[i]"

		state space = (i, remainingTarget)
		dp[n][0] = 0
		dp[i][target] = max(
			dp[i+1][target],
			1 + dp[i+1][target-nums[i]]
		)
		dp[i][r] = longest subsequence starting at i summing to r

		// @todo: why iterate 0..target for every num???
		for i = n-1..0:
			for r = 0..target:
				dp[i][r] = dp[i+1][r]
				if nums[i] <= target:
					dp[i][r] = max(dp[i][r], 1+dp[i+1][r-nums[i]])

		To get the sequence itself, walk dp:
         if dp[i][r] == dp[i+1][r]:
		 	// nums[i] was skipped - move to i+1 with same r
		  else:
		    // nums[i] was taken
			result = append(result, nums[i])
			r -= nums[i]
			// move to i+1

	    ----------------------------------------------------------------------------
		f(arr=nums[1:], t=target-nums[0], seq=[]int{nums[0]}) // don't need the seq itself, just length of longest seq
		if t < 0, remove last item from seq
		if arr == [], -- unfinished recursive def --
*/
func lengthOfLongestSubsequence(nums []int, target int) int {

}
