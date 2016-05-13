// This package provides an implementation of Go's math/rand API that is backed
// by the CSPRNG in crypto/rand.  This allows using the nicer API from
// math/rand, but still ensuring that the random number source is
// cryptographically-secure.
package csmrand
