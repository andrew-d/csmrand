package csmrand

import (
	"math/rand"
	"testing"
)

func TestCSRandSource_Int63(t *testing.T) {
	var f float64
	allocs := testing.AllocsPerRun(1, func() {
		f = Float64()
	})
	if allocs > 1 {
		t.Fatalf("expected 0 allocs, got %f", allocs)
	}
	_ = f
}

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
