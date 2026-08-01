/*
167. Two Sum II - Input Array Is Sorted (Medium)

Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order,
find two numbers such that they add up to a specific target number. Let these two numbers
be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an integer array
[index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same
element twice.

Your solution must use only constant extra space.

Example 1:
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2.

Example 2:
Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore, index1 = 1, index2 = 3.

Example 3:
Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore, index1 = 1, index2 = 2.

Constraints:
- 2 <= numbers.length <= 3 * 10^4
- -1000 <= numbers[i] <= 1000
- numbers is sorted in non-decreasing order.
- -1000 <= target <= 1000
- The tests are generated such that there is exactly one solution.
*/

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSumSorted(nums, target)) // [1, 2]

	nums2 := []int{2, 3, 4}
	target2 := 6
	fmt.Println(twoSumSorted(nums2, target2)) // [1, 3]

	nums3 := []int{-1, 0}
	target3 := -1
	fmt.Println(twoSumSorted(nums3, target3)) // [1, 2]
}

/*
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2.

f(i,j) =

	         | f(i+1,j) if sum < target
	         | f(i,j-1) if sum > target
			 | (i+1, j+1)


			GenPairs(i, 1) = [nums[i]]
			GenPairs(i, n) = [nums[i]] ++ [c | c in GenPairs(nums[i+1:], n-1)] for i: 0->n-1
*/
func twoSumSorted(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i <= j; {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}
		if sum > target {
			j--
		} else {
			i++
		}
	}
	return nil
}
