package main

import "fmt"

func main() {}

/**
s = "AABABBA" k = 1
replace first B with A to get AAAABBA (4 As)
replace third A with B to get AABBBBA (4 Bs)

So, given a string of upper case characters, replace k (or less? Yes, "at most k times") characters
to get the longest sequence of repeating upper case chars possible.
Return the length of the sequence.

Edge cases: k = 0, len(s) = 1, len(s) = 2

O(n^2): 
        start := 0
		for i -> start to len(s):
        1. replace first char X at j that differs from Y (at s[start]) with Y
		2. decrement k and if k > 0 and i+1 < len(s), repeat 1. for s[i+1:]
		3. if k == 0 or (i+1 == len(s) or start == 0), store count of Xs as global max
		   3a. if start > 0, move left flipping chars until i == 0
		   3b. repeat 1. starting from i and counting Y's

		   Compress string s:
		        AABABBA -> 2A 1B 1A 2B 1A -> queue
			
				2A 1B 1A 3B k = 2
				 |     |
				 ------
				2A 1B 1A 3B k = 2
				    |    |
				    ------
				2A 1B 1A 3B k = 2 - no need for last iter
				       |  |
				       ----

				2A 1B 1A 3B k = 3 => 2A 1A 1A 2A 1B => 6
				 |        |
				 ---------
				2A 1B 1A 3B k = 2
				    |    |
				    ------
				???fixed k-width window
				for k <- 2..0 and i <- :
				   

				??? from 1B need to go left if k > 0 and reached end of queue
				??? for the first 1A, don't need to go left if when starting at 2A, 
						1A was scanned/examined - No, if k = 3, should go left to 
							produce max sequence
				prune 2B and last 1A - count can't be greater than max?? can go left
					need to store remaining length at queue[i] in "lengths" array
					build lengths array right to left from compressed string
				max = 0
				peek queue, set count = peeked[0] and traverse remaining items in queue
				O(n^2):
				for start -> 0..len(queue):
					for i -> 0 .. len(queue):
						if queue[i][1] != popped char and k > 0, 
							decrement k by queue[i][0] and add queue[i][0] to count
						if k == 0 and count > max ->  max = count; 
							continue outer most loop
				    if k > 0 && start > 0:
					  for i from start..0:
							if queue[i][1] != popped char and k > 0, 
								decrement k by queue[i][0] and add queue[i][0] to count
							if k == 0 and count > max ->  max = count; 
								continue outer most loop

			    if count > max -> max = count
				Look for char sequences seperated by <= k differing chars
				If none found, s is already a maximally repeating sequence, return its length

*/
func characterReplacement(s string, k int) int {

}
