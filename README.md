This is a problem from the Gopher Slack room


```
florin@sniper [09:10:45] [~/golang/src/github.com/dlsniper/mul/mul] [master]
-> % go test -bench=.                                 
testing: warning: no tests to run
PASS
BenchmarkMul1-8     2000            684466 ns/op
BenchmarkMul2-8     5000            457176 ns/op
BenchmarkMul3-8     5000            212010 ns/op
BenchmarkMul4-8    10000            150611 ns/op
BenchmarkMul5-8    10000            153173 ns/op
BenchmarkMul6-8     3000            976646 ns/op
ok      github.com/dlsniper/mul/mul     10.907s
```

```
florin@sniper [09:09:58] [~/golang/src/github.com/dlsniper/mul/mul] [master *]
-> % go test -benchmem -test.memprofilerate=1 -bench=.
testing: warning: no tests to run
PASS
BenchmarkMul1-8     2000            735461 ns/op           17335 B/op          5 allocs/op
BenchmarkMul2-8     3000            455565 ns/op           24584 B/op          2 allocs/op
BenchmarkMul3-8     5000            210202 ns/op           16384 B/op          1 allocs/op
BenchmarkMul4-8    10000            149474 ns/op           16384 B/op          1 allocs/op
BenchmarkMul5-8    10000            152326 ns/op           16384 B/op          1 allocs/op
BenchmarkMul6-8     2000           1764112 ns/op           26337 B/op        112 allocs/op
ok      github.com/dlsniper/mul/mul     10.744s
```