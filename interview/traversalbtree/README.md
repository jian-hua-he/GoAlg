# Traversal BTree

## Description

Write a function to count the node that equal or bigger than the root value.

For example:

```
     5
   /   \
  2     3
 / \   /
20 21 10
```

Root value is 5. Matched nodes are 5, 20, 21 and 10. Answer is 4.

```
    8
   / \
  3   2
 /
8
```

Root value is 8. Matched nodes are 8, 8. Answer is 2.

Note:
- `V` represents the value of node. `0 <= V <= 100,000`
- `N` represents the count of node. `0 < N <= 50,000`

## Solution 1

We can use recursion to solve the issue. But if we need to handle a lot of nodes. It would cause a stack overflow. So we might think liner way to solve the issue.

## Solution 2

The basic idea of Solution 2: if the node has no children. Compare the node with the root value and cut it from the tree. Then let the pointer go to the previous parent. We need two arrays to remember parents and previous direction.

For instance:

```
Count: 0
Current: 5
Parents: []
Directions: []

   > 5
   /   \
  2     3
 / \   /
20 21 10
```

Left Node is not null. Go left.

```
Count: 0
Current: 2
Parents: [5]
Directions: [left]

     5
   /   \
> 2     3
 / \   /
20 21 10
```

Left Node is not null. Go left.

```
Count: 0
Current: 20
Parents: [5, 2]
Directions: [left, left]

       5
     /   \
    2     3
   / \   /
> 20 21 10
```

No children in this node. Compare the root value and base on the direction to cut it down. Then move the pointer to the parent.

```
Count: 1
Current: 2
Parents: [5]
Directions: [left]

       5
     /   \
  > 2     3
     \   /
     21 10
```

Right Node is not null. Go right.

```
Count: 1
Current: 21
Parents: [5, 2]
Directions: [left, right]

       5
     /   \
    2     3
     \   /
   > 21 10
```

No children in this node. Compare the root value and base on the direction to cut it down. Then move the pointer to the parent.

```
Count: 2
Current: 2
Parents: [5]
Directions: [left]

       5
     /   \
  > 2     3
         /
        10
```

No children in this node. Compare the root value and base on the direction to cut it down. Then move the pointer to the parent.

```
Count: 2
Current: 5
Parents: []
Directions: []

  > 5
     \
      3
     /
    10
```

Right Node is not null. Go right.

```
Count: 2
Current: 3
Parents: [5]
Directions: [right]

    5
     \
    > 3
     /
    10
```

Left Node is not null. Go left.

```
Count: 2
Current: 10
Parents: [5, 3]
Directions: [right, left]

    5
     \
      3
     /
  > 10
```

No children in this node. Compare the root value and base on the direction to cut it down. Then move the pointer to the parent.

```
Count: 3
Current: 3
Parents: [5]
Directions: [right]

    5
     \
    > 3
```

No children in this node. Compare the root value and base on the direction to cut it down. Then move the pointer to the parent.

```
Count: 3
Current: 5
Parents: []
Directions: []

> 5
```

No children in this node. Compare the root value and base on the direction to cut it down. Then move the pointer to the parent.

```
Count: 4
Current: nil
Parents: []
Directions: []

nil
```

Answer is 4