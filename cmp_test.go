package robust

import (
	"testing"
	"github.com/franela/goblin"
)

func TestRobustCmp(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Cmp", func() {
		g.It("test robust cmp", func() {
			  g.Assert(Cmp(ar(5), ar(1, 4)) == 0).IsTrue()
			  g.Assert(Cmp(ar(1e64), ar(-1e-100, 1e64)) > 0).IsTrue()
			  g.Assert(Cmp(ar(1e64), ar(1e-100, 1e64)) < 0).IsTrue()
		})
	})
}
