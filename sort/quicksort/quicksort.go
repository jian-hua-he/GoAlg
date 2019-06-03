package quicksort

func Sort(s []interface{}, f func(interface{}, interface{}) bool) []interface{} {
	partition(s, 0, len(s)-1, f)
	return s
}

func partition(s []interface{}, start int, end int, f func(interface{}, interface{}) bool) {
	if end-start <= 0 {
		return
	}

	j := start
	for i := start; i < end; i += 1 {
		if f(s[end], s[i]) {
			s[i], s[j] = s[j], s[i]
			j += 1
		}
	}

	s[end], s[j] = s[j], s[end]

	partition(s, j+1, end, f)
	partition(s, start, j-1, f)
}
