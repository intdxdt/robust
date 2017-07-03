package cmp

import (
	"robust/diff"
)

//Robust Compare
func RobustCmp(a, b []float64) float64 {
	d := diff.RobustDiff(a, b)
	return d[len(d)-1]
}
