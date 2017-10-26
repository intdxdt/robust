package validate_seq

import (
	to "github.com/intdxdt/robust/test_overlap"
	"math"
)

func ValidateSequence(x []float64) bool {
	var n = len(x)
	if n < 1 {
		return false
	}
	for i := 1; i < n; i++ {
		if math.Abs(x[i-1]) >= math.Abs(x[i]) {
			return false
		}
		if to.TestOverlap(x[i], x[i-1]) {
			return false
		}
	}
	return true
}
