package csmrand

import (
	"crypto/rand"
	"encoding/binary"
	mrand "math/rand"
)

type csRandSource struct{}

func (r csRandSource) Int63() int64 {
	var i int64

	err := binary.Read(rand.Reader, binary.LittleEndian, &i)
	if err != nil {
		panic("could not read random data")
	}

	return (i & 0x7fffffffffffffff)
}

func (r csRandSource) Seed(s int64) {
	// Do nothing
}

var _ mrand.Source = csRandSource{}
