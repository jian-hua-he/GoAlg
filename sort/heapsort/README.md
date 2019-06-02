# Heapsort

Heapsort is a sorting algrithm. Structure like binary tree, for instance:

```
[4, 2, 1, 5, 3]

      4(0)
     /    \
   2(1)  1(2)
  /    \
5(3)   3(4)
```

If you look closely. You can find the rule of the index:
```
Parent(i)      = floor((i-1) / 2)
Left Child(i)  = 2*i + 1
Right Child(i) = 2*i + 2
```

Step of the Heapsort:
1. Build a heapify array
2. Swap first element and last element in array
3. Reduce the length of the array
4. Repeat step 2 until array length is 1
