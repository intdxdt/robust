package two_sum

import (
	"github.com/franela/goblin"
	"testing"
)

func TestTwoSum(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("TwoSum", func() {
		g.It("test fast two sum", func() {
			g.Assert(TwoSum(1e+64, 1)).Eql([]float64{1.0, 1e+64})
			g.Assert(TwoSum(1, 1)).Eql([]float64{0, 2})
			g.Assert(TwoSum(0, -1415)).Eql([]float64{0, -1415})
			g.Assert(TwoSum(1e-64, 1e64)).Eql([]float64{1e-64, 1e64})
			g.Assert(TwoSum(0, 0)).Eql([]float64{0, 0})
			g.Assert(TwoSum(9e15+1, 9e15)).Eql([]float64{1, 18000000000000000})
		})
	})
}
