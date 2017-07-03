package robust

//fast two sum
func TwoSum(a, b float64) []float64 {
	x  := a + b
	bv := x - a
	av := x - bv
	br := b - bv
	ar := a - av
	return []float64{ar + br, x}
}
