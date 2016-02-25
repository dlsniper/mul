This is a problem from the Gopher Slack room


```
florin@sniper [08:29:47] [~/golang/src/github.com/dlsniper/mul/mul] 
-> % go test -bench=.
testing: warning: no tests to run
PASS
BenchmarkMul1-8     2000            654306 ns/op
BenchmarkMul2-8     3000            682292 ns/op
BenchmarkMul3-8     5000            202153 ns/op
BenchmarkMul4-8    10000            146371 ns/op
BenchmarkMul5-8    10000            145110 ns/op
ok      github.com/dlsniper/mul/mul     7.477s
```