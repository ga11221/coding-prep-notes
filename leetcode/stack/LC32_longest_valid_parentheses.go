package main

// valid parentheses: no pop on empty stack and stack is empty at end of string
// length = 2*number of valid pops

// stack for parenthesis: push -- add "("
// pop -- remove last "("
// begin new counter on every first valid push and add 2 to counter for every valid pop until next invalid push/pop - set max = counter reset counter
// repeat til end of string - max = counter if counter > max
// detect where valid pars end and an invalid seq beqins????
// find the unpaired ( to determine where valid pattern ends ???

// save index of first valid open, push/pop indices until next invalid pop and save its index or end of string
// length is index of invalid pop - index of first valid open
// if reached end of string and stack not empty, length = bottom-most index still in stack minus first valid open, second valid string = second-to-bottom index - bottom-index- 1 still in stack, and so on
// invalid push are unpaired opens that remain on stack after reaching end of string
// invalid pop is a pop on empty stack
// invalid push are saved automatically to stack
// invalid pops need to be saved manually
// first push in string interleaved with saved invalid pops and pushes remaining on stack

/*
	    @TODO re-examine recurrence
		Close but it doesn't work in general. The problem is contiguity — valid pairs separated by invalid characters aren't the same substring.

Counterexample: "()xx()"
Your state cumulative over the prefix: unclosed=0, valid=2 → answer = 4. But longest valid substring is "()" (length 2), not 4. The two pairs aren't contiguous.
If you reset on invalid (x), then dp[n][1] only captures the last segment (length 2, correct in this case) but loses the max if a longer segment appeared earlier, e.g. "(()())x()" — first segment length 6, second length 2, reset-on-invalid gives 2.

	s[i] = '(' → dp[i] = 0
	s[i] = ')':
	  s[i-1] = '(' → dp[i] = (i-2 >= 0 ? dp[i-2] : 0) + 2
	  s[i-1] = ')' → if s[i-dp[i-1]-1] == '(' → dp[i] = dp[i-1] + 2 + (i-dp[i-1]-2 >= 0 ? dp[i-dp[i-1]-2] : 0)

	Answer = max(dp). This enforces contiguity structurally — same reason Kadane's uses combine instead of pick/skip.

		(  (   ( ) ( )
	    1      3 2 3 2
			   1 0 1 0
		(  (   ( ) ) )         )   )  ( ) )
	    1  2   3 2 1 0         -1 -2
		(1,0)
		(2,0)
		(3,0)
		(2, 1)
		(1,2)
		(0, 3) -> not clean dp but when unclosed reaches 0, CAS on (globalMax, valid*2)
		(0, 0)
		(0, 0)
		(1, 0)
		(0, 1) -> CAS on (globalMax, valid*2)

		at each idx - increment unclosed if '(' else if unclosed > 0 add to valid and decrement unclosed else reset

		f(i) = (number of unclosed, number of valid)
		f(0) = (1,0)
		f(1) = (2,0)
		f(2) = (3,0)
		f(3) = (2, 1)
		f(4) = (3, 1)
		f(5) = (2,2)
		f(6) = (1,3)

		answer = dp[last][1]*2

		()
		f(-1) = 0
		f(0) = 1
		f(1) = f(0) - 1 if s[0] = '(' and s[1] = ')' = 0; length = idx - lastIdx + 1

		)(
		f(-1) = (0,0)
		f(0) = -1
		f(1) = 1

		(()(
		f(0) = 1
		f(1) = 2
		f(2) = 1
		f(3) = 2

		()()
		f(-1) = 0
		f(0) = 1
		f(1) = f(0) - 1 if s[0] = '(' and s[1] = ')' = 0
		f(2) = 1
		f(3) = f(2) - 1 = 0; j-1+1 where i f(i) = 1 and f(j) = 0 and no -1's occur in [i..j] window

		// global state + accumulator
		f(idx, startingIdx, openingParenCount) with globalMax

		f(idx)
			if idx = length(s)-1 {
				if opening {
					return 1
				}
				return -1
			}
			return 1+f(idx+1) if opening else f(idx+1) -1
*/
func longestValidParentheses(s string) int {
	stack := []int{}
	bs := []byte(s)
	invalidPops := []int{}
	for i, paren := range bs {
		if paren == '(' {
			stack = append(stack, paren)
		} else {
			if len(stack) == 0 {
				invalidPops = append(invalidPops, i)
			} else {
				stack = stack[:len(s)-1]
			}
		}
	}
	if len(stack) == 0 {
	}
	return 0
}
