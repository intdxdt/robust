package robust

import (
	"time"
	"testing"
	"math/rand"
	"github.com/franela/goblin"
	"robust/validate_seq"
)


func TestRobustSegSeg(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("SegIntersection", func() {
		g.It("test robust seg seg intersection", func() {
			var seed = rand.NewSource(time.Now().UnixNano())
			var random = rand.New(seed)

			//Evaluate:
			//
			//  | a[0]  a[1]  1 |
			//  | b[0]  b[1]  1 |
			//  |  x     y    w |
			//
			testPoint := func(a, b, x, y, w []float64) {

				var d0 = Sum(af(a[1]), af(-b[1]))
				var d1 = Sum(af(a[0]), af(-b[0]))
				var d2 = Det2([][]float64{a, b})

				//validate det.RobustDet2
				g.Assert(validate_seq.ValidateSequence(d2)).IsTrue()

				var p0 = Product(x, d0)
				var p1 = Product(y, d1)
				var p2 = Product(w, d2)
				//validate p0
				g.Assert(validate_seq.ValidateSequence(p0)).IsTrue()
				//validate p1
				g.Assert(validate_seq.ValidateSequence(p1)).IsTrue()
				//validate p2
				g.Assert(validate_seq.ValidateSequence(p2)).IsTrue()

				var s = Sum(Subtract(p0, p1), p2)
				//validate s
				g.Assert(validate_seq.ValidateSequence(s)).IsTrue()
				//check point on line
				g.Assert(Cmp(s, []float64{0}) == 0).IsTrue()
			}

			verify := func(a, b, c, d []float64) {
				var x = SegIntersection(a, b, c, d)
				//validate x
				g.Assert(validate_seq.ValidateSequence(x[0])).IsTrue()
				//validate y
				g.Assert(validate_seq.ValidateSequence(x[1])).IsTrue()
				//validate w
				g.Assert(validate_seq.ValidateSequence(x[2])).IsTrue()
				testPoint(a, b, x[0], x[1], x[2])
				testPoint(c, d, x[0], x[1], x[2])

				var p = [][][]float64{{a, b}, {c, d}}
				for s := 0; s < 2; s++ {
					for r := 0; r < 2; r++ {
						for h := 0; h < 2; h++ {
							var y = SegIntersection(
								p[h][s], p[h][s^1],
								p[h^1][r], p[h^1][r^1],
							)
							//validate x
							g.Assert(validate_seq.ValidateSequence(y[0])).IsTrue()
							//validate y
							g.Assert(validate_seq.ValidateSequence(y[1])).IsTrue()
							//validate w
							g.Assert(validate_seq.ValidateSequence(y[2])).IsTrue()
							//check x
							g.Assert(Cmp(Product(y[0], x[2]), Product(x[0], y[2])) == 0).IsTrue()
							//check y
							g.Assert(Cmp(Product(y[1], x[2]), Product(x[1], y[2])) == 0).IsTrue()
						}
					}
				}
			}
			//Fuzz test
			for i := 0; i < 100; i++ {
				verify(
					af(random.Float64(), random.Float64()),
					af(random.Float64(), random.Float64()),
					af(random.Float64(), random.Float64()),
					af(random.Float64(), random.Float64()),
				)
			}


			var isect = SegIntersection(af(-1, 10), af(-10, 1), af(10, 0), af(10, 10))
			//no intersections
            g.Assert(isect[2][0]== 0).IsTrue()
		})
	})
}
