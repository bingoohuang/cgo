package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/bingoohuang/cgo-bench/bench"
	"github.com/bingoohuang/cgo-bench/str"
)

var (
	pNum        *int
	pCoroutines *int
	pJob        *string
)

func init() {
	pNum = flag.Int("n", 1000000, "total times")
	pCoroutines = flag.Int("c", 100, "coroutines")
	pJob = flag.String("job", "base64", "job to benchmark, base64/unbase64/base64both")
	flag.Parse()
}

func main() {
	s := []byte(str.RandStr(10 * 1024))
	var fn func()
	switch *pJob {
	default:
		fn = func() {
			base64.StdEncoding.EncodeToString(s)
		}
	case "base64":
		fn = func() {
			base64.StdEncoding.EncodeToString(s)
		}
	case "unbase64":
		b := base64.StdEncoding.EncodeToString(s)
		fn = func() {
			base64.StdEncoding.DecodeString(b)
		}
	case "base64both":
		fn = func() {
			b := base64.StdEncoding.EncodeToString(s)
			base64.StdEncoding.DecodeString(b)
		}
	}

	out := bench.Bench(fn, &bench.Config{
		N:          *pNum,
		Coroutines: *pCoroutines,
	})
	fmt.Println("TPS:", out)
}
