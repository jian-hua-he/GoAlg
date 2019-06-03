# Heapsort


Step of the Heapsort

1. Build a heapify array
2. Swap first element and last element in array
3. Reduce the length of the array
4. Repeat step 1 until array length is 1

For instance

```
[4, 2, 1, 5, 3]

      4(0)
     /    \
   2(1)  1(2)
  /    \
5(3)   3(4)


The rule of the index

Parent(i)      = floor((i-1) / 2)
Left Child(i)  = 2*i + 1
Right Child(i) = 2*i + 2

Heapify

[4, 2, 1, 5, 3]

    4
   / \
  2   1
 / \
5   3

    ↓

[4, 5, 1, 2, 3]

    4
   / \
  5   1
 / \
2   3

    ↓

[5, 4, 1, 2, 3]

    5
   / \
  4   1
 / \
2   3

Swap first element to last element

[3, 4, 1, 2, 5]

    3
   / \
  4   1
 / \
2  |5|

Heapify again

[3, 4, 1, 2, 5]

    3
   / \
  4   1
 / \
2  |5|

    ↓

[4, 3, 1, 2, 5]

    4
   / \
  3   1
 / \
2  |5|

Swap first element to last element

[2, 3, 1, 4, 5]

     2
    / \
   3   1
  / \
|4| |5|

Heapify again

[2, 3, 1, 4, 5]
     2
    / \
   3   1
  / \
|4| |5|

    ↓

[3, 2, 1, 4, 5]

     3
    / \
   2   1
  / \
|4| |5|

Swap first element to last element

[1, 2, 3, 4, 5]

     1
    / \
   2  |3|
  / \
|4| |5|

Heapify again

[1, 2, 3, 4, 5]

     1
    / \
   2  |3|
  / \
|4| |5|

    ↓
    
[2, 1, 3, 4, 5]

     2
    / \
   1  |3|
  / \
|4| |5|

Swap first element to last element

[1, 2, 3, 4, 5]

     1
    / \
  |2| |3|
  / \
|4| |5|

length equal 1, stop.

```

Complexity

```
+------------------+------------+------------+------------+
|                  |    Best    |   Average  |    Worst   |
+------------------+------------+------------+------------+
| Time complexity  | O(n log n) | O(n log n) | O(n log n) |
| Space complexity |            |            |    O(n)    |
+------------------+------------+------------+------------+
```