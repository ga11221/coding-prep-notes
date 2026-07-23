package main

import (
	"slices"
)

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, v := range nums {
		if val, ok := m[v]; ok {
			m[v] = val + 1
		} else {
			m[v] = 1
		}
	}

	length := 100
	bucket := [100][]int{}
	for k, v := range m {
		for v > length {
			length *= 2
		}
		s := bucket[v]
		bucket[v] = append(s, k)
	}
	ret := []int{}
	for i := len(bucket) - 1; i >= 0; i-- {
		if len(ret) == k {
			break
		}
		if bucket[i] != nil {
			ret = append(bucket[i], ret...)
		}
	}
	return ret[:k]
}

// ordered hash map impl

type OrderedMap struct {
	m           map[int]int
	OrderedKeys []int
	comparator  func(int, int) bool
	length      int
}

func NewOrderedMap(comparator func(int, int) bool, length int) *OrderedMap {
	return &OrderedMap{
		m:           make(map[int]int, length),
		OrderedKeys: []int{},
		comparator:  comparator,
		length:      length,
	}
}

func (om *OrderedMap) put(k int, v int) {
	if count, ok := om.m[k]; ok {
		om.m[k] = count + 1
		v = count + 1
	} else {
		om.m[k] = v
	}
	if len(om.OrderedKeys) == 0 {
		om.OrderedKeys = append(om.OrderedKeys, k)
		return
	}

	j := 0
	found := false
	for j = range om.OrderedKeys {
		if om.OrderedKeys[j] == k {
			found = true
			break
		}
	}
	if found {
		om.OrderedKeys = append(om.OrderedKeys[:j], om.OrderedKeys[j+1:]...)
	}
	i, k1 := 0, 0
	found = false
	for i, k1 = range om.OrderedKeys {
		if om.comparator(v, om.m[k1]) {
			found = true
			break
		}
	}
	if found {
		om.OrderedKeys = slices.Insert(om.OrderedKeys, i, k)
	} else {
		om.OrderedKeys = append(om.OrderedKeys, k)
	}
	l := om.length
	if len(om.OrderedKeys) < l {
		l = len(om.OrderedKeys)
	}
	om.OrderedKeys = om.OrderedKeys[0:l]
}
