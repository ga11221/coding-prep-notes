# Haskell First Session

## Setup
```bash
ghci   # start REPL
:load filename.hs   # load a file
:reload              # reload after edits
:q                   # quit
```

## Warmup (5 min)
```haskell
let square x = x * x
let fact n = if n == 0 then 1 else n * fact (n-1)
let fib n | n <= 1 = n | otherwise = fib (n-1) + fib (n-2)
```

## Guards and branching (10 min)
Write function definitions using guard chains (replaces if-else):
```haskell
sign x
  | x > 0  = "pos"
  | x < 0  = "neg"
  | otherwise = "zero"

findPair :: (Int -> Int -> Bool) -> [Int] -> Maybe (Int, Int)
findPair criteria nums = go 0 1
  where
    go i j
      | i >= length nums = Nothing
      | j >= length nums = go (i+1) (i+2)
      | criteria (nums!!i) (nums!!j) = Just (i, j)
      | otherwise = go i (j+1)
```

## Combinatorial functions (15 min)
Write return-value form recurrences — the code IS the recurrence:
```haskell
-- subsets: f([])=[[]]; f(x:xs)=[x:ys | ys <- subsets xs] ++ subsets xs
subsets :: [a] -> [[a]]
subsets [] = [[]]
subsets (x:xs) = [x:ys | ys <- subsets xs] ++ subsets xs

-- permutations: pick each n, prepend to perms of rest
permutations :: [a] -> [[a]]
permutations [] = [[]]
permutations xs = [x:ys | x <- xs, ys <- permutations (delete x xs)]

-- choose-positions (same-char insertion)
placements :: a -> [a] -> Int -> [[a]]
placements filler orig totalLen = go 0 0 []
  where
    n = length orig
    go pos origIdx acc
      | pos == totalLen = [reverse acc]
      | origIdx == n = [reverse $ replicate (totalLen - pos) filler ++ acc]
      | otherwise = go (pos+1) origIdx (filler:acc)
                 ++ go (pos+1) (origIdx+1) (orig !! origIdx : acc)
```

## Tips
- Write functions in a `.hs` file, then `:load` in ghci
- Use `:t expression` to inspect types
- Guard chains (`| condition = result`) replace if-else
- List comprehensions (`[expr | x <- list, condition]`) replace nested loops
- The return-value form is *the only form* — no mutation, no accumulators needed
