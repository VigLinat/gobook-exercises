package echo

import "testing"

// TODO: replace with random string on each iteration
var benchmarkString = []string{"abc", "def", "ghi", "jkl", "mno", "pqr", "stu", "vwx", "yz", "a", "b", "c", "d", "e", "f", "g", "h"}

func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slow(benchmarkString)
	}
}

func BenchmarkFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fast(benchmarkString)
	}
}
