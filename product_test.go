package robust

import (
	"math"
	"testing"
	"github.com/franela/goblin"
)

func TestRobustProduct(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Product", func() {
		var pow2 = func (n float64) float64 { return math.Pow(2, n) }

		g.It("test robust product", func() {
			for i := -20; i <= 20; i++ {
				for j := -20; j <= 20; j++ {
					fi, fj := float64(i), float64(j)
					g.Assert(Product(af(fi), af(fj))).Eql(af(fi * fj))
				}
			}

			g.Assert(Product(
				af(pow2(-50), pow2(50)),
				af(pow2(-50), pow2(50)),
			)).Eql(
				af(pow2(-100), pow2(1), pow2(100)),
			)
		})
	})
}

