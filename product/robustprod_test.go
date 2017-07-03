package product

import (
	"math"
	"testing"
	"github.com/franela/goblin"
)

func TestTwoProduct(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("RobustProduct", func() {
		g.It("test robust product", func() {
			for i := -20; i <= 20; i++ {
				for j := -20; j <= 20; j++ {
					fi, fj := float64(i), float64(j)
					g.Assert(RobustProduct(ar(fi), ar(fj))).Eql(ar(fi * fj))
				}
			}

			g.Assert(RobustProduct(
				ar(pow2(-50), pow2(50)),
				ar(pow2(-50), pow2(50)),
			)).Eql(
				ar(pow2(-100), pow2(1), pow2(100)),
			)
		})
	})
}

func pow2(n float64) float64 {
	return math.Pow(2, n)
}

func ar(v ...float64) []float64 {
	return v
}
