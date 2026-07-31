package main

import "fmt"

func main() {
	//nums := []int{100, 4, 200, 1, 3, 2}
	//nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	nums := []int{1, 0, 1, 2}
	fmt.Printf("longest consecutive sequence for %v is: %v", nums, longestConsecutive(nums))
}

/*
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.

Example 1:

Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

Example 2:

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9

Example 3:

Input: nums = [1,0,1,2]
Output: 3
*/
func longestConsecutive(nums []int) int { // alternative is to iterate through entire array and assume each elem is start of a new sequence (double-count)
	m := map[int]uint8{}
	for _, n := range nums {
		m[n] = 1
	}

	var max int
	for _, n := range nums {
		countLessThanN := all(0, n, n-1, &m)
		countGreaterThanN := all(0, n, n+1, &m)
		total := countLessThanN + countGreaterThanN + 1
		if total > max {
			max = total
		}
	}
	return max
}

func all(count int, sequenceStart int, num int, m *map[int]uint8) int {
	if _, ok := (*m)[num]; ok {
		count++
		if sequenceStart > num {
			delete(*m, num)
			return all(count, sequenceStart, num-1, m)
		} else {
			delete(*m, num)
			return all(count, sequenceStart, num+1, m)
		}
	}
	return count
}
