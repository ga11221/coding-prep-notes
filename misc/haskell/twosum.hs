import Data.Map (Map, empty, insert, lookup)

--
-- The pattern: a clean 2-arg API wraps a 4-arg recursive function that
--   carries state the outer signature can't express (index + map).
--
-- Two styles, same semantics:
--
--   Style A (go/where):  local binding in where block
--     pro:  fewer top-level names, go is private implementation detail
--     con:  go isn't directly testable or callable
--
--   Style B (delegator): separate top-level functions
--     pro:  each variant independently usable and testable
--     con:  more top-level names
--
-- The label (go / twoSum'' / helper / search) is incidental — what matters
-- is having a function with the right arity to thread extra state through
-- the recursion.
--

-- Style A: go/where
twoSum :: [Int] -> Int -> Maybe (Int, Int)
twoSum xs target = go xs 0 empty
  where
    go [] _ _ = Nothing
    go (x:rest) i seen =
        case Data.Map.lookup (target - x) seen of
            Just j  -> Just (j, i)
            Nothing -> go rest (i + 1) (Data.Map.insert x i seen)

-- Style B: delegator + recursive pair
twoSum' :: [Int] -> Int -> Maybe (Int, Int)
twoSum' xs target = twoSum'' xs target 0 empty

twoSum'' :: [Int] -> Int -> Int -> Map Int Int -> Maybe (Int, Int)
twoSum'' [] _ _ _ = Nothing
twoSum'' (x:rest) target i seen =
    case Data.Map.lookup (target - x) seen of
        Just j  -> Just (j, i)
        Nothing -> twoSum'' rest target (i + 1) (Data.Map.insert x i seen)
