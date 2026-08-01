package main

import (
	"fmt"
	"math"
)

/*
=============== INCOMPLETE =============
In order to add a letter, Alice has to press the key of the corresponding digit i times, where i is the position of the letter in the key.

For example, to add the letter 's', Alice has to press '7' four times. Similarly, to add the letter 'k', Alice has to press '5' twice.
Note that the digits '0' and '1' do not map to any letters, so Alice does not use them.
However, due to an error in transmission, Bob did not receive Alice's text message but received a string of pressed keys instead.

For example, when Alice sent the message "bob", Bob received the string "2266622".
Given a string pressedKeys representing the string received by Bob, return the total number of possible text messages Alice could have sent.

Since the answer may be very large, return it modulo 109 + 7.


*/

func maini1i32ujh() {
	var s = "222223"
	// "2" 		-> 1
	// "22" 	-> 2
	// "222" 	-> aaa, ab, ba, c 4
	// "2222"	-> 1+1+1+1+1+1+1
	// ac, ca, aaaa, aba, baa, aab, bb
	//	"2"+"222"		-> 4
	// "22"+"22" -> 2
	// "222" + "a" -> 1

	// "22222"

	// 1 + 4->7 // aac, aca, aaaaa, aaba, abaa, aaab, abb
	// 2 + 3->4 // baaa, bab,bba, bc
	// 3+2->2 // caa, cb
	// 4+1->1 //

	// "7" -> p 1
	// "77" -> pp,q 2
	// "777" -> ppp, pq, qp, r 4
	// "7777" -> pppp, ppq, qpp, pr, rp,s, pqp, qq 8
	// "77777" ->
	// 1+4 -> 8
	// 2+3 -> 4
	// 3+2 -> 2
	// 4+1 -> 1
	fmt.Print(countTexts(s))
}

var digits = map[int]int{
	1: 1,
	2: 2,
	3: 4,
}

var sevenAndNine = map[int]int{
	1: 1,
	2: 2,
	3: 4,
	4: 8,
}

func countTexts(pressedKeys string) int {
	// 1. separate pressedKeys into chunks of repeating digits and count instances of repeated digits
	// For every chunk:
	// 2. n = length of each chunk
	// 3. for i from 1 to n, find combinations:
	//    if repeated key chunk[i] is 7 or 9, m = sevenAndNine, else m = digits
	//    	if n-1 in map m return, m[n-1]
	//     	else recurse

	if len(pressedKeys) == 1 {
		return 1
	}
	// 1.
	var currIdx = 0
	var curr = string(pressedKeys[currIdx])
	var count = 0
	for i := 1; i < len(pressedKeys); i++ {
		if string(pressedKeys[i]) != curr {
			if curr == "7" || curr == "9" {
				count = count + _countTextsSevensAndNines((i)-currIdx)
			}
			count = count + _countTextsDigits((i)-currIdx)
			currIdx = i
			curr = string(pressedKeys[currIdx])
		}
	}
	// check the last element
	// check if there's only one chunk

	return count % int(math.Pow10(9)+7)
}

func _countTextsSevensAndNines(i int) int {
	if v, ok := sevenAndNine[i]; ok {
		return v
	}
	return _countTextsSevensAndNines(i-1) + sevenAndNine[i]
}

func _countTextsDigits(length int) int {
	if v, ok := digits[length]; ok {
		return v
	}
	var sum = 0
	for j := length; j > 1; j-- {
		textsDigits := _countTextsDigits(j - 1)
		sum = sum + textsDigits
	}
	digits[length] = sum

	return sum
}
