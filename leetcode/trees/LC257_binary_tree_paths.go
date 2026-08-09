package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
LC257 Binary Tree Paths (Tree, DFS, Binary Tree)

Given the root of a binary tree, return all root-to-leaf paths in any order.

A leaf is a node with no children.

Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]

Input: root = [1]
Output: ["1"]

Constraints:
- The number of nodes in the tree is in the range [1, 100].
- -100 <= Node.val <= 100
*/
/*
enumerate: all paths starting from root and ending on leaves
object: set of all paths in tree where start = root and end = any distinct leaf
compress: walk all shared paths and follow each branch to a distinct leaf, then return to branch point and walk next branch
representation: stack of nodes with pending paths
stat: running string of current path, all paths (in any order) in one set
order: pre-order guarantees each node val appended to string representation of path before visiting next node
invariant: is every path valid? are any paths missed or are all valid paths returned?
          yes, check each node for both left and right child before proceeding - if isLeaf, append last node and return up and check remaining paths for unprocessed parents in stack - hence all paths returned

header diffs (see bottom block for full rungs):
- "representation: stack of nodes with pending paths": there is no explicit
  stack - the recursion call stack plays that role; nothing is held "pending"
  beyond the implicit frames
- "compress: walk shared paths ... return to branch point and walk next
  branch": that is traversal mechanics (backtracking = HOW), not a
  compression claim. The real compression is the shared-prefix running string;
  and there is no STATE to compress - the output itself is exponential
  (up to 2^(h-1) paths), so no sub-exponential compression is possible
- rung 4 (aggregation) is absent: the answer is the set of all paths (free/
  constructor semiring, union recurrence at each node) - the largest output
  case, no min/max/count
- invariant is posed as a rhetorical question, not stated as a property that
  yields soundness AND completeness; the "check both children, append at
  leaf" argument is completeness reasoning but is never named as such
- order: correct (pre-order) - though output order is arbitrary, so any DFS
  order would do; pre-order is natural because a node's value must be
  appended before descending into it
*/
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	paths := []string{}
	return *_path(root, "", &paths)
}

func _path(node *TreeNode, path string, paths *[]string) *[]string {
	path += strconv.Itoa(node.Val)
	if isLeaf(node) {
		*paths = append(*paths, path)
		return paths
	}
	if node.Left != nil {
		_path(node.Left, path+"->", paths)
	}
	if node.Right != nil {
		_path(node.Right, path+"->", paths)
	}
	return paths
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

/*
rung 1: enumerate
	all root-to-leaf paths; a leaf is a node with no children
	ground truth: the full list of root->...->leaf chains, any order

rung 2: name object
	ALL - "return all root-to-leaf paths"
	question type: report-all (list of witnesses), not exist/count/optimize
	aggregation: free/constructor semiring - the answer is the SET of all
	solutions, combined by set-union; order arbitrary, no min/max/count

rung 3: compression
	candidates: up to 2^(h-1) paths in a full tree (h = height) - exponential
	NO state compression possible: the output IS every path, so it must be
	materialized in full; there is no statistic to collapse
	the only "compression" is amortizing shared prefixes: a running path string
	built node-by-node instead of re-deriving each root-to-leaf chain from
	scratch (paths through the same ancestor reuse that ancestor's prefix)
	losslessness: the running string preserves the exact root-to-current chain
	in order, so no path information is lost when a prefix is reused

rung 4: aggregation
	statistic per node: the set of complete root-to-leaf paths through it
	recurrence:
	    paths(node) = {chain(root..node)}                if isLeaf(node)
	               = paths(left) U paths(right)          otherwise
	(free/constructor semiring: objects combined by APPEND/union, not min or +)

rung 5: order
	pre-order DFS: append node val before descending; descend into BOTH
	children (single-child case handled: Left != nil and Right != nil are
	checked separately)
	O(h) call stack; the real cost is total output chars
	note: path is passed by value (string concat copies) - O(h) chars per
	frame; a []int accumulator with a final join would avoid the copying

rung 6: invariant -> soundness + completeness
	invariant: at each DFS step, path = the exact root-to-current-node chain
	("->"-joined), and *paths = exactly the root-to-leaf paths fully explored
	so far
	soundness (no fake paths): a string is appended ONLY at isLeaf(node), and
	path is a genuine root-to-current chain -> every emitted string is a true
	root-to-leaf path
	completeness (no missed paths): every node is visited; every non-nil child
	is descended into (both Left and Right), so every chain of non-nil edges
	is traversed; a path can end only at a leaf, and every leaf reached appends
	its chain -> all root-to-leaf paths are emitted
	termination: recursion only moves to children; tree is finite
	base cases: root == nil -> empty; single-node tree -> leaf, appends at once
*/
