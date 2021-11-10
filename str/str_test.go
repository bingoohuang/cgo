package str

import (
	"testing"
)

func BenchmarkEmptyCgo10(b *testing.B) {
	s := RandStr(10)
	for i := 0; i < b.N; i++ {
		Cempty(s)
	}
}

func BenchmarkEmptyCgo1K(b *testing.B) {
	s := RandStr(1024)
	for i := 0; i < b.N; i++ {
		Cempty(s)
	}
}

func BenchmarkEmptyCgo10K(b *testing.B) {
	s := RandStr(10240)
	for i := 0; i < b.N; i++ {
		Cempty(s)
	}
}

func BenchmarkEmptyCgo20K(b *testing.B) {
	s := RandStr(20480)
	for i := 0; i < b.N; i++ {
		Cempty(s)
	}
}

func TestMystr1(t *testing.T) {
	var s = "test cString"
	mystr(s)
}

func BenchmarkMystr20K(b *testing.B) {
	s := RandStr(20480)
	for i := 0; i < b.N; i++ {
		mystr(s)
	}
}

func TestMystr2(t *testing.T) {
	var s = "test c\000String"
	mystr(s)
}

func BenchmarkMystr2K(b *testing.B) {
	s := RandStr(2048)
	for i := 0; i < b.N; i++ {
		mystr(s)
	}
}
