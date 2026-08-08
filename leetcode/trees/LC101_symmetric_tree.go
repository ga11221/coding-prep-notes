package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
LC101 Symmetric Tree (Tree, DFS, BFS)

Given the root of a binary tree, check whether it is a mirror of itself (i.e.,
symmetric around its center).

Input: root = [1,2,2,3,4,4,3]
Output: true

Input: root = [1,2,2,null,3,null,3]
Output: false

Constraints:
- The number of nodes in the tree is in the range [1, 1000].
- -100 <= Node.val <= 100

Follow up: could you solve it both recursively and iteratively?
*/

func isSymmetric(root *TreeNode) bool {
	return _isSymmetric([]*TreeNode{root.Left, root.Right})
}

/*
enumerate over all subtrees - all nodes n
symmetry is not a sub-tree property, so must be evaluated by level. symmetry only holds for root - looking for existence of property on root
state space is all nodes:

	for each node evaluate the children:
		first sibling equals last sibling, 2 sibling = n-2th sibling, etc

compress by level - monotonic in level (using program stack, queue or otherwise)
nodes are equal if both are nil or if both are non-nil and have the same value
because algo progresses by level, if a given level is not symmetric, the property doesn't hold on root - early return
symmetric only if all levels are symmetric
*/
func _isSymmetric(children []*TreeNode) bool {
	// invariant all siblings in left half compared with all in right half
	for i, j := 0, len(children)-1; i <= j; i, j = i+1, j-1 {
		if children[i] != children[j] {
			if children[i] == nil {
				return false
			}
			if children[j] == nil {
				return false
			}
			if children[i].Val != children[j].Val {
				return false
			}
		}
	}
	// post-invariant: this level is symmetric
	newChildren := []*TreeNode{}
	lastLevel := true
	for _, child := range children {
		if child != nil {
			lastLevel = lastLevel && child.Left == nil && child.Right == nil
			newChildren = append(newChildren, []*TreeNode{child.Left, child.Right}...)
		} else {
			newChildren = append(newChildren, []*TreeNode{nil, nil}...)
		}
	}
	if lastLevel {
		return true
	}
	return _isSymmetric(newChildren)
}

/*
6-rung ladder + correctness (see ladder.md - the generalized 6-step ladder).

Approach: BFS by level. For each level, check the node sequence is a mirror
palindrome (first == last, second == second-to-last, ...), then expand to the
next level's children (nil children become nil placeholders). If any level
fails, the tree is asymmetric; the tree is symmetric iff EVERY level passes.

Rung 1 - ENUMERATE (ground truth):

	the tree is a mirror of itself iff for every depth d, the node sequence
	at depth d is a palindrome under the equality "both nil, or both non-nil
	with equal values". Checking all levels is the ground truth.

Rung 2 - NAME THE OBJECT:

	existence question: does the root have the mirror property? The object is
	the level sequence; the property is "level is a palindrome". EXIST ->
	early return on the first asymmetric level is a rung-2 property.

Rung 3 - COMPRESS THE SPACE (representation axis):

	symmetry is NOT a per-subtree property - a subtree being symmetric does
	not imply the whole tree is, so evaluate BY LEVEL: expand the current
	level into its ordered children, keeping nil placeholders so positions
	stay aligned. Mirror pairs at a level are (first, last), (second,
	second-to-last), ... because mirroring flips the order.

Rung 4 - COLLAPSE TO A STATISTIC (aggregation axis):

	per level, the statistic is the boolean "level is a palindrome" under the
	node-equality predicate (nil == nil, otherwise values must match). The
	answer is AND over levels, i.e. OR-aggregation of failure with early
	exits on the first asymmetric level.

Rung 5 - FIX A COMPUTATION ORDER (time axis):

	process levels top-down; each level's nodes feed the next level's slice.
	Every node appears in exactly one frontier -> O(n) time, O(width) space
	(worth O(n) for a skewed/sparse tree).

Rung 6 - PROVE AN INVARIANT:

	invariant: before each level check, children holds exactly the nodes of
	the current depth, left-to-right, with nil placeholders for missing
	children.
	Soundness (no false positives): a level passing the palindrome check is
	genuinely symmetric; false is returned only when a real mirror pair
	differs (nil vs node, or differing values) - nothing is misjudged.
	Completeness (no false negatives): nil placeholders exactly represent the
	true next-level nodes (a nil parent has no children), and the expansion
	preserves order, so every node of every depth is present in exactly one
	checked level. Any asymmetric level is caught when it is processed, and
	all levels are processed until the frontier is all nil - so the tree is
	rejected iff some level is asymmetric. Correctness follows by induction
	on depth.
	Termination: the all-leaf level (next frontier all nil) returns true.

Complexity: O(n) time, O(n) worst-case space for the level frontier.
*/

/*
Analysis of the original header comments vs the proof.

What the header captured (Rungs 1-5):
  - "enumerate over all subtrees" + "symmetry is not a sub-tree property, so
    must be evaluated by level" -> Rung 1/3, and the key insight (the leap).
  - "symmetry only holds for root - looking for existence of property on
    root" -> Rung 2 (existence on the root).
  - "for each node evaluate the children: first sibling equals last sibling,
    2 sibling = n-2th sibling" -> Rung 1/2 predicate: the palindrome check.
  - "compress by level - monotonic in level" -> Rung 3 + Rung 5 (compression
    by level and the frontier computation order).
  - "nodes are equal if both are nil or both non-nil with same value" -> the
    equality predicate, the definition soundness rests on.
  - "if a given level is not symmetric, the property doesn't hold on root -
    early return" -> soundness of the rejection (early exit is safe).
  - "symmetric only if all levels are symmetric" -> completeness (checking
    every level is necessary and sufficient).

What was missing (the actual proof, Rung 6):
  1. The INVARIANT was never stated: children always holds exactly the nodes
     of the current depth, left-to-right, with nil placeholders keeping
     positions aligned. That invariant is what makes the palindrome check
     correct, and it is absent from the header.
  2. COMPLETENESS was asserted, not argued: "symmetric only if all levels
     are symmetric" is the claim; the reason is that nil placeholders exactly
     represent the true next-level nodes (a nil parent has no children), so
     every node of every depth lands in exactly one checked level - no
     asymmetry can hide deeper.
  3. SOUNDNESS was one-sided: the header covers "false is only returned when
     a real mirror pair differs", but never separately justifies "a level
     passing the check is genuinely symmetric" (no false positives).

Terminology slip: "enumerate over all subtrees" - the ground truth is
PER-LEVEL, not per-subtree (corrected on the next line, but the rung-1 anchor
is the level sequence).

Bottom line: the header had all the ingredients (compression, order,
predicate, both halves of the iff) arranged as a mechanism description ("how
the algorithm walks") rather than a proof ("why it can't be wrong"). The gap
is exactly the Rung 6 portion: state the frontier invariant, then derive
soundness and completeness from it.
*/
