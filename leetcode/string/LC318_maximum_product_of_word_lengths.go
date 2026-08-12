package main

import "fmt"

/*
LC318 Maximum Product of Word Lengths (Array, String, Bit Manipulation)

Given a string array words, return the maximum value of
length(word[i]) * length(word[j]) where the two words do not share common
letters. If no such two words exist, return 0.

Input: words = ["abcw","baz","foo","bar","xtfn","abcdef"]
Output: 16

Input: words = ["a","ab","abc","d","cd","bcd","abcd"]
Output: 4

Input: words = ["a","aa","aaaa"]
Output: 0

Constraints:
- 2 <= words.length <= 1000
- 1 <= words[i].length <= 1000
- words[i] consists only of lowercase English letters.


[0 0 0 0 0 0 0]16

[1 3 7 15 15 15 15]0


[1 3 7 8 12 14 15]0
1000 = 8
1100 = 12
0011


[1 3 7 8 12 14 15]4

*/

func main() {
	words := []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}
	fmt.Print(maxProduct(words))
}

func maxProduct(words []string) int {
	bitRepresentation := []int{}
	var bitRep int
	for _, word := range words {
		for i := range word {
			bitRep |= 1 << (word[i] - 'a')
		}
		bitRepresentation = append(bitRepresentation, bitRep)
		bitRep = 0
	}
	fmt.Print(bitRepresentation)
	maxProduct := 0
	for i, bit1 := range bitRepresentation {
		for j, bit2 := range bitRepresentation {
			if bit1&bit2 == 0 {
				product := len(words[i]) * len(words[j])
				if product > maxProduct {
					maxProduct = product
				}
			}
		}
	}
	return maxProduct

}

/*

=== MY analysis: rung-by-rung + proof ===

Approach: per-word 26-bit mask (bit k set iff the word contains 'a'+k).
Two words share no letters <=> maskA & maskB == 0. Max product over all
disjoint pairs.

Rung 1 - ENUMERATE: for each pair (i, j), scan every letter of both words to
test whether any letter appears in both. O(n^2 * len) work.

Rung 2 - NAME THE OBJECT: the answer is a SCALAR maximum over pairs, not a
collection. Question type: OPTIMIZATION (maximize len[i]*len[j] over valid
pairs) with an EXIST gate (return 0 when no valid pair exists).

Rung 3 - COMPRESS THE SPACE: the validity predicate "share no letters" is a
SET-INTERSECTION question over a FIXED 26-letter universe (lowercase
constraint is what fixes the universe - without it the mask would need to
grow per alphabet). Compress each word's letter-set to a 26-bit integer; the
intersection test becomes a single bitwise AND. This is a representation
(equivalence) move: bit k set <=> letter k present, so AND == 0 <=> disjoint,
exactly, for every pair, in O(1) regardless of word length.

Losslessness certificate: the map {letter-set -> 26-bit int} is a bijection -
no membership bit is dropped, no two distinct sets share a mask, and the
universe is bounded so the int fits a machine word. Nothing about the
disjointness predicate is approximated by the compression.

Rung 4 - COLLAPSE TO A STATISTIC: each word contributes exactly two numbers to
the objective: len[i] and its mask. The pair statistic is the product
len[i]*len[j]; the aggregate statistic is the MAX over pairs, seeded at 0 so
the EXIST gate is free (no disjoint pair -> stays 0).

Rung 5 - FIX A COMPUTATION ORDER: two passes on DAG dependencies (masks fully
computed before any comparison):
    pass 1: left-to-right over words, OR in one bit per character. O(sum len).
    pass 2: double loop over masks; if bit1&bit2 == 0, take the max. O(n^2).
n <= 1000 makes 10^6 comparisons trivial; no need for the mask-dedup variant
(collapse masks to their max length, at most min(n, 2^26) distinct) unless n
grows. Starting j at 0 (not i+1) is harmless: the self-pair is excluded
automatically because bit & bit == bit != 0 for any non-empty word
(constraint: len >= 1), and disjointness is symmetric so the duplicate pair
just recomputes the same product.

Rung 6 - PROVE AN INVARIANT:
Invariant after pass 2: maxProduct = max over all pairs with
bit1&bit2 == 0 of len[i]*len[j].
SOUNDNESS: every candidate product is computed from a pair whose masks AND to
zero, which by Rung 3 means the words are letter-disjoint - a feasible
product. maxProduct therefore never exceeds a real optimum.
COMPLETENESS: if (i,j) is the optimal pair, its masks are disjoint, the loop
reaches (i,j), and computes P = len[i]*len[j]; so maxProduct >= P. With
soundness (maxProduct is feasible) and the definition of optimum (>= all
feasible products), maxProduct == optimum.
Correctness = soundness && completeness.

Time O(n^2 + sum len), space O(n).

NOTE: the fmt.Print calls inside maxProduct are debug leftovers - they don't
change the return value (LeetCode ignores stdout) but should be removed before
the final pass if this gets submitted verbatim.
*/
