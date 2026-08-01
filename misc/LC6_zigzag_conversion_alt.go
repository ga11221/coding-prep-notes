package main

import "fmt"

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);


Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I
Example 3:

Input: s = "A", numRows = 1
Output: "A"


Constraints:

1 <= s.length <= 1000
s consists of English letters (lower-case and upper-case), ',' and '.'.
1 <= numRows <= 1000
*/

func main2343() {
	//var nums = []int{2, 7, 11, 15}
	//var target = 9

	//var nums = []int{3, 2, 4}
	//var target = 6

	var s = "AB"
	var numRows = 1

	fmt.Print(convert(s, numRows))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	matrix := make([][]string, numRows)
	for i := range matrix {
		strings := []string{}
		for i := 0; i < len(s); i++ {
			strings = append(strings, "")
		}
		matrix[i] = strings
	}
	rowNum := 0
	colNum := 0
	goingDown := true
	for _, runeVal := range s {
		matrix[rowNum][colNum] = string(runeVal)
		if goingDown {
			if rowNum == numRows-1 {
				goingDown = false
				rowNum--
				colNum++
			} else {
				rowNum++
			}
		} else {
			if rowNum == 0 {
				goingDown = true
				rowNum++
			} else {
				rowNum--
				colNum++
			}
		}

	}
	str := ""
	for _, strings := range matrix {
		for _, s1 := range strings {
			str += s1
		}
	}
	fmt.Printf("%#d\n", len(str))
	return str
}
