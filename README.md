# csmrand

[![Go Reference](https://pkg.go.dev/badge/github.com/andrew-d/csmrand.svg)](https://pkg.go.dev/github.com/andrew-d/csmrand) ![Build Status](https://github.com/andrew-d/csmrand/actions/workflows/go.yml/badge.svg)


Golang's math/rand backed onto crypto/rand's CSPRNG.

This package provides an implementation of Go's `math/rand` API that is backed
by the CSPRNG in `crypto/rand`.  This allows using the nicer API from
`math/rand`, but still ensuring that the random number source is
cryptographically-secure. See the documentation for more details.
