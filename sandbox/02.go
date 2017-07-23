package main

import (
	"fmt"
	"robust"
	"time"
	"math/rand"
)

func main() {
	var seed = rand.NewSource(time.Now().UnixNano())
	var random = rand.New(seed)

	coords := [][]float64{{0, 0}, {1, 1}, {1, 0}, {0, 1}}
	res := robust.ConvexHull2D(coords)
	fmt.Println(res)

	for i := 0; i < 1000; i++ {
		coords = append(coords, af(random.Float64(), random.Float64()))
		coords = append(coords, af(0, random.Float64()))
		coords = append(coords, af(random.Float64(), 0))
		coords = append(coords, af(random.Float64(), 1))
		coords = append(coords, af(1, random.Float64()))
	}

	//g.Assert(ConvexHull2D(h)).Eql(ai())
	res = robust.ConvexHull2D(coords)
	fmt.Println(res)

	or := robust.Orientation2D([]float64{0, 0}, []float64{0.5, 0.5}, []float64{0.7, 0.1})
	fmt.Println(or)

}

func af(v ...float64) []float64 {
	return v
}
