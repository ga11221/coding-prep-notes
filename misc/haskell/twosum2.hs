-- Two Sum II (sorted): pure pick/skip via two pointers
-- The sorted array lets us decide which index to skip:
--   sum < target → skip left (need a larger value)
--   sum > target → skip right (need a smaller value)
--   sum == target → pick the pair

-- Version 1: go/where (inner recursive helper)
twoSumII :: [Int] -> Int -> Maybe (Int, Int)
twoSumII nums target = go 0 (length nums - 1)
  where
    go i j
      | i >= j             = Nothing          -- exhausted
      | nums!!i + nums!!j < target = go (i+1) j     -- skip left
      | nums!!i + nums!!j > target = go i (j-1)     -- skip right
      | otherwise          = Just (i+1, j+1)  -- 1-indexed

-- Version 2: delegator (clean 2-arg → exposed 4-arg)
--   Initial left/right passed explicitly.
--   twoSumII' is the recursive function; twoSumII just sets up the pointers.
twoSumII_simple :: [Int] -> Int -> Maybe (Int, Int)
twoSumII_simple nums target = twoSumII' nums target 0 (length nums - 1)

twoSumII' :: [Int] -> Int -> Int -> Int -> Maybe (Int, Int)
twoSumII' _ _ i j
  | i >= j = Nothing
twoSumII' nums target i j
  | sum < target = twoSumII' nums target (i+1) j
  | sum > target = twoSumII' nums target i (j-1)
  | otherwise    = Just (i+1, j+1)
  where sum = nums!!i + nums!!j
