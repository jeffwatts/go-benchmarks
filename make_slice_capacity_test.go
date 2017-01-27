package go_benchmarks

import (
	"testing"
)

const (
	testStr = "This is the string that will be copied into the tests"
	sliceLen = 5
)

func BenchmarkMakeWithCapacity(b *testing.B) {
	var result []TextRange // Just to avoid the compiler optimizing the method call out

	for i := 0; i < b.N; i++ {
		MakeWithCapacity(sliceLen, testStr, sliceLen + 1)
	}

	_ = result
}

func BenchmarkMakeWithoutCapacity(b *testing.B) {
	var result []TextRange // Just to avoid the compiler optimizing the method call out

	for i := 0; i < b.N; i++ {
		MakeWithoutCapacity(sliceLen, testStr)
	}

	_ = result
}
