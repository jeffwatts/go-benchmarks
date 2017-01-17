package go_benchmarks

import "testing"

const (
	sampleText = `This is the text we will test against, over and over and over again.
A possible match is 3528 3464 4649 6691, but it should not matter.`
	timesToMatch = 4
)

func BenchmarkCachedRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		compiledObject := NewCanonicalCompileOnceObject()

		for j := 0; j < timesToMatch; j++ {
			_ = compiledObject.FindMatchIndexes(sampleText)
		}
	}
}

func BenchmarkOnDemandRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < timesToMatch; j++ {
			_ = OnDemandFindMatchIndexes(sampleText)
		}
	}
}
