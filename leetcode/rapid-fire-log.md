# Rapid Fire LC Medium Approaches (Jul 25, 2026)

## LC15 — 3Sum
**User approach:** Map of target - nums[i], then two pointers on sorted array.
**Refined approach:** Sort the array. Fix one element nums[i], then two-pointer the remaining subarray for the pair that sums to -nums[i]. Skip duplicates at both the fixed element and the two pointers to avoid duplicate triplets.
**Time:** O(n²)
**Pattern:** Sorting + Two Pointers
**Key insight:** Reduces 3Sum to repeated 2Sum on sorted subarrays. Duplicate skipping is what makes it clean.
**Discrete Math:** Sorting/Total Order, Set theory (unique triplets), Combinatorics (choose 3 from n)

## LC128 — Longest Consecutive Sequence
**User approach:** Sliding window — expand right while elements are consecutive, shrink and start from next element that breaks sequence.
**Refined approach:** Hash set. For each element, check if it's the start of a sequence (num-1 not in set), then count forward. Each element visited at most twice.
**Time:** O(n)
**Pattern:** Hash Set
**Key insight:** Only start counting from sequence beginnings. Skipping elements that are mid-sequence avoids revisiting.
**Discrete Math:** Set theory (membership), Integer sequences

## LC238 — Product of Array Except Self
**User approach:** Find product for first element (24) and carry forward, product for next element at i is product/array[i] * array[i-1].
**Refined approach:** Two passes — no division. Left pass builds prefix products, right pass multiplies in suffix products. Suffix: iterate right to left, starting with 1, multiply by next element.
**Time:** O(n)
**Pattern:** Prefix/Suffix Products
**Key insight:** Each element = product of everything left × everything right. Two independent passes, no division needed.
**Discrete Math:** Monoid (prefix/suffix product under multiplication), Associativity

## LC200 — Number of Islands
**User approach:** DFS from each unvisited '1', mark everything reachable, count.
**Time:** O(m*n)
**Pattern:** Graph DFS/BFS
**Key insight:** Classic flood fill. Visit each cell once.
**Discrete Math:** Graph theory (connected components, undirected grid graph)

## LC139 — Word Break
**User approach:** Trie for dictionary lookup.
**Refined approach:** DP + Trie (or hash set). dp[i] = true if any s[j:i] is in dict and dp[j] is true. Trie helps lookup but DP tracks valid prefixes.
**Time:** O(n² × k) where k = max word length, or O(n²) with hash set
**Pattern:** Dynamic Programming
**Key insight:** Must track which prefixes are valid, not just look up substrings.
**Discrete Math:** Formal languages (string decomposition from dictionary), Recurrence relations

## LC152 — Maximum Product Subarray
**User approach:** Sliding window, move left past negative number. / Prefix products - carry forward min and max products at every index.
**Refined approach:** DP tracking both max and min product at each position. new_max = max(nums[i], prev_max * nums[i], prev_min * nums[i]), same for min. Double negatives flip sign — min can become max.
**Time:** O(n)
**Pattern:** Dynamic Programming
**Key insight:** Negative × negative = positive. Track both extremes at each position.
**Discrete Math:** Real numbers (sign analysis), DP recurrence, Monotonicity

## LC5 — Longest Palindromic Substring
**User approach:** Expand around center. Odd length: expand from single char. Even length: find repeating chars and expand from pair.
**Time:** O(n²)
**Pattern:** Two Pointers / Expand Around Center
**Key insight:** Every palindrome has a center. Check all centers, expand outward while characters match.
**Discrete Math:** Symmetry (palindrome = invariant under reversal), Enumeration (2n-1 centers)

## LC33 — Search in Rotated Sorted Array
**User approach:** BST to find rotation point, then BST on right half if target between rotation point and end, or left half.
**Refined approach:** Single pass. At each mid, one half is always sorted. Check if target is in the sorted half, go there. Otherwise go to the other half.
**Time:** O(log n)
**Pattern:** Binary Search
**Key insight:** Rotation guarantees one half is always sorted. Use that to decide which half to discard.
**Discrete Math:** Binary search, Invariant (sortedness of at least one half), Total order

## LC73 — Set Matrix Zeroes
**User approach:** DFS.
**Refined approach:** Use first row/column as markers. Scan for zeros, mark their row/col, then zero based on marks. O(1) space.
**Time:** O(m×n)
**Pattern:** In-place Matrix Manipulation
**Key insight:** First row/column can double as marker storage. Process markers last to avoid overwriting.
**Discrete Math:** Matrix theory (in-place transformation), Set theory (marker sets)

## LC11 — Container With Most Water
**User approach:** Monotonically increasing stack.
**Refined approach:** Two pointers from both ends. Move the shorter one inward — it's the bottleneck, moving the taller can never help.
**Time:** O(n)
**Pattern:** Two Pointers
**Key insight:** The shorter line limits the area. Moving the taller line inward can only decrease width without increasing height.
**Discrete Math:** Invariant (shorter line = bottleneck), Optimization, Two-pointer reasoning

