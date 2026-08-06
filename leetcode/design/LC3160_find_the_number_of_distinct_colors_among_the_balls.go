package main

import "fmt"

/*
LC3160 Find the Number of Distinct Colors Among the Balls (Design, Hash Table)

Input: limit = 4, queries = [[1,4],[2,5],[1,3],[3,4]]

Output: [1,2,2,3]

Explanation:

After query 0, ball 1 has color 4.
After query 1, ball 1 has color 4, and ball 2 has color 5.
After query 2, ball 1 has color 3, and ball 2 has color 5.
After query 3, ball 1 has color 3, ball 2 has color 5, and ball 3 has color 4.

Input: limit = 4, queries = [[0,1],[1,2],[2,2],[3,4],[4,5]]

Output: [1,2,2,3,4]

Explanation:

After query 0, ball 0 has color 1.
After query 1, ball 0 has color 1, and ball 1 has color 2.
After query 2, ball 0 has color 1, and balls 1 and 2 have color 2.
After query 3, ball 0 has color 1, balls 1 and 2 have color 2, and ball 3 has color 4.
After query 4, ball 0 has color 1, balls 1 and 2 have color 2, ball 3 has color 4, and ball 4 has color 5.

rung 1: Enumerate
after each query, count distinct colors by scanning balls 0..limit
O(q * limit), q = len(queries)

rung 2: Name the object
state after query i = coloring of balls (ball -> color)
answer[i] = |image of the coloring| = distinct colors held by some ball

rung 3: Compress the space (representation)
do not materialize limit+1 balls; track only colored balls:

	ballColors:     ball -> color, only colored balls (comma-ok = uncolored)
	distinctColors: color -> count of distinct balls holding it

state = O(#colored balls), not O(limit)

rung 4: Collapse to a statistic (aggregation)
answer = len(distinctColors); a color is in the set iff count >= 1
set-vs-multiset trap: to know when a color becomes unused you need the COUNT,
not a flag - a bare set cannot answer "is old still held by another ball?"

rung 5: Fix a computation order (time axis)
queries left-to-right; each query updates state in O(1) - exhaustion (single
pass), since query i's answer never depends on later queries

rung 6: Prove an invariant
after query i, distinctColors = {c | count[c] >= 1} matches exactly the balls
that hold c; decrement to 0 only on the last ball, so len is the true count
*/
func queryResults(limit int, queries [][]int) []int {
	distinctColors := map[int]int{}
	ballColors := map[int]int{}
	output := []int{}
	for _, query := range queries {
		if currColor, ok := ballColors[query[0]]; ok {
			if currColor != query[1] {
				distinctColors[query[1]]++
				distinctColors[currColor]--
				if distinctColors[currColor] == 0 {
					delete(distinctColors, currColor)
				}
			}
		} else {
			distinctColors[query[1]]++
		}
		ballColors[query[0]] = query[1]
		output = append(output, len(distinctColors))
	}
	return output
}

func main() {
	limit := 4
	queries := [][]int{{1, 4}, {2, 5}, {1, 3}, {3, 4}}
	fmt.Printf("for limit %v and queries: %v, output: %v", limit, queries, queryResults(limit, queries))
}
