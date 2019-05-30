package medianunsort

func FindMedian(s []float64) float64 {
	sum := sum(s)
	avg := sum / float64(len(s))
	diff := absVal(avg, s[0])

	result := 0.0
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
