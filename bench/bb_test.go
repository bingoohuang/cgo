package bench

import (
	"encoding/base64"
	"fmt"
	"github.com/bingoohuang/cgo-bench/str"
	"strings"
	"testing"
	"time"
	"unicode"
)

func TestBench(t *testing.T) {
	tps := Bench(func() {
		time.Sleep(10 * time.Millisecond)
	}, &Config{
		N:          1000,
		Coroutines: 10,
	})
	fmt.Println("TPS ", tps)
}

func BenchmarkBench(b *testing.B) {
	s := str.RandStr(10 * 1024)
	bs := base64.StdEncoding.EncodeToString([]byte(s))
	for i := 0; i < b.N; i++ {
		base64.StdEncoding.DecodeString(bs)
	}
}

var (
	s  = str.RandStr(10 * 1024)
	bs = base64.StdEncoding.EncodeToString([]byte(s))
)

func BenchmarkBase64Decode原生(b *testing.B) {
	for i := 0; i < b.N; i++ {
		base64.RawStdEncoding.DecodeString(bs)
	}
}

func BenchmarkBase64Decode当前(b *testing.B) {
	for i := 0; i < b.N; i++ {
		base64.RawStdEncoding.DecodeString(strings.TrimRight(StripUnprintable(bs), "="))
	}
}

func BenchmarkBase64Decode改造1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		base64.RawStdEncoding.DecodeString(strings.TrimRight(bs, "="))
	}
}

func StripUnprintable(str string) string {
	var b strings.Builder

	b.Grow(len(str))

	for _, ch := range str {
		if unicode.IsPrint(ch) && !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
