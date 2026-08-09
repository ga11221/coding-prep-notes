package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
LC112 Path Sum (Tree, DFS, Binary Tree)

Given the root of a binary tree and an integer targetSum, return true if the
tree has a root-to-leaf path such that adding up all the values along the path
equals targetSum.

A leaf is a node with no children.

Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
Output: true
Explanation: The root-to-leaf path with the target sum is shown.

Input: root = [1,2,3], targetSum = 5
Output: false
Explanation: There are two root-to-leaf paths in the tree:
(1 --> 2): The sum is 3.
(1 --> 3): The sum is 4.
There is no root-to-leaf path with sum = 5.

Input: root = [], targetSum = 0
Output: false
Explanation: Since the tree is empty, there are no root-to-leaf paths.

Constraints:
- The number of nodes in the tree is in the range [0, 5000].
- -1000 <= Node.val <= 1000
- -1000 <= targetSum <= 1000
*/
/*
 rung 1: enumerate
   all paths from root to leaves, leaf has no children

 rung 2: name object
 	exists(sum(P0-Pend)==targetSum) where P0 == root and Pend is a leaf in tree

  rung 3: compress (representation axis)
    rather than find all paths first, check running sum during DFS against targetSum
	exists => early exit

  rung 4: collapse to stat (aggregation axis)
  	running total that can be carried into 1 or 2 branches dependeding on number of children

  rung 5: FIX A COMPUTATION ORDER (time axis)
    pre-order - on the way to finding all paths, find first that sums to targetSum - pruning not possible as there can be negative numbers

  rung 6: INVARIANT
     if/when terminates, is the result it returns correct?
	 	summing until leaf reached, and comparing with target
		BFS will reach all leaves - no invalid result returned
	 does it return all possible valid results?
	    DFS will try all possible paths before determining targetSum does not exist
		always returns true on encountering first witness

*/
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if isLeaf(root) {
		return root.Val == targetSum
	}
	if root.Left == nil {
		return dfs(root.Right, targetSum-root.Val)
	}
	if root.Right == nil {
		return dfs(root.Left, targetSum-root.Val)
	}
	return dfs(root.Left, targetSum-root.Val) || dfs(root.Right, targetSum-root.Val)
}

