package prioritysort

func Sort(s []interface{}, f func(interface{}, interface{}) bool) []interface{} {
	arr := s

	for i := len(s) - 1; i > 0; i -= 1 {
		arr = heapify(arr, i, i+1, f)
		arr[0], arr[i] = arr[i], arr[0]
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

	s[root], s[swap] = s[swap], s[root]

	return heapify(s, root-1, end, f)
}
