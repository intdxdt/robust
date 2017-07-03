package robust

import (
	"math"
	"time"
	"testing"
	"math/rand"
	"github.com/franela/goblin"
)

func init() {
}

func TestRobustSubtract(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Subtract", func() {
		g.It("test robust diff", func() {
			var seed = rand.NewSource(time.Now().UnixNano())
			var random = rand.New(seed)

			g.Assert(Subtract(ar(1), ar(1))).Eql(ar(0))

			var s = ar(0)
			for i := 0; i < 100; i++ {
				s = Subtract(s, ar(random.Float64()*math.Pow(2, random.Float64()*1000)))
				//t.ok(validate(s))
			}
		})
	})
}


