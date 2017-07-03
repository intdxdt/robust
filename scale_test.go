package robust

import (
	"testing"
	"github.com/franela/goblin"
)


func TestRobustScale(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Sclae", func() {
		g.It("test robust scale", func() {
			g.Assert(Scale(ar(4), 2)).Eql(ar(8))
			g.Assert(Scale(ar(1, 1e64), 2)).Eql(ar(2, 2e64))
			g.Assert(Scale(ar(1), 1)).Eql(ar(1))
			s := Scale(ar(-2.4707339790384e-144, -1.6401064715739963e-142, 2e-126), -10e-64)
			g.Assert(s[len(s)-1] < 0)

			for i := -50; i < 50; i++ {
				for j := -50; j < 50; j++ {
					g.Assert(Scale(ar(float64(i)), float64(j))).Eql(ar(float64(i * j)))
				}
			}
		})
	})
}
