goos: windows
goarch: amd64
pkg: bad-comment-server
cpu: 11th Gen Intel(R) Core(TM) i5-11500 @ 2.70GHz
BenchmarkGetDistricts
BenchmarkGetDistricts-12                    2617            517866 ns/op
17928 B/op        125 allocs/op
BenchmarkParse
BenchmarkParse-12                           2120            667093 ns/op
18851 B/op        133 allocs/op
BenchmarkValidate
BenchmarkValidate-12                        1796            710589 ns/op
18724 B/op        133 allocs/op
BenchmarkParallelValidate
BenchmarkParallelValidate-12                5308            392678 ns/op
16233 B/op        121 allocs/op
BenchmarkParallelParse
BenchmarkParallelParse-12                  27702             44516 ns/op
8212 B/op        103 allocs/op
BenchmarkParallelGetDistricts
BenchmarkParallelGetDistricts-12           26863             47922 ns/op
7059 B/op         94 allocs/op