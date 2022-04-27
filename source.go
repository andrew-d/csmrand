package csmrand

import (
	"crypto/rand"
	"encoding/binary"
	mrand "math/rand"
)

// CSRandSource is a math/rand.Source that wraps crypto/rand.Reader.
type CSRandSource struct{}

// readUint64Bytes returns 8-bytes, suitable for pulling an uint64 out of.
func readUint64Bytes() []byte {
	buf := make([]byte, 8)
	n, err := rand.Reader.Read(buf)
	if err != nil || n != 8 {
		panic("could not read random data")
	}

	return buf
}

// Int63 implements the math/rand.Source interface.  Will panic if an
// error occurs when reading the underlying Reader.
func (r CSRandSource) Int63() int64 {
	buf := readUint64Bytes()
	i := int64(binary.LittleEndian.Uint64(buf))

	return i & 0x7fffffffffffffff
}

// Seed implements the math/rand.Source interface.  Does nothing.
func (r CSRandSource) Seed(s int64) {
	// Do nothing
}

// Uint64 implements the math/rand.Source64 interface.  Will panic if an
// error occurs when reading the underlying Reader.
func (r CSRandSource) Uint64() uint64 {
	buf := readUint64Bytes()
	return binary.LittleEndian.Uint64(buf)
}

var _ mrand.Source = CSRandSource{}
var _ mrand.Source64 = CSRandSource{}
