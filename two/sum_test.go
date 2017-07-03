package two

import (
	"testing"
	"github.com/franela/goblin"
)

func TestTwoSum(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Sum", func() {
		g.It("test fast two sum", func() {
			g.Assert(Sum(1e+64, 1)).Eql([]float64{1.0, 1e+64})
			g.Assert(Sum(1, 1)).Eql([]float64{0, 2})
			g.Assert(Sum(0, -1415)).Eql([]float64{0, -1415})
			g.Assert(Sum(1e-64, 1e64)).Eql([]float64{1e-64, 1e64})
			g.Assert(Sum(0, 0)).Eql([]float64{0, 0})
			g.Assert(Sum(9e15+1, 9e15)).Eql([]float64{1, 18000000000000000})
		})
	})
}
