package main

import (
	"math"
	"slices"
)

/*
You are part of a university admissions office
and need to keep track of the kth highest test score from applicants in real-time.
This helps to determine cut-off marks for interviews and admissions dynamically as new applicants submit their scores.

You are tasked to implement a class which, for a given integer k,
maintains a stream of test scores and continuously returns the kth highest test score
after a new score has been submitted.
More specifically, we are looking for the kth highest score in the sorted list of all scores.

Implement the KthLargest class:

KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of test scores nums.
int add(int val) Adds a new test score val to the stream and

	returns the element representing the kth largest element in the pool of test scores so far.

Example 1:

Input:
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]

Output: [null, 4, 5, 5, 8, 8]

Explanation:

KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]); [2,4,5,8]
kthLargest.add(3); // return 4  [2,3,4,5,8]
kthLargest.add(5); // return 5 [2,3,4,5,5,8]
kthLargest.add(10); // return 5 [2,3,4,5,5,8,10]
kthLargest.add(9); // return 8 [2,3,4,5,5,8,9,10]
kthLargest.add(4); // return 8 [2,3,4,4,5,5,8,9,10]

Example 2:

Input:
["KthLargest", "add", "add", "add", "add"]
[[4, [7, 7, 7, 7, 8, 3]], [2], [10], [9], [9]]

Output: [null, 7, 7, 7, 8]

Explanation:

KthLargest kthLargest = new KthLargest(4, [7, 7, 7, 7, 8, 3]);[3,7,7,7,7,8] -> [7,7,7,8]
kthLargest.add(2); // return 7 [2,3,7,7,7,7,8] -> [7,7,7,8] ->7
kthLargest.add(10); // return 7 [2,3,7,7,7,7,8,10] [7,7,8,10] -> 7
kthLargest.add(9); // return 7 [2,3,7,7,7,7,8,9,10] [7,8,9,10]-> 7
kthLargest.add(9); // return 8 [2,3,7,7,7,7,8,9,9,10] [8,8,9,10] -> 8
*/
func mainjiai2hi91jajaaaa() {
	/*
		var kthLargest = Constructor(3, []int{4, 5, 8, 2})
		println(kthLargest.Add(3))  // return 4  [2,3,4,5,8]
		println(kthLargest.Add(5))  // return 5 [2,3,4,5,5,8]
		println(kthLargest.Add(10)) // return 5 [2,3,4,5,5,8,10]
		println(kthLargest.Add(9))  // return 8 [2,3,4,5,5,8,9,10]
		println(kthLargest.Add(4))  // return 8 [2,3,4,4,5,5,8,9,10]
	*/

	// [[1,[]],[-3],[-2],[-4],[0],[4]]
	/*
		var kthLargest = Constructor(1, []int{})
		println(kthLargest.Add(-3)) // return -3  [-3]
		println(kthLargest.Add(-2)) // return -2 [-3,-2]
		println(kthLargest.Add(-4)) // return -2 [-4, -3,-2]
		println(kthLargest.Add(0))  // return  0 [-4, -3,-2,0]
		println(kthLargest.Add(4))  // return 4 [-4, -3,-2,0,4]
	*/
	// [[2,[0]],[-1],[1],[-2],[-4],[3]]
	/*
		var kthLargest = Constructor(2, []int{0}) //[0]
		println(kthLargest.Add(-1))               // return -1  [-1, 0]
		println(kthLargest.Add(1))                // return 0 [-1, 0, 1]
		println(kthLargest.Add(-2))               // return 0 [-2, -1, 0, 1]
		println(kthLargest.Add(-4))               // return  0 [-4, -2, -1, 0, 1]
		println(kthLargest.Add(3))                // return 1 [-4, -2, -1, 0, 1, 3]
	*/
	// [[3,[1,1]],[1],[1],[3],[3],[3],[4],[4],[4]]
	/*
		var kthLargest = Constructor(3, []int{1, 1}) //[1,1]
		println(kthLargest.Add(1))                   // return 1  [1,1,1]
		println(kthLargest.Add(1))                   // return 1 [1,1,1,1]
		println(kthLargest.Add(3))                   // return 1 [1,1,1,1,3]
		println(kthLargest.Add(3))                   // return  1 [1,1,1,1,3,3]
		println(kthLargest.Add(3))                   // return 3 [1,1,1,1,3,3,3]
		println(kthLargest.Add(4))                   // return 3 [1,1,1,1,3,3,3,4]
		println(kthLargest.Add(4))                   // return 3 [1,1,1,1,3,3,3,4,4]
		println(kthLargest.Add(4))                   // return 4 [1,1,1,1,3,3,3,4,4,4]
	*/

}

type KthLargest struct {
	k    int
	nums []int
}

func Constructor1(k int, nums []int) KthLargest {
	largest := KthLargest{
		k: k,
	}
	if len(nums) == 0 {
		largest.nums = nil
		return largest
	}
	slices.Sort(nums)
	largest.nums = make([]int, k)
	for i, j := len(nums)-1, k-1; i >= len(nums)-k && i >= 0; i, j = i-1, j-1 {
		largest.nums[j] = nums[i]
	}
	for i := 0; i < k-len(nums); i++ {
		largest.nums[i] = math.MinInt
	}
	return largest
}

func (this *KthLargest) Add(val int) int {
	if len(this.nums) == 0 {
		this.nums = make([]int, this.k)
		this.nums[0] = val
	} else if val > this.nums[this.k-1] {
		for i, j := 0, 1; j < this.k; i, j = i+1, j+1 {
			this.nums[i] = this.nums[j]
		}
		this.nums[this.k-1] = val
	} else if val > this.nums[0] {
		var insertIdx = this.binaryInsert(val, 0, this.k)
		for i, j := 0, 1; j < insertIdx; i, j = i+1, j+1 {
			this.nums[i] = this.nums[j]
		}
		this.nums[insertIdx-1] = val
	}
	return this.nums[0]
}

func (this *KthLargest) binaryInsert(val, start, end int) int {
	if end == start {
		return start
	}
	if val < this.nums[start+(end-start)/2] {
		return this.binaryInsert(val, start, start+(end-start)/2)
	}
	return this.binaryInsert(val, start+((end-start)/2)+1, end)
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
