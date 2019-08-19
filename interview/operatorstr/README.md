# Operator String

Write a function that analyzes the string and returns a number. The string contains the operator and the number:

- `POP`: Get the most top element of the stack and remove the element from the stack
- `DUP`: Duplicate the most top element and put it back to the stack
- `+`: Pop two elements and add up. Put the added up number back to the top of the stack
- `-`: Pop two elements and subtract the second element from the first element. Put the result back to the top of the stack
- `${Num}`: Put the number into the stack

For example:

```
"13 DUP 13 7 + 9 POP -"

+----------+-------------------+
| Operator |       Stack       |
+----------+-------------------+
|     13   | [13]              |
|    DUP   | [13, 13]          |
|     13   | [13, 13, 13]      |
|      7   | [13, 13, 13, 7]   |
|      +   | [13, 13, 20]      |
|      9   | [13, 13, 20, 9]   |
|    POP   | [13, 13, 20]      |
|      -   | [13, 7]           |
+----------+-------------------+
```

The result is `7`.

Note:
- If the string has an invalid operator. Return `-1`
- If stack elements less than 1. `+` and `-` will be considered as an invalid operator.