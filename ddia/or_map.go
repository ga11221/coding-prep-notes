package main

// OR-Map (Observed-Remove Map with nested CRDT values)
// Each key maps to a (CRDT value, causal context: set of tags)
// put(key, crdt): replace the value and add a new tag to the key's causal context
// remove(key): add all current tags for key to tombstone set
// merge: union of live key-context pairs; for overlapping keys, recursively merge values

type CRDT interface {
	Merge(other CRDT)
}

type ORMap struct {
	values    map[string]CRDT       // key -> nested CRDT value
	context   map[string][]string   // key -> causal context (tags)
}

func NewORMap() *ORMap {
	return nil
}

func (m *ORMap) Put(key string, value CRDT) {
}

func (m *ORMap) Remove(key string) {
}

func (m *ORMap) Get(key string) CRDT {
	return nil
}

func (m *ORMap) Merge(other *ORMap) {
}
