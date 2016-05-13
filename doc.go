// This package provides an implementation of Go's math/rand API that is backed
// by the CSPRNG in crypto/rand.  This allows using the nicer API from
// math/rand, but still ensuring that the random number source is
// cryptographically-secure.
//
// If there is a case where this package cannot obtain random data from
// crypto/rand, it will panic.  This is largely the solution you want, since
// this should only occur (on Linux) when the /dev/urandom pool hasn't been
// initialized yet.  In this case, the calling code can't do anything sensible
// with the error anyway.  Also, panicking on this error allows us to maintain
// the same API as math/rand.
package csmrand
