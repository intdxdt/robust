package robust

import (
	"sort"
)

//Convex Monotone Hull 2D
func ConvexHull2D(points [][]float64) []int {
	var n = len(points)

	if n < 3 {
		var result = make([]int, n)

		for i := 0; i < n; i++ {
			result[i] = i
		}

		if n == 2 && points[0][0] == points[1][0] && points[0][1] == points[1][1] {
			return []int{0}
		}

		return result
	}

	//Sort point indices along x-axis
	var sorted = make([]int, n)
	for i := 0; i < n; i++ {
		sorted[i] = i
	}

	cvx := convexSort{points: points, indxs: sorted}
	sort.Sort(cvx)

	//Construct upper and lower hulls
	var lower = []int{sorted[0], sorted[1]}
	var upper = []int{sorted[0], sorted[1]}

	for i := 2; i < n; i++ {
		var idx = sorted[i]
		var p = points[idx]

		//Insert into lower list
		var m = len(lower)
		for m > 1 && Orientation2d(points[lower[m-2]], points[lower[m-1]], p) <= 0 {
			m -= 1
			pop(&lower)
		}
		push(&lower, idx)

		//Insert into upper list
		m = len(upper)
		for m > 1 && Orientation2d(points[upper[m-2]], points[upper[m-1]], p) >= 0 {
			m -= 1
			pop(&upper)
		}
		push(&upper, idx)
	}

	//Merge lists together
	var result = make([]int, len(upper)+len(lower)-2)
	var ptr = 0
	for i, nl := 0, len(lower); i < nl; i++ {
		result[ptr] = lower[i]
		ptr++
	}

	for j := len(upper) - 2; j > 0; j-- {
		result[ptr] = upper[j]
		ptr++
	}

	//Return result
	return result
}

type convexSort struct {
	points [][]float64
	indxs  []int
}

func (o convexSort) Len() int {
	return len(o.points)
}

func (o convexSort) Swap(i, j int) {
	o.indxs[i], o.indxs[j] = o.indxs[j], o.indxs[i]
}

func (o convexSort) Less(i, j int) bool {
	a, b := o.points[i], o.points[j]
	if a[0] < b[0] {
		return true
	}
	return a[1] < b[1]
}

func pop(ptr *[]int) {
	var a = *ptr
	n := len(a) - 1
	a[n] = 0
	*ptr = a[:n]
}

func push(ptr *[]int, v int) {
	var a = *ptr
	a = append(a, v)
	*ptr = a
}
