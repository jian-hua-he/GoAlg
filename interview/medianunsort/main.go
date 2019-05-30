package medianunsort

func FindMedian(s []float64) float64 {
	sum := sum(s)
	length := len(s)
	avg := sum / float64(length)

	if length%2 == 1 {
		return findInOdd(s, avg)
	}

	return findInEven(s, avg)
}

func sum(s []float64) float64 {
	sum := 0.0

	for _, v := range s {
		sum += v
	}

	return sum
}

func absVal(n1 float64, n2 float64) float64 {
	if n1 > n2 {
		return n1 - n2
	}

	return n2 - n1
}

func findInOdd(s []float64, avg float64) float64 {
	result := 0.0
	diff := absVal(avg, s[0])

	for _, v := range s {
		d := absVal(avg, v)
		if d >= diff {
			continue
		}

		result = v
		diff = d
	}

	return result
}

func findInEven(s []float64, avg float64) float64 {
	arr := s

	index := 0
	diff1 := absVal(avg, arr[0])
	n1 := arr[0]
	for i, v := range arr {
		d := absVal(avg, v)
		if d >= diff1 {
			continue
		}

		n1 = v
		index = i
		diff1 = d
	}

	arr = append(arr[0:index], arr[index+1:]...)
	diff2 := absVal(avg, arr[0])
	n2 := arr[0]
	for _, v := range arr {
		d := absVal(avg, v)
		if d >= diff2 {
			continue
		}

		n2 = v
		diff2 = d
	}

	return (n1 + n2) / 2
}
