package maths

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func Min(xs []float64) float64 {
	for i := 1; i < len(xs); i++ {
		for j := 1; j < len(xs); j++ {
			var k int = j - 1
			if xs[j] < xs[k] {
				jVal := xs[j]
				kVal := xs[k]
				xs[j] = kVal
				xs[k] = jVal
			}
		}
	}
	return xs[0]
}

func Max(xs []float64) float64 {
	for i := 1; i < len(xs); i++ {
		for j := 1; j < len(xs); j++ {
			var k int = j - 1
			if xs[j] < xs[k] {
				jVal := xs[j]
				kVal := xs[k]
				xs[j] = kVal
				xs[k] = jVal
			}
		}
	}
	return xs[len(xs)-1]
}
