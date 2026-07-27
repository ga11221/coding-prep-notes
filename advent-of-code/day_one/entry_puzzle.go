package main

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
)

func main() {
	f, _ := os.ReadFile("puzzle_input.txt")
	scanner := bufio.NewScanner(bytes.NewReader(f))
	start := 50
	zeroCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		turn, _ := strconv.Atoi(line[1:])
		if turn > 100 { //checked file - turns are no more than 3 digits long
			div := turn / 100
			turn = turn - (div * 100)
			zeroCount += div
		}

		if line[0:1] == "R" {
			end := start + turn
			if end > 100 {
				zeroCount++
			}
			start = end % 100
		} else {
			left := start - turn
			if left < 0 {
				if start > 0 {
					zeroCount++
				}
				start = 100 + left
			} else {
				start = left
			}
		}
		if start == 0 {
			zeroCount++
		}
	}
	println()
	println(zeroCount)
}
