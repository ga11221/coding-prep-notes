package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	n := 5
	/*
		edges := [][]int{
			{0, 1},
			{0, 2},
			{0, 3},
			{1, 4},


			{2,3}, // makes cycle
		}
	*/
	edges := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{1, 3},
		{1, 4},
	}
	fmt.Printf("for edges: %v, is valid tree: %v", edges, validTree(n, edges))
	/*
		[[0,1],[2,3],[1,2]] - sort first

		[[0,1],[1,4],[2,3],[3,4]]
	*/
}

/*
Example 1:
Input: n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]
Output: true

                                0
						1      2      3
						4


								0
						1      2       3   4
						4
Example 2:
Input: n = 5, edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]
Output: false

       					parents = {0,1}
i=1;  edge = [1,2]      parents = {0,1,2}
i=2;  edge = [2,3]      parents = {0,1,2,3}
								0
								1
							2   4    3
							3



Constraints: 1 ≤ n ≤ 2000, 0 ≤ edges.length ≤ 5000
*/

func validTree(n int, edges [][]int) bool {
	slices.SortFunc(edges, func(a, b []int) int {
		c := cmp.Compare(a[0], b[0])
		if c == 0 {
			return cmp.Compare(a[1], b[1])
		}
		return c
	})
	return _validTree(n, 0, edges, 1, [][]int{})

}

/*
f(parents, edge) = true if edge is nil; false if edge[0] in parents and edge[1] in parents
f([0,1]) = f(parents=[0,1], next edge in edges) if edge[0] not in parents and edge[1] not in parents else false

cycle = any two non-adjacent nodes in a single path are connected or any two distinct paths with a common root are connected

	[[0,1],[1,4],[2,3],[3,4]]
	0        2
	1		 3
	4
*/
/*

	{0, 1},
	{0, 2},
	{0, 3},
	{1, 4},


	{2,3}, // makes cycle

*/
func _validTree(numberOfEdges int, i int, edges [][]int, pathNumber int, paths [][]int) bool {
	if i == len(edges) {
		return true
	}
	edge := edges[i]

	if pathNumber > len(paths) {
		paths = append(paths, make([]int, numberOfEdges+1))
	}

	// cycle check: both endpoints already in the same component
	for _, p := range paths {
		if p[edge[0]] == 1 && p[edge[1]] == 1 {
			return false
		}
	}

	// find which components contain each endpoint
	mergeablePaths := []int{}
	for i, p := range paths {
		if p[edge[0]] == 1 || p[edge[1]] == 1 {
			mergeablePaths = append(mergeablePaths, i)
		}
	}

	if len(mergeablePaths) == 2 {
		// endpoint in separate components → merge them
		firstPath := paths[mergeablePaths[0]]
		for i, n := range paths[mergeablePaths[1]] {
			if n == 1 {
				firstPath[i] = 1
			}
		}
		paths = append(paths[0:mergeablePaths[1]], paths[mergeablePaths[1]+1:]...)
	} else if paths[pathNumber-1][edge[0]] == 1 || paths[pathNumber-1][edge[1]] == 1 {
		// one endpoint in current component → extend it
		path := paths[pathNumber-1]
		path[edge[0]] = 1
		path[edge[1]] = 1
	} else {
		// neither endpoint seen → start a new component
		paths = append(paths, make([]int, numberOfEdges+1))
		pathNumber++
		paths[pathNumber-1][edge[0]] = 1
		paths[pathNumber-1][edge[1]] = 1
	}

	return _validTree(numberOfEdges, i+1, edges, pathNumber, paths)

}

/*

The sorting + multi-component tracking essentially implements Union-Find manually using bit arrays instead of parent pointers.
The redundant empty path0 that accumulates pathNumber but never gets used is harmless.


| Your current code                                                          | Union-Find                                                                                    |
|----------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------|
| `paths [][]int` — array per component                                      | `parent []int` — single array, tree of pointers                                               |
| `p[u] == 1 && p[v] == 1` → cycle                                           | `find(u) == find(v)` → cycle                                                                  |
| `mergeablePaths` — find which paths u and v belong to                      | `find(u)`, `find(v)` — find respective roots                                                  |
| Merging two paths (copy bits)                                              | `parent[root_u] = root_v` — O(1) pointer assignment                                           |
| "Extend path" + "Start new path" cases                                     | Just `union(u, v)` — each vertex starts as its own component, no special cases needed         |

Plan:
1. Replace `paths [][]int, pathNumber int` with `parent []int` initialized as `parent[i] = i`
2. Write `find(x int) int` — walk while `parent[x] != x`
3. Write `union(x, y int) bool` — if `find(x) == find(y)` return false (cycle), else `parent[find(x)] = find(y)` and return true
4. `_validTree` reduces to: iterate edges, `if !union(u, v) return false`, and after loop check `len(edges) == n-1`
		(tree must have n-1 edges and all nodes reachable — union ensures connectivity)

*/
