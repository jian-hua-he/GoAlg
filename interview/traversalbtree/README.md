# Traversal BTree

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
