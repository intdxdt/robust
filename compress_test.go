package robust

import (
	"time"
	"math"
	"testing"
	"math/rand"
	"github.com/franela/goblin"
)


func TestRobustCompress(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Compress", func() {
		g.It("test robust Compress", func() {
			var seed = rand.NewSource(time.Now().UnixNano())
			var random = rand.New(seed)
			g.Assert(Compress(af(0))).Eql(af(0))
			g.Assert(Compress(af(1))).Eql(af(1))

			for i := 0; i < 10; i++ {
				var h = random.Float64()
				g.Assert(Compress(af(h))).Eql(af(h))
				h = -h
				g.Assert(Compress(af(h))).Eql(af(h))
			}

			g.Assert(Compress(af(1, 2))).Eql(af(3))
			g.Assert(Compress(af(math.Pow(2, -52), 1))).Eql(af(1 + math.Pow(2, -52)))

			verify := func(seq []float64) {
				var rseq = Compress(seq[:len(seq):len(seq)])
				//must Compress:  + rseq.length +  <=  + seq.length
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
					seq = Scale(seq, 2*random.Float64()-1.0)
				}
				verify(seq)
			}

		})
	})
}
