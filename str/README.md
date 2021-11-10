```sh
🕙[2021-11-11 07:02:56.890] ❯ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/bingoohuang/cgo-bench/str
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkEmptyCgo10-12           4039773               293.6 ns/op            16 B/op          1 allocs/op
BenchmarkEmptyCgo1K-12           3695788               322.1 ns/op            16 B/op          1 allocs/op
BenchmarkEmptyCgo10K-12          1969735               601.4 ns/op            16 B/op          1 allocs/op
BenchmarkEmptyCgo20K-12           915693              1128 ns/op              16 B/op          1 allocs/op
BenchmarkMystr20K-12            18098246                66.14 ns/op            0 B/op          0 allocs/op
BenchmarkMystr2K-12             17770120                63.43 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/bingoohuang/cgo-bench/str    8.328s

cgo-bench/str on  main [!?] via 🐹 v1.17.3 via C base took 9s 
```

1. [谈谈cgo字符串传递过程中的一些优化_wx_kingstone的博客-程序员宅基地](https://www.cxyzjd.com/article/qq_25341531/105616985)
