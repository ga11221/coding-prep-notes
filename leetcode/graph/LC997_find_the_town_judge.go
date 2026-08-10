package main

/*
LC997 Find the Town Judge (Graph, Array, Counting)

In a town, there are n people labeled from 1 to n. There is a rumor that one
of these people is secretly the town judge.

If the town judge exists, then:
1. The town judge trusts nobody.
2. Everybody (except for the town judge) trusts the town judge.
3. There is exactly one person that satisfies properties 1 and 2.

You are given an array trust where trust[i] = [a_i, b_i] representing that
the person labeled a_i trusts the person labeled b_i. If a trust relationship
does not exist in trust, then such a trust relationship does not exist.

Return the label of the town judge if the town judge exists and can be
identified, or return -1 otherwise.

Input: n = 2, trust = [[1,2]]
Output: 2

Input: n = 3, trust = [[1,3],[2,3]]
Output: 3

Input: n = 3, trust = [[1,3],[2,3],[3,1]]
Output: -1

Constraints:
- 1 <= n <= 1000
- 0 <= trust.length <= 10^4
- trust[i].length == 2
- All the pairs of trust are unique.
- a_i != b_i
- 1 <= a_i, b_i <= n
*/

/*

1. all pairs of relationships where [A,B] indicates A r B

2. list of unique pairs (A, B) where A or B might exist in multiple pairs. WITNESS question - find B such that for all A, A r B and no pair B r A exists

3. adjacency list: 1 -> [2, 3] for (1, 2) and (1, 3)
    len(adj[A]) = out-degree

4. aggregate by in-degrees/out-degrees - the node of interest has in-degree = n-1 and out-degree = 0

5. for each vertex with out-degree = 0, find the first that's present in the adj list of all other vertices

6. only for vertices with zero out-degree, looks for their presence in all adj lists, that is, all other vertices should have at least one edge directed at the given vertex with zero out-degree
*/

/*

=== DIFF: my analysis (see bottom of file) vs theirs (above) ===

Representation (rung 3):
- mine: two int counters per vertex. in[v] = # people trusting v; out[v] =
  # people v trusts. Both judge conditions ARE the counters.
- theirs: a full adjacency SET per vertex (out-neighborhoods). The "everyone
  trusts judge" test is deferred to a membership re-scan at query time.

Verification (rungs 5-6):
- mine: one O(n) sweep over 1..n testing in[v] == n-1 && out[v] == 0 - the
  test is arithmetic; the graph is fully summarized at build time.
- theirs: two passes - find the unique zero-out-degree vertex, then verify
  membership in every OTHER vertex's adjacency set.

Correctness gap (the load-bearing diff):
- theirs iterates `range adjList` - only vertices that appear in the trust
  array. A vertex with NO incident edges is invisible to both passes, so the
  "everybody except the judge trusts the judge" condition is never checked
  for it -> false positives. n=3, trust=[[1,2]] returns 2; correct is -1
  (person 3 trusts nobody and isn't the judge, so 3 must trust the judge).
  This is the missing-losslessness-certificate failure from ladder.md: the
  adjacency-map compression dropped the empty neighborhoods of isolated
  vertices and nothing restored them.
- mine: arrays are preallocated over ALL n vertices (in/out = 0), so an
  isolated vertex is represented and fails the n-1 test.

n == 1:
- theirs: special-cases it (return 1).
- mine: falls out of the general invariant (in[1] == 0 == n-1, out[1] == 0).

Space: theirs O(n + m) with per-vertex map overhead; mine O(n) two arrays.
*/

func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}
	adjList := map[int]map[int]uint8{}
	for _, pair := range trust {
		if _, ok := adjList[pair[0]]; !ok {
			adjList[pair[0]] = map[int]uint8{}
		}
		adjList[pair[0]][pair[1]] = 1
		if _, ok := adjList[pair[1]]; !ok {
			adjList[pair[1]] = map[int]uint8{}
		}
	}
	zeroOutDegreeVertexCount := 0
	zeroOutDegreeVertex := 0
	for v, adjMap := range adjList {
		if len(adjMap) == 0 {
			zeroOutDegreeVertex = v
			zeroOutDegreeVertexCount++
		}
	}
	if zeroOutDegreeVertexCount == 0 || zeroOutDegreeVertexCount > 1 {
		return -1
	}
	for v, adjMap := range adjList {
		if _, ok := adjMap[zeroOutDegreeVertex]; v != zeroOutDegreeVertex && !ok {
			return -1
		}
	}
	return zeroOutDegreeVertex
}

