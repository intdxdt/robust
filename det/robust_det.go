package det

import (
	"robust/sum"
	"robust/scale"
	"robust/compress"
	"robust/two"
)

var prod = two.Product
var rsum = sum.RobustSum
var rscale = scale.RobustScale

func RobustDet2(m [][]float64) []float64 {
	return compress.RobustCompress(
		rsum(prod(m[0][0], m[1][1]), prod(-m[0][1], m[1][0])),
	)
}

func RobustDet3(m [][]float64) []float64 {
	return compress.RobustCompress(
		rsum(
			rscale(rsum(prod(m[1][1], m[2][2]), prod(-m[1][2], m[2][1])), m[0][0]),
			rsum(
				rscale(rsum(prod(m[1][0], m[2][2]), prod(-m[1][2], m[2][0])), -m[0][1]),
				rscale(rsum(prod(m[1][0], m[2][1]), prod(-m[1][1], m[2][0])), m[0][2]),
			),
		))
}
