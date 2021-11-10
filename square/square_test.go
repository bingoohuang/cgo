package square

import "testing"

func BenchmarkSumNSquare(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sumFirstNSquares(10000)
	}
}

func BenchmarkGoSumNSquare(b *testing.B) {
	for n := 0; n < b.N; n++ {
		goSumFirstNSquares(10000)
	}
}

func BenchmarkBetterSumNSquare(b *testing.B) {
	for n := 0; n < b.N; n++ {
		betterSumFirstNSquares(10000)
	}
}
