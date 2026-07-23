package main

// step 0 Restate problem + confirm assumptions
// So we have two lists of intervals.
// Both lists are sorted in that the start of one interval is strictly greatly than the end of the preceding interval.
// At most, one list can be empty
// the start and end of any interval in either list can be equal, or the end must be greater than the start of any given interval
// the intersection of any two intervals is "where they overlap" in 1D
// an interval from one list may intersect zero or more intervals of the other and the intersection is commutative
// "each list of intervals is pairwise disjoint" - there's no overlap between intervals in the same list or, intervals in the same list don't intersect

// step 1 Ask about edge cases upfront
// both lists cannot be empty - at most one list can be empty
//
//	if one list is empty, then intersection is empty
//
// the intersection of two intervals might be a single point (ie ok to have a resulting interval [a,b] where a==b)
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	// step 2 + 3 Talk through approach + ask if sounds good before writing code
	// scenario 1:
	// interval1 [     ]
	// interval2     [      ]
	// intersection  [ ]
	// i2.start >= i1.start and i2.start <= i1.end
	// intersection = [i2.start, i1.end]
	// intersection = [max(starts), min(ends)]

	// scenario 2:
	// interval1 	   [     ]
	// interval2 [     ]
	// intersection    []
	// i2.end >= i1.start and i2.end <= i1.end
	// intersection = [i1.start, i2.end]
	// intersection = [max(starts), min(ends)]

	// scenario 3:
	// interval1 	   [     ]
	// interval2         [ ]
	// intersection    	 [ ]
	// i2.start >= i1.start and i2.end <= i1.end
	// intersection = [i2.start, i2.end]
	// intersection = [max(starts), min(ends)]
	// --------OR the inverse-----------
	// interval1      [ ]
	// interval2    [     ]
	// intersection   [ ]
	// i1.start>=i2.start and i1.end <= i2.end
	// intersection = [i1.start, i1.end]
	// intersection = [max(starts), min(ends)]

	// no intersection
	// interval1      [   ]
	// interval2  [ ]
	// i2.end < i1.start
	// max(starts) >  min(ends)
	// ----OR----
	// interval1      [   ]
	// interval2  			[  ]
	// i1.end < i2.start
	// max(starts) >  min(ends)

	// list arguments are unmodified - no need for defensive-copy
	// traverse both lists in parallel:
	//   advance the list that ends earlier

}
