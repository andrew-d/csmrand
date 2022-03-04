package csmrand

// A source of random numbers.  Mimics the type of the same name in math/rand.
// This is useful if you want to use this package as a drop-in replacement for
// code that uses math/rand.
//
// Note that since this package obtains all state from the crypto/rand.Reader,
// all methods of this struct use a single shared implementation.
type Rand struct{}

func (r Rand) ExpFloat64() float64 {
	return globalRand.ExpFloat64()
}

func (r Rand) Float32() float32 {
	return globalRand.Float32()
}

func (r Rand) Float64() float64 {
	return globalRand.Float64()
}

func (r Rand) Int() int {
	return globalRand.Int()
}

func (r Rand) Int31() int32 {
	return globalRand.Int31()
}

func (r Rand) Int31n(n int32) int32 {
	return globalRand.Int31n(n)
}

func (r Rand) Int63() int64 {
	return globalRand.Int63()
}

func (r Rand) Int63n(n int64) int64 {
	return globalRand.Int63n(n)
}

func (r Rand) Intn(n int) int {
	return globalRand.Intn(n)
}

func (r Rand) NormFloat64() float64 {
	return globalRand.NormFloat64()
}

func (r Rand) Perm(n int) []int {
	return globalRand.Perm(n)
}

func (r Rand) Read(p []byte) (n int, err error) {
	return globalRand.Read(p)
}

// Note: does nothing.
func (r Rand) Seed(seed int64) {
	globalRand.Seed(seed)
}

func (r Rand) Uint32() uint32 {
	return globalRand.Uint32()
}

func (r Rand) Shuffle(n int, swap func(i, j int)) {
	globalRand.Shuffle(n, swap)
}
