package robust

//Robust Product
func Product(a, b []float64) []float64 {
	if len(a) == 1 {
		return Scale(b, a[0])
	}
	if len(b) == 1 {
		return Scale(a, b[0])
	}
	if len(a) == 0 || len(b) == 0 {
		return []float64{0}
	}
	var r = []float64{0}
	if len(a) < len(b) {
		for i := 0; i < len(a); i++ {
			r = Sum(r, Scale(b, a[i]))
		}
	} else {
		for i := 0; i < len(b); i++ {
			r = Sum(r, Scale(a, b[i]))
		}
	}
	return r
}
