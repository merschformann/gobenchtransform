package example_test

import (
	"testing"

	"github.com/merschformann/gobenchtransform/example"
)

func generateRanges(ranges int) []example.Range {
	r := make([]example.Range, ranges)
	width, offset := 10, 20
	for i := 0; i < ranges; i++ {
		r[i] = example.Range{Min: i * (width + offset), Max: i*(width+offset) + width}
	}
	return r
}

func benchmarkRangeChecker(b *testing.B, rc example.RangeChecker) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rc.InRange(i)
	}
}

func BenchmarkSlice(b *testing.B) {
	benchmarkRangeChecker(b, example.NewRangeCheckerSlice(generateRanges(100)))
}

func BenchmarkMap(b *testing.B) {
	benchmarkRangeChecker(b, example.NewRangeCheckerMap(generateRanges(100)))
}
