package main

func main() {}

/*
LC835 Image Overlap (see ladder.md - the generalized 6-step ladder).

Problem: given two binary n x n matrices, translate one image (slide up/down/
left/right any number of units, no rotation; bits translated out of bounds are
erased) and count positions where both have a 1. Return the largest possible
overlap. Example: [[1,1,0],[0,1,0],[0,1,0]] vs [[0,0,0],[0,1,1],[0,0,1]] -> 3.

Rung 1 - ENUMERATE (ground truth):

	all (2n-1)^2 translations, each costing O(n^2) to count overlap -> O(n^4).
	Correctness anchor only.

Rung 2 - NAME THE OBJECT:

	maximize the count of coinciding 1-pairs (a in img1, b in img2) over all
	rigid translations. An OPTIMIZE problem: the answer is one scalar, the max
	frequency over the translation space.

Rung 3 - COMPRESS THE SPACE (representation axis):

	represent as differences: a translation is a rigid vector s, so a coinciding
	pair (a, b) IS a translation: s = b - a. Overlap(s) = the number of pairs
	with b - a = s. So don't try translations - count difference vectors.
	Same move as subarray-sum's P[r] - P[l-1] = K: the quantity of interest is
	the difference, and frequency counting collapses the enumeration.

Rung 4 - COLLAPSE TO A STATISTIC (aggregation axis):

	freq[s]++ for every pair of 1s; answer = max(freq). Each overlapping 1-pair
	at shift s maps to exactly one vector - pairs, not flags.

Rung 5 - FIX A COMPUTATION ORDER (time axis):

	one pass over all pairs (m1 * m2), each O(1). n <= 30 so at most 810k pairs.
	No early exit: max needs the full frequency map.

Rung 6 - PROVE AN INVARIANT:

	overlap(s) = |{ (a, b) : a in ones1, b in ones2, b - a = s }| = freq[s], so
	max over s of overlap(s) = max over s of freq[s]. Every 1-pair contributes to
	exactly one vector; nothing is double-counted.

Complexity: O(m1 * m2) time, O(m1 * m2) space.
*/
func largestOverlap(img1 [][]int, img2 [][]int) int {
	img1Ones := findOnes(img1)
	img2Ones := findOnes(img2)
	freq := map[[2]int]int{}
	if len(img1Ones) == 0 || len(img2Ones) == 0 {
		return 0
	}
	for p1 := range img1Ones {
		for p2 := range img2Ones {
			freq[[2]int{p2[0] - p1[0], p2[1] - p1[1]}]++
		}
	}
	var best int
	for _, count := range freq {
		if count > best {
			best = count
		}
	}
	return best
}

func findOnes(image [][]int) map[[2]int]uint8 {
	ones := map[[2]int]uint8{}
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[0]); j++ {
			if image[i][j] == 1 {
				ones[[2]int{i, j}] = 1
			}
		}
	}
	return ones
}
