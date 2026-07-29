package main

// OR-Set (Observed-Remove Set)
// add(element): add with unique tag (replicaID + counter)
// remove(element): add all current tags for element to tombstone set
// merge: union of both live adds and tombstoned tags
// lookup: element is present if |adds| > |tombstoned tags for that element|

type ORSet struct {
	adds map[string]map[string]bool      // element -> set of tags
}

func NewORSet() *ORSet {
	return nil
}

func (s *ORSet) Add(element string, tag string) {
}

func (s *ORSet) Remove(element string) {
}

func (s *ORSet) Contains(element string) bool {
	return false
}

func (s *ORSet) Merge(other *ORSet) {
}
