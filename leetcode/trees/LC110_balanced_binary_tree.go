package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
LC110 Balanced Binary Tree (Tree, DFS, Binary Tree)

Given a binary tree, determine if it is height-balanced.

A height-balanced binary tree is a binary tree in which the depth of the two
subtrees of every node never differs by more than one.

Input: root = [3,9,20,null,null,15,7]
Output: true

Input: root = [1,2,2,3,3,null,null,4,4]
Output: false

Input: root = []
Output: true

Constraints:
- The number of nodes in the tree is in the range [0, 5000].
- -10^4 <= Node.val <= 10^4
*/

/*
Rung 1 - ENUMERATE (ground truth):
	 what is the (uncompressed) candidate space?
		all subtrees with `root` as its root
	 what are you enumerating over?
		a tree and its subtrees
	 ALL subtrees must be balanced for root to be balanced

Rung 2 - NAME THE OBJECT:
	ALL question: do ALL subtrees satisfy predicate |h(L) - h(R)| <= 1 (or, and) semiring (?)
	Early exit if a single subtree does not satisfy predicate
	Scalar - true (is balanced) vs false

Rung 3 - COMPRESS THE SPACE (representation axis):
	recursive representation - bottom-up WRT tree - compressed into stack of pending parent nodes
	all parents are evaluated, and a parents subtrees must be balanced for it to be balanced - no losslessness

Rung 4 - COLLAPSE TO A STATISTIC (aggregation axis):
	predicate check per subtree - early exit indicates falsehood

Rung 5 - FIX A COMPUTATION ORDER (time axis):
	in-order: compute height/depth of left subtree, then right, and evaluate predicate
	each subtree height is computed exactly once
	if false for a sub-tree than false for `root` in ground truth
	if true for sub-tree, must hold for containing sub-trees

Rung 6 - PROVE AN INVARIANT:

*/

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight, leftBalanced := _height(root.Left)
	rightHeight, rightBalanced := _height(root.Right)

	if !leftBalanced || !rightBalanced {
		return false
	}

	if leftHeight > rightHeight {
		return leftHeight-rightHeight <= 1
	}
	return rightHeight-leftHeight <= 1
}

func _height(node *TreeNode) (height int, balanced bool) {
	if node == nil {
		return 0, true
	}

	leftHeight, leftBalanced := _height(node.Left)
	rightHeight, rightBalanced := _height(node.Right)
	if !leftBalanced || !rightBalanced {
		return -1, false
	}
	if leftHeight > rightHeight {
		return leftHeight + 1, leftHeight-rightHeight <= 1
	}
	return rightHeight + 1, rightHeight-leftHeight <= 1

}

