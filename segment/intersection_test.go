package segment

import (
	"time"
	"testing"
	"math/rand"
	"robust/det"
	"robust/sum"
	"robust/cmp"
	"robust/diff"
	"robust/product"
	"github.com/franela/goblin"
)


func TestRobustSegSeg(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("Robust Segment Intersection", func() {
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

				var d0 = sum.RobustSum(ar(a[1]), ar(-b[1]))
				var d1 = sum.RobustSum(ar(a[0]), ar(-b[0]))
				var d2 = det.RobustDet2([][]float64{a, b})

				//validate det.RobustDet2
				//g.Assert(validate(d2)).IsTrue()

				var p0 = product.RobustProduct(x, d0)
				var p1 = product.RobustProduct(y, d1)
				var p2 = product.RobustProduct(w, d2)
				//validate p0
				//t.ok(validate(p0))
				//validate p1
				//t.ok(validate(p1))
				//validate p2
				//t.ok(validate(p2))

				var s = sum.RobustSum(diff.RobustDiff(p0, p1), p2)
				//validate s
				//t.ok(validate(s))
				//check point on line
				g.Assert(cmp.RobustCmp(s, []float64{0}) == 0).IsTrue()
			}

			verify := func(a, b, c, d []float64) {
				var x = RobustIntersection(a, b, c, d)
				//validate x
				//t.ok(validate(x[0]))
				//validate y
				//t.ok(validate(x[1]))
				//validate w
				//t.ok(validate(x[2]))
				testPoint(a, b, x[0], x[1], x[2])
				testPoint(c, d, x[0], x[1], x[2])

				var p = [][][]float64{{a, b}, {c, d}}
				for s := 0; s < 2; s++ {
					for r := 0; r < 2; r++ {
						for h := 0; h < 2; h++ {
							var y = RobustIntersection(
								p[h][s], p[h][s^1],
								p[h^1][r], p[h^1][r^1],
							)
							//validate x
							//t.ok(validate(y[0]))
							//validate y
							//t.ok(validate(y[1]))
							//validate w
							//t.ok(validate(y[2]))
							//check x
							g.Assert(cmp.RobustCmp(product.RobustProduct(y[0], x[2]), product.RobustProduct(x[0], y[2])) == 0).IsTrue()
							//check y
							g.Assert(cmp.RobustCmp(product.RobustProduct(y[1], x[2]), product.RobustProduct(x[1], y[2])) == 0).IsTrue()
						}
					}
				}
			}
			//Fuzz test
			for i := 0; i < 100; i++ {
				verify(
					ar(random.Float64(), random.Float64()),
					ar(random.Float64(), random.Float64()),
					ar(random.Float64(), random.Float64()),
					ar(random.Float64(), random.Float64()),
				)
			}


			var isect = RobustIntersection(ar(-1, 10), ar(-10, 1), ar(10, 0), ar(10, 10));
			//no intersections
            g.Assert(isect[2][0]== 0).IsTrue()
		})
	})
}
