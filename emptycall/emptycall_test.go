package emptycall_test

import (
	"testing"

	"github.com/bingoohuang/cgo-bench/emptycall"
)

func BenchmarkEmptyCgoCalls(b *testing.B) {
	for i := 0; i < b.N; i++ {
		emptycall.Cempty()
	}
}

func BenchmarkEmptyGoCalls(b *testing.B) {
	for i := 0; i < b.N; i++ {
		emptycall.Empty()
	}
}
