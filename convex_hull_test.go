package robust

import (
	"testing"
	"github.com/franela/goblin"
	"time"
	"math/rand"
)

func TestConvexHull(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("ConvexHull2D", func() {
		g.It("test convex hull 2d", func() {
			var seed = rand.NewSource(time.Now().UnixNano())
			var random = rand.New(seed)

			g.Assert(TwoSum(1e+64, 1)).Eql([]float64{1.0, 1e+64})
			var res = ConvexHull2D([][]float64{{0, 0}, {1, 1}, {1, 0}, {0.5, 0.5}, {0.7, 0.1}})
			g.Assert(res).Eql(ai(0, 1, 2))

			var h [][]float64
			h = [][]float64{{0, 0}, {1, 0}, {1, 1}, {0, 1}}
			g.Assert(ConvexHull2D(h)).Eql(ai(0, 3, 2, 1))

			h = [][]float64{{0, 0}, {1, 1}, {1, 0}, {0, 1}}
			g.Assert(ConvexHull2D(h)).Eql(ai(0, 3, 1, 2))

			for i := 0; i < 1000; i++ {
				h = append(h, af(random.Float64(), random.Float64()))
				h = append(h, af(0, random.Float64()))
				h = append(h, af(random.Float64(), 0))
				h = append(h, af(random.Float64(), 1))
				h = append(h, af(1, random.Float64()))
			}

			//g.Assert(ConvexHull2D(h)).Eql(ai(0, 3, 1, 2))
			  //Degenerate cases
			  g.Assert(ConvexHull2D([][]float64{{0,0}})).Eql(ai(0))
			  g.Assert(ConvexHull2D([][]float64{})).Eql([]int{})
			  g.Assert(ConvexHull2D([][]float64{{0,0}, {1,1}})).Eql(ai(0,1))
			  g.Assert(ConvexHull2D([][]float64{{0,0}, {0,0}})).Eql(ai(0))
		})
	})
}
