# Output from benchmarks	

Using a Macbook Pro (Retina, 15-inch, Mid 2014), 2.5Ghz Intel Core i7, 16 GB 1600 MHz DDR3

go1.9

```bash
$ go test -cpuprofile=cpu.out -memprofile=mem.out -run=xxx -bench=. -benchmem .
goarch: amd6 goos: darwin4
pkg: github.com/kklipsch/mapbench
BenchmarkMapInit-8               3000000               494 ns/op
BenchmarkMapSet-8                3000000               489 ns/op
BenchmarkMapFunctions-8          3000000               552 ns/op
BenchmarkMapInitVars-8           5000000               359 ns/op
BenchmarkMapSetVars-8            5000000               352 ns/op
BenchmarkMapFunctionsVars-8      5000000               359 ns/op
BenchmarkMapSetAlloc-8           5000000               297 ns/op
BenchmarkMapFunctionsAlloc-8     5000000               344 ns/op
PASS
ok      github.com/kklipsch/mapbench    10.105s
```

