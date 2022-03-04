package csmrand

import (
	"crypto/rand"
	"encoding/binary"
	mrand "math/rand"
)

// CSRandSource is a math/rand.Source that wraps crypto/rand.Reader.
type CSRandSource struct{}

// Int63 implements the math/rand.Source interface.  Will panic if an
// error occurs when reading the underlying Reader.
func (r CSRandSource) Int63() int64 {
	var i int64

	err := binary.Read(rand.Reader, binary.LittleEndian, &i)
	if err != nil {
		panic("could not read random data")
	}

	return (i & 0x7fffffffffffffff)
}

// Seed implements the math/rand.Source interface.  Does nothing.
func (r CSRandSource) Seed(s int64) {
	// Do nothing
}

// Uint64 implements the math/rand.Source64 interface.  Will panic if an
// error occurs when reading the underlying Reader.
func (r CSRandSource) Uint64() uint64 {
	var i uint64

	err := binary.Read(rand.Reader, binary.LittleEndian, &i)
	if err != nil {
		panic("could not read random data")
	}

	return i
}

var _ mrand.Source = CSRandSource{}
var _ mrand.Source64 = CSRandSource{}
