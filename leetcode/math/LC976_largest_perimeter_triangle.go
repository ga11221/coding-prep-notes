package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/*
Given an integer array nums, return the largest perimeter of a triangle with a non-zero area, formed from three of these lengths. If it is impossible to form any triangle of a non-zero area, return 0.

Example 1:

Input: nums = [2,1,2]
sorted = [1,2,2]
Output: 5
Explanation: You can form a triangle with three side lengths: 1, 2, and 2.
Example 2:

Input: nums = [1,2,1,10]
sorted = [1,1,2,10]
Output: 0
Explanation:
You cannot use the side lengths 1, 1, and 2 to form a triangle.
You cannot use the side lengths 1, 1, and 10 to form a triangle.
You cannot use the side lengths 1, 2, and 10 to form a triangle.
As we cannot use any three side lengths to form a triangle of non-zero area, we return 0.

(a, b, c) make triangle if a+b>c AND b+c>a AND a+c>b
if sorted so that a <= b <= c, then a+b > c

C(n, 3)
C(3,3) = 1

pick(nums, k) = for i <= length(nums) - k:

	{ {nums[i]} U c for all c in pick(nums[i+1:], k-1) }

pick(nums, 0) = {{}}

pick([2,1,10], 2) = 2 U 1 or 2 U 10

f(0) = 0
f(1) = 0
f(2) = 0
f(3) = a + b + c if a + b > c where c = max(nums) else 0
f(4) =
*/
func largestPerimeter(nums []int) int {

}
