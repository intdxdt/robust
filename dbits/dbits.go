package dbits

import (
	"bytes"
	"encoding/binary"
)

func DoubleBits(n float64) [2]uint32 {
	var a, b uint32
	var buf bytes.Buffer

	if err := binary.Write(&buf, binary.LittleEndian, n); err != nil {
		panic(err)
	}

	if err := binary.Read(&buf, binary.LittleEndian, &a); err != nil {
		panic(err)
	}

	if err := binary.Read(&buf, binary.LittleEndian, &b); err != nil {
		panic(err)
	}

	return [2]uint32{a, b}
}

//pack low and higher uints of float as float 
func Pack(lo uint32, hi uint32) float64 {
	var buf bytes.Buffer
	var f float64
	if err := binary.Write(&buf, binary.LittleEndian, lo); err != nil {
		panic(err)
	}
	if err := binary.Write(&buf, binary.LittleEndian, hi); err != nil {
		panic(err)
	}
	if err := binary.Read(&buf, binary.LittleEndian, &f); err != nil {
		panic(err)
	}
	return f
}

func Lo(n float64) uint32 {
	var buf bytes.Buffer
	var a uint32
	if err := binary.Write(&buf, binary.LittleEndian, n); err != nil {
		panic(err)
	}

	if err := binary.Read(&buf, binary.LittleEndian, &a); err != nil {
		panic(err)
	}
	return a
}

func Hi(n float64) uint32 {
	var buf bytes.Buffer
	var a, b uint32
	if err := binary.Write(&buf, binary.LittleEndian, n); err != nil {
		panic(err)
	}
	binary.Read(&buf, binary.LittleEndian, &a)
	if err := binary.Read(&buf, binary.LittleEndian, &b); err != nil {
		panic(err)
	}
	return b
}

func Sign(n float64) uint32 {
	return Hi(n) >> 31
}

func Exponent(n float64) int32 {
	var b = Hi(n)
	return int32((b<<1)>>21) - 1023
}

func Fraction(n float64) [2]uint32 {
	var l = Lo(n)
	var h = Hi(n)
	var b = h & ((1 << 20) - 1)
	if (h & 0x7ff00000) != 0 {
		b += 1 << 20
	}
	return [2]uint32{l, b}
}

func Denormalized(n float64) bool {
	var h = Hi(n)
	return (h & 0x7ff00000) == 0
}
