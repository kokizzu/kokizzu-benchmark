
# map/struct serialization deserialization benchmark

- encoding/json
- github.com/kokizzu/json5b/encoding/json5b
- github.com/goccy/go-json
- github.com/vmihailenco/msgpack/v5
- github.com/fxamacker/cbor/v2

## TL;DR

- `goccy/go-json` the fastest

## How to run?

```bash
go test -bench=.
go test -bench=. -benchmem
go test -bench=. -benchtime 3s
```

## How to add new test?

- create a function
- add on `TestVerify`
- create `Benchmark_X2X_RepoName_MethodName(b *testing.B)`
- add to `README.md`

```bash
go mod tidy
go test .
```

## Last Result?

```bash
go test -bench=. -benchmem -benchtime=3s
goos: linux
goarch: amd64
pkg: ser_deser         
Benchmark_M2S_FxamackerCbor_MarshalUnmarshal-32  3801982    853.6 ns/op   112 B/op   8 allocs/op
Benchmark_S2M_FxamackerCbor_MarshalUnmarshal-32  3251618   1084 ns/op     444 B/op  11 allocs/op
Benchmark_M2S_EncodingJson_MarshalUnmarshal-32   2056944   1780 ns/op     600 B/op  16 allocs/op
Benchmark_S2M_EncodingJson_MarshalUnmarshal-32   1914165   1782 ns/op     688 B/op  18 allocs/op
Benchmark_M2S_KokizzuJson5b_MarshalUnmarshal-32  1987240   1711 ns/op     632 B/op  16 allocs/op
Benchmark_S2M_KokizzuJson5b_MarshalUnmarshal-32  1888774   1884 ns/op     960 B/op  20 allocs/op
Benchmark_M2S_GoccyGoJson_MarshalUnmarshal-32    6661932    517.4 ns/op    80 B/op   3 allocs/op
Benchmark_S2M_GoccyGoJson_MarshalUnmarshal-32    4716292    769.7 ns/op   513 B/op  12 allocs/op
Benchmark_M2S_Vmihailenco_MarhsalUnmarshal-32    4572597    741.3 ns/op   188 B/op   5 allocs/op
Benchmark_S2M_Vmihailenco_MarhsalUnmarshal-32    3700621    951.5 ns/op   606 B/op  12 allocs/op
PASS
```