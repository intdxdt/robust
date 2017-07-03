package robust

import (
	"github.com/franela/goblin"
	"testing"
)


func TestRobustIntersects(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("SegIntersects", func() {
		g.It("test robust seg seg intersects", func() {
			g.Assert(SegIntersects(ar(-1, 0), ar(1, 0), ar(0, -1), ar(0, 1))).IsTrue()
			g.Assert(SegIntersects(ar(0.5, 0), ar(1, 0), ar(0, -1), ar(0, 1))).IsFalse()
			g.Assert(SegIntersects(ar(0, 0), ar(1, 0), ar(0, -1), ar(0, 1))).IsTrue()
			g.Assert(SegIntersects(ar(0, 0), ar(100000000000000020000, 1e-12),
				ar(1, 0), ar(1e20, 1e-11))).IsTrue()
			g.Assert(SegIntersects(ar(0, 0), ar(1e20, 1e-11),
				ar(1, 0), ar(100000000000000020000, 1e-12))).IsFalse()

			//collinear, no intersect
			g.Assert(SegIntersects(ar(0, 1), ar(0, 2), ar(0, -1), ar(0, -2))).IsFalse()
			//collinear, intersect
			g.Assert(SegIntersects(ar(0, 1), ar(0, 2), ar(0, 1.5), ar(0, -2))).IsTrue()
			//collinear, endpoint touch
			g.Assert(SegIntersects(ar(0, 1), ar(0, 2), ar(0, 1), ar(0, -2))).IsTrue()
			//endpoint touches
			g.Assert(SegIntersects(ar(0, 1), ar(0, -1), ar(0, 0), ar(0, 1))).IsTrue()

		})
	})
}

