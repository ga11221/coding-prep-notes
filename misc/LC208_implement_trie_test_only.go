package main

import (
	"testing"
)

func TestTrie(t *testing.T) {
	t.Run("Implement Trie", func(t *testing.T) {
		trie := Constructor()
		trie.Insert("apple")
		t.Log(trie.Search("apple"))   // return True
		t.Log(trie.Search("app"))     // return False
		t.Log(trie.StartsWith("app")) // return True
		trie.Insert("app")
		t.Log(trie.Search("app")) // return True

		//got := longestConsecutive(tt.nums)
		//if got != tt.want {
		//t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
		//return
		//}
	})
}
