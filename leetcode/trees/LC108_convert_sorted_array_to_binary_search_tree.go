package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
LC108 Convert Sorted Array to Binary Search Tree (Array, Divide and Conquer,
Tree, BST, Binary Search)

Given an integer array nums where the elements are sorted in ascending order,
convert it to a height-balanced binary search tree.

A height-balanced binary tree is a binary tree in which the depth of the two
subtrees of every node never differs by more than one.

Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: [0,-10,5,null,-3,null,9] is also accepted.

Input: nums = [1,3]
Output: [3,1]
Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.

Constraints:
- 1 <= nums.length <= 10^4
- -10^4 <= nums[i] <= 10^4
- nums is sorted in a strictly increasing order.
*/

/*
height-balanced = every non-leaf level is full - every non-leaf node has a left/right child => no, that's a complete tree (which is also height-balanced)
complete = only leaf level may have missing children
insert in pre-order - ie  node, then it's left, then it's right
enumerate over nums - n integers sorted asc
state space is nums with pointer moving recursively through midpoint, then mid of left half, mid of right half
node is at midpoint of start and end, left child taken from |start - mid|/2 + start and right child taken from |end-mid|/2 + mid

[a, b, c, d, e]
start = 0 (0,2,4)
end = 4
mid = 2
N = c
L = |mid-start|/2 + start = |2-0|/2 + 0 = 1=> b
R = |end - mid+1|/2 + mid = |4-2+1|/2 + 2 = 3=> d

[a,e]

LH: (lower middle for LH) (0,0,1) last iter: start == mid
start = 0
end = mid -1 = 1
mid = |1-0|/2 + start  = 0+0 = 0
N = a
start == mid

RH: (upper middle for RH) (3,4,4) last iter: end == mid
start = mid+1 = 3
end = 4
mid = |4-3+1|/2 + start  = 1 + 3 = 4
*/

func sortedArrayToBST(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, lo, hi int) *TreeNode {
	if lo > hi {
		return nil
	}
	mid := lo + (hi-lo)/2
	return &TreeNode{
		Val:   nums[mid],
		Left:  build(nums, lo, mid-1),
		Right: build(nums, mid+1, hi),
	}
}

func sortedArrayToBST_alt(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return &TreeNode{
		Val:   nums[len(nums)/2],
		Left:  sortedArrayToBST_alt(nums[0 : len(nums)/2]),
		Right: sortedArrayToBST_alt(nums[len(nums)/2+1 : len(nums)]),
	}
}

func sortedArrayToBST_alt2(nums []int) *TreeNode {
	start := 0
	end := len(nums) - 1
	mid := lowerMid(start, end)
	this := nums[mid]
	root := &TreeNode{Val: this}
	_sortedArrayToBST_alt2(nums, root, "L", start, mid-1)
	_sortedArrayToBST_alt2(nums, root, "R", mid+1, end)
	return root
}

func _sortedArrayToBST_alt2(nums []int, parent *TreeNode, child string, start, end int) {
	if end < 0 || start > end {
		return
	}
	var mid int
	if child == "L" {
		mid = lowerMid(start, end)
	} else {
		mid = upperMid(start, end)
	}
	this := nums[mid]
	node := &TreeNode{Val: this}
	if child == "L" {
		parent.Left = node
	} else {
		parent.Right = node
	}
	_sortedArrayToBST_alt2(nums, node, "L", start, mid-1)
	_sortedArrayToBST_alt2(nums, node, "R", mid+1, end)
}

func sortedArrayToBST_alt3(nums []int) *TreeNode {
	start := 0
	end := len(nums) - 1
	return _sortedArrayToBST_alt3(nums, start, end)
}

func _sortedArrayToBST_alt3(nums []int, start, end int) *TreeNode {
	if end < 0 || start > end {
		return nil
	}
	mid := lowerMid(start, end)
	this := nums[mid]
	node := &TreeNode{Val: this}
	node.Left = _sortedArrayToBST_alt3(nums, start, mid-1)
	node.Right = _sortedArrayToBST_alt3(nums, mid+1, end)
	return node
}

func lowerMid(start, end int) int {
	return (end-start)/2 + start
}

func upperMid(start, end int) int {
	return (end-start+1)/2 + start
}
