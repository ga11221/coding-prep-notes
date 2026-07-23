# Recursion Rapid Fire — Answer Log

## Q1: Climbing Stairs (n steps, 1 or 2 at a time)

**User answer:**
Permutation done linearly — f(i, n) = count({1+f(i-1, n-1)}) + count({2+f(i-1, n-2)}) and f(_,0) = 1 and f(_, 1)=1

**Feedback:**
- i parameter is unnecessary — pure function of n only
- Recurrence right in spirit: f(n) = f(n-1) + f(n-2)
- Base cases correct: f(0)=1, f(1)=1
- Pattern: linear recursion (branching factor 2), not permutation

---

## Q2: Subsets (return all subsets of array)

**User answer:**
Backtracking — f(array) = {array[i]} U { c | c in f(array[i+1:]) } and f([]) = {{}} with 2^length(array) leaves

**Feedback:**
- Pattern: tree recursion / combinatorial recursion, not backtracking (no constraint undoing)
- Recurrence missing skip case — formulation only includes take, never skips array[i]
- Correct recurrence: f(i) = f(i+1) ∪ { [nums[i]] + c for c in f(i+1) }
- { [nums[i]] + c for c in f(i+1) } is implicit map over recursive result
- This is return-value form (composition up the stack)
- Leaves: 2^n — correct

**Follow-up questions from user:**
- Recognized { [nums[i]] + c for c in f(i+1) } is an implicit map — confirmed yes
- Asked if skip/take recurrence is return-value form — confirmed yes

---

## Q3: Valid Parentheses (n pairs, return all valid combinations)

**User answer:**
Permutation — f(a, b) = { '(' + f(a-1, b) if a > 0 } U { ')' + f(a, b-1) if b > a } and f(0, 0) = {} where both a,b start at n

**Feedback:**
- Pattern: permutation — correct
- Recurrence: correct (a=remaining open, b=remaining close)
- Constraint b > a: correct (means unmatched open paren exists to close)
- Base case nit: f(0,0) = {""} (set with empty string), not {} (empty set)
- Clean answer overall

---

## Q4: Merge Sort Time Complexity

**Question:** What's the recurrence for time complexity? What theorem solves it?

**User answer:** (answered via reference — no verbal answer)

**Feedback:** T(n) = 2T(n/2) + O(n), T(1) = O(1). Master Theorem: a=2, b=2, d=1 → O(n log n).

---

## Q5: Binary Search Time Complexity

**Question:** What's the recurrence? How does it differ from merge sort?

**User answer:** (answered via reference — no verbal answer)

**Feedback:** T(n) = T(n/2) + O(1), T(1) = O(1). Master Theorem: a=1, b=2, d=0 → O(log n). Key diff: merge sort = 2 subproblems + O(n) work; binary search = 1 subproblem + O(1) work.

---

## Q6: 3-way Split Recurrence Derivation

**Question:** Function splits into 3 subproblems, each 1/3 the size, with O(n) work to split and combine. Write the recurrence and solve it.

**User answer:** (answered verbally — requested answer)

**Feedback:** T(n) = 3T(n/3) + O(n), T(1) = O(1). Master Theorem: a=3, b=3, d=1 → log_3(3) = 1 = d → O(n log n). Same complexity as merge sort — branching factor and size reduction cancel out.

---

## Q7: Karatsuba-style Recurrence

**Question:** Karatsuba multiplication splits into 3 subproblems of n/2 size with O(n) combine work. Write the recurrence and solve it.

**User answer:** (answered verbally — requested answer)

**Feedback:** T(n) = 3T(n/2) + O(n). Master Theorem: a=3, b=2, d=1 → log_2(3) ≈ 1.585 > 1 → O(n^1.585). Faster than naive O(n^2) multiplication. Key: subproblems are n/2 (not n/3), so the tree is shallower but wider.

---

# LeetCode-Style Recurrence Pattern Recognition

## Q8: Binary Search (recursive, return index)

**Question:** Given a sorted array and a target, write the recurrence for binary search that returns the index (or -1). What's the state? What's the base case?

**User answer:** f(array, a, b, target) = f(array, b-a/2 + 1, target) if target > b-a/2, f(array, a, b-a/2, target) if target < b-a/2, else array[b-a/2]. Base: -1 if a==b. State is two pointers a, b starting at 0 and len(array)-1.

**Feedback:**
- State correct: two pointers a, b
- Bug: `b-a/2` is ambiguous (operator precedence). Should be `mid = a + (b-a)/2`
- Bug: branches compare against `mid` value (`array[mid]`), not `mid` index
- Bug: base case `a == b` should be `a > b` (not found). `a == b` still has one element to check
- Bug: return `array[b-a/2]` should return `mid` (index), not the value
- Correct form: f(a, b) → if a > b: return -1; mid = a + (b-a)/2; if target < array[mid]: f(a, mid-1); if target > array[mid]: f(mid+1, b); else return mid

---

## Q9: Combination Sum (elements reusable)

**Question:** Given an array of distinct integers and a target, return all combinations that sum to target (elements reusable). Write the recurrence, identify the pattern, and state the branching factor.

**User answer:** Combinatorial recursion — order doesn't matter. f(nums, i, target) = { [nums[i]] + c | c ∈ f(nums, i, target - nums[i]) } ∪ f(nums, i+1, target). Base cases: target == 0 → [[]], target < 0 → []. Append to combos when target == 0.

**Feedback:** Correct recurrence. Two branches at each level: (1) **take** nums[i], stay at i (reusable), recurse with target - nums[i]. (2) **skip** nums[i], advance to i+1, same target. This differs from Q9-like problems where elements are used once — here the "take" branch recurses on same index, not i+1. Pattern: combinatorial / tree recursion with constraint (target ≥ 0). Branching factor: 2 per node (take or skip), but take can repeat — unbounded depth on the take branch until target ≤ 0. Leaves: depends on values, not purely 2^n. Key insight: staying at same index on take is what makes elements reusable — skipping advances, taking doesn't.

---

## Q10: Permutations (swap-based)

**Question:** Given an array, generate all permutations using the swap-based approach. Write the recurrence, trace the recursion tree, state branching factor and leaves.

**User answer:** P(array, i) = { P(swap(array, i, j), i+1) | j ∈ [i+1..n-1] }. P(array, n) = copy(array). Swap-back is implementation detail, not part of the recurrence — recurrence describes what the function computes, not how it mutates state.

**Feedback:** Correct. Key points: (1) Loop starts at i+1, not i — j=i is noop. (2) Branching factor: n-i at level i (decreases as i increases). (3) Total leaves: n! — product of (n-1)(n-2)...1. (4) Copy only at base case — O(n) per permutation, O(n) space. (5) Swap-back is mechanism, not math — each branch sees correct state for its j. Recurrence captures the WHAT (all permutations), swap-back is the HOW (implementation detail). Tree for [1,2,3]: 3 branches at root (j=1,2), 2 at next level, 1 at deepest → 3! = 6 leaves.

