package main

import "sort"

type TrieNode struct {
	children    [26]*TrieNode
	suggestions []string
}

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)

	root := &TrieNode{}
	for _, p := range products {
		node := root
		for _, ch := range p {
			idx := ch - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &TrieNode{}
			}
			node = node.children[idx]
			if len(node.suggestions) < 3 {
				node.suggestions = append(node.suggestions, p)
			}
		}
	}

	n := len(searchWord)
	dp := make([][]string, n+1)
	dp[0] = []string{}

	node := root
	for i := 0; i < n; i++ {
		idx := searchWord[i] - 'a'
		if node != nil {
			node = node.children[idx]
		}
		if node != nil {
			dp[i+1] = node.suggestions
		} else {
			dp[i+1] = []string{}
		}
	}

	return dp[1:]
}

// combinatorial DP approach (without trie):
//   sort(products)
//   dp[0] = products (all match empty prefix)
//   for i := 1; i <= len(searchWord); i++ {
//       dp[i] = filter(dp[i-1], func(p string) bool {
//           return i <= len(p) && p[i-1] == searchWord[i-1]
//       })
//   }
//   result[i-1] = dp[i][:3] if len(dp[i]) >= 3 else dp[i]
//
// This is O(n * m) where n = len(searchWord), m = len(products)
// and does not use a trie. The hybrid (trie for prefix aggregation +
// linear DP for traversal) is O(n + total chars in products) and
// avoids re-filtering all products at each step.
