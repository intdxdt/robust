package twiddle

import (
	"testing"
	"github.com/franela/goblin"
	"fmt"
	"strings"
)

func TestTwiddle(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Twiddle", func() {
		g.It("not", func() {
			g.Assert(not(170) == -171).IsTrue()
			g.Assert(not(0) == -1).IsTrue()
			g.Assert(not(-3) == 2).IsTrue()
		})
		g.It("sign", func() {
			g.Assert(Sign(-100) == -1).IsTrue()
			g.Assert(Sign(100) == 1).IsTrue()
			g.Assert(Sign(0) == 0).IsTrue()
			g.Assert(Sign(INT_MAX) == 1).IsTrue()
			g.Assert(Sign(INT_MIN) == -1).IsTrue()
		})

		g.It("abs", func() {
			g.Assert(Abs(0) == 0)
			g.Assert(Abs(1) == 1)
			g.Assert(Abs(-1) == 1)
			g.Assert(Abs(INT_MAX) == INT_MAX)
			g.Assert(Abs(-INT_MAX) == INT_MAX)
			//abs(-INT_MIN) -- overflow
		})
		g.It("min", func() {
			g.Assert(Min(0, 0) == 0).IsTrue()
			g.Assert(Min(-1, 1) == -1).IsTrue()
			g.Assert(Min(INT_MAX, INT_MAX) == INT_MAX).IsTrue()
			g.Assert(Min(INT_MIN, INT_MIN) == INT_MIN).IsTrue()
			g.Assert(Min(INT_MAX, INT_MIN) == INT_MIN).IsTrue()
		})

		//    #[test]
		g.It("max", func() {
			g.Assert(Max(0, 0) == 0)
			g.Assert(Max(-1, 1) == 1)
			g.Assert(Max(INT_MAX, INT_MAX) == INT_MAX)
			g.Assert(Max(INT_MIN, INT_MIN) == INT_MIN)
			g.Assert(Max(INT_MAX, INT_MIN) == INT_MAX)
		})
		//
		//    #[test]
		g.It("is pow2", func() {
			g.Assert(!IsPow2(0))
			for i := 0; i < 31; i++ {
				g.Assert(IsPow2(1 << uint32(i))).IsTrue()
			}
			g.Assert(!IsPow2(100))
			g.Assert(!IsPow2(0x7fffffff))
			g.Assert(!IsPow2(-1000000))
		})

		g.It("log2", func() {
			for i := 0; i < 31; i++ {
				if i > 0 {
					g.Assert(Log2((1<<uint32(i))-1) == uint32(i-1))
					g.Assert(Log2((1<<uint32(i))+1) == uint32(i))
				}
				g.Assert(Log2(1<<uint32(i)) == uint32(i))
			}
		})

		g.It("log10", func() {
			g.Assert(Log10(1) == 0).IsTrue()
			g.Assert(Log10(10) == 1).IsTrue()
			g.Assert(Log10(100) == 2).IsTrue()
			g.Assert(Log10(1000) == 3).IsTrue()
			g.Assert(Log10(10000) == 4).IsTrue()
			g.Assert(Log10(100000) == 5).IsTrue()
			g.Assert(Log10(1234007) == 6).IsTrue()
			g.Assert(Log10(10004659) == 7).IsTrue()
			g.Assert(Log10(100046598) == 8).IsTrue()
			g.Assert(Log10(1000465983) == 9).IsTrue()
		})

		g.It("pop_count", func() {
			g.Assert(PopCount(0) == 0).IsTrue()
			g.Assert(PopCount(1) == 1).IsTrue()
			//g.Assert(pop_count(-1), 32)
			for i := 0; i < 31; i++ {
				g.Assert(PopCount(1<<uint32(i)) == 1).IsTrue()
				g.Assert(PopCount((1<<uint32(i))-1) == uint32(i)).IsTrue()
			}
			g.Assert(PopCount(0xf0f00f0f) == uint32(16)) //overflow for int32
		})

		//#[test]
		g.It("count_trailing_zeros", func() {
			g.Assert(CountTrailingZeros(0) == 32).IsTrue()
			g.Assert(CountTrailingZeros(1) == 0).IsTrue()
			//    g.Assert(CountTrailingZeros(-1), 0)
			for i := 0; i < 31; i++ {
				g.Assert(CountTrailingZeros(1<<uint32(i)) == uint32(i)).IsTrue()
				if i > 0 {
					g.Assert(CountTrailingZeros((1<<uint32(i))-1) == 0).IsTrue()
				}
			}
			g.Assert(CountTrailingZeros(0xf81700) == 8).IsTrue()
		})

		//#[test]
		g.It("next_pow2", func() {
			for i := 0; i < 31; i++ {
				if i != 1 {
					g.Assert(NextPow2((1<<uint32(i))-1) == 1<<uint32(i)).IsTrue()
				}
				g.Assert(NextPow2(1<<uint32(i)) == 1<<uint32(i)).IsTrue()
				if i < 30 {
					g.Assert(NextPow2((1<<uint32(i))+1) == 1<<(uint32(i)+1)).IsTrue()
				}
			}
		})

		//#[test]
		g.It("prev_pow2", func() {
			fmt.Printf("%2s    %10s    %10s\n", "i", "((1 << i) + 1)", "PrevPow2")
			fmt.Println(strings.Repeat("-", 34))
			for i := 0; i < 31; i++ {
				if i > 0 {
					g.Assert(PrevPow2((1<<uint32(i))-1) == 1<<(uint32(i)-1))
				}
				g.Assert(PrevPow2(1<<uint32(i)) == 1<<uint32(i))

				if 0 < i && i < 30 {
					fmt.Printf("%2d .. %10v .. %10v\n", i, (1<<uint32(i))+1, PrevPow2((1<<uint32(i))+1))
					g.Assert(PrevPow2((1<<uint32(i))+1) == 1<<uint32(i))
				}
			}
		})

		//#[test]
		g.It("parity", func() {
			g.Assert(Parity(1) == 1).IsTrue()
			g.Assert(Parity(0) == 0).IsTrue()
			g.Assert(Parity(0xf) == 0).IsTrue()
			g.Assert(Parity(0x10f) == 1).IsTrue()
		})

		g.It("reverse", func() {
			g.Assert(Reverse(0) == 0).IsTrue()
			//g.Assert(Reverse(-1)== -1).IsTrue() //overflow
		})

		g.It("next_combination", func() {
			g.Assert(NextCombination(1) == 2).IsTrue()
			g.Assert(NextCombination(0x300) == 0x401).IsTrue()
		})
		//
		//#[test]
		g.It("interleave2", func() {
			for x := 0; x < 100; x++ {
				for y := 0; y < 100; y++ {
					var h = Interleave2(uint32(x), uint32(y))
					g.Assert(Deinterleave2(uint32(h), 0) == uint32(x))
					g.Assert(Deinterleave2(uint32(h), 1) == uint32(y))
				}
			}
		})

		//#[test]
		g.It("interleave3", func() {
			for x := 0; x <= 25; x++ {
				for y := 0; y <= 25; y++ {
					for z := 0; z <= 25; z++ {
						var h = Interleave3(uint32(x), uint32(y), uint32(z))
						g.Assert(Deinterleave3(uint32(h), 0) == uint32(x))
						g.Assert(Deinterleave3(uint32(h), 1) == uint32(y))
						g.Assert(Deinterleave3(uint32(h), 2) == uint32(z))
					}
				}
			}
		})

	})
}
