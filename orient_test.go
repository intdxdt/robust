package robust

import (
	"testing"
	"github.com/franela/goblin"
)

func TestRobustOrient(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Orientation", func() {
		g.It("test robust orient 2D", func() {
			g.Assert(Orientation2d(ar(0.1, 0.1), ar(0.1, 0.1), ar(0.3, 0.7)) == 0).IsTrue()
			g.Assert(Orientation2d(ar(0, 0), ar(-1e-64, 0), ar(0, 1)) > 0).IsTrue()

			g.Assert(Orientation2d(ar(0, 0), ar(1e-64, 1e-64), ar(1, 1)) == 0).IsTrue()
			g.Assert(Orientation2d(ar(0, 0), ar(1e-64, 0), ar(0, 1)) < 0).IsTrue()

			x := 1e-64
			for i := 0; i < 200; i++ {
				g.Assert(Orientation2d(ar(-x, 0), ar(0, 1), ar(x, 0)) > 0).IsTrue()
				g.Assert(Orientation2d(ar(-x, 0), ar(0, 0), ar(x, 0)) == 0).IsTrue()
				g.Assert(Orientation2d(ar(-x, 0), ar(0, -1), ar(x, 0)) < 0).IsTrue()
				g.Assert(Orientation2d(ar(0, 1), ar(0, 0), ar(x, x)) < 0).IsTrue()
				x *= 10
			}
		})
		g.It("test robust orient 3D", func() {
			g.Assert(Orientation3d(ar(0, 0, 0), ar(1, 0, 0), ar(0, 1, 0), ar(0, 0, 1)) < 0).IsTrue()
			g.Assert(Orientation3d(ar(0, 0, 0), ar(1, 0, 0), ar(0, 0, 1), ar(0, 1, 0)) > 0).IsTrue()
			g.Assert(Orientation3d(ar(0, 0, 0), ar(1e-64, 0, 0), ar(0, 0, 1), ar(0, 1e64, 0)) > 0).IsTrue()
		})
	})
}
