# go-example-tar-hash-benchmark

Create tgz

```sh
$ base64 -i /dev/random 2>/dev/null | fold -w 100 | head -c 20400000 > 20M.txt
$ tar -czvf example.tgz 20M.txt
```

Run benchmark

```sh
$ go test -bench BenchmarkCalcHashAndCopyFile -memprofile mem.pprof
goos: darwin
goarch: arm64
pkg: go-example-tar-hash-benchmark
BenchmarkCalcHashAndCopyFile-10               82          15065315 ns/op
PASS
ok      go-example-tar-hash-benchmark   2.254s
```

Profiling memory usage

```sh
$ go tool pprof mem.pprof
Type: alloc_space
Time: Mar 21, 2024 at 4:39pm (JST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 11.35MB, 100% of 11.35MB total
      flat  flat%   sum%        cum   cum%
   11.35MB   100%   100%    11.35MB   100%  io.copyBuffer
         0     0%   100%    11.35MB   100%  go-example-tar-hash-benchmark.BenchmarkCalcHashAndCopyFile
         0     0%   100%    11.35MB   100%  go-example-tar-hash-benchmark.calcHashAndCopyFile
         0     0%   100%    11.35MB   100%  io.Copy (inline)
         0     0%   100%     7.74MB 68.18%  os.(*File).ReadFrom
         0     0%   100%     7.74MB 68.18%  os.genericReadFrom
         0     0%   100%    11.35MB   100%  testing.(*B).launch
         0     0%   100%    11.35MB   100%  testing.(*B).runN
(pprof)
```
