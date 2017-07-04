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
			var coords = [][]float64{{0, 0}, {1, 1}, {1, 0}, {0.5, 0.5}, {0.7, 0.1}}
			g.Assert(ConvexHull2D(coords)).Eql([][]float64{ {0, 0},  {1, 0}, {1, 1}})

			coords = [][]float64{{0, 0}, {1, 0}, {1, 1}, {0, 1}}
			g.Assert(ConvexHull2D(coords)).Eql([][]float64{{0, 0}, {1, 0}, {1, 1}, {0, 1}})

			coords = [][]float64{{0, 0}, {1, 1}, {1, 0}, {0, 1}}
			g.Assert(ConvexHull2D(coords)).Eql([][]float64{{0, 0}, {1, 0}, {1, 1}, {0, 1}})

			for i := 0; i < 1000; i++ {
				coords = append(coords, af(random.Float64(), random.Float64()))
				coords = append(coords, af(0, random.Float64()))
				coords = append(coords, af(random.Float64(), 0))
				coords = append(coords, af(random.Float64(), 1))
				coords = append(coords, af(1, random.Float64()))
			}

			g.Assert(ConvexHull2D(coords)).Eql([][]float64{ {0, 0},  {1, 0}, {1, 1}, {0, 1}})
			//Degenerate cases
			g.Assert(ConvexHull2D([][]float64{{0, 0}})).Eql([][]float64{{0, 0}})
			g.Assert(ConvexHull2D([][]float64{})).Eql([][]float64{})
			g.Assert(ConvexHull2D([][]float64{{0, 0}, {1, 1}})).Eql([][]float64{{0, 0}, {1, 1}})
			g.Assert(ConvexHull2D([][]float64{{0, 0}, {0, 0}})).Eql([][]float64{{0, 0}})
		})
	})
}
