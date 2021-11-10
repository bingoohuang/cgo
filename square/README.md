```sh
🕙[2021-11-10 23:35:43.259] ❯ go test -bench=. -benchmem       
goos: darwin
goarch: amd64
pkg: github.com/bingoohuang/cgo-bench/square
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkSumNSquare-12                      1912            612889 ns/op               0 B/op          0 allocs/op
BenchmarkGoSumNSquare-12                  472579              2559 ns/op               0 B/op          0 allocs/op
BenchmarkBetterSumNSquare-12            17852760                60.06 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/bingoohuang/cgo-bench/square 3.620s

Downloads/cgo-bench/square via 🐹 v1.17.3 via C base took 4s 
```
