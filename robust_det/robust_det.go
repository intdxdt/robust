package robust_det

import (
	"robust/robust_sum"
	"robust/robust_scale"
	"robust/robust_compress"
	"robust/two_product"
)

var prod = two_product.TwoProduct
var sum = robust_sum.RobustSum
var scale = robust_scale.RobustScale
var compress = robust_compress.RobustCompress

func RobustDet2(m [][]float64) []float64 {
	return compress(
		sum(prod(m[0][0], m[1][1]), prod(-m[0][1], m[1][0])),
	)
}

func RobustDet3(m [][]float64) []float64 {
	return compress(
		sum(
			scale(sum(prod(m[1][1], m[2][2]), prod(-m[1][2], m[2][1])), m[0][0]),
			sum(
				scale(sum(prod(m[1][0], m[2][2]), prod(-m[1][2], m[2][0])), -m[0][1]),
				scale(sum(prod(m[1][0], m[2][1]), prod(-m[1][1], m[2][0])), m[0][2]),
			),
		))
}
