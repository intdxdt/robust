package det

import (
	"testing"
	"github.com/franela/goblin"
)

func init() {
}

func TestRobustDet(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("RobustDet", func() {
		g.It("test robust determinant", func() {
			g.Assert(RobustDet3([][]float64{
				{1, 2, 3}, {3, 4, 5}, {6, 7, 8},
			})).Eql(ar(0))

			g.Assert(RobustDet2([][]float64{
				{1, 2}, {3, 4},
			})).Eql(ar(-2))

			g.Assert(RobustDet3([][]float64{
				{1, 2, 3}, {3, 4, 5}, {6, 7, 8},
			})).Eql(ar(0))
			g.Assert(RobustDet2([][]float64{
				{1, 2}, {3, 4},
			})).Eql(ar(-2))
		})
	})
}

func ar(v ...float64) []float64 {
	return v
}
