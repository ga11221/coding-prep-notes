package main

/*
LC1403 Minimum Subsequence in Non-Increasing Order (Array, Greedy, Sorting)

Given the array nums, obtain a subsequence of the array whose sum of elements
is strictly greater than the sum of the non-included elements in the array.

If there are multiple solutions, return the subsequence with the minimum size.
If there still exist multiple solutions, return the subsequence with the
maximum total sum (which is necessarily the one with the minimum size among
the valid answers).

A subsequence of an array can be obtained by erasing some (possibly zero)
elements from the array. Note that the solution is guaranteed to be unique
under the given constraints.

Return the answer as a subsequence sorted in non-increasing order.

Input: nums = [4,3,10,9,8]
Output: [10,9]
Explanation: The subsequences [10,9] and [10,8] have a sum strictly greater
than the sum of the non-included elements (4 + 3 + 8 = 15, 4 + 3 + 9 = 16).
[10,9] is returned because it has the minimum size (2) and maximum sum (19).

Input: nums = [4,4,7,6,7]
Output: [7,7,6]
Explanation: The subsequences [7,7] and [7,6] have a sum strictly greater
than the sum of the non-included elements. [7,7,6] is returned because it
has the maximum sum (20) among the minimum-size valid subsequences (2 -> fails
strict inequality; the valid minimum size is 3).

Constraints:
- 1 <= nums.length <= 500
- 1 <= nums[i] <= 100
*/
func minSubsequence(nums []int) []int {
	total := 0
	bucket := make([]int, 501)
	for _, n := range nums {
		total += n
		bucket[n]++
	}
	retVal := []int{}
	subArraySum := 0
	for i := len(bucket) - 1; i >= 0; i-- {
		for bucket[i] > 0 {
			subArraySum += i
			total -= i
			retVal = append(retVal, i)
			if subArraySum > total {
				return retVal
			}
			bucket[i]--
		}
	}
	return retVal
}
