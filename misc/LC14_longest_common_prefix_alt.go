package main

/*

longestCommonPrefix

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters if it is non-empty.
*/

func main() {
	//var strs = []string{"flower", "flow", "flight"}
	//var strs = []string{"dog","racecar","car"}
	var strs = []string{"flower", "flower", "flower", "flower"}
	println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	var common = ""
	var commonStr = ""
	for length := 0; ; length++ {
		for i := 0; i < len(strs); i++ {
			if len(strs[i]) == length {
				return commonStr
			}
			if i == 0 {
				common = strs[i][length : length+1]
			} else if common != strs[i][length:length+1] {
				return commonStr
			}
		}
		commonStr = commonStr + common
	}
}
