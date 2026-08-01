package main

// So, given two integers, n and k, n determines the size of the array
// the array contains non-repeating (distinct) positive numbers (>= 1)
// find the smallest set of n numbers that don't contain a pair of numbers
// that sum to k - return the sum of that smallest set

// n = 4 k = 8
// [1,2,3,4] - there are no two numbers in array that sum to k=8
// the sum of this array is 10

// n = 1 k = 1
// [2] = 2

// n = 1 k = 2
// [1] = 1

// n = 2 k = 2
// [1,2]

// n = 1000 k = 100
// [...48, 49, 50, 51, 52, ...]
// [...,48, 49, 50, skip the next 49, 100, 101, 102, ..., 1050]
// fill array halfway to k, then from k to n/2
// 1 <= n, k <= 50
// edge cases (?): n=1, n>k, k>n, n=k
func minimumSum(n int, k int) int {
	// greedy? always start with smallest num,
	// choose next smallest that doesn't sum to k
	// would have to check each new number with every
	// existing element in array O(n^2)

}
