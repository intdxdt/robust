package robust

func Det2(m [][]float64) []float64 {
	return Compress(
		rsum(tprod(m[0][0], m[1][1]), tprod(-m[0][1], m[1][0])),
	)
}

func Det3(m [][]float64) []float64 {
	return Compress(
		rsum(
			rscale(rsum(tprod(m[1][1], m[2][2]), tprod(-m[1][2], m[2][1])), m[0][0]),
			rsum(
				rscale(rsum(tprod(m[1][0], m[2][2]), tprod(-m[1][2], m[2][0])), -m[0][1]),
				rscale(rsum(tprod(m[1][0], m[2][1]), tprod(-m[1][1], m[2][0])), m[0][2]),
			),
		))
}
