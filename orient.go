package robust

import (
	"math"
)


const EPSILON = 1.1102230246251565e-16
const ERRBOUND3 = (3.0 + 16.0*EPSILON) * EPSILON
const ERRBOUND4 = (7.0 + 56.0*EPSILON) * EPSILON

//orientation in 2d space
// < 0 if ccw - c is on left of segment(a, b)
// > 0 if cw - c is on right of segment(a, b)
// = 0 if a, b, and c are coplanar
func Orientation2D(a, b, c []float64) float64 {
	var l = (a[1] - c[1]) * (b[0] - c[0])
	var r = (a[0] - c[0]) * (b[1] - c[1])
	var det = l - r
	var s float64
	if l > 0 {
		if r <= 0 {
			return det
		} else {
			s = l + r
		}
	} else if l < 0 {
		if r >= 0 {
			return det
		} else {
			s = -(l + r)
		}
	} else {
		return det
	}
	tol := ERRBOUND3 * s
	if det >= tol || det <= -tol {
		return det
	}
	return orientation3Exact(a, b, c)
}

func Orientation3D(a, b, c, d []float64) float64 {
	var adx = a[0] - d[0]
	var bdx = b[0] - d[0]
	var cdx = c[0] - d[0]
	var ady = a[1] - d[1]
	var bdy = b[1] - d[1]
	var cdy = c[1] - d[1]
	var adz = a[2] - d[2]
	var bdz = b[2] - d[2]
	var cdz = c[2] - d[2]
	var bdxcdy = bdx * cdy
	var cdxbdy = cdx * bdy
	var cdxady = cdx * ady
	var adxcdy = adx * cdy
	var adxbdy = adx * bdy
	var bdxady = bdx * ady

	var det = adz*(bdxcdy-cdxbdy) +
		bdz*(cdxady-adxcdy) + cdz*(adxbdy-bdxady)

	var permanent = (math.Abs(bdxcdy)+math.Abs(cdxbdy))*math.Abs(adz) +
		(math.Abs(cdxady)+math.Abs(adxcdy))*math.Abs(bdz) +
		(math.Abs(adxbdy)+math.Abs(bdxady))*math.Abs(cdz)

	var tol = ERRBOUND4 * permanent
	if (det > tol) || (-det > tol) {
		return det
	}
	return orientation4Exact(a, b, c, d)
}

//orientation 2d exact
func orientation3Exact(m0, m1, m2 []float64) float64 {
	p := rsum(
		rsum(tprod(m1[1], m2[0]), tprod(-m2[1], m1[0])),
		rsum(tprod(m0[1], m1[0]), tprod(-m1[1], m0[0])),
	)
	n := rsum(tprod(m0[1], m2[0]), tprod(-m2[1], m0[0]))
	d := rdiff(p, n)
	return d[len(d)-1]
}

func orientation4Exact(m0, m1, m2, m3 []float64) float64 {
	p := rsum(
		rsum(
			rscale(rsum(tprod(m2[1], m3[0]), tprod(-m3[1], m2[0])), m1[2]),
			rsum(
				rscale(rsum(tprod(m1[1], m3[0]), tprod(-m3[1], m1[0])), -m2[2]),
				rscale(rsum(tprod(m1[1], m2[0]), tprod(-m2[1], m1[0])), m3[2]),
			),
		),
		rsum(
			rscale(rsum(tprod(m1[1], m3[0]), tprod(-m3[1], m1[0])), m0[2]),
			rsum(
				rscale(rsum(tprod(m0[1], m3[0]), tprod(-m3[1], m0[0])), -m1[2]),
				rscale(rsum(tprod(m0[1], m1[0]), tprod(-m1[1], m0[0])), m3[2]),
			),
		),
	)

	n := rsum(
		rsum(
			rscale(rsum(tprod(m2[1], m3[0]), tprod(-m3[1], m2[0])), m0[2]),
			rsum(
				rscale(rsum(tprod(m0[1], m3[0]), tprod(-m3[1], m0[0])), -m2[2]),
				rscale(rsum(tprod(m0[1], m2[0]), tprod(-m2[1], m0[0])), m3[2]),
			),
		),
		rsum(
			rscale(rsum(tprod(m1[1], m2[0]), tprod(-m2[1], m1[0])), m0[2]),
			rsum(
				rscale(rsum(tprod(m0[1], m2[0]), tprod(-m2[1], m0[0])), -m1[2]),
				rscale(rsum(tprod(m0[1], m1[0]), tprod(-m1[1], m0[0])), m2[2]),
			),
		),
	)
	d := rdiff(p, n)
	return d[len(d)-1]

}
