package test_overlap

import (
	"math"
	"testing"
	"github.com/franela/goblin"
)

func TestFloatOverlap(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Test Float Overlap", func() {
		g.It("test double bits", func() {
			g.Assert(TestOverlap(1.5, 0.5))
			g.Assert(TestOverlap(math.Pow(2, -52), 1.0+math.Pow(2, -52)))
			g.Assert(!TestOverlap(1.0, 0.5))

			//Test 0
			g.Assert(!TestOverlap(0.0, 1.0))
			g.Assert(!TestOverlap(0.0, 0.0))

			//test denormalized numbers
			//underflow - in rust  float64MIN_EXP == -1021
			//
			//g.Assert(!TestOverlap(2float64.powi(-1024), 2float64.powi(-1023)));
			//g.Assert(!TestOverlap(2float64.powi(-1023), 2float64.powi(-1022)));
			//g.Assert(TestOverlap( 2float64.powi(-1024) + 2float64.powi(-1040), 2float64.powi(-1030)));
			//g.Assert(!TestOverlap(2float64.powi(-1030) - 2float64.powi(-1031), 2float64.powi(-1030)));

		})
	})
}
