# Question Description

Consider an array contains the data with the format below:

```
[
    {
        name: 'Foo',
        priority: 1,
    },
    {
        name: 'Bar',
        priority: 2,
    },
    {
        name: 'Fuzz',
        priority: 1,
    },
    {
        name: 'Buzz',
        priority: 1,
    },
]
```

Priority is higher if the number is lower. If there is the same priority number. Original one should not move. In the example above, The result should be:

```
[
    {
        name: 'Foo',
        priority: 1,
    },
    {
        name: 'Fuzz',
        priority: 1,
    },
    {
        name: 'Buzz',
        priority: 1,
    },
    {
        name: 'Bar',
        priority: 2,
    },
]
```