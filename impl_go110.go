// +build go1.10

package csmrand

func Shuffle(n int, swap func(i, j int)) {
	globalRand.Shuffle(n, swap)
}

func (r Rand) Shuffle(n int, swap func(i, j int)) {
	globalRand.Shuffle(n, swap)
}
