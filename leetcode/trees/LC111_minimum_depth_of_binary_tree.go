package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
LC111 Minimum Depth of Binary Tree (Tree, DFS, BFS, Binary Tree)

Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root
node down to the nearest leaf node.

Note: A leaf is a node with no children.

Input: root = [3,9,20,null,null,15,7]
Output: 2

Input: root = [2,null,3,null,4,null,5,null,6]
Output: 5

Constraints:
- The number of nodes in the tree is in the range [0, 10^5].
- -1000 <= Node.val <= 1000
*/

/*
Rung 1 - ENUMERATE (ground truth):

	what is the space?
		subtrees of given tree
	what are we enumerating over?
		nodes of tree

Rung 2 - NAME THE OBJECT:

	min(leftHeight, rightHeight)
	prune longer paths

Rung 3 - COMPRESS THE SPACE (representation axis):

	post-order - return the smaller of the 2 heights between left/right children
	because the space is compressed into log n, discarding the longer depth will not return invalid result

Rung 4 - COLLAPSE TO A STATISTIC (aggregation axis):

	returned height is minimum of left and right subtree heights

Rung 5 - FIX A COMPUTATION ORDER (time axis):

	post-order: each subtree height computed exactly once and the minimum is propagated

Rung 6 - PROVE AN INVARIANT:

	invariant: each subtree is represented as a single scalar - it's minimum height
	heights/depths only increment as each unprocessed parent is encountered in stack - longer path will remain longer throughout
	every tree has a min height - whether containing 0, 1, or more nodes -
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return _minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return _minDepth(root.Left) + 1
	}
	leftMinDepth := _minDepth(root.Left) + 1
	rightMinDepth := _minDepth(root.Right) + 1
	if leftMinDepth > rightMinDepth {
		return rightMinDepth
	}
	return leftMinDepth
}

func _minDepth(node *TreeNode) int {
	if node.Left == nil && node.Right == nil {
		return 1
	}
	if node.Left == nil {
		return _minDepth(node.Right) + 1
	}
	if node.Right == nil {
		return _minDepth(node.Left) + 1
	}
	leftMinDepth := _minDepth(node.Left) + 1
	rightMinDepth := _minDepth(node.Right) + 1
	if leftMinDepth > rightMinDepth {
		return rightMinDepth
	}
	return leftMinDepth
}

/*
6-rung ladder + correctness (see ladder.md - the generalized 6-step ladder).

Approach: bottom-up post-order DFS. The minimum depth of a node's subtree is
the length of the shortest root-to-leaf path within it. The single critical
rule: a node with ONE child is NOT a leaf, so that branch must be continued
(1 + min-depth of the existing child) - never collapsed to 1.

Rung 1 - ENUMERATE (ground truth):

	the candidates are the root-to-leaf PATHS of the tree, not the subtrees or
	nodes per se. The predicate: "this path ends at a leaf (no children)".
	The answer is min over all such paths of (path length). Nodes are the
	recursion domain; paths are what the answer aggregates over.
	(Header: "subtrees"/"nodes" conflates the two - the answer is a min over
	paths, and the recursion over nodes is the vehicle.)

Rung 2 - NAME THE OBJECT:

	question type: OPTIMIZE (minimize). Object: the length of the shortest
	root-to-leaf path. Semiring: (min, +) - aggregate the best path length,
	transition is +1 per edge. No early exit (you must establish the min, so
	the whole tree is processed). (Header: "min(leftHeight, rightHeight)" is
	the right aggregation; "prune longer paths" is NOT what this algorithm
	does - it fully explores both branches. Real pruning would be BFS's
	first-leaf-found-at-depth-d, or branch-and-bound; the post-order version
	prunes nothing.)

Rung 3 - COMPRESS THE SPACE (representation axis):

	the representation is subtree -> its min-depth scalar, computed
	bottom-up. The losslessness certificate: min-depth(node) = 1 + min over
	EXISTING children of min-depth(child) - the parent needs only each child's
	scalar, never its structure, and the min is a sufficient statistic for the
	shortest-path question. The header's "compressed into log n" is wrong:
	the compression is to one scalar per subtree, and the recursion depth is
	the tree height, which is O(n) on a skew tree. "Discarding the longer
	depth will not return invalid result" is right in spirit but only because
	each child's scalar is its TRUE min depth - which itself requires the
	single-child rule below.

Rung 4 - COLLAPSE TO A STATISTIC (aggregation axis):

	the statistic is the integer min-depth; aggregation is (min, +1).
	CRITICAL: "min of left and right subtree heights" is valid ONLY when both
	children exist. With one child, the min is forced to that child's path -
	this is where the naive 1 + min(left, right) produces the false depth 1
	for [2,null,3,...] (the completeness bug this problem is famous for).
	The header's Rung 4 states only the two-child case; the one-child case is
	what actually separates a correct solution from the classic trap.

Rung 5 - FIX A COMPUTATION ORDER (time axis):

	post-order: children before parent, each node visited exactly once ->
	O(n) time, O(height) stack space. The min is propagated up one level at a
	time. (Header: correct - "each subtree height computed exactly once and
	the minimum is propagated.")

Rung 6 - PROVE AN INVARIANT:

	invariant: _minDepth(node) returns the EXACT length of the shortest
	root-to-leaf path within node's subtree (leaf defined as no children).
	Soundness (no false positives): every returned value is a real path
	length - `+1` is applied only when descending into an existing child, and
	the base case is a true leaf (both children nil). No path is ever
	synthesized from a nil child, so a returned depth always corresponds to an
	actual root-to-leaf path.
	Completeness (no false negatives): every root-to-leaf path is explored -
	both children are recursed when they exist, and the one-child case forces
	descent into the only branch instead of terminating early. Hence the true
	shortest path is found, never undercut by a fake 1 from a single-child
	node, and never missed. The min over the (correctly computed) child
	scalars is exact.
	Termination: recursion depth bounded by tree height (<= n <= 10^5).
	Correctness by induction on subtree size: leaves return 1; a one-child
	node returns 1 + its child's min; a two-child node returns 1 + the smaller
	of the two. The answer for the root follows directly.

Complexity: O(n) time, O(height) stack space.

Diff vs the header comments (rung by rung):

- Rung 1: conflates nodes with the true candidates (root-to-leaf paths).
  Harmless for the recursion, but the ground truth is a min over paths.
- Rung 2: right aggregation (min), but "prune longer paths" misdescribes the
  algorithm - post-order prunes nothing; the only real pruning variant is
  BFS stopping at the first leaf.
- Rung 3: the compression is subtree -> scalar, NOT "into log n" (a skew tree
  is O(n) deep); the losslessness claim is correct but only via the
  single-child rule.
- Rung 4: states only the two-child aggregation and silently omits THE
  one-child case - the exact spot the famous false-depth-1 bug lives. Your
  implementation handles it; the header doesn't name it.
- Rung 5: correct as written.
- Rung 6: has base-case awareness ("every tree has a min height") and the
  compression restated, but no soundness/completeness split and no invariant
  stating what _minDepth actually guarantees. The single-child case is the
  missing completeness argument.
*/
