package main

/*
LC1791 Find Center of Star Graph (Graph, Counting)

There is an undirected star graph consisting of n nodes labeled from 1 to n.
A star graph is a graph where there is one center node and exactly n - 1 edges
that connect the center node with every other node.

You are given a 2D integer array edges where each edges[i] = [u_i, v_i]
indicates that there is an edge between the nodes u_i and v_i. Return the
center of the given star graph.

Input: edges = [[1,2],[2,3],[4,2]]
Output: 2

Input: edges = [[1,2],[5,1],[1,3],[1,4]]
Output: 1

Constraints:
- 3 <= n <= 10^5
- edges.length == n - 1
- edges[i].length == 2
- 1 <= u_i, v_i <= n
- u_i != v_i
- The given edges represent a valid star graph.
*/

func findCenter(edges [][]int) int {
	distinctVertices := map[int]int{}

	for _, edge := range edges {
		distinctVertices[edge[0]]++
		distinctVertices[edge[1]]++
	}
	verticesCount := len(distinctVertices)
	for vertex, count := range distinctVertices {
		if count == verticesCount-1 {
			return vertex
		}
	}
	return -1
}
