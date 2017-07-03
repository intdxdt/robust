package robust

import (
	"math"
	"testing"
	"github.com/franela/goblin"
)

func TestRobustProduct(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Product", func() {
		g.It("test robust product", func() {
			for i := -20; i <= 20; i++ {
				for j := -20; j <= 20; j++ {
					fi, fj := float64(i), float64(j)
					g.Assert(Product(ar(fi), ar(fj))).Eql(ar(fi * fj))
				}
			}

			g.Assert(Product(
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
