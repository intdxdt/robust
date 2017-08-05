package dbits

import (
	"math"
	"testing"
	"github.com/franela/goblin"
)

func TestDoubleBits(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("DoubleBits", func() {
		g.It("test double bits", func() {
			g.Assert(Lo(1.0)).Equal(uint32(0))
			g.Assert(Hi(1.0)).Equal(uint32(0x3ff00000))
			g.Assert(Pack(0, 0x3ff00000)).Equal(1.0)
			g.Assert(DoubleBits(1.0)).Equal([2]uint32{0, 0x3ff00000})

			g.Assert(Fraction(1.)).Equal([2]uint32{0, 1 << 20})
			g.Assert(Exponent(1.)).Equal(int32(0))
			g.Assert(Sign(1.)).Equal(uint32(0))
			g.Assert(Sign(-1.)).Equal(uint32(1))
			g.Assert(Exponent(0.5)).Equal(int32(-1))

			g.Assert(Denormalized(math.Pow(2, -1024))).IsTrue()
			g.Assert(Denormalized(1.)).IsFalse()
			g.Assert(Denormalized(math.Pow(2, -1023))).IsTrue()
			g.Assert(Denormalized(math.Pow(2, -1022))).IsFalse()
		})
	})
}
