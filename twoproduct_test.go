package robust

import (
	"math"
	"time"
	"testing"
	"math/rand"
	"github.com/franela/goblin"
)

var testValues []float64

func init() {
	var seed = rand.NewSource(time.Now().UnixNano())
	var random = rand.New(seed)
	testValues = []float64{
		0, 1,
		math.Pow(2, -52),
		math.Pow(2, -53),
		1.0 - math.Pow(2, -53),
		1.0 + math.Pow(2, -52),
		math.Pow(2, -500),
		math.Pow(2, 500),
		2, 0.5, 3, 1.5, 0.1, 0.3, 0.7,
	}
	for i := 0; i < 20; i++ {
		testValues = append(testValues, random.Float64())
		testValues = append(testValues, random.Float64()*math.Pow(2, 1000*random.Float64()-500))
	}
	for i := len(testValues) - 1; i > 0; i-- {
		testValues = append(testValues, -testValues[i])
	}
}

func TestTwoProduct(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("TwoProduct", func() {
		g.It("test fast two product", func() {
			g.Assert(
				TwoProduct(1.0+math.Pow(2, -52), 1.0+math.Pow(2, -52))).Eql(
				[]float64{math.Pow(2, -104), 1.0 + math.Pow(2, -51)})

			for j := 0; j < len(testValues); j++ {
				a := testValues[j]
				g.Assert(a*a < math.Inf(1)).IsTrue()
				g.Assert(TwoProduct(0, a)).Eql([]float64{0, 0})
				g.Assert(TwoProduct(1, a)).Eql([]float64{0, a})
				g.Assert(TwoProduct(-1, a)).Eql([]float64{0, -a})
			}
		})
	})
}
