/*
LC200. Number of Islands

Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1

Example 2:
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3

Constraints:
- m == grid.length
- n == grid[i].length
- 1 <= m, n <= 300
- grid[i][j] is '0' or '1'.
*/

package main

/*
Certificate (Rung 6):

Soundness: Visited set prevents counting the same island twice — BFS from a new cell can't connect to already-visited cells.

Completeness: Row-wise scan checks every cell — no island can exist without a cell being visited.

Determinism: Same cell + same visited set → same BFS neighbors (4 directions, boundary check) → same result. Visited set only grows, never shrinks.
*/

func numIslands(grid [][]byte) int {
	// Rung 1: Brute force — enumerate all connected components of 1s
	// Rung 2: Question type — count (how many islands?)
	// Rung 3: Compression — visited set (skip already-processed cells)
	// Rung 4: Collapse — increment count for each new component
	// Rung 5: Exhaustion — each cell visited once, no cycles in traversal
	// Rung 6: Certificate — see above

	visited := map[[2]int]bool{}
	count := 0
	for i, row := range grid {
		for j, cell := range row {
			if _, ok := visited[[2]int{i, j}]; !ok && cell == '1' {
				count++
				mapIsland(visited, i, j, grid)
			}
		}
	}
	return count
}

/*
["1","1","1","1","0"],
["1","1","0","1","0"],
["1","1","0","0","0"],
["0","0","0","0","0"]
*/
func mapIsland(visited map[[2]int]bool, i, j int, grid [][]byte) {
	if i < len(grid) && i >= 0 && j < len(grid[i]) && j >= 0 {
		cell := grid[i][j]
		if _, ok := visited[[2]int{i, j}]; !ok && cell == '1' {
			visited[[2]int{i, j}] = true
			mapIsland(visited, i+1, j, grid) //down
			mapIsland(visited, i-1, j, grid) // up
			mapIsland(visited, i, j+1, grid) // right
			mapIsland(visited, i, j-1, grid) // left
		}
	}
}
