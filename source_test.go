package csmrand

import (
	"math/rand"
	"testing"
)

func BenchmarkGlobalFloat64(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(8)
	b.RunParallel(func(pb *testing.PB) {
		var f float64
		for pb.Next() {
			f = Float64()
		}
		_ = f
	})
}

func BenchmarkMathGlobalFloat64(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(8)
	b.RunParallel(func(pb *testing.PB) {
		var f float64
		for pb.Next() {
			f = rand.Float64()
		}
		_ = f
	})
}

func BenchmarkLocalFloat64(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(8)
	b.RunParallel(func(pb *testing.PB) {
		rng := rand.New(CSRandSource{})
		var f float64
		for pb.Next() {
			f = rng.Float64()
		}
		_ = f
	})
}

func BenchmarkMathLocalFloat64(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(8)
	b.RunParallel(func(pb *testing.PB) {
		rng := rand.New(rand.NewSource(42))
		var f float64
		for pb.Next() {
			f = rng.Float64()
		}
		_ = f
	})
}
