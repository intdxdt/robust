package robust

//Robust Cmp
func Cmp(a, b []float64) float64 {
	d := Subtract(a, b)
	return d[len(d)-1]
}