func findJudge_alt(n int, trust [][]int) int {
	// count out degrees
	// count in degrees
	outDegrees := make([]int, n+1) // n vertices, any one of them can be ommitted from trust list; ie n = 1 and trust = []
	inDegrees := make([]int, n+1)
	for _, pair := range trust {
		outDegrees[pair[0]]++
		inDegrees[pair[1]]++
	}
	for i := 1; i < len(outDegrees); i++ {
		if outDegrees[i] == 0 && inDegrees[i] == n-1 {
			return i
		}
	}
	return -1
}

/*

=== MY analysis: rung-by-rung + proof ===

Approach: two counters per vertex. For each edge (a, b): out[a]++, in[b]++.
Judge = the unique vertex v with in[v] == n-1 && out[v] == 0.

Rung 1 - ENUMERATE: for each candidate v in 1..n, scan all m pairs to check
(a) v never trusts anyone, (b) every other person trusts v. O(n*m).

Rung 2 - NAME THE OBJECT: the object is ONE vertex (the judge). Question
type: WITNESS (return it) with an EXIST gate (return -1 if none). As a set
statement: judge j is the vertex with out(j) = {} and in(j) = V \ {j}.
Because the object is a single vertex - not the collection of all judges -
rung 4 collapses to scalars, not multisets (no set-vs-multiset trap here).

Rung 3 - COMPRESS THE SPACE: both conditions are GLOBAL statements over the
whole edge list, not local to any vertex. Compression: edge (a,b) is +1 to
out[a] and +1 to in[b]; the conditions become two LOCAL scalar tests:
    "v trusts nobody"        <=> out[v] == 0
    "everyone else trusts v" <=> in[v] == n-1
n-1 is exact, not approximate, because trust pairs are unique and a_i != b_i:
the n-1 others each contribute at most one edge to v, so in[v] == n-1 iff all
of them trust v. Transition is LOCAL (an edge touches only its two endpoints'
counters) and the state is SMALL (O(1) per vertex). This is a representation
(equivalence) move: the pair of scalars reproduces the truth of both
conditions exactly.

Losslessness certificate: arrays are allocated over ALL n vertices, so an
isolated vertex is represented as (in=0, out=0) and can be checked. The
adjacency-set variant in findJudge above drops isolated vertices from the map
entirely - their empty neighborhoods are real data, and skipping them yields
the false positive below. Every bug we have caught is a missing certificate;
this is the same shape.

Rung 4 - COLLAPSE TO A STATISTIC: collapse each vertex's whole neighborhood to
the two scalars (out[v], in[v]). WITNESS aggregation: we keep the identity of
the vertex whose pair is (0, n-1) - no frequency, no multiset.

Rung 5 - FIX A COMPUTATION ORDER: two exhaustion passes, both on DAG
dependencies (counters accumulate and never read each other, so every value is
final when computed):
    pass 1: left-to-right over the edge list - +1 to the two endpoints.
    pass 2: left-to-right over 1..n - return the first (and only) vertex
            with in == n-1 && out == 0.
Their solution needs a second STRUCTURAL pass (membership re-scan of every
adjacency set) because it deferred the "everyone trusts judge" test to query
time; the counter representation moves it to build time, so pass 2 is pure
arithmetic.

Rung 6 - PROVE AN INVARIANT:
Invariant after pass 1: in[v] = # distinct people trusting v; out[v] = # of
people v trusts.
SOUNDNESS: if in[v] == n-1 && out[v] == 0, then v is trusted by all n-1
others and trusts nobody - both judge conditions hold. No second vertex can
pass: two candidates would each need out-degree 0 (so neither trusts the
other) and in-degree n-1 (so each is trusted by the other) - contradiction.
The "exactly one person" clause therefore holds by construction.
COMPLETENESS: if judge j exists, in[j] == n-1 && out[j] == 0, so pass 2
reaches j and returns it; no vertex (isolated or not) is invisible to the
arrays.
Correctness = soundness && completeness. (n == 1 needs no special case:
in[1] == 0 == n-1 and out[1] == 0.)

Concrete divergence: n=3, trust=[[1,2]] -> vertex 3 is isolated, (0,0); fails
the n-1 test; returns -1. findJudge above returns 2.
*/
