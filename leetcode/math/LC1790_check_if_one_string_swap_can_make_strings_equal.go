package main

// Given two strings, if I can make one character swap or less in either s1 or s2,
// so that s1 = s2, return true
// if it was possible to make one or no swaps in one string, it doesn't matter
// which string, s1 or s2, is manipulated

// either try O(n^2) swaps or find the two letters in s1 that differ from s2
// and exchange places

// eg banki kanba false
// eg bank kanb true
// eg bank kank false
// eg chill (differ by 3)
// eg ihcll

// strings of unequal length, equal strings, empty strings, strings that differ by 1 char or 3 or more
func areAlmostEqual(s1 string, s2 string) bool {
	// iterate over both strings in parallel
	// 1. for every index where s1 and s2 differ in chars, store
	// 2. if indices stored do not count to two exactly or zero, then false
	// 3. else swap chars and check those indices again
	// bank kanb
	if len(s1) == 0 && len(s2) == 0 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	var c1, c2 string
	diffs := []int{}
	for i := range len(s1) { // i <- 0; 1; 2; 3
		c1 = string(s1[i]) // c1 <- b; a; n; k
		c2 = string(s2[i]) // c2 <- k; a; n; b
		if c1 != c2 {
			diffs = append(diffs, i) // [0] // [0,3]
			if len(diffs) > 2 {
				return false
			}
		}
	}
	if len(diffs) == 0 {
		return true
	}
	if len(diffs) == 1 {
		return false
	}
	// c1[0] == c2[3] && c1[3] == c2[0] => b == b && k == k
	return s1[diffs[0]] == s2[diffs[1]] &&
		s1[diffs[1]] == s2[diffs[0]]
}

func main() {}