/*
6-rung ladder + correctness (see ladder.md - the generalized 6-step ladder).

Approach: bottom-up post-order DFS. _height(node) returns the height of node's
subtree AND whether every node in it is balanced; the moment either child is
unbalanced, the whole computation short-circuits to (-1, false). isBalanced
applies the same predicate to the root's two children and to the root itself.

Rung 1 - ENUMERATE (ground truth):

	the candidates are the n subtrees - one per node, each rooted at that
	node. The predicate on a candidate subtree is "|h(L) - h(R)| <= 1" applied
	at ITS root AND inherited down its children. A tree is balanced iff ALL n
	subtree-predicates hold. (Header: correct - "all subtrees must be balanced
	for root to be balanced".)

Rung 2 - NAME THE OBJECT:

	question type: ALL (universal quantification) - does EVERY subtree satisfy
	the predicate? Semiring: the dual of existence - (AND, OR): aggregation
	over the n subtrees is AND (all must pass), failure transition is OR (one
	bad child poisons the node). Early exit is structural: the first false
	somewhere in the tree makes the answer false, so the answer is a scalar
	boolean with early termination. (Header: right question type and early
	exit; the "(or, and) semiring (?)" is backwards - ALL uses (AND, OR).)

Rung 3 - COMPRESS THE SPACE (representation axis):

	the representation is subtree -> (height, balanced) pair, computed
	bottom-up. The compression is the COORDINATE CHANGE: a parent needs only
	each child's height to judge its own balance, not the child's whole
	structure. The losslessness certificate: height is a SUFFICIENT STATISTIC
	for the balance predicate - any two subtrees with equal height are
	indistinguishable for their parent's balance check, and every node's own
	balance is decided locally from its two children's heights. (Header: "no
	losslessness" is half right - no candidate is discarded, since every node
	is visited exactly once; but the subtree->height collapse is itself a
	compression and it needs exactly this sufficiency certificate.)

Rung 4 - COLLAPSE TO A STATISTIC (aggregation axis):

	the statistic is the (height, balanced) pair; height aggregates by
	max(L,R)+1, balance aggregates by AND of both children plus the local
	|hL - hR| <= 1. Early exit on (-1, false) is the AND-aggregation short
	circuit. (Header: matches, though it only names the predicate check, not
	the height aggregation that makes the statistic well-defined.)

Rung 5 - FIX A COMPUTATION ORDER (time axis):

	the order is POST-order (children before parent), not in-order - the
	header's "in-order: left, then right" is the traversal per node but the
	depth-first order that matters is bottom-up. Each node contributes to
	exactly one pair, so total work is O(n) and height is computed exactly
	once. The asymmetric propagation the header intuits (lines 54-55) is the
	load-bearing insight: FALSE propagates up unconditionally (an unbalanced
	subtree can never be rebalanced by an ancestor, so early exit is safe),
	while TRUE does NOT (a balanced subtree does not make its ancestor
	balanced - each ancestor must recheck its own |hL - hR|). The header's
	line 55 wording is backwards but the intuition is exactly right.

Rung 6 - PROVE AN INVARIANT:

	invariant: _height(node) returns (h, true) where h is the EXACT height of
	node's subtree and every node in that subtree is balanced; otherwise it
	returns (-1, false).
	Soundness (no false positives): a "true" answer propagates only through
	nodes where both children returned true AND the local |hL - hR| <= 1 held
	at each step, so "true" at the root means every node in the tree passed
	the predicate - nothing is misreported as balanced.
	Completeness (no false negatives): any node whose subtree is unbalanced
	has its imbalance discovered exactly when that node is processed (either
	an unbalanced child short-circuits, or its own |hL - hR| > 1 is caught),
	and every node IS processed (post-order visits all n). So no unbalanced
	tree can return true.
	Termination: the recursion is bounded by tree depth (<= n <= 5000).
	Correctness follows by induction on subtree height: leaves trivially
	return (0, true); an internal node's pair is correct given correct pairs
	for both children.

Complexity: O(n) time, O(height) stack space - the compression is what makes
the O(1) per-node collapse possible (vs. O(height) height recomputation per
node for a naive check).

Diff vs the header comments (rung by rung):

- Rung 1: correct and well anchored ("ALL subtrees must be balanced"). Only
  nit: candidates are "each node's subtree", not "subtrees with root as its
  root".
- Rung 2: the strongest part - universal question type, early exit, scalar
  answer all identified. One correction: the semiring is (AND, OR), the dual
  of exist's (OR, AND); the trailing "(?)" marks the right doubt.
- Rung 3: the representation is the (height, balanced) pair (a coordinate
  change), and there IS a losslessness certificate - height is a sufficient
  statistic for the balance predicate. "no losslessness" only holds for
  candidate-dropping, which is a different axis.
- Rung 4: names the predicate but not the height aggregation that defines
  the statistic.
- Rung 5: the order is post-order, not in-order; the false/true asymmetric
  propagation (the real insight) is intuited but line 55's direction is
  reversed.
- Rung 6: EMPTY - the gap. The invariant above is the missing piece; it is
  what makes the early exit provably sound and complete rather than just
  "feeling right".
*/