## LC207 — Course Schedule
**User approach:** BFS with set for visited nodes.
**Refined approach:** Topological sort (Kahn's). Track in-degrees, only visit nodes with 0 prerequisites remaining. If not all nodes visited, cycle exists.
**Time:** O(V + E)
**Pattern:** Graph / Topological Sort
**Key insight:** BFS alone can't detect cycles. In-degree tracking ensures prerequisites are satisfied before visiting.
**Discrete Math:** Graph theory (DAG), Partial order (prerequisites), Topological order

## LC146 — LRU Cache
**User approach:** Map for key-values, LIFO queue to track recency.
**Refined approach:** Map + doubly linked list. Map points to nodes, list tracks recency. Move accessed item to head in O(1). Evict from tail.
**Time:** O(1) per operation
**Pattern:** Design / Hash Map + Linked List
**Key insight:** Doubly linked list allows O(1) removal from middle. Stack can't do that.
**Discrete Math:** Data structure design, Map + linked list

## LC102 — Binary Tree Level Order Traversal
**User approach:** BFS with queue.
**Refined approach:** Same. Track level size at each step, process that many nodes per level.
**Time:** O(n)
**Pattern:** BFS / Tree Traversal
**Key insight:** Queue size at each iteration = number of nodes at that level.
**Discrete Math:** Tree traversal, BFS, Level sets (partition of nodes by depth)

## LC322 — Coin Change
**User approach:** Combo recursion.
**Refined approach:** Bottom-up DP. dp[i] = min(dp[i], dp[i-coin] + 1) for each coin. Build from 0 to amount.
**Time:** O(amount × coins)
**Pattern:** Dynamic Programming
**Key insight:** Each amount depends on smaller amounts. Build up, don't recurse down.
**Discrete Math:** DP, Unbounded knapsack, Optimization (minimization)

## LC98 — Validate Binary Search Tree
**User approach:** In-order traversal.
**Refined approach:** Same. Track previous value during in-order traversal. If current ≤ previous, invalid BST.
**Time:** O(n)
**Pattern:** Tree / In-order Traversal
**Key insight:** In-order of BST is strictly increasing. One pass, one check.
**Discrete Math:** BST, In-order traversal, Total order (strictly increasing sequence)

## LC133 — Clone Graph
**User approach:** Build adjacency lists and copy.
**Refined approach:** BFS or DFS with a map (original → clone). Avoid cycles by checking if node already cloned.
**Time:** O(V + E)
**Pattern:** Graph BFS/DFS
**Key insight:** Map prevents infinite loops on cycles. Clone on first visit, reuse on revisit.
**Discrete Math:** Graph theory (undirected graph, cycles), DFS/BFS

## LC153 — Find Minimum in Rotated Sorted Array
**User approach:** BST on right half — if right half larger than midpt, BST on left.
**Refined approach:** Same. Compare mid with right boundary. If mid > right, min is right half. Otherwise left half (or mid).
**Time:** O(log n)
**Pattern:** Binary Search
**Key insight:** One comparison with right boundary is enough to decide which half to keep.
**Discrete Math:** Binary search, Invariant (comparison with right boundary)

## LC416 — Partition Equal Subset Sum
**User approach:** Find any combination that sums to sum(array)/2.
**Refined approach:** 0/1 knapsack DP. dp[i] = true if sum i is achievable. dp[0] = true, iterate elements, update backwards.
**Time:** O(n × sum/2)
**Pattern:** Dynamic Programming / Knapsack
**Key insight:** Partition equal sum = subset sum to total/2. One boolean array, iterate backwards to avoid reuse.
**Discrete Math:** Subset sum, 0/1 knapsack, Partition problem (NP-complete, pseudo-polynomial)

## LC236 — Lowest Common Ancestor of Binary Tree
**User approach:** Prefix traversal.
**Refined approach:** Recursion. If current is p or q, return it. If both left and right return non-nil, current is LCA. If one side returns non-nil, propagate it up.
**Time:** O(n)
**Pattern:** Tree Recursion
**Key insight:** LCA is the first node where p and q diverge into different subtrees. Recursion naturally finds it.
**Discrete Math:** Tree theory, Ancestor relation (partial order), Recursion

## LC19 — Remove Nth Node From End of List
**User approach:** Fast pointer starts n ahead of slow pointer.
**Refined approach:** Same. Move fast n steps, then both together. When fast hits end, slow is at node before target.
**Time:** O(n)
**Pattern:** Two Pointers / Linked List
**Key insight:** One pass with two pointers spaced n apart. No need to compute list length.
**Discrete Math:** Linked list, Two-pointer technique, Offset (n steps apart)

## LC3 — Longest Substring Without Repeating Characters
**User approach:** Sliding window with map to track seen. / Sliding window - extend right while char uniq, shrink left while left matches newly added char.
**Refined approach:** Sliding window with a hash map tracking last seen index of each char. Expand right; when duplicate hit, jump left past previous occurrence (not incrementally). Update max length.
**Time:** O(n)
**Pattern:** Sliding Window / Hash Map
**Key insight:** Map stores last seen index of each character. Jump left directly past the duplicate, no incremental shrink.
**Discrete Math:** Sliding window, Invariant (window has all unique chars), Map (last seen index)

## LC208 — Implement Trie
**User approach:** Tree with [26]char at each level.
**Refined approach:** Same. Each node: array of 26 children + isEnd flag. Insert traverses/creates, search traverses checking isEnd, startsWith traverses without checking isEnd.
**Time:** O(m) per operation where m = word length
**Pattern:** Design / Trie
**Key insight:** Fixed-width array at each node gives O(1) child lookup. isEnd distinguishes prefixes from complete words.
**Discrete Math:** Trie (prefix tree), Formal languages, Deterministic finite automaton (DFA)

## LC76 — Minimum Window Substring
**User approach:** Sliding window — expand right until all chars found, shrink from left to second match, continue expanding until dropped match found again.
**Refined approach:** Expand right until window valid, shrink left while still valid, update min. Repeat. No need to track specific matches — just check counts.
**Time:** O(n)
**Pattern:** Sliding Window / Hash Map
**Key insight:** Track required vs available counts. Shrink while all requirements met, expand when broken.
**Discrete Math:** Sliding window, Counting/multiset, Invariant (window contains all required chars)

## LC148 — Sort List
**User approach:** Fast pointer 2 ahead of slow to find midpoint, merge sort.
**Refined approach:** Same. Split at midpoint, recursively sort halves, merge sorted halves.
**Time:** O(n log n)
**Pattern:** Linked List / Merge Sort
**Key insight:** Merge sort is the natural fit for linked lists — split is O(1) with fast/slow pointers, merge is O(n) with dummy head.
**Discrete Math:** Divide and conquer, Merge sort, Comparison-based sorting (Ω(n log n) lower bound)

## LC1481 — Least Number of Unique Integers after K Removals
**User approach:** Freq map and priority queue ordered by freq asc, remove from head, return length of remaining queue.
**Time:** O(n log n)
**Pattern:** Greedy / Heap
**Key insight:** Remove rarest first. Sort or min-heap on frequencies.
**Discrete Math:** Greedy algorithm, Frequency (multiset), Sorting

## LC22 — Generate Parentheses
**User approach:** Generate permutations with 2ary state for open/close, only close if fewer opens remain.
**Time:** O(4^n / √n) — Catalan number
**Pattern:** Backtracking
**Key insight:** Two counters constrain the tree. close < open to add `)`, open < n to add `(`. Prunes invalid branches early.
**Discrete Math:** Catalan numbers, Dyck words, Combinatorics, Backtracking

## LC39 — Combination Sum
**User approach:** Combo recursion.
**Refined approach:** Backtracking. Try each candidate, recurse with remaining target, backtrack when target < 0 or == 0. Reuse allowed.
**Time:** O(n^(target/min)) — exponential, not n²
**Pattern:** Backtracking
**Key insight:** Reuse means recursion tree depth isn't bounded by array length. Prune when current sum exceeds target.
**Discrete Math:** Combinatorics (combinations with repetition/unbounded), Backtracking, Pruning

## LC48 — Rotate Image
**User approach:** Swap corners going clockwise with dummy var, move inward and repeat.
**Refined approach:** Transpose (swap across diagonal), then reverse each row. Same O(n²), easier to implement.
**Time:** O(n²)
**Pattern:** Matrix Manipulation
**Key insight:** Rotate = transpose + reverse. Two simple ops instead of tracking four-way swaps.
**Discrete Math:** Matrix transformation (transpose), Permutation (rotation = composition of two involutions)

## LC84 — Largest Rectangle in Histogram
**User approach:** Monotonically increasing stack of indices - compute area with smallest element in stack when invariant breaks.
**Refined approach:** Monotonic stack of indices. Push each index. When current height < stack top, pop and compute area: height[popped] × width (current index - new top - 1, or full width if stack empty). Push current index. Add sentinel (height 0) at end to flush remaining.
**Time:** O(n)
**Pattern:** Monotonic Stack
**Key insight:** Each popped index defines a rectangle that extends left (previous smaller) and right (current smaller). The stack maintains "next smaller to the left" — when it breaks, you know the rectangle's boundaries.
**Discrete Math:** Monotonic stack, Next smaller element, Cartesian product (height × width)

## LC56 — Merge Intervals
**User approach:** Sort intervals, sweep line to merge potentially multiple overlapping intervals.
**Refined approach:** Sort by start time. Sweep left to right — if current start ≤ previous end, merge (extend end to max). Otherwise push new interval.
**Time:** O(n log n)
**Pattern:** Intervals / Sweep Line
**Key insight:** Sorting guarantees you only need to look at the previous interval — no backtracking needed. Merge or push, never skip.
**Discrete Math:** Interval graph, Sweep line, Partial order (overlap relation), Sorting

## LC31 — Next Permutation
**User approach:** Permutations/backtracking starting from last index.
**Refined approach:** Scan right to left for first index where nums[i] < nums[i+1]. Swap nums[i] with the smallest element to its right that's larger than it. Reverse the suffix after i.
**Time:** O(n)
**Pattern:** Array Manipulation
**Key insight:** Backtracking gives all permutations in O(n!) — overkill. One linear scan + one swap + one reverse gives the next permutation directly.
**Discrete Math:** Permutations, Lexicographic order, Combinatorics (n! total, next in order)

## LC394 — Decode String
**User approach:** Recursively find inner most char - duplicate as necessary - and repeat going up the call stack - do this for all distinct chunks/groups of encoded chars.
**Refined approach:** Two stacks (or recursion). Push numbers and strings onto stacks when hitting `[`. Pop and repeat when hitting `]`. Iterative version uses two stacks — cleaner than recursion for this.
**Time:** O(n × max_repeat)
**Pattern:** Stack
**Key insight:** `[` saves state (current string and repeat count) to stacks, `]` restores and applies. Same idea as recursion but explicit.
**Discrete Math:** Context-free grammar (nested brackets), Stack automaton, Recursion

## LC49 — Group Anagrams
**User approach:** Freq_map of chars per string - O(n^2) iteration over all maps for full intersection of frequencies.
**Refined approach:** Hash map with sorted string as key (or tuple of char counts as key). Same anagram → same key → group them. O(nk log k) with sorting, O(nk) with fixed-size count array.
**Time:** O(nk log k) or O(nk)
**Pattern:** Hash Map
**Key insight:** Your approach works but is O(n²) because you compare every pair. Sorting or counting as the key groups in one pass — no pairwise comparison needed.
**Discrete Math:** Equivalence relation (anagram ∼), Partition into equivalence classes, Hash map (canonical form)

## LC155 — Min Stack
**User approach:** Linked list -> pop removes head and repoints -> push updates head to point to old head -> top gets head -> track min separately in encounter order.
**Refined approach:** Two stacks — main stack + min stack. Push to min stack whenever new value ≤ current min. Pop from both. Min stack top is always current minimum.
**Time:** O(1) all ops
**Pattern:** Design / Stack
**Key insight:** Your linked list idea works too — essentially the same thing. Min stack approach is cleaner to implement. The ≤ on push (not <) handles duplicates correctly.
**Discrete Math:** Data structure design, Stack (LIFO), Min-tracking (auxiliary data structure)

## LC287 — Find the Duplicate Number
**User approach:** Cycle detection - treating nums as a linked list but not sure how to walk the array.
**Refined approach:** Treat nums[i] as next pointer from node i. Start at index 0, use Floyd's cycle detection (slow = nums[slow], fast = nums[nums[fast]]). Cycle entrance = duplicate.
**Time:** O(n)
**Pattern:** Array / Floyd's Cycle Detection
**Key insight:** No extra structure needed — indices ARE your pointers. In-place, O(1) space.
**Discrete Math:** Pigeonhole principle, Functional graph (each node has exactly one outgoing edge), Cycle detection

## LC54 — Spiral Matrix
**User approach:** Nested traversal?
**Refined approach:** Four boundaries (top, bottom, left, right). Traverse top row, right col, bottom row, left col. Shrink boundaries inward, repeat.
**Time:** O(m×n)
**Pattern:** Matrix Manipulation
**Key insight:** Each traversal shrinks the "ring" — boundaries are your state, not recursion.
**Discrete Math:** Matrix traversal, Invariant (boundary coordinates), Loop invariant

## LC78 — Subsets
**User approach:** Combo recursion - no duplicate subsets - pruned naturally if going left to right in nums.
**Refined approach:** Backtracking. At each index, include or exclude. Or simpler: at each step, pick from remaining elements to the right, recurse. Natural dedup because you never look left.
**Time:** O(2^n)
**Pattern:** Backtracking
**Key insight:** Same as combo/permutation backtracking — just don't track a target. Every node in the tree is a valid subset.
**Discrete Math:** Power set, Combinatorics (2ⁿ subsets), Binary representation (include/exclude bit)

## LC46 — Permutations
**User approach:** Recursion for nums i to end.
**Refined approach:** Backtracking. At each position, swap each remaining element into it, recurse, swap back. Or: build a new list by inserting current element at every position in the recursive result.
**Time:** O(n × n!)
**Pattern:** Backtracking
**Key insight:** Two clean implementations — swap-in-place or build-new-list. Both give the same result. Swap-in-place is O(1) space, build-new-list is cleaner to reason about.
**Discrete Math:** Combinatorics (n! permutations), Symmetric group, Backtracking

## LC104 — Maximum Depth of Binary Tree
**User approach:** In-order traversal - return 1+ (next level) - return total to parent call.
**Refined approach:** Recursion. return 1 + max(maxDepth(left), maxDepth(right)). Or BFS counting levels.
**Time:** O(n)
**Pattern:** Tree Recursion
**Key insight:** In-order isn't required — any traversal works. DFS is simplest: base case is null (return 0), otherwise 1 + max of children.
**Discrete Math:** Tree recursion, Height (longest root-to-leaf path)

## LC136 — Single Number
**User approach:** Compare left and right pointers with center elem -> scan inward from left/right.
**Refined approach:** XOR all elements. Pairs cancel (a ^ a = 0), single number remains. result ^= nums[i] for all i.
**Time:** O(n)
**Pattern:** Bit Manipulation
**Key insight:** XOR is its own inverse. a ^ a = 0, a ^ 0 = a. Order doesn't matter.
**Discrete Math:** XOR, Group theory (abelian group under XOR, inverse property), Bit operations

## LC2 — Add Two Numbers
**User approach:** Traverse lists in parallel -> propagate carry number to next and create new tail if carry causes "overflow".
**Refined approach:** Same. Iterate both lists, sum digits + carry, create new node with sum % 10, carry = sum / 10. Continue until both lists exhausted and carry is 0.
**Time:** O(max(m, n))
**Pattern:** Linked List / Math
**Key insight:** Reverse order is a gift — you process least significant digit first, which is exactly what addition needs.
**Discrete Math:** Linked list traversal, Base-10 arithmetic, Carry propagation

## LC198 — House Robber
**User approach:** Prefix sum that skips an index O(n).
**Refined approach:** DP. At each house, max of (skip it = prev) or (rob it = prev_prev + nums[i]). Two variables, no array needed.
**Time:** O(n), O(1) space
**Pattern:** Dynamic Programming
**Key insight:** You only need the last two states. Not prefix sums — it's a decision at each step, not a guaranteed include.
**Discrete Math:** DP, Recurrence relation (Fibonacci-like: f(n) = max(f(n-1), f(n-2) + aₙ))

## LC300 — Longest Increasing Subsequence
**User approach:** DP - state is length of sequence at i and last number in sequence.
**Refined approach:** O(n²) DP: dp[i] = length of LIS ending at i, check all j < i. O(n log n): patience sorting with binary search on tails array.
**Time:** O(n²) or O(n log n)
**Pattern:** Dynamic Programming / Binary Search
**Key insight:** Your state description is correct — that's the O(n²) DP. The O(n log n) optimization uses binary search on tails instead of scanning all previous elements.
**Discrete Math:** Partial order (increasing subsequence), DP, Patience sorting

## LC141 — Linked List Cycle
**User approach:** Fast pointer 1 ahead of slow - fast falls behind slow in cycle - but how to detect?
**Refined approach:** Slow moves 1 step, fast moves 2 steps. If fast or fast.next is null → no cycle. If they meet → cycle.
**Time:** O(n), O(1) space
**Pattern:** Fast & Slow Pointers
**Key insight:** In a cycle, fast gains 1 step per iteration on slow. It's guaranteed to lap slow within one cycle length.
**Discrete Math:** Functional graph, Floyd's cycle detection (tortoise and hare), Invariant (fast gains 1 step/iteration)

## LC560 — Subarray Sum Equals K
**User approach:** Combo recursion.
**Refined approach:** Prefix sum + hash map. Track cumulative sum, for each index check if sum - k existed before. Map stores count of each prefix sum.
**Time:** O(n)
**Pattern:** Prefix Sum / Hash Map
**Key insight:** Combo recursion is O(2^n). Prefix sum reduces it to one pass — each position asks "have I seen the sum that makes this subarray equal k?"
**Discrete Math:** Prefix sum, Congruence (cumulative sum - k), Map (count of prefix values)

## LC494 — Target Sum
**User approach:** dp where dp[i] = [s | s in dp[i-1]] ++ [s-dp[i]] ++ [s+dp[i]] O(2^n).
**Refined approach:** Convert to subset sum. P = (total + target) / 2. Count subsets that sum to P. Recurrence: dp[0] = 1; for each num, for j from P down to num: dp[j] += dp[j - num]. Backwards iteration prevents reuse.
**Time:** O(n × sum)
**Pattern:** Dynamic Programming / Knapsack
**Key insight:** Assigning +/- is the same as picking a "positive" subset. Reframing as subset sum reduces exponential to O(n × sum). **FLAG: Review before interview — user wants to revisit.**
**Discrete Math:** Subset sum, 0/1 knapsack, Transformation (+/- assignment → positive subset)

## LC1091 — Shortest Path in Binary Matrix
**User approach:** DFS - prioritize moving down vertically or diagonally.
**Refined approach:** BFS from (0,0), explore all 8 neighbors. Track visited cells. Return distance when you reach (n-1, n-1).
**Time:** O(n²)
**Pattern:** BFS
**Key insight:** BFS explores level by level — the first time you reach the destination is guaranteed shortest. DFS doesn't have that guarantee.
**Discrete Math:** Graph theory (grid graph, 8-directional), BFS, Unweighted shortest path

## LC981 — Time Based Key-Value Store
**User approach:** Sorted values by timestamp desc - priority queue/heap - get returns max for key.
**Refined approach:** Hash map of key → sorted list of (timestamp, value) pairs. On get, binary search the list for largest timestamp ≤ target.
**Time:** O(log n) per get, O(1) amortized per set
**Pattern:** Hash Map + Binary Search
**Key insight:** Each key's values are naturally inserted in timestamp order, so the list is already sorted. Binary search is the right tool for "largest ≤ target."
**Discrete Math:** Total order (timestamps), Binary search (predecessor query), Map of sorted lists

## LC116 — Populating Next Right Pointers in Each Node
**User approach:** BFS - each child points to next in level.
**Refined approach:** BFS works, O(n) space. But O(1) space possible — use already-established next pointers to traverse current level while setting up next level. No queue needed.
**Time:** O(n), O(1) space optimal
**Pattern:** Tree / BFS
**Key insight:** Perfect binary tree means the next pointers you've already set give you a free linked list to traverse when setting up the next level.
**Discrete Math:** Tree theory (perfect binary tree), BFS, Level-order traversal, Linked list

## LC347 — Top K Frequent Elements
**User approach:** Freq_map by linear scan - bucket array to get top k.
**Refined approach:** Same. Count frequencies in O(n). Bucket sort — bucket[i] = list of nums with frequency i. Walk buckets from highest frequency down, collect k elements.
**Time:** O(n)
**Pattern:** Bucket Sort / Hash Map
**Key insight:** Bucket sort avoids O(n log n) sorting. Frequency is bounded by n, so you get O(n) guaranteed.
**Discrete Math:** Counting sort / Bucket sort, Frequency distribution, Selection

## LC994 — Rotting Oranges
**User approach:** BFS starting at cells with 2's and flipping 1's - when no children left O(m*n) scan for 1s.
**Refined approach:** Multi-source BFS from all 2s simultaneously. Track fresh orange count. Each time a 1 rots, decrement count. When BFS ends, if count > 0 return -1, else return minutes elapsed.
**Time:** O(m×n)
**Pattern:** BFS
**Key insight:** The O(m×n) scan at the end is unnecessary — just track the fresh count as you go. When BFS finishes, one comparison tells you if it's possible.
**Discrete Math:** Graph theory (grid graph, 4-directional), Multi-source BFS, Parallel propagation

## LC230 — Kth Smallest Element in a BST
**User approach:** In-order traversal to find smallest, decrement k on the way up.
**Refined approach:** Same. In-order traversal (left, root, right). Decrement k each time you visit a node. When k hits 0, that's your answer.
**Time:** O(n) worst case, O(k) average
**Pattern:** Tree / In-order Traversal
**Key insight:** In-order of BST is sorted. You don't need to traverse the whole tree — stop as soon as k hits 0.
**Discrete Math:** BST, In-order traversal, Total order (sortedness), Order statistics (kth smallest)

## LC105 — Construct Binary Tree from Preorder and Inorder Traversal
**User approach:** Start with preorder[0] add inorder[0] to left child skip preorder[1] add preorder[2] as right child, alternate skipping arrays for left and right children.
**Refined approach:** First element of preorder is always the root. Find that root in inorder — everything left is left subtree, everything right is right subtree. Split both arrays at that point and recurse.
**Time:** O(n)
**Pattern:** Tree Recursion
**Key insight:** Preorder gives you the root. Inorder tells you the size of left/right subtrees. Use the left subtree size to split the preorder array correctly.
**Discrete Math:** Tree reconstruction, Invariant (preorder = root first, inorder = left/root/right), Recursion

## LC55 — Jump Game
**User approach:** Create intervals from each index and sort by end of intervals, look for disjoint intervals.
**Refined approach:** Greedy. Track max_reachable. At each index i, if i > max_reachable, return false. Otherwise, update max_reachable = max(max_reachable, i + nums[i]).
**Time:** O(n)
**Pattern:** Greedy
**Key insight:** You don't need to track which path — just track the farthest you can possibly reach. If every index up to the end is reachable, you're good.
**Discrete Math:** Greedy algorithm, Reachability (in directed path graph), Monotonic (max reachable)

## LC75 — Sort Colors
**User approach:** DNF 3-way partition.
**Refined approach:** Dutch National Flag. Three pointers: low, mid, high. Regions: 0 to low-1 are all 0s, low to mid-1 are all 1s, mid to high are unexamined, high+1 to end are all 2s. If nums[mid] == 0, swap with low, advance both. If nums[mid] == 1, advance mid. If nums[mid] == 2, swap with high, decrement high.
**Time:** O(n), one pass
**Pattern:** Two Pointers / In-place Partition
**Key insight:** Three-way partition in a single pass. Mid walks forward, low and high converge from ends.
**Discrete Math:** Dutch national flag partitioning (3-way), Loop invariant (three regions maintained)

## LC210 — Course Schedule II
**User approach:** BFS with set to track seen.
**Refined approach:** Build adjacency list and in-degree array. Start with all 0-in-degree nodes. BFS: pop node, add to result, decrement neighbors' in-degrees. If neighbor hits 0, enqueue. If result length != numCourses, cycle exists.
**Time:** O(V + E)
**Pattern:** Topological Sort / BFS
**Key insight:** Regular BFS with a visited set can revisit nodes. In-degree tracking ensures you only visit each node once prerequisites are satisfied.
**Discrete Math:** Graph theory (DAG), Topological sort (linear extension of partial order), Kahn's algorithm

## LC62 — Unique Paths
**User approach:** f(m, n) = f(m+1, n) U f(m, n+1) with boundary checks.
**Refined approach:** Same idea as recursion, but bottom-up DP. dp[i][j] = dp[i-1][j] + dp[i][j-1]. Or compress to 1D — just one row, update in place.
**Time:** O(m×n), O(n) space with compression
**Pattern:** Dynamic Programming
**Key insight:** Recursion + memoization works but bottom-up avoids stack overhead. 1D compression works because each cell only depends on top and left.
**Discrete Math:** DP, Combinatorics (C(m+n-2, m-1) paths in grid), Pascal's triangle (additive recurrence)

## LC131 — Palindrome Partitioning
**User approach:** Recursively find palindromes left to right - i from 0 to end.
**Refined approach:** Backtracking. At each position, try all substrings starting at i. If palindrome, add to path, recurse from i+len. When you reach end of string, record the partition.
**Time:** O(n × 2^n) worst case
**Pattern:** Backtracking
**Key insight:** Same tree structure as subsets/combo — just with a palindrome check as the gate for valid branches.
**Discrete Math:** Palindrome (symmetry), Set partition (of string indices), Backtracking

## LC199 — Binary Tree Right Side View
**User approach:** DFS following right children to leaf.
**Refined approach:** DFS right-first, track depth. First node at each depth is the rightmost. Or BFS level-order, take last node per level.
**Time:** O(n)
**Pattern:** Tree DFS/BFS
**Key insight:** Just following right children misses nodes in left subtree when right subtree is shorter. Track depth to capture deeper visible nodes.
**Discrete Math:** Tree traversal (right-first DFS), Depth tracking, Level sets

## LC235 — Lowest Common Ancestor of a Binary Search Tree
**User approach:** Find p and q and return them - the node higher up in the recursive call stack that gets non-nil vals for both is the LCA.
**Refined approach:** Use BST property. If both p and q are less than root, go left. If both greater, go right. Otherwise, root is the LCA — you've found the split point.
**Time:** O(h)
**Pattern:** Tree / BST
**Key insight:** BST ordering means you don't need to search both subtrees — one comparison tells you which direction to go. No need to check for nulls.
**Discrete Math:** BST, Binary search, Total order (split point where p and q diverge)

## LC79 — Word Search
**User approach:** Backtracking.
**Refined approach:** DFS + backtracking. At each cell, try all 4 directions. Mark visited, recurse, unmark on return. Early exit when all chars matched.
**Time:** O(m × n × 4^L) where L = word length
**Pattern:** Backtracking / Matrix DFS
**Key insight:** The exponential factor is 4^L not 4^(m×n) because you stop at word length. Unmarking (backtracking) is what allows revisiting cells on different paths.
**Discrete Math:** Graph theory (grid graph, 4-directional), Backtracking, DFS, Path finding

## LC221 — Maximal Square
**User approach:** Backtracking - match 1x1 - then look for 2x2 etc.
**Refined approach:** DP. dp[i][j] = side length of largest square ending at (i,j). If matrix[i][j] == '1', dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1. Track max side, return area.
**Time:** O(m×n)
**Pattern:** Dynamic Programming
**Key insight:** A square at (i,j) can only exist if there are squares above, left, and diagonally above-left. The minimum of those three + 1 gives the largest square ending here. Use bottom-up 1D DP to avoid stack overflow.
**Discrete Math:** DP, Min of three neighbors (optimal substructure), Inclusion–exclusion principle

## LC91 — Decode Ways
**User approach:** Count combinations where s is traversed in order and a single digit can be reused - pruning for numbers 1-26.
**Refined approach:** DP. dp[i] = number of ways to decode s[i:]. If s[i] != '0', add dp[i+1]. If s[i:i+2] is between 10-26, add dp[i+2]. Build from end to start.
**Time:** O(n)
**Pattern:** Dynamic Programming
**Key insight:** Backtracking is exponential. DP with two states back is O(n) — at each position you have at most two choices (1 digit or 2 digits), and you only need the next two values.
**Discrete Math:** DP, Recurrence relation (Fibonacci-like), String encoding (prefix codes)

## LC437 — Path Sum III
**User approach:** DFS - flag which child (R/L) following for current path - tally count if sum = target.
**Refined approach:** DFS with prefix sum map. At each node, current_sum - target tells you if a path ending here exists. Add current_sum to map, recurse left and right, remove from map (backtrack).
**Time:** O(n)
**Pattern:** DFS / Prefix Sum
**Key insight:** Same idea as LC560 (Subarray Sum Equals K) — but on a tree. The map tracks prefix sums along the current path. Backtracking the map when you leave a node ensures you only count paths in the current branch.
**Discrete Math:** Prefix sum, Tree DFS, Backtracking (map state restoration), Congruence (current - target)

## LC1029 — Two City Scheduling
**User approach:** (corrected example) Sort by difference between city costs, send top n to cheaper city.
**Refined approach:** Sort by the difference cost_A - cost_B. Send the first n (biggest A savers) to A, rest to B.
**Time:** O(n log n)
**Pattern:** Greedy
**Key insight:** The person with the biggest savings going to A (vs B) should go to A. Sort by difference, take top n for A.
**Discrete Math:** Greedy algorithm, Sorting by difference (comparator), Optimization (minimization of sum)

## LC138 — Copy List with Random Pointer
**User approach:** 2 pass copy - second pass sets random pointers in clone.
**Refined approach:** Same. Pass 1: create cloned nodes, map original → clone. Pass 2: set clone.next = map[original.next] and clone.random = map[original.random].
**Time:** O(n)
**Pattern:** Hash Map / Linked List
**Key insight:** Map eliminates the need to search for random targets. Two passes — one to build nodes, one to wire pointers.
**Discrete Math:** Linked list (with random pointers), Bijection (original ↔ clone via map)

## LC1143 — Longest Common Subsequence
**User approach:** Map chars to positions for text2 - iterate over text1, looking up chars for text2 in map and backtrack for all shared chars with increasing indices.
**Refined approach:** DP. dp[i][j] = LCS of text1[i:] and text2[j:]. If chars match, dp[i][j] = 1 + dp[i+1][j+1]. Else, dp[i][j] = max(dp[i+1][j], dp[i][j+1]).
**Time:** O(m×n)
**Pattern:** Dynamic Programming
**Key insight:** Backtracking over all shared chars is exponential. DP with two indices is O(m×n) — at each position, you either match or skip one char from either string.
**Discrete Math:** DP, Sequence alignment, Subsequence relation (partial order on positions)

## LC739 — Daily Temperatures
**User approach:** Monotonically increasing stack of indices - for any item left in stack - set their temperatures[i] to zero.
**Refined approach:** Monotonic stack of indices. Push each index. When current temp > stack top, pop and compute answer[popped] = i - popped. Items left in stack at the end already have 0 (default).
**Time:** O(n)
**Pattern:** Monotonic Stack
**Key insight:** Same idea as LC84 — stack maintains "next greater element" to the right. When you find one, compute the distance and pop.
**Discrete Math:** Monotonic stack, Next greater element, Order statistics (distance to next greater)

## LC100 — Same Tree
**User approach:** Traverse in parallel and compare values - order doesn't matter but should be consistent for both trees.
**Refined approach:** Same. Recursion: return p.val == q.val && sameTree(p.left, q.left) && sameTree(p.right, q.right). Base case: both null → true, one null → false.
**Time:** O(n)
**Pattern:** Tree Recursion
**Key insight:** Order doesn't matter as long as both trees use the same traversal. Pre-order is simplest — check current, then recurse left and right.
**Discrete Math:** Tree theory, Structural equality (recursive definition), Isomorphism of rooted trees

## LC121 — Best Time to Buy and Sell Stock
**User approach:** Generating sequences where each elem starts a new sequence if it can't be appended to an existing because it's smaller than last of every sequence so far.
**Refined approach:** Track min price so far. At each day, calculate prices[i] - min_so_far, update max profit.
**Time:** O(n), O(1) space
**Pattern:** Greedy / Single Pass
**Key insight:** You only need to know the cheapest price you've seen up to now. One pass, two variables.
**Discrete Math:** Greedy algorithm, Min/max tracking, Order (buy before sell → causality constraint)

## LC463 — Island Perimeter
**User approach:** DFS keeping count of sides, subtract one for every shared border below or to left/right (level-wise).
**Refined approach:** Each land cell starts with 4 sides. For each land cell, check 4 neighbors — if neighbor is also land, subtract 1 (shared border). Total = sum of all adjusted sides.
**Time:** O(m×n)
**Pattern:** Matrix DFS
**Key insight:** 4 minus number of adjacent land cells, summed over all land cells.
**Discrete Math:** Grid graph, Counting (edge contributions), Adjacency (4-neighbor relation)

## LC226 — Invert Binary Tree
**User approach:** Depth-first - swap children on the way up.
**Refined approach:** Same. Recursion: invert left, invert right, swap left and right. Base case: null → return null.
**Time:** O(n)
**Pattern:** Tree Recursion
**Key insight:** Post-order works naturally — children are inverted before you swap them. Pre-order also works (swap then recurse). Either way, every node gets exactly one swap.
**Discrete Math:** Tree recursion, Post-order traversal, Involution (inverting twice = identity)

## LC1024 — Video Stitching
**User approach:** Sort by start time, greedy backtracking.
**Refined approach:** Sort by start time. Greedy: always extend to the clip that reaches furthest while staying within current coverage. Count clips used.
**Time:** O(n log n)
**Pattern:** Greedy / Intervals
**Key insight:** No backtracking needed — pure greedy. Extending coverage as far as possible at each step is always optimal. Picking shorter leaves less room for future clips.
**Discrete Math:** Greedy algorithm, Interval covering (set cover special case), Sorting by start time

## LC1249 — Minimum Remove to Make Valid Parentheses
**User approach:** Use stack to track invalid parens positions.
**Refined approach:** Stack for unmatched ( positions. Also track unmatched ) — just skip them (don't push). At the end, remove any remaining unmatched ( from stack. Build result excluding those indices.
**Time:** O(n)
**Pattern:** Stack
**Key insight:** You only need to track ( positions — unmatched ) are simply ignored during the scan. Stack tells you which ( to remove at the end.
**Discrete Math:** Stack automaton, Context-free language (Dyck language), Parenthesis matching

## LC1539 — Kth Missing Positive Number
**User approach:** Iteration - add differences between neighboring elems until the kth is found or the next gap that contains k is found.
**Refined approach:** Binary search. At index i, the number of missing integers before arr[i] is arr[i] - i - 1. Find the index where missing count crosses k.
**Time:** O(log n)
**Pattern:** Binary Search
**Key insight:** Gap-walking is O(n). Binary search on the missing count formula gives O(log n). Missing before arr[i] = arr[i] - (i + 1).
**Discrete Math:** Binary search, Natural numbers, Gaps between consecutive integers

## LC334 — Increasing Triplet Subsequence
**User approach:** DP where dp[i][triplet_pos] = true if dp[i-1][triplet_pos-1] = true.
**Refined approach:** Greedy. Track `first` (smallest so far) and `second` (smallest value that is end of an increasing pair). If nums[i] > first → update first. Else if nums[i] > second → update second. Else → return true. First update after second is fine because second being set guarantees some smaller element existed before it at the time of update.
**Time:** O(n), O(1) space
**Pattern:** Greedy
**Key insight:** second and first don't need to come from the same subsequence — second being updated means a smaller element existed before it. So finding anything > second always yields a valid triplet.
**Discrete Math:** Partial order (increasing subsequence), Greedy, LIS special case (k=3)

## LC1423 — Maximum Points You Can Obtain from Cards
**User approach:** Start window at center, track sums on left and right, expand in direction of smaller sum.
**Refined approach:** Sliding window of "cards NOT taken." Take all k from right initially (window = first n-k cards). Slide window right: drop leftmost card, add next card. Current score = total - window_sum. Track max across all k+1 possible splits.
**Time:** O(n), O(1) space
**Pattern:** Sliding Window
**Key insight:** Taking i from left and k-i from right = skipping a contiguous block of n-k. Sliding that block left to right explores all valid splits in one pass.
**Discrete Math:** Sliding window, Complement (reverse thinking), Combinatorics (k from n ends)

## LC438 — Find All Anagrams in a String
**User approach:** Resetting freq_map of p, move window past char not in freq map, move window forward by one if window length matches p but no anagram found.
**Refined approach:** Fixed-size sliding window of len(p). Maintain window freq map and compare to p's freq map. Use a `matches` counter (how many of 26 letters have equal counts) to avoid full comparison each step. When a char not in p enters the window, that window is definitely not an anagram until it fully exits.
**Time:** O(n), O(1) space
**Pattern:** Sliding Window
**Key insight:** Window always slides by one — never skips. But chars not in p act as blockers that keep matches < 26 until they slide out. The `matches` counter makes each step O(1) instead of O(26).
**Discrete Math:** Sliding window, Frequency counting, Permutation (anagram = character multiset equality)

## LC735 — Asteroid Collision
**User approach:** Stack where top is popped if mag(top) < mag(current) and top and current have opposite signs, remove both if equal.
**Refined approach:** Push each asteroid. When current is negative and stack top is positive: if |top| < |current| → pop and repeat, if equal → pop both, if |top| > |current| → discard current. Collision only happens when positive is left of negative (moving toward each other). Negative top + positive current = moving apart, no collision.
**Time:** O(n), O(n) space
**Pattern:** Stack / Simulation
**Key insight:** The while loop is the tricky part — after popping a smaller positive, the incoming negative might still collide with the next positive on the stack. Each asteroid pushed/popped at most once.
**Discrete Math:** Stack, State machine (direction + magnitude), Simulation (sequential events)

## LC64 — Minimum Path Sum
**User approach:** DFS.
**Refined approach:** Bottom-up DP. dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]. First row/col only have one direction to come from. Space optimize to 1D array since each cell only depends on above and left.
**Time:** O(m×n), O(n) space
**Pattern:** Dynamic Programming / Grid
**Key insight:** DFS is O(2^(m+n)) without memo. DP with 1D compression — each cell depends only on cell above and cell to the left, update in place left to right, top to bottom.
**Discrete Math:** DP, Grid path, Optimization (minimization), 1D space compression

## LC1048 — Longest String Chain
**User approach:** Sort by word length asc, dp[i] = 0 if length 1 else 1+dp[i-1] if freq maps differ by one key or value.
**Refined approach:** Sort by length. dp[word] = longest chain ending at word. For each word, try removing each character — if result exists in the set, dp[word] = max(dp[word], dp[result] + 1). Track global max.
**Time:** O(n × L²) where n = words, L = max word length
**Pattern:** Dynamic Programming / String
**Key insight:** Predecessor = remove exactly one character to get the other word. Frequency map check doesn't work — "ab" and "ba" have same frequencies but neither is a predecessor. Removing one char at a time and checking existence is the correct test.
**Discrete Math:** DP, Partial order (predecessor relation by deleting one char), Sorting by length (topological order)

## LC1347 — Minimum Number of Steps to Make Two Strings Anagram
**User approach:** The difference of frequency maps.
**Refined approach:** Build freq maps for both strings. Sum of positive differences: for each character, if s has more than t, the excess must be replaced. Since lengths are equal, this equals half the sum of all absolute differences.
**Time:** O(n), O(1) space
**Pattern:** Hash Map / Counting
**Key insight:** Two strings of equal length are anagrams iff their frequency maps are identical. Count the gap between maps — that's the answer.
**Discrete Math:** Frequency counting, Anagram (multiset equality), L₁ distance (sum of absolute differences)

## LC380 — Insert Delete GetRandom O(1)
**User approach:** Map of val to index in positional array, tombstones in pos array, look-up table of tombstone intervals.
**Refined approach:** Map + dynamic array. Insert: append to array, store map[val] = index. Remove: swap element to remove with last element, update swapped element's index in map, pop from end, delete map entry. getRandom: random index in dense array.
**Time:** O(1) all ops, O(n) space
**Pattern:** Design / Hash Map + Array
**Key insight:** Swap-and-pop keeps the array dense — no tombstones, no gaps to track. Remove is O(1) because the map always has correct indices after the swap.
**Discrete Math:** Data structure design, Map + dynamic array, Swap-and-pop (compactify)

## LC567 — Permutation in String
**User approach:** Fixed window over s2 of size s1, count matching chars from s1.
**Refined approach:** Same as LC438. Fixed window of len(s1) sliding over s2. Maintain freq maps and compare. Use matches counter for O(1) per step. Return true on first match found.
**Time:** O(n), O(1) space
**Pattern:** Sliding Window
**Key insight:** Same problem as LC438 — just return true/false instead of all start indices.
**Discrete Math:** Sliding window, Frequency, Permutation (multiset equality), Anagram

## LC904 — Fruit Into Baskets
**User approach:** Longest subarray with at most two distinct integers, dp[i] = index after i where third unique integer is encountered.
**Refined approach:** Sliding window with a map. Expand right, add to map. When map size > 2, shrink left until map size ≤ 2. Track max window size. Map holds at most 3 entries.
**Time:** O(n), O(1) space
**Pattern:** Sliding Window
**Key insight:** "Two baskets" = at most 2 distinct values in the window. Classic variable-size sliding window — expand until constraint breaks, shrink to restore.
**Discrete Math:** Sliding window, At most 2 distinct (cardinality constraint), Multiset

## LC71 — Simplify Path
**User approach:** Split by / and push to stack, pop logic for .. and . gets complex.
**Refined approach:** Split by /. For each component: empty or "." → skip, ".." → pop stack if not empty, anything else → push. Join stack with "/" and prepend "/".
**Time:** O(n), O(n) space
**Pattern:** Stack
**Key insight:** Each component is independent — just push, pop, or skip. No counting or running string needed. Join at the end.
**Discrete Math:** Stack automaton, Formal language (canonical path grammar), Canonical form (unique representation)

## LC424 — Longest Repeating Character Replacement
**User approach:** Sliding window that expands to include k+1 varying characters, update freq map with every move.
**Refined approach:** Sliding window. Expand right, track max_freq of any character in window. If window_size - max_freq > k, shrink from left. max_freq never needs to decrease on shrink — answer only gets bigger with larger windows.
**Time:** O(n), O(1) space
**Pattern:** Sliding Window
**Key insight:** Replacements needed = window_size - most_frequent_count. Keep expanding while that's ≤ k. max_freq is monotonically non-decreasing — no need to shrink it.
**Discrete Math:** Sliding window, Frequency, Invariant (window - max_freq ≤ k), Monotonic (max_freq never decreases)

## LC1300 — Sum of Mutated Array Closest to Target
**User approach:** Sort array, if arr[0] >= target => val = target/len(arr), else multiply num[i] * len(array).
**Refined approach:** Sort + prefix sums. For each index i, try value = arr[i]. Sum = prefix[i] + (n - i) × arr[i]. Track closest to target. Or binary search value from 1 to max(arr), compute sum for each candidate.
**Time:** O(n log n), O(n) space for prefix sums
**Pattern:** Binary Search / Sorting
**Key insight:** After sorting, each candidate value replaces all elements >= it with itself. Prefix sums let you compute the resulting sum in O(1) per candidate.
**Discrete Math:** Sorting, Prefix sum, Binary search (find optimal mutation value), Step function (sum after capping)

## LC309 — Best Time to Buy and Sell Stock with Cooldown
**User approach:** dp[i] = max(nums[i] - max of prior prices) or -1 for cooldown if sold on prior day.
**Refined approach:** Three states: hold[i], sold[i], rest[i]. hold = max(hold, rest - price), sold = hold + price, rest = max(rest, prev_sold). Answer = max(sold, rest). Cooldown enforced by rest transitioning from sold[i-1].
**Time:** O(n), O(1) space
**Pattern:** Dynamic Programming
**Key insight:** Cooldown is captured by the state machine — after selling, you must pass through rest before buying again. Three states keep the transitions clean.
**Discrete Math:** DP, State machine (finite automaton with 3 states), Optimization (maximization)

Three states: hold (holding a stock), sold (just sold, must cooldown tomorrow), rest (not holding, free to buy). Cooldown is the transition from sold → rest, not a state itself.

Single array approach: dp[i][state] where state = 0 (hold), 1 (sold), 2 (rest).
- dp[i][0] = max(dp[i-1][0], dp[i-1][2] - prices[i])  // keep holding, or buy from rest
- dp[i][1] = dp[i-1][0] + prices[i]                     // sell today
- dp[i][2] = max(dp[i-1][2], dp[i-1][1])                // keep resting, or finish cooldown
Answer: max(dp[n-1][1], dp[n-1][2]). O(n) time, O(n) space. Three separate variables compress to O(1) since only i-1 is needed.

## LC24 — Swap Nodes in Pairs
**User approach:** Slow/fast pointers with saved references, pointer reassignments get tangled.
**Refined approach:** Dummy head. prev points to node before the pair. Each iteration: first = prev.next, second = prev.next.next. Rewire: first.next = second.next, second.next = first, prev.next = second. Advance prev = first.
**Time:** O(n), O(1) space
**Pattern:** Linked List
**Key insight:** Dummy head eliminates edge cases for head swaps. Each pair swap is the same four pointer reassignments — prev always anchors the operation.
**Discrete Math:** Linked list (singly linked), Pointer manipulation, Pairwise permutation

## LC176 — Second Highest Salary
**User approach:** select * from Employee order by salary desc limit 2 offset 1
**Refined approach:** SELECT DISTINCT salary FROM Employee ORDER BY salary DESC LIMIT 1 OFFSET 1. Or SELECT MAX(salary) FROM Employee WHERE salary < (SELECT MAX(salary) FROM Employee). Second version handles duplicates and returns null naturally.
**Time:** O(n)
**Pattern:** SQL
**Key insight:** LIMIT/OFFSET is close but needs DISTINCT for duplicates. Subquery version is cleaner — find max, then find max below it.
**Discrete Math:** SQL, Set theory (max of complement), Order statistics (second largest with duplicates)

## LC47 — Permutations II
**User approach:** f(nums) = [nums[i]] ++ f({nums[i]} \ nums) for all i where nums[i] not in seen.
**Refined approach:** Sort nums. Backtrack at each position, try each unused element. Skip if nums[i] == nums[i-1] and nums[i-1] was not used at this level — prevents starting same permutation from same position twice.
**Time:** O(n × n!), O(n) space
**Pattern:** Backtracking
**Key insight:** Same as LC46 but with duplicate handling. Sort + "previous not used at this level" check eliminates duplicate permutations without needing a set per recursion level.
**Discrete Math:** Combinatorics (multiset permutations, n! / Π(kᵢ!)), Backtracking, Duplicate pruning via ordering

## LC1177 — Can Make Palindrome from Substring
**User approach:** Count differences between s[left] and s[right] as they move inward until they pass each other.
**Refined approach:** Prefix sum of character counts. For each query, get substring frequencies in O(26). Count odd-frequency characters. Changes needed = odd_count / 2. Return odd_count / 2 <= k.
**Time:** O(26 × n) precompute + O(26 × q) queries
**Pattern:** Prefix Sum / Hash Map
**Key insight:** In a palindrome, at most one character has odd frequency. Each change fixes two odd counts, so odd_count / 2 is minimum changes needed. Prefix sums make substring frequency queries O(26).
**Discrete Math:** Prefix sum, Mod 2 counting (parity), Palindrome (at most one odd frequency)

## LC515 — Find Largest Value in Each Tree Row
**User approach:** BFS, keep track of level to determine number of nodes to compare for max, reset max after each level.
**Refined approach:** Same. BFS level-order, track level size at each step, find max among that level's nodes.
**Time:** O(n), O(n) space
**Pattern:** BFS / Tree
**Key insight:** Queue size at each iteration = number of nodes at that level. Track max per level, reset for next.
**Discrete Math:** Tree, BFS, Level sets (partition of nodes by depth), Maximum over level

## LC318 — Maximum Product of Word Lengths
**User approach:** Freq_map for each word, O(n²) comparisons while tracking max product if maps don't intersect.
**Refined approach:** Bitmask per word (26-bit int, bit i set if word contains i-th letter). O(n²) comparisons: mask[i] & mask[j] == 0 means no common letters. Track max len[i] × len[j].
**Time:** O(n² × L) build masks + O(n²) comparisons, O(n) space
**Pattern:** Bit Manipulation / String
**Key insight:** Bitmask AND is O(1) vs O(26) for freq map intersection. Same complexity, better constant factor.
**Discrete Math:** Set theory (intersection ⊆ via bitmask AND), Bit operations (26-bit mask as characteristic vector)

## LC725 — Split Linked List in Parts
**User approach:** k+1 pointers pointing to first k+1 nodes, advancing from tail with decreasing counts.
**Refined approach:** Count n. base = n/k, extra = n%k. First extra parts get base+1 nodes, rest get base. Walk list, split at each part's size.
**Time:** O(n), O(1) space
**Pattern:** Linked List
**Key insight:** Simple division math — extra nodes distributed one each to the first extra parts. Walk once, split at computed sizes.
**Discrete Math:** Linked list traversal, Division algorithm (quotient + remainder), Distribution (extra = n mod k)

## LC1701 — Average Waiting Time
**User approach:** Sort by arrival time, carry over finish time, store wait durations in separate array and take avg.
**Refined approach:** Sort by arrival. Track finish time. For each customer: start = max(finish, arrival), wait = start - arrival + time, finish = start + time. Average the waits.
**Time:** O(n log n), O(1) space
**Pattern:** Greedy / Sorting
**Key insight:** Customer waits if chef is still busy (finish > arrival). Otherwise starts immediately. Greedy — serve in order, carry finish time forward.
**Discrete Math:** Greedy algorithm, Sorting by arrival time, Simulation (sequential processing)

## LC1823 — Find the Winner of the Circular Game
**User approach:** Array simulation — advance (pos+k)%n, skip eliminated positions, repeat n-1 times.
**Refined approach:** Josephus recurrence: f(n,k) = (f(n-1,k) + k) % n, f(1,k) = 0. Iterate from 2 to n, return result + 1.
**Time:** O(n), O(1) space
**Pattern:** Math / Recursion
**Key insight:** No array needed. Each iteration computes the survivor's position in a circle of size i from the winner in size i-1. Elimination offset is always k.
**Discrete Math:** Josephus problem, Recurrence relation, Modular arithmetic

## LC274 — H-Index
**User approach:** Bucket array, count non-zeros from right; first bucket[i] <= count.
**Refined approach:** Bucket sort — bucket[i] = count of papers with exactly i citations. Citations > n get capped at n (h can never exceed n). Scan from right, accumulate count. The first index where accumulated count >= i is h.
**Time:** O(n), O(n) space
**Pattern:** Bucket Sort
**Key insight:** Unlike sorting (O(n log n)), bucket sort exploits the fact that citations are bounded by n. Capping at n is the trick — h can't exceed the number of papers.
**Discrete Math:** Bucket sort, Counting, Capping (bounded by n), Order statistics

## LC1658 — Minimum Operations to Reduce X to Zero
**User approach:** Fix i, two-pointer L/R with backtracking if sum > x or R < L.
**Refined approach:** Reverse the problem. target = sum(nums) - x. Sliding window to find longest subarray summing to target. Answer = n - max_window_len. Works because all elements are positive — window sum is monotonic, each left maps to at most one right. Left and right each advance n times total = O(n). If target < 0, impossible.
**Time:** O(n), O(1) space
**Pattern:** Sliding Window
**Key insight:** Removing from ends = keeping a contiguous middle. Minimize removals = maximize kept subarray. Positivity guarantees monotonic sum — right advances when sum < target, left advances when sum > target, record when sum == target. Each element added once (right passes it) and removed at most once (left passes it) — 2n total moves, amortized O(n).
**Discrete Math:** Sliding window, Complement (reverse thinking), Monotonic (positive elements → prefix monotonicity)

## LC478 — Generate Random Point in a Circle
**User approach:** Split circle into 4 quadrants, bound each by its square and randomize.
**Refined approach:** Polar coordinates with sqrt correction. Let θ = uniform(0, 2π), r = R·√(uniform(0,1)). Return (cx + r·cosθ, cy + r·sinθ).
**Time:** O(1)
**Pattern:** Inverse CDF / Probability
**Key insight:** Area scales with r², so radius density must be proportional to r (not uniform). √(uniform) transforms uniform U[0,1] via inverse CDF of F(r) = r²/R². Rejection sampling also works: proposal = uniform over bounding square [−R,R]², target = uniform over circle. c = sup target(x)/proposal(x) = (1/πR²)/(1/4R²) = 4/π ≈ 1.273. Accept ratio = 1/c ≈ 78.5% = area of circle / area of square.
**Discrete Math:** Probability (uniform distribution), Inverse CDF transform, Rejection sampling (Monte Carlo method), Ratio of areas (π/4)

## LC3531 — Count Covered Buildings
**User approach:** DFS row-by-row checking column neighbors.
**Refined approach:** Group by row (x) and column (y), sort each group. For each building, binary search its row group to check existence on both left and right, and its column group to check both above and below.
**Time:** O(m log m)
**Pattern:** Grouping + Binary Search
**Key insight:** "Four directions" decomposes into independent row and column existence checks. Binary search in sorted groups determines whether neighbors exist on both sides.
```
rows = { 1→[2], 2→[1,2,3], 3→[2] }
cols = { 1→[2], 2→[1,2,3], 3→[2] }
(2,2): rows[2] idx=1 → left(1) right(3) ✓, cols[2] idx=1 → above(1) below(3) ✓ → covered
(1,2): rows[1] len=1 → no horizontal neighbors ✗ → not covered
```
**Discrete Math:** Set theory (existence in groups), Binary search, Relation (directional adjacency), Decomposition (row × col independence)

## LC3796 — Find Maximum Value in a Constrained Sequence ⚑ REVISIT
**User approach:** Greedy — pair smallest output with largest diff.
**Refined approach:** Two-pass constraint propagation. Initialize upper = INF, apply restrictions. Forward: upper[i] = min(upper[i], upper[i-1] + diff[i-1]). Backward: upper[i] = min(upper[i], upper[i+1] + diff[i]). Answer = max(upper).
**Time:** O(n)
**Pattern:** Constraint Propagation / Shortest Path on Line
**Key insight:** |a[i]-a[i+1]| ≤ diff[i] gives two directed inequalities. Chain is a DAG — forward pass propagates left-to-right increases, backward pass propagates right-to-left decreases from restrictions. The tightest bound is the intersection of both passes.
**Discrete Math:** Constraint propagation, Inequalities (|a[i]-a[i+1]| ≤ diff[i] ⇔ a[i+1] ≤ a[i]+diff[i] ∧ a[i] ≤ a[i+1]+diff[i]), Shortest path on DAG (chain graph), Bellman-Ford on line (two passes)

## LC3546 — Equal Sum Grid Partition I
**User approach:** Row sum array + col sum array, prefix sums in both directions for both.
**Refined approach:** Compute total sum. If odd → false. Prefix sum rows until target = total/2, same for columns. Single direction suffices (remaining = total - prefix).
**Time:** O(m·n)
**Pattern:** Prefix Sum
**Key insight:** Single prefix check works for each direction — no need for both forward and backward. If total is odd, impossible immediately.
**Discrete Math:** Prefix sum, Partition, Parity (odd sum → impossible, integer partition constraint)

## LC16 — 3Sum Closest
**User approach:** Same as 3Sum but stop checking as sum moves away from target.
**Refined approach:** Sort, fix i, two-pointer L/R. Track closest by abs difference. If sum < target, advance L; if sum > target, decrement R; exact match → return. Skip duplicates, early break if nums[i] * 3 > target.
**Time:** O(n²)
**Pattern:** Sorting + Two Pointers
**Key insight:** Same skeleton as 3Sum but replace "sum == 0" with "track min absolute difference." Two-pointer direction always moves sum toward target.
**Discrete Math:** Sorting/Total order, Invariant (two-pointer convergence), L₁ distance (minimization to target)

## LC3377 — Digit Operations to Make Two Integers Equal
**User approach:** Prime sieve, modify unit digit first, inc/dec adjacent to avoid primes.
**Refined approach:** Dijkstra on state graph. Sieve primes up to 10⁴. Each state = current number, edge weight = target value. Generate neighbors by ±1 per digit, skip primes. dist[v] = dist[u] + v.
**Time:** O(V log V + E), V ≤ 10⁴
**Pattern:** Graph / Dijkstra / Number Theory
**Key insight:** Edge weights depend on target node (not source), so Dijkstra still works — cost accumulates additively. Prime constraint makes greedy fail; graph is small enough for full search.
**Discrete Math:** Graph theory (shortest path on weighted graph), Number theory (prime sieve, digit manipulation)

## LC848 — Shifting Letters
**User approach:** Preprocess shifts array with suffix sum, apply shifts in O(n).
**Refined approach:** Same. Compute suffix sum (mod 26) from right. Each s[i] gets shifted by total[i] = shifts[i] + total[i+1].
**Time:** O(n)
**Pattern:** Suffix Sum
**Key insight:** Naively applying each shift to i+1 letters is O(n²). Each letter's total shift = sum of shifts[i..n-1] — one pass from right.
**Discrete Math:** Mod 26 (modular arithmetic), Prefix/suffix sums (associativity of addition)

## LC319 — Bulb Switcher
**User approach:** Multiples cause extra toggles; n=4 leaves bulb 2 off, toggles 4 back on.
**Refined approach:** Toggle count = number of divisors of bulb i. Odd divisor count ↔ perfect square. Answer = floor(√n).
**Time:** O(1)
**Pattern:** Math / Number Theory
**Key insight:** Divisors come in pairs unless i is a perfect square. Only perfect-square bulbs are toggled an odd number of times → left on.
**Discrete Math:** Number theory (divisor parity, perfect squares)

## LC2592 — Maximize Greatness of an Array
**User approach:** Sort, for each i find least greater untaken element, count matches.
**Refined approach:** Sort nums. Two pointers i, j. For each i, advance j until nums[j] > nums[i]. Count match, advance both. Stop when j exhausts.
**Time:** O(n log n)
**Pattern:** Greedy / Two Pointers
**Key insight:** Pair smallest beater with smallest target — never waste a large element. Greedy on sorted arrays is optimal because ordering preserves the "greater than" relation transitively.
**Discrete Math:** Sorting/Total order, Greedy algorithm, Injective matching (each beater used at most once)

## Pattern Coverage Gap Analysis (Jul 28, 2026)

### COMPLETELY MISSING (0 problems)
- Queue / Deque / Priority Queue
- Union-Find / Disjoint Set Union
- Segment Tree / Fenwick Tree / BIT
- Reservoir Sampling
- Quickselect
- Rolling Hash / Rabin-Karp
- KMP / String Matching

### UNDERREPRESENTED (1-2 problems)
- Trie (1)
- Heap / Top K (1)
- Simulation (1)
- Interval / Sweep Line (2)
- Bit Manipulation (2)
- Recursion / Divide and Conquer (2)
- Strings / String Manipulation (0 — covered tangentially only)

## LC1438 — Longest Continuous Subarray With Absolute Diff ≤ Limit
**User approach:** Sliding window tracking min and its index; expand right, if diff > limit shrink left past the breakpoint.
**Refined approach:** Sliding window with two monotonic deques (max-deque decreasing, min-deque increasing). Expand r, maintain deques. While max front - min front > limit, shrink l and pop expired elements. Track max window size.
**Time:** O(n)
**Pattern:** Queue / Sliding Window / Monotonic Queue
**Key insight:** Two deques maintain running min and max in O(1) amortized per operation. Condition is any two elements, so both min and max must be tracked.
**Discrete Math:** Order statistics, Min/max in sliding window, Monotonic sequences

## LC621 — Task Scheduler
**User approach:** Partition tasks, append to queue, track window length; if window < n+1, pull from next partition or idle.
**Refined approach (simulation):** Max heap by frequency + cooldown queue (count, readyTime). At each time, pop cooldown queue if ready. Pop max from heap, decrement, push to cooldown if remaining. If heap empty but queue not, idle. O(n log m) time, O(m) space.
**Formula approach:** maxFreq, idleSlots = (maxFreq-1)*(n+1), result = max(len(tasks), idleSlots + count of tasks with max freq). O(n) time, O(1) space.
**Time:** O(n log m) / O(n)
**Pattern:** Queue / Heap / Greedy
**Key insight:** Always schedule the most frequent available task. Formula skips simulation entirely — idle slots are the fixed overhead from the most frequent task. Only tasks with max frequency extend the total beyond that.
**Discrete Math:** Scheduling theory (cooldown constraints), Frequency distribution, Pigeonhole principle (idle slots)

## LC684 — Redundant Connection
**User approach:** Union-Find with int array parent tracking. Iterate edges, if find(u) == find(v) return edge (creates cycle), else union(u,v).
**Refined approach:** Same — standard Union-Find. Path compression + union by rank for near-O(1) amortized.
**Time:** O(n α(n))
**Pattern:** Union-Find
**Key insight:** Cycle detection in undirected graph reduces to checking if endpoints are already connected before adding edge. First edge where find(u)==find(v) is the redundant one.
**Discrete Math:** Equivalence relations, Connected components in undirected graphs, Forest/tree characterization (n nodes, n-1 edges, no cycles)

## LC11 — Container With Most Water
**User approach:** Two pointers, advance the min of height[L], height[R] inward, track global max of min(h[L],h[R])*(R-L).
**Refined approach:** Same — two pointers at ends. Compute area = min(h[L],h[R])*(R-L), track max. Move the pointer with smaller height inward (moving taller one can't increase area — height capped by shorter, width only decreases).
**Time:** O(n)
**Pattern:** Two Pointers / Greedy
**Key insight:** The limiting factor is the shorter line. Moving the taller one inward cannot yield a larger area since the height is bounded by the shorter line and width always shrinks.
**Discrete Math:** Optimization, Monotonicity

## LC167 — Two Sum II (Input Array Is Sorted)
**User approach:** Two pointers, move R inward if sum > target, move L inward if sum < target.
**Refined approach:** Same — standard sorted two-sum.
**Time:** O(n)
**Pattern:** Two Pointers / Binary Search
**Key insight:** Sorted input enables linear two-pointer instead of hash map. The monotonicity of sum with respect to pointer movement guarantees finding the pair.
**Discrete Math:** Monotonic sequences, Order theory

## LC875 — Koko Eating Bananas
**User approach:** Binary search between min(piles) and avg(piles).
**Refined approach:** Binary search k between 1 and max(piles). canEat(k) = sum(ceil(pile/k)) ≤ h. If can eat, try lower; else try higher.
**Time:** O(n log max(pile))
**Pattern:** Binary Search
**Key insight:** Feasibility function is monotonic — if k works, any k' > k also works. Binary search on the threshold.
**Discrete Math:** Monotone predicate, Binary search on continuous space

## LC200 — Number of Islands
**User approach:** DFS in 4 directions.
**Refined approach:** Iterate grid. On hitting '1', increment count, DFS to sink all adjacent land (mark '0'). Count = number of DFS calls.
**Time:** O(m*n)
**Pattern:** Graph DFS / Matrix DFS
**Key insight:** Sinking visited land avoids needing a visited set — in-place modification. Each cell visited at most twice.
**Discrete Math:** Connected components in undirected graph, 4-connectivity vs 8-connectivity

## LC2029 — Stone Game IX
**User approach:** If len%2==0 and sum%3==0 Alice wins; if len%2==1 and sum%3==0 Bob wins.
**Refined approach:** Only remainders mod 3 matter (c0, c1, c2). Running sum mod 3 cycles: Alice loses starting with 0-stone (sum=0 → loses). If both c1,c2==0 → Alice loses. If abs(c1-c2)>2 → Alice wins. If c0==0 → Alice wins if both c1,c2>0. Otherwise depends on c0 parity.
**Time:** O(n)
**Pattern:** Math / Game Theory / Greedy
**Key insight:** Only mod-3 classes matter, not actual values. Game reduces to counting remainders and reasoning about forced sequences.
**Discrete Math:** Modular arithmetic, Combinatorial game theory (normal play)

## LC3255 — Find the Power of K-Size Subarrays II
**User approach:** Running stack, empty on non-consecutive or decreasing element.
**Refined approach:** Track consecutive-ascending run length ending at each index. If nums[i]==nums[i-1]+1 → len++, else len=1. If len≥k, result[max(0,i-k+1)] = nums[i] (max of valid window = last element), else -1.
**Time:** O(n)
**Pattern:** Sliding Window / Array
**Key insight:** "Consecutive ascending" means +1 exactly, not monotonic. Valid window's max is its last element. Single pass with running counter — no stack needed.
**Discrete Math:** Integer sequences, Consecutive integers

## LC1404 — Number of Steps to Reduce a Number in Binary Representation to One
**User approach:** DP on integer value: dp[1]=0, dp[even]=log(even), dp[odd]=1+dp[i-1].
**Refined approach:** Process LSB to MSB with carry. For each bit (skip leading 1): if odd → 2 steps (subtract+divide), carry=1; if even → 1 step (divide), carry unchanged. O(n) time, O(1) space. String too long (500 bits) for int conversion.
**Time:** O(n)
**Pattern:** String / Bit Manipulation / Simulation
**Key insight:** No conversion to int — binary string is already the number. Operations map directly to string manipulation: divide = drop LSB, subtract = flip LSB 1→0 (with carry propagation). Carry tracks pending subtraction across bits.
**Discrete Math:** Binary representation, Division algorithm

## LC1915 — Number of Wonderful Substrings
**User approach:** Combinatorics with repeated chars/combinations.
**Refined approach:** 10-bit parity mask per prefix (bit = odd/even count for letters a-j). For each prefix i, wonderful substrings ending at i = count[mask[i]] (all even) + sum_{b=0..9} count[mask[i] ^ (1<<b)] (exactly one odd). count[0]=1 for empty prefix.
**Time:** O(n * 10) = O(n)
**Pattern:** Bit Manipulation / Prefix XOR / Hash Map
**Key insight:** Parity is the only thing that matters for "at most one odd" condition. XOR tracks parity toggles. Wonderful = mask has 0 or 1 set bits. 10 letters max enables O(10) per position.
**Discrete Math:** Parity, XOR as toggle, Combinatorics (count of substrings with given parity)
