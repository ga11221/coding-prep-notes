package main

import "fmt"

/*
You are given an array of integers nums,
there is a sliding window of size k which is moving from the very left of the array to the very right.
You can only see the k numbers in the window. Each time the sliding window moves right by one position.

Return the max sliding window.

Example 1:

Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
when to update max?
populate bitmap using first window and find max in first window
in next window, is dropped number equal to max?
set bitmap[droppedNumber]=0
set bitmap[newNumber]=1
   yes
		is added number greater than or equal to max?
			yes -> update max set bitmap[newNumber] = 1
			no -> find max [starting at old max in bitmap, move left until you find first entry = 1]
	no
		is added number greater than max?
			yes -> update max
			no -> advance window
Explanation:
Window position                Max
---------------               -----
[1  3  -1] -3  -5  3  6  7       3 // min = -1 max = 3 , place 1, 3, -1 in bitmap
 1 [3  -1  -3] -5  3  6  7       3 // drop 1 is it greater than equal to max -> no less than min -> no
								  // add -3 is it greater than max -> no less than min -> yes
									// return max
 1  3 [-1  -3  -5] 3  6  7       5 // drop 3 is it greater than or equal to max -> yes
								  // add 5 is it greater than or equal to max -> yes less than min -> yes
									// return max
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
Example 2:

Input: nums = [1], k = 1
Output: [1]


Constraints:

1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length
*/

func main() {
	//fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1}, 1))
}

/*
Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
when to update max?
populate bitmap using first window and find max in first window
in next window, is dropped number equal to max?
set bitmap[droppedNumber]=0
set bitmap[newNumber]=1

	   yes
			is added number greater than or equal to max?
				yes -> update max
				no -> find max [starting at old max in bitmap, move left until you find first entry = 1]
		no
			is added number greater than max?
				yes -> update max
				no -> advance window
*/

func maxSlidingWindow(nums []int, k int) []int {
	var bitmap = [20001]int{}
	var m = nums[0]
	bitmap[m+10000]++
	for i := 1; i < k; i++ {
		if nums[i] > m {
			m = nums[i]
		}
		bitmap[nums[i]+10000]++
	}
	var ms = []int{m}
	for i := 1; i < len(nums)-k+1; i++ {
		droppedNumber := nums[i-1]
		bitmap[droppedNumber+10000]--

		addedNumber := nums[i-1+k]
		bitmap[addedNumber+10000]++

		if droppedNumber == m {
			if addedNumber >= m {
				m = addedNumber
				ms = append(ms, m)
			} else {
				for j := m; ; j-- {
					if bitmap[j+10000] > 0 {
						m = j
						ms = append(ms, m)
						break
					}
				}
			}
		} else {
			if addedNumber > m {
				m = addedNumber

			}
			ms = append(ms, m)
		}
	}
	return ms
}
