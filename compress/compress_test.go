package compress

import (
	"time"
	"math"
	"testing"
	rs "robust/scale"
	"math/rand"
	"github.com/franela/goblin"
)

func init() {
}


func TestRobustCompress(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("RobustCompress", func() {
		g.It("test robust RobustCompress", func() {
			var seed = rand.NewSource(time.Now().UnixNano())
			var random = rand.New(seed)
			g.Assert(RobustCompress(ar(0))).Eql(ar(0))
			g.Assert(RobustCompress(ar(1))).Eql(ar(1))

			for i := 0; i < 10; i++ {
				var h = random.Float64()
				g.Assert(RobustCompress(ar(h))).Eql(ar(h))
				h = -h
				g.Assert(RobustCompress(ar(h))).Eql(ar(h))
			}

			g.Assert(RobustCompress(ar(1, 2))).Eql(ar(3))
			g.Assert(RobustCompress(ar(math.Pow(2, -52), 1))).Eql(ar(1 + math.Pow(2, -52)))

			verify := func(seq []float64) {
				var rseq = RobustCompress(seq[:len(seq):len(seq)])
				//must RobustCompress:  + rseq.length +  <=  + seq.length
				g.Assert(len(rseq) <= len(seq)).IsTrue()
				//t.same(toFixed(rseq).toString(16), toFixed(seq).toString(16), "verifying sequence")
			}

			//Bigger sequences
			for i := 0; i < 10; i++ {
				var seq = make([]float64, 0)
				for j := 0; j < 18; j++ {
					seq = append(seq, math.Pow(2, -900+100*float64(j))*(random.Float64()-0.5))
				}
				verify(seq)
			}

			//Multiply a bunch of random numbers
			for i := 0; i < 10; i++ {
				var seq = []float64{1}
				for j := 0; j < 20; j++ {
					seq = rs.RobustScale(seq, 2*random.Float64()-1.0)
				}
				verify(seq)
			}

		})
	})
}

func ar(v ...float64) []float64 {
	return v
}
