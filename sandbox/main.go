package main

import (
    "robust"
    "encoding/json"
    "math"
    "fmt"
    "os"
)

var NX = 512
var NY = 512

func naiveLeftRight(a, b, c []float64) float64{
    var abx = c[0] - a[0]
    var aby = c[1] - a[1]
    var acx = b[0] - a[0]
    var acy = b[1] - a[1]
    return abx*acy - aby*acx
}

func plotPredicate(pred func ([]float64, []float64, []float64 ) float64, out string) {
    res := make([]float64, 0)
    for i := 0; i < NX; i++ {
        for j := 0; j < NY; j++ {
            px := 0.5 + float64(i)*math.Pow(2, -53)
            py := 0.5 + float64(j)*math.Pow(2, -53)

            o := pred(ar(px, py), ar(12, 12), ar(24, 24))
            res = append(res, o)
        }
    }
    fmt.Println("len or arr = ", len(res))
    fid, _ := os.Create(out)
    vals , _ := json.Marshal(res)
    fid.WriteString("module.exports=" + string(vals) + "\n")
}

func main(){
    plotPredicate(naiveLeftRight, "go-naive.js")
    plotPredicate(robust.Orientation2D, "go-robust.js")
}



func ar(v ...float64) []float64 {
	return v
}
