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
