package main

import "fmt"

/*
347. Top K Frequent Elements (Medium)

Given an integer array nums and an integer k, return the k most frequent elements.
You may return the answer in any order.

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Approach: bucket sort — O(n)
- freq map: O(n)
- bucket array of size n+1 (index = frequency): O(n)
- collect from bucket[n] downward until k elements: O(n)
*/

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))      // [1 2]
	fmt.Println(topKFrequent([]int{1}, 1))                     // [1]
	fmt.Println(topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2)) // [-1 2] or [2 -1]
}

func topKFrequent(nums []int, k int) []int {
	freq_map := map[int]int{}
	for _, n := range nums {
		freq_map[n]++
	}
	buckets := make([][]int, len(nums)+1)
	for n, freq := range freq_map {
		buckets[freq] = append(buckets[freq], n)
	}
	topK := []int{}
	for i := len(nums); k > 0; i-- {
		if buckets[i] != nil {
			topK = append(topK, buckets[i]...)
			k -= len(buckets[i])
		}
	}
	return topK
}
