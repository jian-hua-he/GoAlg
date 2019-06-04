package mergesort

import (
	"math"
)

func Sort(s []interface{}, f func(interface{}, interface{}) bool) []interface{} {
	arr := append([]interface{}{}, s...)
	return mergeSort(arr, f)
}

func mergeSort(s []interface{}, f func(interface{}, interface{}) bool) []interface{} {
	if len(s) <= 1 {
		return s
	}

	middle := int(math.Floor(float64(len(s) / 2)))
	left := mergeSort(s[:middle], f)
	right := mergeSort(s[middle:], f)
	result := merge(left, right, f)

	return result
}

func merge(left []interface{}, right []interface{}, f func(interface{}, interface{}) bool) []interface{} {
	result := make([]interface{}, 0)

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if f(right[j], left[i]) {
			result = append(result, left[i])
			i += 1
		} else {
			result = append(result, right[j])
			j += 1
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
