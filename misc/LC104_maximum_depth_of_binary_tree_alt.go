package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	globalMax := 0
	return _maxDepth(root, &globalMax) + 1
}

/*
*

		     3
		    / \
	       9  20
	 	      / \
	         15  7
*/
func _maxDepth(node *TreeNode, globalMax *int) int {
	var maxLeft, maxRight int
	if node.Left == nil && node.Right == nil {
		return 0 // 6: return 0; 5: return 0
	}
	if node.right == nil {
		maxLeft = 1 + _maxDepth(node.Left, globalMax) // 1: maxLeft = 1+_maxDepth(2, globalMax); 1b: 1+2 = 3
	} else if node.left == nil {
		maxRight = 1 + _maxDepth(node.Right, globalMax) // 4a: maxRight = 1 + _maxDepth(6, globalMax) => 4b: 1 + (0) = 1
	} else {
		maxLeft = 1 + _maxDepth(node.Left, globalMax)   // 2: maxLeft = 1+_maxDepth(4, globalMax); 2b: maxLeft = 1+1 =2
		maxRight = 1 + _maxDepth(node.Right, globalMax) // 2: maxRight = 1+_maxDepth(5, globalMax); 2b: maxRight = 1+0=1
	}
	if maxLeft > maxRight {
		if maxLeft > *globalMax {
			globalMax = &maxLeft // 2b: globalMax = 2;1b: globalMax = 3
		}
		return maxLeft // 2b: return 2;1b: return 3
	} else {
		if maxRight > *globalMax {
			globalMax = &maxRight // 4b: globalMax = 1
		}
		return maxRight // 4b: return 1
	}

}
