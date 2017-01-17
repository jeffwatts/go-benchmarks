package go_benchmarks

import (
	"testing"
	"strings"
)

var (
	str1 = strings.Repeat("a", 1000000)
	str2 = strings.Repeat("b", 1000000)
	str3 = strings.Repeat("c", 1000000)
	str4 = strings.Repeat("d", 1000000)
	times = 4
)

func BenchmarkValueContainer(b *testing.B) {
	v := ValueContainer{str1, str2, str3, str4}

	for i := 0; i < b.N; i++ {
		PingInspector(times, v)
	}
}

func BenchmarkPointerContainer(b *testing.B) {
	p := &PointerContainer{str1, str2, str3, str4}

	for i := 0; i < b.N; i++ {
		PingInspector(times, p)
	}
}
