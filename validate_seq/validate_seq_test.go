package validate_seq

import (
	"testing"
	"github.com/franela/goblin"
)

func TestValidateSeq(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Validate Non-Overlaping Seq", func() {
		g.It("test validate seq", func() {
			g.Assert(ValidateSequence([]float64{1e-16, 1.})).IsTrue()
			g.Assert(ValidateSequence([]float64{0.5, 1.5})).IsFalse()
			g.Assert(ValidateSequence([]float64{0.})).IsTrue()
			g.Assert(ValidateSequence([]float64{})).IsFalse()
		})
	})
}
