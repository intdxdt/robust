package robust


//Robust segment intersection of line segments
func SegIntersection(a, b, c, d []float64) [][]float64 {
	return exactIntersect(a, b, c, d)
}

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
func exactIntersect(a, b, c, d []float64) [][]float64 {

	if !SegIntersects(a, b, c, d) {
		return [][]float64{{0}, {0}, {0}}
	}

	var x1 = rsum([]float64{c[1]}, []float64{-d[1]})
	var y1 = rsum([]float64{-c[0]}, []float64{d[0]})

	var denom = rsum(
		rsum(rscale(y1, a[1]), rscale(y1, -b[1])),
		rsum(rscale(x1, a[0]), rscale(x1, -b[0])),
	)

	var w0 = rsum(tprod(-a[0], b[1]), tprod(a[1], b[0]))
	var w1 = rsum(tprod(-c[0], d[1]), tprod(c[1], d[0]))

	//Calculate nX, nY
	var nX = rsum(
		rsum(rscale(w1, a[0]), rscale(w1, -b[0])),
		rsum(rscale(w0, -c[0]), rscale(w0, d[0])),
	)

	var nY = rsum(
		rsum(rscale(w1, a[1]), rscale(w1, -b[1])),
		rsum(rscale(w0, -c[1]), rscale(w0, d[1])),
	)

	return [][]float64{
		Compress(nX), Compress(nY), Compress(denom),
	}
}
