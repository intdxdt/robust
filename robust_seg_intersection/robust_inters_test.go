package robust_seg_intersection

import (
	"time"
	"testing"
	"math/rand"
	"robust/robust_det"
	"robust/robust_sum"
	"robust/robust_cmp"
	"robust/robust_diff"
	"robust/robust_product"
	"github.com/franela/goblin"
)



func TestRobustSegSeg(t *testing.T) {
	var det = robust_det.RobustDet2
	var robustSum = robust_sum.RobustSum
	var robustDiff = robust_diff.RobustDiff
	var robustCompare = robust_cmp.RobustCmp
	var robustProduct = robust_product.RobustProduct

	g := goblin.Goblin(t)

	g.Describe("Robust Seg Intersection", func() {
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

				var d0 = robustSum(ar(a[1]), ar(-b[1]))
				var d1 = robustSum(ar(a[0]), ar(-b[0]))
				var d2 = det([][]float64{a, b})

				//validate det
				//g.Assert(validate(d2)).IsTrue()

				var p0 = robustProduct(x, d0)
				var p1 = robustProduct(y, d1)
				var p2 = robustProduct(w, d2)
				//validate p0
				//t.ok(validate(p0))
				//validate p1
				//t.ok(validate(p1))
				//validate p2
				//t.ok(validate(p2))

				var s = robustSum(robustDiff(p0, p1), p2)
				//validate s
				//t.ok(validate(s))
				//check point on line
				g.Assert(robustCompare(s, []float64{0}) == 0).IsTrue()
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
							g.Assert(robustCompare(robustProduct(y[0], x[2]), robustProduct(x[0], y[2])) == 0).IsTrue()
							//check y
							g.Assert(robustCompare(robustProduct(y[1], x[2]), robustProduct(x[1], y[2])) == 0).IsTrue()
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

func ar(v ...float64) []float64 {
	return v
}
