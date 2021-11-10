#!/bin/bash

gcc -O2 c/emptycall_bench.c -o empty_c
go test -bench=. -count=5 -cpu=1

for i in `seq 1 5`; do
    ./empty_c
done
rm empty_c