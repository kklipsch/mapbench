# Output from benchmarks	

Using a Macbook Pro (Retina, 13-inch, Late 2013), 2.4Ghz Intel Core i5, 8 GB 1600 MHz DDR3

go1.9

```bash
$ go test -cpuprofile=cpu.out -memprofile=mem.out -run=xxx -bench=. -benchmem .
goos: darwin
goarch: amd64
pkg: github.com/kklipsch/mapbench
BenchmarkMapInit-4               2000000               625 ns/op             672 B/op          4 allocs/op
BenchmarkMapSet-4                2000000               615 ns/op             672 B/op          4 allocs/op
BenchmarkMapFunctions-4          2000000               689 ns/op             692 B/op          6 allocs/op
BenchmarkMapSetAlloc-4           5000000               375 ns/op             336 B/op          2 allocs/op
BenchmarkMapFunctionsAlloc-4     3000000               460 ns/op             356 B/op          4 allocs/op
PASS
ok      github.com/kklipsch/mapbench    10.105s
```

