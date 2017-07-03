package robust

import (
	"math"
	"testing"
	"github.com/franela/goblin"
)


func TestRobustSum(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Sum", func() {
		g.It("test robust sum", func() {
			g.Assert(Sum(
				ar(1, 64), ar(-1e-64, 1e64),
			)).Eql(ar(-1e-64, 65, 1e64), )

			g.Assert(Sum(ar(0), ar(0))).Eql(ar(0))
			g.Assert(Sum(ar(0), ar(1))).Eql(ar(1))

			g.Assert(Sum(
				ar(1, 1e64), ar(1e-64, 2),
			)).Eql(ar(1e-64, 3, 1e64))

			g.Assert(Sum(
				ar(1), ar(1e-64, 1e-16),
			)).Eql(ar(1e-64, 1e-16, 1))

			g.Assert(Sum(ar(0), ar(1))).Eql(ar(1))

			for i := -10; i <= 10; i++ {
				for j := -10; j <= 10; j++ {
					g.Assert(Sum(ar(float64(i)), ar(float64(j)))).Eql(ar(float64(i + j)))
				}
			}

			// t.ok(validate(sum([ 5.711861227349496e-133, 1e-116 ], [ 5.711861227349496e-133, 1e-116 ])))

			nois := make([]float64, 10)
			expect := make([]float64, 10)
			for i := 0; i < 10; i++ {
				nois[i] = math.Pow(2, float64(-1000+53*i))
				expect[i] = math.Pow(2, float64(-999+53*i))
			}
			x := Sum(nois, nois)
			g.Assert(x).Eql(expect)
			// t.ok(validate(x))

			g.Assert(Sum(ar(0), ar(1, 1e64))).Eql(ar(1, 1e64))

			// var s = [0]
			// for(var i=0; i<1000; ++i) {
			// s = sum(s, [Math.random() * Math.pow(2, Math.random()*1800-900)])
			// t.ok(validate(s))
			// }

		})
	})
}
