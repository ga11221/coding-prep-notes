package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	results := make(chan map[int32][2]int32, k)
	for id, minMaxPair := range splitInto64() {
		wg.Add(1)
		go func(id int, minMaxPair [2]int32) {
			defer wg.Done()
			pairHashes := map[int32][2]int32{}
			for i := minMaxPair[0]; i <= minMaxPair[1]; i++ {
				for j := minMaxPair[0]; j <= minMaxPair[1]; j++ {
					pairHashes[hashPair(i, j)] = [2]int32{int32(i), int32(j)}
				}
			}
			results <- pairHashes
		}(id, minMaxPair)
	}
	i := 0
	for hashes := range results {
		if i == 0 {
			for k, v := range hashes {
				fmt.Printf("hash: %d for values: %v\n", k, v)
			}
		}
		fmt.Printf("%v: computed %v hashes\n", i, len(hashes))
		i++
	}
}

// Cantor pairing function: bijection from N² → N (collision-free for non-negatives).
// Order: diagonal index = a+b, triangular number = index*(index+1)/2, offset = b.
func hashPair(a, b int32) int32 {
	if a > b {
		a, b = b, a
	}
	return (a+b)*(a+b+1)/2 + b
}

const k = 65536

func splitInto64() [][2]int32 {
	total := int64(int(int(math.MaxInt32-2)/10e3)) + 1
	size := total / k
	fmt.Printf("splitting into %d numbers into %d parts, each of length %d\n", total, k, size)

	parts := make([][2]int32, k)
	for i := range k {
		start := int32(int64(i) * size)
		end := start + int32(size) - 1
		parts[i] = [2]int32{start, end}
	}
	return parts
}
