package robust

import (
	"testing"
	"github.com/franela/goblin"
)

func TestRobustDet(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Det", func() {
		g.It("test robust determinant", func() {
			g.Assert(Det3([][]float64{
				{1, 2, 3}, {3, 4, 5}, {6, 7, 8},
			})).Eql(ar(0))

			g.Assert(Det2([][]float64{
				{1, 2}, {3, 4},
			})).Eql(ar(-2))

			g.Assert(Det3([][]float64{
				{1, 2, 3}, {3, 4, 5}, {6, 7, 8},
			})).Eql(ar(0))
			g.Assert(Det2([][]float64{
				{1, 2}, {3, 4},
			})).Eql(ar(-2))
		})
	})
}
