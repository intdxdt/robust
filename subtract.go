package robust

import "math"

//Robust subtraction of two floats
func Subtract(e, f []float64) []float64 {
	return linearExpansionSub(e, f)
}

//linear expansion
func linearExpansionSub(e, f []float64) []float64 {
	var ne = len(e)
	var nf = len(f)
	if ne == 1 && nf == 1 {
		return scalar_scalar(e[0], -f[0])
	}
	var n = ne + nf
	var g = make([]float64, n)
	var count = 0
	var eptr  = 0
	var fptr  = 0
	var abs   = math.Abs
	var ei = e[eptr]
	var ea = abs(ei)
	var fi = -f[fptr]
	var fa = abs(fi)
	var a, b float64
	if ea < fa {
		b = ei
		eptr += 1
		if eptr < ne {
			ei = e[eptr]
			ea = abs(ei)
		}
	} else {
		b = fi
		fptr += 1
		if fptr < nf {
			fi = -f[fptr]
			fa = abs(fi)
		}
	}
	if (eptr < ne && ea < fa) || (fptr >= nf) {
		a = ei
		eptr += 1
		if eptr < ne {
			ei = e[eptr]
			ea = abs(ei)
		}
	} else {
		a = fi
		fptr += 1
		if fptr < nf {
			fi = -f[fptr]
			fa = abs(fi)
		}
	}
	var x = a + b
	var bv = x - a
	var y = b - bv
	var q0 = y
	var q1 = x
	var _x, _bv, _av, _br, _ar float64
	for eptr < ne && fptr < nf {
		if ea < fa {
			a = ei
			eptr += 1
			if eptr < ne {
				ei = e[eptr]
				ea = abs(ei)
			}
		} else {
			a = fi
			fptr += 1
			if fptr < nf {
				fi = -f[fptr]
				fa = abs(fi)
			}
		}
		b = q0
		x = a + b
		bv = x - a
		y = b - bv
		if y != 0 {
			g[count] = y
			count += 1
		}
		_x = q1 + x
		_bv = _x - q1
		_av = _x - _bv
		_br = x - _bv
		_ar = q1 - _av
		q0 = _ar + _br
		q1 = _x
	}
	for eptr < ne {
		a = ei
		b = q0
		x = a + b
		bv = x - a
		y = b - bv
		if y != 0 {
			g[count] = y
			count += 1
		}
		_x = q1 + x
		_bv = _x - q1
		_av = _x - _bv
		_br = x - _bv
		_ar = q1 - _av
		q0 = _ar + _br
		q1 = _x
		eptr += 1
		if eptr < ne {
			ei = e[eptr]
		}
	}
	for fptr < nf {
		a = fi
		b = q0
		x = a + b
		bv = x - a
		y = b - bv
		if y != 0 {
			g[count] = y
			count += 1
		}
		_x = q1 + x
		_bv = _x - q1
		_av = _x - _bv
		_br = x - _bv
		_ar = q1 - _av
		q0 = _ar + _br
		q1 = _x
		fptr += 1
		if fptr < nf {
			fi = -f[fptr]
		}
	}
	if q0 != 0 {
		g[count] = q0
		count += 1
	}
	if q1 != 0 {
		g[count] = q1
		count += 1
	}
	if count == 0 {
		g[count] = 0.0
		count += 1
	}

	return g[:count:count]
}

//scalar sum: easy case: add two scalars
func scalar_scalar(a, b float64) []float64 {
	var x  = a + b
	var bv = x - a
	var av = x - bv
	var br = b - bv
	var ar = a - av
	var y  = ar + br
	if y != 0 {
		return []float64{y, x}
	}
	return []float64{x}
}
