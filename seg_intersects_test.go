package robust

import (
	"github.com/franela/goblin"
	"testing"
)

func TestRobustIntersects(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("SegIntersects", func() {
		g.It("test robust seg seg intersects", func() {
			g.Assert(SegIntersects(af(-1, 0), af(1, 0), af(0, -1), af(0, 1))).IsTrue()
			g.Assert(SegIntersects(af(0.5, 0), af(1, 0), af(0, -1), af(0, 1))).IsFalse()
			g.Assert(SegIntersects(af(0, 0), af(1, 0), af(0, -1), af(0, 1))).IsTrue()
			g.Assert(SegIntersects(af(0, 0), af(100000000000000020000, 1e-12),
				af(1, 0), af(1e20, 1e-11))).IsTrue()
			g.Assert(SegIntersects(af(0, 0), af(1e20, 1e-11),
				af(1, 0), af(100000000000000020000, 1e-12))).IsFalse()

			//collinear, no intersect
			g.Assert(SegIntersects(af(0, 1), af(0, 2), af(0, -1), af(0, -2))).IsFalse()
			//collinear, intersect
			g.Assert(SegIntersects(af(0, 1), af(0, 2), af(0, 1.5), af(0, -2))).IsTrue()
			//collinear, endpoint touch
			g.Assert(SegIntersects(af(0, 1), af(0, 2), af(0, 1), af(0, -2))).IsTrue()
			//endpoint touches
			g.Assert(SegIntersects(af(0, 1), af(0, -1), af(0, 0), af(0, 1))).IsTrue()

		})
	})
}
