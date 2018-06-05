package test_overlap

import (
	"math"
	"testing"
	"github.com/franela/goblin"
)

func TestFloatOverlap(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Test Float Overlap", func() {
		g.It("test float overlap", func() {
			g.Assert(TestOverlap(1.5, 0.5))
			g.Assert(TestOverlap(math.Pow(2, -52), 1.0+math.Pow(2, -52)))
			g.Assert(!TestOverlap(1.0, 0.5))

			//Test 0
			g.Assert(!TestOverlap(0.0, 1.0))
			g.Assert(!TestOverlap(0.0, 0.0))

			//test denormalized numbers
			g.Assert(!TestOverlap(math.Pow(2, -1024), math.Pow(2,-1023)))
			g.Assert(!TestOverlap(math.Pow(2,-1023),  math.Pow(2,-1022)))
			g.Assert(TestOverlap( math.Pow(2,-1024) + math.Pow(2,-1040),  math.Pow(2,-1030)))
			g.Assert(!TestOverlap(math.Pow(2,-1030) - math.Pow(2,-1031),  math.Pow(2,-1030)))

		})
	})
}
