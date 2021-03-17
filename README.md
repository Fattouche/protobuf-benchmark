# protobuf-benchmark
Benchmarking protobuf vs json

## Running
`go test -bench .`

## Generation
`protoc -I pb test.proto --go_out=plugins=grpc:pb --go_opt=paths=source_relative pb/test.proto`

## Results(Locally)


```
goos: darwin
goarch: amd64
pkg: github.com/protobuf-benchmark
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz

                                     Num Iterations   Time/Function Call   Memory Allocation   Rate of Memory Allocation

BenchmarkMarshalJSON/Small-12            2662459         441.5 ns/op           80 B/op            1 allocs/op
BenchmarkMarshalJSON/Medium-12           1474044         793.8 ns/op           672 B/op           2 allocs/op
BenchmarkMarshalJSON/Large-12            247953          4315 ns/op            5250 B/op          2 allocs/op

BenchmarkMarshalProto/Small-12           6676744         178.5 ns/op           48 B/op            1 allocs/op
BenchmarkMarshalProto/Medium-12          4704790         235.8 ns/op           288 B/op           1 allocs/op
BenchmarkMarshalProto/Large-12           1983378         591.9 ns/op           3072 B/op          1 allocs/op

BenchmarkUnMarshalJSON/Small-12          707073          2192 ns/op            240 B/op           7 allocs/op
BenchmarkUnMarshalJSON/Medium-12         385875          3160 ns/op            480 B/op           7 allocs/op
BenchmarkUnMarshalJSON/Large-12          53250           22647 ns/op           3456 B/op          7 allocs/op

BenchmarkUnMarshalProto/Small-12         12717336        93.67 ns/op           0 B/op             0 allocs/op
BenchmarkUnMarshalProto/Medium-12        11514138        102.7 ns/op           0 B/op             0 allocs/op
BenchmarkUnMarshalProto/Large-12         12027140        96.40 ns/op           0 B/op             0 allocs/op
```