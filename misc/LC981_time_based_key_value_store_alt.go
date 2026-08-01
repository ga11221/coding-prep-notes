package main

// @TODO
type TimeMapEntry struct {
	value string
	next  int
}

type TimeMap struct {
	bucketMap map[string][]TimeMapEntry
}

func Constructor() TimeMap {
	return TimeMap{
		lastIndex: map[string]int{},
		bucketMap: map[string][]TimeMapEntry{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	entry := TimeMapEntry{
		value: value,
	}
	if timeEntries, ok := this.bucketMap[key]; ok {
		next := this.lastIndex[key]
		entry.next = next
		timeEntries[key][timestamp] = entry
	} else {
		entries := make(TimeMapEntry, 10e7)
		entry.next = -1
		entries[timestamp] = entry
		this.bucketMap[key] = entries
	}
	this.lastIndex[key] = timestamp
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if lastIndex, ok := this.lastIndex[key]; !ok || lastIndex == -1 {
		return ""
	}
	timeEntries := this.bucketMap[key]
	// what if timestamp is not in there???
	// find max(time) < timestamp in timeEntries in O(1)???
	// bin search for non-empty entry from [0, timestamp] O(log(n))
	// store intervals for non-empty???
	thisEntry := timeEntries[timestamp]
	if thisEntry == nil {
		//bin search for next
	} else {
		nextIdx := timeEntries[timestamp].next
		return timeEntries[nextIdx].value
	}
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);


  */
