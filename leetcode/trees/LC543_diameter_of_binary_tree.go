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

func diameterOfBinaryTree(root *TreeNode) int {
	globalMax := 0
	// max path that goes through node N is max depth of N.left + max depth of N.right

	if root == nil {
		return 0
	}
	maxDepth(root, &globalMax)
	return globalMax
}

func maxDepth(node *TreeNode, globalMax *int) int {
	if node.Right == nil && node.Left == nil {
		return 0
	}
	var maxLeft, maxRight int
	if node.Right == nil {
		maxLeft = 1 + maxDepth(node.Left, globalMax)
	} else if node.Left == nil {
		maxRight = 1 + maxDepth(node.Right, globalMax)
	} else {
		maxLeft = 1 + maxDepth(node.Left, globalMax)
		maxRight = 1 + maxDepth(node.Right, globalMax)
	}
	pathLength := maxLeft + maxRight
	if pathLength > *globalMax {
		*globalMax = pathLength
	}
	if maxLeft > maxRight {
		return maxLeft
	}
	return maxRight

}
