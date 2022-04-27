package csmrand

import (
	"crypto/rand"
	"encoding/binary"
	mrand "math/rand"
)

// randBufferCapacity is the amount of random bytes we will read from rand.Reader
// at once.  It is sized to amortize the cost of the getrandom syscall without
// being an outrageously large buffer.
const randBufferCapacity = 4096

// CSStatefulRandSource is a math/rand.Source that wraps crypto/rand.Reader that
// is not safe for use by concurrent goroutines.  Similar to math/rand.Source,
// protect its usage by using a sync.Pool (recommended) or with a sync.Mutex.
type CSStatefulRandSource struct {
	buf []byte
}

// refill reads 'randBufferCapacity' bytes of random into r.buf
func (r *CSStatefulRandSource) refill() {
	if cap(r.buf) == randBufferCapacity {
		// if we've already read in
		r.buf = r.buf[0:randBufferCapacity]
	} else {
		r.buf = make([]byte, randBufferCapacity)
	}

	bytesRead, err := rand.Reader.Read(r.buf)
	if err != nil {
		panic("could not read random data")
	}
	if bytesRead != randBufferCapacity {
		panic("short read of random data")
	}

	if len(r.buf) != randBufferCapacity {
		panic("invariant broken")
	}
}

func (r *CSStatefulRandSource) consume8Bytes() []byte {
	if len(r.buf) == 0 {
		r.refill()
	}

	size := len(r.buf)
	if size < 8 {
		panic("invariant broken")
	}

	// consume from the end so that r.buf is always pointing at the
	// start of the backing array (and when we've consumed all of our
	// randomness we can just re-enlarge)
	consumed := r.buf[size-8:]
	r.buf = r.buf[:size-8]

	return consumed
}

// Int63 implements the math/rand.Source interface.  Will panic if an
// error occurs when reading the underlying Reader.
func (r *CSStatefulRandSource) Int63() int64 {
	buf := r.consume8Bytes()
	i := int64(binary.LittleEndian.Uint64(buf))
	return i & 0x7fffffffffffffff
}

// Seed implements the math/rand.Source interface.  Does nothing.
func (r *CSStatefulRandSource) Seed(s int64) {
	// Do nothing
}

// Uint64 implements the math/rand.Source64 interface.  Will panic if an
// error occurs when reading the underlying Reader.
func (r *CSStatefulRandSource) Uint64() uint64 {
	buf := r.consume8Bytes()
	return binary.LittleEndian.Uint64(buf)
}

var _ mrand.Source = &CSStatefulRandSource{}
var _ mrand.Source64 = &CSStatefulRandSource{}
