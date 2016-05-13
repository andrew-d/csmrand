package csmrand

import (
	mrand "math/rand"
)

var (
	source   mrand.Source
	randInst *mrand.Rand
)

func init() {
	source = csRandSource{}
	randInst = mrand.New(source)
}

func ExpFloat64() float64 {
	return randInst.ExpFloat64()
}

func Float32() float32 {
	return randInst.Float32()
}

func Float64() float64 {
	return randInst.Float64()
}

func Int() int {
	return randInst.Int()
}

func Int31() int32 {
	return randInst.Int31()
}

func Int31n(n int32) int32 {
	return randInst.Int31n(n)
}

func Int63() int64 {
	return randInst.Int63()
}

func Int63n(n int64) int64 {
	return randInst.Int63n(n)
}

func Intn(n int) int {
	return randInst.Intn(n)
}

func NormFloat64() float64 {
	return randInst.NormFloat64()
}

func Perm(n int) []int {
	return randInst.Perm(n)
}

func Read(p []byte) (n int, err error) {
	return randInst.Read(p)
}

func Seed(seed int64) {
	randInst.Seed(seed)
}

func Uint32() uint32 {
	return randInst.Uint32()
}
