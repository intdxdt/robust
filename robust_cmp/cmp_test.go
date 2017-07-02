package robust_cmp

import (
	"testing"
	"github.com/franela/goblin"
)

func TestRobustCmp(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("RobustCmp", func() {
		g.It("test robust cmp", func() {
			  g.Assert(RobustCmp(ar(5), ar(1, 4)) == 0).IsTrue()
			  g.Assert(RobustCmp(ar(1e64), ar(-1e-100, 1e64)) > 0).IsTrue()
			  g.Assert(RobustCmp(ar(1e64), ar(1e-100, 1e64)) < 0).IsTrue()
		})
	})
}

func ar(v ...float64) []float64 {
	return v
}
