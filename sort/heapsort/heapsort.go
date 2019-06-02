package heapsort

import (
	"fmt"
)

func Sort(s []interface{}, f func(interface{}, interface{}) bool) []interface{} {
	arr := s

	for i := len(s) - 1; i > len(s)-2; i -= 1 {
		fmt.Println("=======")
		fmt.Println(i)
		fmt.Println("Before")
		fmt.Println(arr)

		arr = heapify(arr, i, i+1, f)

		fmt.Println("After")
		fmt.Println(arr)

		arr[0], arr[i] = arr[i], arr[0]

		fmt.Println("Swap")
		fmt.Println(arr)
		fmt.Println("=======")
	}

	return arr
}

func heapify(s []interface{}, root int, end int, f func(interface{}, interface{}) bool) []interface{} {
	if root < 0 {
		return s
	}

	left := root*2 + 1
	right := root*2 + 2
	swap := root

	if left < end && f(s[left], s[root]) {
		swap = left
	}

	if right < end && f(s[right], s[swap]) {
		swap = right
	}

	if root != swap {
		s[root], s[swap] = s[swap], s[root]
	}

	return heapify(s, root-1, end, f)
}
