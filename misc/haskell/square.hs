-- Version 1: clean 2-arg findPair delegates to a separate recursive function.
--   findPair predicate xs = findPair' predicate xs 0
--   The delegator is a one-liner. The real recursion lives in findPair'.
--   pro: both functions are independently usable, no where shadowing
--   con: two top-level names, caller needs to know about findPair'

findPair :: (a -> a -> Bool) -> [a] -> Maybe (Int, Int)
findPair pred xs = findPair' pred xs 0

findPair' :: (a -> a -> Bool) -> [a] -> Int -> Maybe (Int, Int)
findPair' _ [] _ = Nothing
findPair' _ [_] _ = Nothing
findPair' predicate (x:rest) i =
    case [ j | (j, y) <- zip [0..] rest, predicate x y ] of
        []    -> findPair' predicate rest (i + 1)
        (j:_) -> Just (i, i + j + 1)

-- Version 2: counter exposed directly, caller must pass 0.
--   findPair predicate list 0
--   findPair IS the recursive function — no delegation, no helper.
--   pro: single definition, fully self-referential
--   con: caller must pass 0 for start index

findPairDirect :: (a -> a -> Bool) -> [a] -> Int -> Maybe (Int, Int)
findPairDirect _ [] _ = Nothing
findPairDirect _ [_] _ = Nothing
findPairDirect predicate (x:rest) i =
    case [ j | (j, y) <- zip [0..] rest, predicate x y ] of
        []    -> findPairDirect predicate rest (i + 1)
        (j:_) -> Just (i, i + j + 1)
