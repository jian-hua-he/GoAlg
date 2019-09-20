# Minimum queries

## Description

generates the minimum amount of queries of given length.

```
function generateQueries (keywords[], maxCharLength)
output: queries[]
```

### Constraint

- It should be fast
- you cannot exceed the maximum maxCharLength, or the query would not be accepted

### Examples

```
keywords = ["hello", "happy", "great"]
maxCharLength = 50

output = "hello OR happy OR great"
```

```
keywords = ["hello", "happy", "great"]
maxCharLength = 15

output query 1: "hello OR happy"
output query 2: "great"
```

```
keywords = ["hello", "happy", "great"]
maxCharLength = 6

output query 1: "hello"
output query 2: "happy"
output query 3: "great"
```

```
keywards = ["foo", "bar", "fuzz", "buzz"]
maxCharLength = 6

output query 1: "foo OR bar"
output query 2: "fuzz"
output query 3: "buzz"
```

```
keywards = ["a", "bb", "ccc", "dddd", "eeeee", "ffffff", "ggggggg", "hhhhhhhh", "iiiiiii", "jjjjjj", "kkkkk", "llll", "mmm", "nn", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"]
maxCharLength = 13

output query 1: hhhhhhhh OR a
output query 2: ggggggg OR bb
output query 3: iiiiiii OR nn
output query 4: ffffff OR ccc
output query 5: jjjjjj OR mmm
output query 6: eeeee OR dddd
output query 7: kkkkk OR llll
output query 8: o OR p OR q
output query 9: r OR s OR t
output query 10: u OR v OR w
output query 11: x OR y OR z
```