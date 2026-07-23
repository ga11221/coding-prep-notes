package main

import (
	"fmt"
	"iter"
)

/**
TreeSet - hash map with ordered list of keys
@TODO sortedSet ops need to be implemented
*/

type TreeSet[K comparable, V any] struct {
	m          map[K]V
	sortedKeys []K
}

func (t *TreeSet[K, V]) IterEntrySet() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, key := range t.sortedKeys {
			if !yield(key, t.m[key]) {
				return
			}
		}
	}
}

func (t *TreeSet[K, V]) IterKeySet() iter.Seq[K] {
	return func(yield func(K) bool) {
		for _, key := range t.sortedKeys {
			if !yield(key) {
				return
			}
		}
	}
}

func (t *TreeSet[K, V]) IterValueSet() iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, key := range t.sortedKeys {
			if !yield(t.m[key]) {
				return
			}
		}
	}
}

func (t *TreeSet[K, V]) Put(key K, value V) (V, bool) {
	oldVal, ok := t.m[key]
	// @todo: insert in sorted order
	if !ok {
		t.sortedKeys = append(t.sortedKeys, key)
	}
	t.m[key] = value
	return oldVal, ok
}

func (t *TreeSet[K, V]) Get(key K) (V, bool) {
	val, ok := t.m[key]
	return val, ok
}

func (t *TreeSet[K, V]) Delete(key K) bool {
	_, had := t.m[key]
	if had {
		delete(t.m, key)
		// @todo remove from sortedKeys
	}
	return had
}

func (t *TreeSet[K, V]) DeleteByCondition(cond func(k K, v V) bool) []K {
	var deleted []K
	for k, v := range t.m {
		if cond(k, v) {
			delete(t.m, k)
			// @todo remove from sortedKeys
			deleted = append(deleted, k)
		}
	}
	return deleted
}

func main() {
	fmt.Println("vim-go")
}
