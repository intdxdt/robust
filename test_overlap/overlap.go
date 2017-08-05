package test_overlap

import (
	"robust/twiddle"
	"robust/dbits"
	"math"
)

//use bits{count_trailing_zeros, log2};
//use db{fraction, denormalized, exponent};
func tz(f []uint32) int32 {
	if f[0] != 0 {
		return int32(twiddle.CountTrailingZeros(f[0]))
	} else if f[1] != 0 {
		return 32 + int32(twiddle.CountTrailingZeros(f[1]))
	}
	return 0
}

func lz(f []uint32) int32 {
	if f[1] != 0 {
		return 20 - int32(twiddle.Log2(f[1]))
	} else if f[0] != 0 {
		return 52 - int32(twiddle.Log2(f[0]))
	}
	return 52

}

func lo(n float64) int32 {
	var e = dbits.Exponent(n)
	var f = dbits.Fraction(n)
	var z = tz(f[:])
	return e - (52 - z)
}

func hi(n float64) int32 {
	if dbits.Denormalized(n) {
		v := dbits.Fraction(n)
		return -(1023 + lz(v[:]))
	}
	return dbits.Exponent(n)

}

func TestOverlap(a, b float64) bool {
	if math.Abs(b) > math.Abs(a) {
		a, b = b, a
	}
	if a == 0.0 || b == 0.0 {
		return false
	}
	var a0 = hi(a)
	var a1 = lo(a)
	var b0 = hi(b)
	var b1 = lo(b)
	//[a1------a0]
	//     [b1-----b0]
	//---------or----------
	//    [a1-------a0]
	//[b1-------b0]
	return (b1 <= a0) && (a1 <= b0)
}
