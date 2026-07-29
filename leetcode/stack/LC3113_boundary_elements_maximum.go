package main

import "fmt"

func main() {
	nums := []int{1, 4, 3, 3, 2}
	fmt.Println(numberOfSubarrays(nums)) // 6

	nums2 := []int{3, 3, 3}
	fmt.Println(numberOfSubarrays(nums2)) // 6

	nums3 := []int{1}
	fmt.Println(numberOfSubarrays(nums3)) // 1
}

/*
Given an array of positive integers nums, return the number of subarrays
where the first and last elements of the subarray are equal to the largest
element in the subarray.

duplicates in O(n)/O(n) dupe -> indices
for each dupe:
	pairwise scan for elem > boundaries O(n^3)

sliding window: O(n^2)/O(1)
	1:i <- 0..len(nums)
	j <- i+1
	if nums[j] < nums[i] advance j
	if nums[j] = nums[1] increment count
	else goto 1

Monotonic stack??
pop if next is larger...

f(i) = subarrays ending at i that begin with nums[i] and don't contain elems > nums[i] (how to carry state?)
f(0) =  []{elem}
f(1) = append nums[1] for all sub in f(0) where nums[1] < sub[first] .... (append but also track if sub[first] == nums[i] -> copy sub as it might continue eg 4,1,2,4,3,4)
Monotonic stack approach:
- Maintain a stack decreasing from bottom to top
- Each stack element is [value, count] — count of valid subarrays ending here with this value as boundary/max
- For each num:
  - Pop all elements smaller than num (they can't be max if num is in the subarray)
  - If stack is empty or top != num: push [num, 1]
  - If top == num: increment count (extend existing subarrays)
  - Add count to answer

O(n) time, O(n) space.
*/

func numberOfSubarrays(nums []int) int {
	// TODO
	return 0
}
