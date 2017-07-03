package robust


func Scale(e []float64, scale float64) []float64{
	return linearExpansionScale(e, scale)
}

func linearExpansionScale(e []float64, scale float64) []float64{
	var n = len(e)
	if n == 1 {
		var ts = TwoProduct(e[0], scale)
		if ts[0] != 0 {
			return ts
		}
		return []float64{ts[1]}
	}
	var g     = make([]float64, 2 * n)
	var q     = []float64{0.1, 0.1}
	var t     = []float64{0.1, 0.1}
	var count = 0
	q = TwoProduct(e[0], scale)
	if q[0] != 0 {
		g[count] = q[0]
		count++
	}
	for i := 1; i < n; i++ {
		t = TwoProduct(e[i], scale)
		var pq = q[1]
		q = TwoSum(pq, t[0])
		if q[0] != 0 {
			g[count] = q[0]
			count++
		}
		var a  = t[1]
		var b  = q[1]
		var x  = a + b
		var bv = x - a
		var y  = b - bv
		q[1]   = x
		if y !=0 {
			g[count] = y
			count++
		}
	}
	if q[1] != 0 {
		g[count] = q[1]
		count++
	}
	if count == 0 {
		g[count] = 0.0
		count++
	}
	return g[:count:count]
}