# Benchmark: CGO vs GO vs C in Empty Calls

## Run

```bash
sh run.sh
```

## Results

```sh
🕙[2021-11-10 23:17:28.000] ❯ ./run.sh 
goos: darwin
goarch: amd64
pkg: github.com/bingoohuang/cgo-bench/emptycall
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkEmptyCgoCalls  20005183                59.12 ns/op
BenchmarkEmptyCgoCalls  18326498                57.97 ns/op
BenchmarkEmptyCgoCalls  20268393                59.62 ns/op
BenchmarkEmptyCgoCalls  20969628                59.63 ns/op
BenchmarkEmptyCgoCalls  19783921                59.55 ns/op
BenchmarkEmptyGoCalls   1000000000               0.2533 ns/op
BenchmarkEmptyGoCalls   1000000000               0.2528 ns/op
BenchmarkEmptyGoCalls   1000000000               0.2524 ns/op
BenchmarkEmptyGoCalls   1000000000               0.2542 ns/op
BenchmarkEmptyGoCalls   1000000000               0.2535 ns/op
PASS
ok      github.com/bingoohuang/cgo-bench/emptycall      8.522s
BenchmarkEmptyCCalls    1000000000      0.00 ns/op
BenchmarkEmptyCCalls    1000000000      0.00 ns/op
BenchmarkEmptyCCalls    1000000000      0.00 ns/op
BenchmarkEmptyCCalls    1000000000      0.00 ns/op
BenchmarkEmptyCCalls    1000000000      0.00 ns/op
```

## Conclusions

- Pure Go call is `(57.31-0.2462)/0.2462 = 231.78` faster than Cgo call.
- Pure C call is (0.2462 - 0.00) / 0.00 = infinity faster than Go call.
- Pure C call is (57.31 - 0.00) / 0.00 = infinity faster than Cgo call.

## Related researches

- https://github.com/draffensperger/go-interlang
