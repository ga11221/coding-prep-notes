package main

/*
LC2596 Check Knight Tour Configuration (Graph, Matrix)
There is a knight on an n x n chessboard. In a valid configuration, the knight starts at the top-left cell of the board and visits every cell on the board exactly once.

You are given an n x n integer matrix grid consisting of distinct integers from the range [0, n * n - 1] where grid[row][col] indicates that the cell (row, col) is the grid[row][col]th cell that the knight visited. The moves are 0-indexed.

Return true if grid represents a valid configuration of the knight's movements or false otherwise.

Note that a valid knight move consists of moving two squares vertically and one square horizontally, or two squares horizontally and one square vertically. The figure below illustrates all the possible eight moves of a knight from some cell

Input: grid = [[0,11,16,5,20],[17,4,19,10,15],[12,1,8,21,6],[3,18,23,14,9],[24,13,2,7,22]]
Output: true
Explanation: The above diagram represents the grid. It can be shown that it is a valid configuration

rung 1: O(nxn) grid, starting at (0,0) find all valid moves ie dx =+/-2 and dy = +/-1 or dx =+/-1 and dy = +/-2, determine if next move is in this set of valid moves - for each cell that's 2 <= |{valid_moves}| <= 8

	2n^2 <= total cmps <= 8n^2

rung 2:

	Object = the ordered visit sequence: cell_0, cell_1, ..., cell_{n*n-1}
	where grid[cell_k] = k. Ground-truth move graph: cell' is reachable
	from cell iff (|dr|,|dc|) == (1,2) or (2,1), bounded to
	0 <= row,col < n (strict, not <=). Question: is the given sequence
	a walk in this graph? (sets only — no meet/join; "lattice" is a misnomer)

rung 3:

	traverse grid to find and order cells/moves => collapse into linear representation of ordered moves
	Lattice = array of ordered moves, missing move or dupe => invalid
	no fixpoint - no meet/join/sup/inf - exhaustion of moves
	but there is a combine? move2->move3 validity depends on move1->move2

rung 4:

	Anchor: cell_0 must be (0,0), i.e. grid[0][0] == 0.
	Statistic = meet (AND-fold / GLB) over k = 1..n*n-1 of
	isKnightMove(cell_{k-1}, cell_k), where isKnightMove:
	|dr|,|dc| == {1,2} in either order (|dr|+|dc| == 3, both >= 1).
	Dual to reachability's join, so one-shot exhaustion, no fixpoint iteration.
	ALL question type; early-exit on the first false pair (witness).

rung 5:

	left-to-right in array of ordered moves

rung 6:

	Invariant: after processing cell_0..cell_k, the prefix is a valid
	tour prefix (starts at (0,0), every consecutive pair is a knight move).
	Valid config <=> invariant holds at k = n*n-1.
*/
func checkValidGrid(grid [][]int) bool {
	if grid[0][0] != 0 {
		return false
	}
	orderedMoves := make([][2]int, len(grid)*len(grid))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			orderedMoves[grid[i][j]] = [2]int{i, j}
		}
	}
	for i, j := 0, 1; j < len(orderedMoves); i, j = i+1, j+1 {
		if !isKnightMove(orderedMoves[i], orderedMoves[j]) {
			return false
		}
	}
	return true
}

func isKnightMove(pos1, pos2 [2]int) bool {
	dr, dc := pos1[0]-pos2[0], pos1[1]-pos2[1]
	return dr*dr+dc*dc == 5
}
