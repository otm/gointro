package tests

import "testing"

func BenchmarkConvert10(b *testing.B) {
	// run the convert function b.N times
	for n := 0; n < b.N; n++ {
		convert(10)
	}
}
