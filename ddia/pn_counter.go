package main

// PN Counter: stateful CRDT with separate P (increment) and N (decrement) vectors
// merge = max(P1, P2) + max(N1, N2), value = sum(P) - sum(N)

type PNCounter struct {
	P []int // increments per replica
	N []int // decrements per replica
}

func NewPNCounter(replicaCount int) *PNCounter {
	return nil
}

func (c *PNCounter) Increment(replicaIdx int, delta int) {
}

func (c *PNCounter) Decrement(replicaIdx int, delta int) {
}

func (c *PNCounter) Value() int {
	return 0
}

func (c *PNCounter) Merge(other *PNCounter) {
}
