# Quicksort

The steps of the Quicksort

1. Determine pivot in array
2. Given index i and j to Iterate array and compare pivot
3. If compare result is true, swap element i with element j and move index j to next
4. Move pivot element to the position j
5. Split two array with j element and repeat step 2

P.S

Pivot can be:
- The first element of the array
- The last element of the array
- The middle element of the array

For example

```
[9, 4, 1, 6, 7, 3, 8, 2, 5]

We determine the pivot is 5. Start to iterate array.

[9, 6, 1, 4, 7, 3, 8, 2, |5|]
 i
 j

9 is bigger then 5. Index i Move to next.

[9, 6, 1, 4, 7, 3, 8, 2, |5|]
    i
 j

6 is bigger than 5. Index i Move to next.

[9, 6, 1, 4, 7, 3, 8, 2, |5|]
       i
 j

1 is smaller than 5. Switch element i with element j. Index i and j move to next.

[1, 6, 9, 4, 7, 3, 8, 2, |5|]
          i
    j

4 is smaller than 5. Switch i and j then move to next.

[1, 4, 9, 6, 7, 3, 8, 2, |5|]
             i
       j

7 is bigger than 5. Index i Move to next

[1, 4, 9, 6, 7, 3, 8, 2, |5|]
                i
       j

3 is smaller than 5. Switch i and j then move to next.

[1, 4, 3, 6, 7, 9, 8, 2, |5|]
                   i
          j

8 is bigger than 5. Index i move to next.

[1, 4, 3, 6, 7, 9, 8, 2, |5|]
                      i
          j

2 is smaller than 5. Switch i and j then move to next.

[1, 4, 3, 2, 7, 9, 8, 6, |5|]
                          i
             j

Swap the pivot element to index j

[1, 4, 3, 2, |5|, 7, 9, 8, 6]

Split two array then countinue step 1

[1, 4, 3, |2|]
[7, 9, 8, |6|]

...

```

Complexity

```
+------------------+------------+------------+------------------+
|                  |    Best    |   Average  |       Worst      |
+------------------+------------+------------+------------------+
| Time complexity  | O(n log n) | O(n log n) |      O(n^2)      |
| Space complexity |            |            | O(n) or O(log n) |
+------------------+------------+------------+------------------+
```