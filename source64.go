// +build go1.8

package csmrand

import (
	"crypto/rand"
	"encoding/binary"
	mrand "math/rand"
)

// Implement the math/rand.Source64 interface.  Will panic if an error occurs
// when reading the underlying Reader.
func (r CSRandSource) Uint64() uint64 {
	var i uint64

	err := binary.Read(rand.Reader, binary.LittleEndian, &i)
	if err != nil {
		panic("could not read random data")
	}

	return i
}

func Uint64() uint64 {
	return globalRand.Uint64()
}

var _ mrand.Source64 = CSRandSource{}
