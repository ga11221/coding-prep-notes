package main

// step 0:
// So, I'm given an array of positive/negative integers and a number k that can be zero or greater
// I'm to find two elements in the array that are equal in value and are apart by k or less indices
// if any two such elements exist, return true else false
// inferring from the fact that k can be 0, i can be equal to j - no, i != j
// integer array must be at least of length 1

// edge cases: k = 0 return true; for any k if len(nums) = 1 return true;

// nums = [1,2,3,1], k = 3
// nums = [3,4] k=1
// duplicate 1's at i=0 and j=3, abs(i-j)=3 <= k return true
func containsNearbyDuplicate(nums []int, k int) bool {
	// 1. find duplicates
	// 2. for each pair of duplicates, compare their distance to k

	// 1. populate freq map - elem value -> [indices in nums]
	// 2. as duplicates are pushed to the map, compute difference with last occurence

	// linear in time and space

	if k == 0 || len(nums) == 1 {
		return false
	}
	frequencies := map[int][]int{}
	// nums= [1,0,1,1], k = 1
	for i, val := range nums {
		if indices, ok := frequencies[val]; ok { //i=2,indices=[0]
			if i-indices[len(indices)-1] <= k { // 2-indices[1-1]=2
				return true
			}
			frequencies[val] = append(frequencies[val], i)
		} else {
			frequencies[val] = []int{i} // i = 0, val = 1 {1->[0]}; // i=1, val=0 {1->[0], 0->[1]}; // i =2,val=1 {1->[0], 0->[1]}; // after first three iters: {1->[0];2->[1];3->[2}}
		}
	}
	return false
}

func main() {}
