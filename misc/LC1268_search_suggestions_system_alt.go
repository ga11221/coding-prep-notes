package main

import "slices"

/*
*
@TODO rework this - rushed through it
Brute force is O(n^2)
So I'm given a list of distinct strings, all of which are at least 1 character in length
For each prefix [0,i] for i going from 0 to len(searchWord) in searchWord, find all
words in products sharing that prefix. If there are more than three words that match the
prefix, return the three lexicographically smallest words

Edge cases: searchWord empty => return three lexicographically smallest words in products
searchWord[0] not in products[i][0] -> return emtpy lists for each character in searchWord

Brute force:
1. sort products
2. let "matches" be the list of all strings in products with prefix searchWord[0:i]
3. for each next prefix in searchWord, find matching prefixes in the last added sublist of "matches"
4. add the first three or less matches to the returned array


Example 1:

Input: products = ["mobile","mouse","moneypot","monitor","mousepad"], searchWord = "mouse"
Output: [["mobile","moneypot","monitor"],["mobile","moneypot","monitor"],["mouse","mousepad"],["mouse","mousepad"],["mouse","mousepad"]]
Explanation: products sorted lexicographically = ["mobile","moneypot","monitor","mouse","mousepad"].
After typing m and mo all products match and we show user ["mobile","moneypot","monitor"].
After typing mou, mous and mouse the system suggests ["mouse","mousepad"].

f(idx) = all words starting with char at searchWord[idx]
f(idx+1) = all words from f(idx) with char at idx+1 matching searchWord[idx+1]

state space is O(len(searchWord)) or O(len(searchWord) * len(products))?

f(0) = {product in products | product[0] == searchWord[0]}


*/
func suggestedProducts(products []string, searchWord string) [][]string {
	slices.Sort(products)
	if searchWord == "" {
		subList := []string{}
		for i := 0; i < 3 && i < len(products); i++ {
			sublist = append(subList, products[i])
		}
		matches = append(matches, subList)
		return [][]string{subList}
	}
	matches := [][]string{}
	subList := []string{}
	subMatches := []string{}
	for _, product := range products {
		if product[0] == searchWord[0] {
			if len(subList) < 3 {
				subList = append(subList, product)
			}
			subMatches = append(subMatches, product)
		}
	}
	matches = append(matches, subMatches)
	findProducts(matches, searchWord, 1)
	return matches
}

func findProducts(matches [][]string, searchWord string, i int) {
	if index == len(searchWord) {
		return
	}
	subList := []string{}
	subMatches := []string{}
	for _, product := range matches[len(matches)-1] {
		if i < len(product) && product[i] == searchWord[i] {
			if len(subList) < 3 {
				subList = append(subList, product)
			}
			subMatches = append(subMatches, product)
		}
	}
	matches = append(matches, subMatches)

}

func main() {}
