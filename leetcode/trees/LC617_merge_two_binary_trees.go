package main

/*
LC617 Merge Two Binary Trees (Tree, DFS, BFS)

You are given two binary trees root1 and root2.

Imagine that when you put one of them to cover the other, some nodes of the two
trees are overlapped while the others are not. You need to merge the two trees
into a new binary tree. The merge rule is that if two nodes overlap, then sum
node values up as the new value of the merged node. Otherwise, the NOT null node
will be used as the node of the new tree.

Return the merged tree.

Note: The merging process must start from the root nodes of both trees.

Input: root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
Output: [3,4,5,5,4,null,7]

Input: root1 = [1], root2 = [1,2]
Output: [2,2]

Constraints:
- The number of nodes in both trees is in the range [0, 2000].
- -10^4 <= Node.val <= 10^4

Enumerate: all n+m nodes of root1 and root2
state space: { (N1, N2) | N1 in root1, N2 in root2 and N1 and N2 occupy same positions in their respective trees,

	ie same depth and order in sibling sequence }

compress space: repoint parent in tree1 or tree2 with nil child to corresponding node in other tree and do not proceed further
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	mergedRoot := &TreeNode{Val: root1.Val + root2.Val}
	_mergeTrees(root1, root2, mergedRoot)
	return mergedRoot
}

func _mergeTrees(node1 *TreeNode, node2 *TreeNode, mergedNode *TreeNode) {
	if node1.Left == nil {
		mergedNode.Left = node2.Left
	} else if node2.Left == nil {
		mergedNode.Left = node1.Left
	} else {
		mergedNode.Left = &TreeNode{Val: node1.Left.Val + node2.Left.Val}
		_mergeTrees(node1.Left, node2.Left, mergedNode.Left)
	}

	if node1.Right == nil {
		mergedNode.Right = node2.Right
	} else if node2.Right == nil {
		mergedNode.Right = node1.Right
	} else {
		mergedNode.Right = &TreeNode{Val: node1.Right.Val + node2.Right.Val}
		_mergeTrees(node1.Right, node2.Right, mergedNode.Right)
	}
}

/*
6-rung ladder + correctness (see ladder.md - the generalized 6-step ladder).

Approach: build a NEW tree (originals untouched). Where both nodes exist, the
merged node is a fresh node whose value is the sum; where only one exists, the
merged child is a REPOINT to that existing subtree (no new nodes, O(1)).

Rung 1 - ENUMERATE (ground truth):

	for every position occupied in either tree: if both trees have a node
	there, the result node = sum of the two; if exactly one has a node, the
	result keeps it; if neither, nothing. The result has exactly the union
	structure of the two trees.

Rung 2 - NAME THE OBJECT:

	the merged tree - a new tree whose nodes are PAIRS (n1, n2) of nodes at
	the same position in their respective trees, with value n1.Val + n2.Val
	(or the single existing node). Question type ALL (construction).

Rung 3 - COMPRESS THE SPACE (representation axis):

	merging a subtree with nil IS the identity: f(n, nil) = n and
	f(nil, n) = n. So instead of copying a one-sided subtree node by node,
	REPOINT to it - the whole subtree becomes the answer at O(1) cost.
	State = the current pair (n1, n2); the recurrence is local to the pair.

Rung 4 - COLLAPSE TO A STATISTIC (aggregation axis):

	the aggregation is SUM over the overlapping node values. Where one side
	is nil the "statistic" degenerates to the identity of the surviving
	subtree - reused wholesale, not summed node-by-node.

Rung 5 - FIX A COMPUTATION ORDER (time axis):

	preorder: root pair, then left pair, then right pair. Every pair is
	visited at most once (each node belongs to one pair), so one pass,
	O(n1 + n2) time, O(height) recursion stack.

Rung 6 - PROVE AN INVARIANT:

	invariant: at every call, mergedNode is the exact merge of node1 and
	node2 at the same position (its value = sum where both exist, its
	children = the merged children).
	Soundness: every created node's value is n1.Val + n2.Val of real nodes -
	no invented values - and every repointed child IS the correct merge of a
	one-sided position (identity rule). Nothing is fabricated.
	Completeness: the recurrence covers both children of every pair; the
	one-sided cases are terminal (no further nodes exist there), and the
	both-exist cases recurse. Every original node is either folded into a
	summed node or reused by pointer - none is dropped, none is visited
	twice. Hence the returned root is the correct merged tree by induction
	on tree structure.

Complexity: O(n1 + n2) time, O(height) stack space; new nodes allocated only
where both trees have a node.
*/