func dfs(node *TreeNode, targetSum int) bool {
	if isLeaf(node) {
		return node.Val == targetSum
	}
	if node.Left == nil {
		return dfs(node.Right, targetSum-node.Val)
	}
	if node.Right == nil {
		return dfs(node.Left, targetSum-node.Val)
	}
	return dfs(node.Left, targetSum-node.Val) || dfs(node.Right, targetSum-node.Val)
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

/*
6-rung ladder + correctness (see ladder.md - the generalized 6-step ladder).

Approach: pre-order DFS carrying a residual target. hasPathSum(nil) = false;
a leaf passes iff its value equals the residual; a one-child node must
continue down its only branch (a node with one child is NOT a leaf - the
path can only end at a real leaf). Both children are OR-ed. No pruning is
possible because node values can be negative.

Rung 1 - ENUMERATE (ground truth):

	the candidates are the root-to-leaf PATHS (path = a sequence of nodes from
	the root ending at a leaf, i.e. a node with no children). The predicate on
	a path: sum of its values == targetSum. The answer is EXIST: does at least
	one candidate pass? (Header: correct - "all paths from root to leaves".)

Rung 2 - NAME THE OBJECT:

	question type: EXIST (exists sum(P0..Pend) == targetSum, P0 = root, Pend
	a leaf). Semiring: (OR, AND) - aggregation over the n paths is OR, the
	transition along a path is AND (every step counts). Early exit is
	structural: the first witness suffices, so a "true" can short-circuit the
	whole traversal. (Header: correct - the exists question and early exit
	are both identified.)

Rung 3 - COMPRESS THE SPACE (representation axis):

	instead of materializing all O(#paths) paths and summing each, carry the
	RESIDUAL target down the DFS: path-sum == targetSum iff at each step the
	accumulated subtraction reaches exactly 0 at a leaf. The losslessness
	certificate: the residual fully captures the partial path - it is a
	sufficient statistic (two different partial paths reaching a node with the
	same residual are indistinguishable for the rest of the search), and the
	leaf check preserves the "must end at a leaf" constraint exactly. No path
	is dropped: every descent represents exactly one partial path. (Header:
	correct in spirit - "check running sum during DFS against targetSum".)

Rung 4 - COLLAPSE TO A STATISTIC (aggregation axis):

	the statistic is the integer residual target; the transition is subtract
	(node.Val). The collapse: O(1) state per node instead of an explicit path
	accumulator. The single-child rule belongs here too - the residual is
	carried into 1 branch (one child) or 2 (both), never into a nonexistent
	one. (Header: "running total carried into 1 or 2 branches" - right.)

Rung 5 - FIX A COMPUTATION ORDER (time axis):

	pre-order: check the node, descend. ORDER of exploration is irrelevant to
	correctness (EXIST), so any traversal order works. NO pruning: negative
	values mean the running sum can dip below target and recover, so a branch
	can never be abandoned on the sum alone - the whole tree is explored in
	the worst (false) case. O(n) time, O(height) stack. (Header: the sharpest
	rung - "pruning not possible as there can be negative numbers" is exactly
	right, and it is the correct contrast to LC111 where a bounded search IS
	possible.)

Rung 6 - PROVE AN INVARIANT:

	invariant: dfs(node, targetSum) returns true iff node's subtree contains a
	root-to-node-subtree-leaf path with sum == targetSum (equivalently, the
	residual is reachable at a leaf). For the top call, node's subtree IS the
	whole tree.
	Soundness (no false positives): "true" is returned ONLY at a genuine leaf
	when its value equals the residual - i.e. when the accumulated subtraction
	lands exactly on targetSum - so every witness is a real root-to-leaf path
	with the true sum. A node with one child never reports itself as a leaf,
	so the path always terminates at a true leaf.
	The leaf gate is what makes soundness hold, and it is load-bearing:
	root 1 -> left 2 -> 2's right child 5, targetSum = 3. The path 1->2 sums
	to 3, but 2 is NOT a leaf (it has child 5), so it is not a witness; the
	only real path sums to 8. Answer: FALSE. A buggy "compare at every node"
	version returns true at node 2; this impl returns false because it gates
	the comparison behind isLeaf.
	Completeness (no false negatives): every root-to-leaf path is explored -
	both children are descended when they exist, the one-child case forces the
	only branch, and the false case only concludes after every leaf has been
	checked. A witness anywhere is found, so "false" is returned only when no
	path sums to targetSum.
	Termination: recursion depth bounded by tree height (<= n <= 5000).
	Correctness by induction on subtree size: leaves compare exactly; internal
	nodes reduce the target by their value and hand the residual to their
	(real) children, which by induction answer correctly; the OR of the
	children's answers is the existential over the node's paths.

Complexity: O(n) time, O(height) stack space - no pruning is possible, so the
worst case (returning false) visits every node.

Diff vs the header comments (rung by rung):

- Rung 1: correct - paths from root to leaves; leaf definition stated.
- Rung 2: correct - EXISTS, early exit on witness.
- Rung 3: correct in spirit (running sum); could name the residual as a
  sufficient statistic, but the idea is fully there.
- Rung 4: correct - residual carried into 1 or 2 branches; the one-child
  rule is implicit ("depending on number of children") and the impl handles
  it.
- Rung 5: the strongest rung - negative-values-rule-out-pruning is precisely
  the right call, and it is the sharp contrast with LC111.
- Rung 6: the best Rung 6 attempt so far - both halves are attempted. Gaps:
  (a) the invariant is not stated (dfs's exact guarantee: "true iff a path
  from node to a LEAF in its subtree sums to targetSum"); (b) "BFS will
  reach all leaves" is a mislabel - this is DFS; (c) the leaf-required
  soundness point - a true witness must END at a leaf, so a running sum that
  hits target at an internal node is not a witness - is enforced by the impl
  (the isLeaf gate) but never named as a condition. Worked example: root 1,
  left 2, and 2 has a right child 5, targetSum = 3. 1->2 sums to 3 but 2 is
  not a leaf, so the answer is false; only the isLeaf gate prevents a false
  positive at node 2; (d) completeness is asserted ("tries all possible
  paths") but not tied to the one-child forcing rule.
*/
