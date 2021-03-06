package robust

import "math"

//Checks if two Segments Intersect
func SegIntersects(a0, a1, b0, b1 []float64) bool {

	var x0 = Orientation2D(a0, b0, b1)
	var y0 = Orientation2D(a1, b0, b1)
	if (x0 > 0 && y0 > 0) || (x0 < 0 && y0 < 0) {
		return false
	}

	var x1 = Orientation2D(b0, a0, a1)
	var y1 = Orientation2D(b1, a0, a1)
	if (x1 > 0 && y1 > 0) || (x1 < 0 && y1 < 0) {
		return false
	}

	//Check for degenerate collinear case
	if x0 == 0 && y0 == 0 && x1 == 0 && y1 == 0 {
		return checkCollinear(a0, a1, b0, b1)
	}
	return true
}

func checkCollinear(a0, a1, b0, b1 []float64) bool {
	for d := 0; d < 2; d++ {
		var x0 = a0[d]
		var y0 = a1[d]
		var l0 = math.Min(x0, y0)
		var h0 = math.Max(x0, y0)

		var x1 = b0[d]
		var y1 = b1[d]
		var l1 = math.Min(x1, y1)
		var h1 = math.Max(x1, y1)

		if h1 < l0 || h0 < l1 {
			return false
		}
	}
	return true
}
