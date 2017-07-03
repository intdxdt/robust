package robust

func Compress(e []float64) []float64{
	return compressExpansion(e)
}

func compressExpansion(e []float64) []float64 {
	var m = len(e)
	var Q = e[len(e)-1]
	var bottom = m
	for i := m - 2; i >= 0; i-- {
		var a = Q
		var b = e[i]
		Q = a + b
		var bv = Q - a
		var q = b - bv
		if q != 0 {
			bottom--
			e[bottom] = Q
			Q = q
		}
	}
	var top = 0
	for i := bottom; i < m; i++ {
		var a = e[i]
		var b = Q
		Q = a + b
		var bv = Q - a
		var q = b - bv
		if q != 0 {
			e[top] = q
			top++
		}
	}
	e[top] = Q
	top++
	return e[:top:top]
}
