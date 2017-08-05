package robust

import (
	"math"
	"testing"
	"robust/validate_seq"
	"github.com/franela/goblin"
	"time"
	"math/rand"
)

func TestRobustSum(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Sum", func() {
		g.It("test robust sum", func() {
			var seed = rand.NewSource(time.Now().UnixNano())
			var random = rand.New(seed)
			g.Assert(Sum(
				af(1, 64), af(-1e-64, 1e64),
			)).Eql(af(-1e-64, 65, 1e64), )

			g.Assert(Sum(af(0), af(0))).Eql(af(0))
			g.Assert(Sum(af(0), af(1))).Eql(af(1))

			g.Assert(Sum(af(1, 1e64), af(1e-64, 2), )).Eql(af(1e-64, 3, 1e64))

			g.Assert(Sum(af(1), af(1e-64, 1e-16), )).Eql(af(1e-64, 1e-16, 1))

			for i := -10; i <= 10; i++ {
				for j := -10; j <= 10; j++ {
					g.Assert(Sum(af(float64(i)), af(float64(j)))).Eql(af(float64(i + j)))
				}
			}

			g.Assert(validate_seq.ValidateSequence(Sum(
				[]float64{5.711861227349496e-133, 1e-116},
				[]float64{5.711861227349496e-133, 1e-116},
			)))

			nois := make([]float64, 10)
			expect := make([]float64, 10)
			for i := 0; i < 10; i++ {
				nois[i] = math.Pow(2, float64(-1000+53*i))
				expect[i] = math.Pow(2, float64(-999+53*i))
			}
			x := Sum(nois, nois)
			g.Assert(x).Eql(expect)
			g.Assert(validate_seq.ValidateSequence(x))

			g.Assert(Sum(af(0), af(1, 1e64))).Eql(af(1, 1e64))

			var s = []float64{0}
			for i := 0; i < 1000; i++ {
				s = Sum(s, []float64{random.Float64() * math.Pow(2, random.Float64()*1800-900)})
				g.Assert(validate_seq.ValidateSequence(s))
			}

		})
	})
}
