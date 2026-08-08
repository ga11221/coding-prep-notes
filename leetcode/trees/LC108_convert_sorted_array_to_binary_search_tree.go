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

/*
6-rung ladder + correctness (see ladder.md - the generalized 6-step ladder).

Approach: build(nums, lo, hi) makes the middle element of the sorted segment
[lo..hi] the root of the segment's subtree, then recurses on the left half
[lo..mid-1] and right half [mid+1..hi]. One uniform rule (mid = lo + (hi-lo)/2)
is used everywhere - no per-side bookkeeping needed.

Rung 1 - ENUMERATE (ground truth):

	the ground truth is any BST over the n distinct keys whose inorder
	traversal equals nums and whose every node has subtree heights differing
	by <= 1. Many such trees exist (the problem accepts multiple outputs);
	we need to produce one, so the enumeration is over valid root choices,
	and any middle pick is a valid one.

Rung 2 - NAME THE OBJECT:

	existence/construction question: build ONE height-balanced BST for nums.
	The object is the recursive decomposition in which a sorted segment
	[lo..hi] maps to a node, its left subtree to the smaller keys, its right
	subtree to the larger keys.

Rung 3 - COMPRESS THE SPACE (representation axis):

	the strict ordering of nums is the ordering device: for any root chosen
	from the segment, the BST property forces exactly the keys on one side
	of the root into the left subtree and the rest into the right subtree.
	The segment [lo..hi] IS the candidate set in compact form - no per-node
	validation needed, the sorted order is the losslessness certificate
	(the left half holds precisely the keys smaller than the root, the right
	half precisely the larger ones, disjoint and exhaustive).

Rung 4 - COLLAPSE TO A STATISTIC (aggregation axis):

	per subtree, the statistic is its HEIGHT. Balance requires |hL - hR| <= 1
	at every node. Picking mid = lo + (hi-lo)/2 makes the two halves differ
	in size by at most 1, and since the rule is applied recursively to both
	halves, every subtree is balanced by induction - height is O(log n).

Rung 5 - FIX A COMPUTATION ORDER (time axis):

	pre-order: place the root of a segment, then build its left subtree, then
	its right. The recursion consumes each element exactly once and never
	revisits it -> O(n) time, O(log n) stack depth. Ordering by segment
	length gives the top-down mid-first placement.

Rung 6 - PROVE AN INVARIANT:

	invariant: build(nums, lo, hi) returns a height-balanced BST whose inorder
	traversal is exactly nums[lo..hi].
	Soundness (no false positives): the returned tree is a valid BST - its
	inorder is nums, which is strictly increasing, so the BST ordering holds -
	and it is height-balanced by Rung 4's size argument, inductively. The
	output therefore always satisfies the required property.
	Completeness (no false negatives): the recursion partitions the segment
	into root + left segment + right segment, disjoint and exhaustive, with
	base case lo > hi. Hence every element nums[i] is placed in exactly one
	node - none can be dropped (the historical failure mode) and none can be
	duplicated. Termination: each recursive call operates on a strictly
	smaller segment (mid is strictly inside for lo < hi), so the depth is
	O(log n) and the tree is complete in the sense of covering all n keys.
	Correctness follows by induction on the segment length.
	The lower/upper middle choice is cosmetic: either pick still yields two
	halves differing by <= 1 element, so both are height-balanced and both
	accepted.

Complexity: O(n) time, O(log n) stack space (excluding the output tree).
*/
