package product

import (
	sum "robust/sum"
	scale "robust/scale"
)

//Robust Product
func RobustProduct(a, b []float64) []float64 {
	if len(a) == 1 {
		return scale.RobustScale(b, a[0])
	}
	if len(b) == 1 {
		return scale.RobustScale(a, b[0])
	}
	if len(a) == 0 || len(b) == 0 {
		return []float64{0}
	}
	var r = []float64{0}
	if len(a) < len(b) {
		for i := 0; i < len(a); i++ {
			r = sum.RobustSum(r, scale.RobustScale(b, a[i]))
		}
	} else {
		for i := 0; i < len(b); i++ {
			r = sum.RobustSum(r, scale.RobustScale(a, b[i]))
		}
	}
	return r
}
