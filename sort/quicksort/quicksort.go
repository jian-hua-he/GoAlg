package quicksort

import "fmt"

func Sort(s []interface{}, f func(interface{}, interface{}) bool) []interface{} {
	pivot := len(s) - 1

	j := 0
	for i := 0; i < len(s)-1; i += 1 {
		if f(s[pivot], s[i]) {
			s[i], s[j] = s[j], s[i]
			j += 1
		}
	}

	s[pivot], s[j] = s[j], s[pivot]

	fmt.Println(s)

	return s
}
