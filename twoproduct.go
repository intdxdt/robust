package robust

const SPLITTER = 134217729.0//math.Pow(2, 27) + 1.0

//two product
func TwoProduct(a, b float64)[]float64{
    x := a * b

    c := SPLITTER * a
    abig := c - a
    ahi  := c - abig
    alo  := a - ahi

    d    := SPLITTER * b
    bbig := d - b
    bhi  := d - bbig
    blo  := b - bhi

    err1 := x - (ahi * bhi)
    err2 := err1 - (alo * bhi)
    err3 := err2 - (ahi * blo)

    y := alo * blo - err3

    return []float64{y, x}
}
