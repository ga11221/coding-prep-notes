package main

import "fmt"

/**
type Partition struct {
	start, end     int
	containedChars map[string]uint8
}
*/

func main() {
	s := "ababcbacadefegdehijhklij"
	fmt.Printf("Valid window sizes for %s: %v\n", s, partitionLabels(s))
}

func partitionLabels(s string) []int {

	partitionChars := map[string]*[2]int{}
	partitionQ := []string{}
	for i, c := range s {
		str := string(c)
		if _, ok := partitionChars[str]; ok {
			partitionChars[str][1] = i
		} else {
			newWindow := [2]int{i, i}
			partitionChars[str] = &newWindow
			partitionQ = append(partitionQ, str)
		}
	}
	// for each partition, determine if any of its contained chars
	// occur outside the partition
	// every window must be compared with all other windows for overlap
	// ie find all other partitions that overlap this one and merge them
	// merge(w1, w2) = (w1i,w1j) + (w2i, w2j) = (w1i,w2j) where w1i<w2i and w2j>w1j
	validWindows := [][2]int{}
	for _, startingWindowChar := range partitionQ {
		window := partitionChars[startingWindowChar]
		// only check if window size > 1
		if window[1] > window[0] {
			// @todo besides merging windows, need to remove contained windows
			for _, nextWindow := range partitionChars {
				/**
				[     ]
				    [           ]
				*/
				if window[0] < nextWindow[0] && nextWindow[0] < window[1] && window[1] < nextWindow[1] {
					window[1] = nextWindow[1]
				}
			}
			validWindows = append(validWindows, *window)
		} else {
			validWindows = append(validWindows, *window)
		}
	}
	validWindowSizes := []int{}
	for _, validWindow := range validWindows {
		validWindowSizes = append(validWindowSizes, validWindow[1]-validWindow[0]+1)
	}
	return validWindowSizes

}
