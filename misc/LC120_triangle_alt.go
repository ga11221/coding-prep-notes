package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/*
Given a triangle array, return the minimum path sum from top to bottom.

For each step, you may move to an adjacent number of the row below. More formally, if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.


Example 1:

Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
Output: 11
Explanation: The triangle looks like:
   2
  3 4
 6 5 7
4 1 8 3

    2
  3   4
 6  5 -107
4  1   8   3

The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).
Example 2:

Input: triangle = [[-10]]
Output: -10


combinations/permutations/subsets/sequences

sequence = pick i vs pick i+1, order matters, optimize sum to be minimum

f(row) = f(row) + f(row-1)

f(first_row) = first_row[0]
f(last_row) = for each i in last_row:
					val1 = i < len(triangle[last_row-1]) ? triangle[last_row][i] + f(triangle[last_row-1][i]) : Sentinel
					val2 = i - 1 >= 0 ? triangle[last_row][i] + f(triangle[last_row-1][i-1]) : Sentinel
					min(val1, val2)
for any two consecutive elements in row[i], they share exactly one common parent in row[i-1] and 1-2 disjoint parents (1 if at the row[i][0] or row[i][end])

*/

func minimumTotal(triangle [][]int) int {

}
