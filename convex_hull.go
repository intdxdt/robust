package robust

import "sort"

//Computes the convex hull of a set of 2D points.
//Input   : an sequence of (x, y) pairs representing the input points.
//Output  : a of vertices of the convex hull in counter-clockwise order,
//  starting from the vertex with the lexicographically smallest coordinates.
//Ref: Andrew's monotone chain algorithm. O(n log n) complexity.
func ConvexHull2D(coordinates [][]float64) [][]float64 {
	upper   := make([][]float64, 0)
	lower   := make([][]float64, 0)
	coords  := make([][]float64, len(coordinates))

	copy(coords, coordinates)

	n := len(coords)
	if n < 3 {
	    if n == 2 &&
		    coords[0][0] == coords[1][0] &&
		    coords[0][1] == coords[1][1] {
	      return [][]float64{coords[0]}
	    }
	    return coords
    }

	sort.Sort(coordSlice(coords))

	for _, pt := range coords {
		// should go clockwise
		// if counter or on line pop
		for len(upper) > 1 && Orientation2D(upper[len(upper)-2], upper[len(upper)-1], pt) <= 0 {
			pop(&upper)
		}

		// should go counter clock
		// if clockwise or on line pop
		for len(lower) > 1 && Orientation2D(lower[len(lower)-2], lower[len(lower)-1], pt) >= 0 {
			pop(&lower)
		}

		push(&upper, pt)
		push(&lower, pt)
	}

	// or upper = [o for o in upper]
	reverse(&upper)

	//end points are repeated top hull & down hull
	// return lower[:-1] + upper[:-1], lower, or upper
	if len(lower) > 0 {
		lower = lower[:len(lower)-1]
	}
	if len(upper) > 0 {
		upper = upper[:len(upper)-1]
	}

	return concat(lower, upper)
}

type coordSlice [][]float64
func (o coordSlice) Len() int { return len(o) }
func (o coordSlice) Swap(i, j int) { o[i], o[j] = o[j], o[i] }
//2d sort
func (o coordSlice) Less(i, j int) bool {
	a, b := o[i], o[j]
	d := a[0] - b[0]
	if d == 0 {
		d = a[1] - b[1]
	}
	return d < 0
}

func pop(ptr *[][]float64) {
	var a = *ptr
	n := len(a) - 1
	a[n] = nil
	*ptr = a[:n]
}

func push(ptr *[][]float64, v []float64) {
	var a = *ptr
	a = append(a, v)
	*ptr = a
}

func reverse(ptr *[][]float64) {
	coords := *ptr
	for i, j := 0, len(coords)-1; i < j; i, j = i+1, j-1 {
		coords[i], coords[j] = coords[j], coords[i]
	}
	*ptr = coords
}

func concat(a, b [][]float64) [][]float64 {
	n := len(a) + len(b)
	res := make([][]float64, n, n)
	copy(res, a)
	copy(res[len(a):], b)
	return res
}
