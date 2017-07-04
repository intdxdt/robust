package robust

import (
	"testing"
	"github.com/franela/goblin"
)

func TestRobustCmp(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Cmp", func() {
		g.It("test robust cmp", func() {
			  g.Assert(Cmp(af(5), af(1, 4)) == 0).IsTrue()
			  g.Assert(Cmp(af(1e64), af(-1e-100, 1e64)) > 0).IsTrue()
			  g.Assert(Cmp(af(1e64), af(1e-100, 1e64)) < 0).IsTrue()
		})
	})
}
