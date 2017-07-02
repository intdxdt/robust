package robust_seg_intersection

import (
	prod "robust/two_product"
	sum "robust/robust_sum"
	scale "robust/robust_scale"
	comp "robust/robust_compress"
	intersects "robust/robust_segseg"
)
var twoProduct      = prod.TwoProduct
var robustSum       = sum.RobustSum
var robustScale     = scale.RobustScale
var compress        = comp.RobustCompress
var robustIntersect = intersects.SegIntersects

// Find solution to system of two linear equations
//
//  | a[0]  a[1]   1 |
//  | b[0]  b[1]   1 |  =  0
//  |  x      y    1 |
//
//  | c[0]  c[1]   1 |
//  | d[0]  d[1]   1 |  =  0
//  |  x      y    1 |
//
func exactIntersect(a, b, c, d []float64)[][]float64 {

	if !robustIntersect(a, b, c, d) {
		return [][]float64{{0}, {0}, {0}}
	}

	var x1    = robustSum([]float64{c[1]}, []float64{-d[1]})
	var y1    = robustSum([]float64{-c[0]}, []float64{d[0]})

	var denom = robustSum(
		robustSum(robustScale(y1, a[1]), robustScale(y1, -b[1])),
		robustSum(robustScale(x1, a[0]), robustScale(x1, -b[0])),
	)

	var w0 = robustSum(twoProduct(-a[0], b[1]), twoProduct(a[1], b[0]))
	var w1 = robustSum(twoProduct(-c[0], d[1]), twoProduct(c[1], d[0]))

	//Calculate nX, nY
	var nX = robustSum(
		robustSum(robustScale(w1, a[0]), robustScale(w1, -b[0])),
		robustSum(robustScale(w0, -c[0]), robustScale(w0, d[0])),
	)

	var nY = robustSum(
		robustSum(robustScale(w1, a[1]),  robustScale(w1, -b[1])),
		robustSum(robustScale(w0, -c[1]), robustScale(w0,  d[1])),
	)

	return [][]float64{compress(nX), compress(nY), compress(denom)}
}
