# Discrete Math Rapid Fire (Jul 27, 2026)

## Q1: Poset vs Total Order
**Q:** Difference between a partially ordered set and a totally ordered set? Example of each.
**User:** Poset doesn't have min/max, total order does. Poset = ℝ, Total order = ℕ.
**Correction:** Both ℝ and ℕ under ≤ are total orders (all pairs comparable). A proper poset example is (P({a,b}), ⊆) where {a} and {b} are incomparable. The distinguishing property is total comparability, not existence of min/max. ℤ is totally ordered but has no min/max.
**Key insight:** Comparability of all elements is the defining property of a total order. Min/max existence is a separate property (well-ordering requires it).

## Q2: Duality Principle
**Q:** State the duality principle for posets.
**User:** The dual of any theorem in a Lattice is also a theorem.
**Correction:** Holds for all posets, not just lattices. Reversing every ≤ in a true statement about *all* posets yields another true statement. Lattices have additional structure (∨, ∧) but the duality principle is already valid at the poset level.

## Q3: Extra Structure of Lattices vs Posets
**Q:** What's the extra structure of lattices?
**User:** Asked the question.
**Refined answer:** A lattice is a poset where every pair has both a join (∨, sup) and meet (∧, inf). Posets may have missing sup/inf for some pairs — those aren't lattices. Example of a non-lattice: a 4-element "V" shape where {a,b} has no sup.

## Q4: Join ↔ Least Upper Bound
**Q:** Why is join interchangeable with least upper bound?
**Answer:** They're the same concept described two ways — order-theoretic (lub) vs algebraic (binary op ∨). The correspondence theorem: poset with all pair sups defines ∨; conversely, define ≤ by a∨b=b and ∨ becomes sup.

## Q5: Lattice ⇒ Well-ordering?
**Q:** If a lattice has both sup and inf, is it also a well-ordering?
**Answer:** No. Lattice = property of pairs; well-ordering = property of all subsets. [0,1] is a complete lattice but not well-ordered ((0,1] has no min).

## Q6: Unsorted integer array — best classification?
**Q:** An unsorted integer array — what's the best order-theoretic classification?
**Answer:** Totally ordered set. The values under ≤ are all comparable. "Unsorted" describes the arrangement, not the underlying order type.

