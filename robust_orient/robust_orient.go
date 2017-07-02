package robust_orient

import (
	"math"
	"robust/robust_diff"
	"robust/robust_scale"
	"robust/robust_sum"
	"robust/two_product"
)

const EPSILON = 1.1102230246251565e-16
const ERRBOUND3 = (3.0 + 16.0*EPSILON) * EPSILON
const ERRBOUND4 = (7.0 + 56.0*EPSILON) * EPSILON

var sub = robust_diff.RobustDiff
var sum = robust_sum.RobustSum
var prod = two_product.TwoProduct
var scale = robust_scale.RobustScale

//orientation 1d space
func Orientation1d(a, b []float64) float64 {
	return b[0] - a[0]
}

//orientation in 2d space
func Orientation2d(a, b, c []float64) float64 {
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

func Orientation3d(a, b, c, d []float64) float64 {
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
		bdz*(cdxady-adxcdy) +
		cdz*(adxbdy-bdxady)

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
	p := sum(
		sum(prod(m1[1], m2[0]), prod(-m2[1], m1[0])),
		sum(prod(m0[1], m1[0]), prod(-m1[1], m0[0])),
	)
	n := sum(prod(m0[1], m2[0]), prod(-m2[1], m0[0]))
	d := sub(p, n)
	return d[len(d)-1]
}

func orientation4Exact(m0, m1, m2, m3 []float64) float64 {
	p := sum(
		sum(
			scale(sum(prod(m2[1], m3[0]), prod(-m3[1], m2[0])), m1[2]),
			sum(
				scale(sum(prod(m1[1], m3[0]), prod(-m3[1], m1[0])), -m2[2]),
				scale(sum(prod(m1[1], m2[0]), prod(-m2[1], m1[0])), m3[2]),
			),
		),
		sum(
			scale(sum(prod(m1[1], m3[0]), prod(-m3[1], m1[0])), m0[2]),
			sum(
				scale(sum(prod(m0[1], m3[0]), prod(-m3[1], m0[0])), -m1[2]),
				scale(sum(prod(m0[1], m1[0]), prod(-m1[1], m0[0])), m3[2]),
			),
		),
	)

	n := sum(
		sum(
			scale(sum(prod(m2[1], m3[0]), prod(-m3[1], m2[0])), m0[2]),
			sum(
				scale(sum(prod(m0[1], m3[0]), prod(-m3[1], m0[0])), -m2[2]),
				scale(sum(prod(m0[1], m2[0]), prod(-m2[1], m0[0])), m3[2]),
			),
		),
		sum(
			scale(sum(prod(m1[1], m2[0]), prod(-m2[1], m1[0])), m0[2]),
			sum(
				scale(sum(prod(m0[1], m2[0]), prod(-m2[1], m0[0])), -m1[2]),
				scale(sum(prod(m0[1], m1[0]), prod(-m1[1], m0[0])), m2[2]),
			),
		),
	)
	d := sub(p, n)
	return d[len(d)-1]

}
