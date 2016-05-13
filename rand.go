package csmrand

import (
	mrand "math/rand"
)

var (
	globalRand *mrand.Rand
)

func init() {
	globalRand = mrand.New(CSRandSource{})
}

func ExpFloat64() float64 {
	return globalRand.ExpFloat64()
}

func Float32() float32 {
	return globalRand.Float32()
}

func Float64() float64 {
	return globalRand.Float64()
}

func Int() int {
	return globalRand.Int()
}

func Int31() int32 {
	return globalRand.Int31()
}

func Int31n(n int32) int32 {
	return globalRand.Int31n(n)
}

func Int63() int64 {
	return globalRand.Int63()
}

func Int63n(n int64) int64 {
	return globalRand.Int63n(n)
}

func Intn(n int) int {
	return globalRand.Intn(n)
}

func NormFloat64() float64 {
	return globalRand.NormFloat64()
}

func Perm(n int) []int {
	return globalRand.Perm(n)
}

func Read(p []byte) (n int, err error) {
	return globalRand.Read(p)
}

func Seed(seed int64) {
	globalRand.Seed(seed)
}

func Uint32() uint32 {
	return globalRand.Uint32()
}
