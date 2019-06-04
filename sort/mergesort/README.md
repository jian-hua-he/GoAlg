# Merge Sort

The steps of the Merge Sort

1. Divide the array until one element left
2. Compare two elements and merge
3. Repeat step one until all element merged

For example

```
Divide: |
Merge: ()

[6, 2, 1, 5, 3, 4]
        ↓
[6, 2, 1] | [5, 3, 4]
        ↓
[6] | [2, 1]  [5] | [3, 4]
        ↓
[6]  [2] | [1]  [5]  [3] | [4]
        ↓
[6]  ([1, 2])  [5] ([3, 4])
        ↓
([1, 2, 6]) ([3, 4, 5])
        ↓
([1, 2, 3, 4, 5, 6])


Merge process

Current element: {}

Merged: array: []
Divided array: [1, 2, 6] [3, 4, 5]
                ^         ^

1 < 3. Put 1 to the merged array and move index of the left array to next

Merged: array: [1]
Divided array: [1, 2, 6] [3, 4, 5]
                   ^      ^

2 < 3. Put 2 to the merged array and move index of the left array to next

Merged: array: [1, 2]
Divided array: [1, 2, 6] [3, 4, 5]
                      ^   ^

6 > 3. Put 3 to the merged array and move index of the right array to next

Merged: array: [1, 2, 3]
Divided array: [1, 2, 6] [3, 4, 5]
                      ^      ^

6 > 4. Put 4 to the merged array and move index of the right array to next

Merged: array: [1, 2, 3, 4]
Divided array: [1, 2, 6] [3, 4, 5]
                      ^         ^

6 > 5. Put 5 to the merged array.

Merged: array: [1, 2, 3, 4, 5]
Divided array: [1, 2, 6] [3, 4, 5]
                      ^         ^

Index reaches the maximum. Put elements from the left array then put elements from the right array into the merged array.

Merged array: [1, 2, 3, 4, 5, 6]

```

Complexity

```
+------------------+------------+------------+-------------------+
|                  |    Best    |   Average  |       Worst       |
+------------------+------------+------------+-------------------+
| Time complexity  | O(n log n) | O(n log n) |     O(n log n)    |
+------------------+------------+------------+-------------------+
| Space complexity |            |            |       O(n)        |
|                  |            |            | O(1) linked list  |
+------------------+------------+------------+-------------------+
```
