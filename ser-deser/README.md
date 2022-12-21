
# map/struct serialization deserialization benchmark

- encoding/json
- github.com/kokizzu/json5b/encoding/json5b
- github.com/goccy/go-json
- github.com/vmihailenco/msgpack/v5
- github.com/fxamacker/cbor/v2
- gopkg.in/yaml.v3
- github.com/ghodss/yaml
- github.com/goccy/go-yaml
- github.com/ugorji/go/codec
- github.com/json-iterator/go
- github.com/shamaton/msgpack/v2
- github.com/pquerna/ffjson

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
Benchmark_M2S_FxamackerCbor_MarshalUnmarshal-32        3801982     853.6 ns/op   112 B/op   8 allocs/op
Benchmark_S2M_FxamackerCbor_MarshalUnmarshal-32        3251618    1084 ns/op     444 B/op  11 allocs/op
Benchmark_M2S_EncodingJson_MarshalUnmarshal-32         2056944    1780 ns/op     600 B/op  16 allocs/op
Benchmark_S2M_EncodingJson_MarshalUnmarshal-32         1914165    1782 ns/op     688 B/op  18 allocs/op
Benchmark_M2S_KokizzuJson5b_MarshalUnmarshal-32        1987240    1711 ns/op     632 B/op  16 allocs/op
Benchmark_S2M_KokizzuJson5b_MarshalUnmarshal-32        1888774    1884 ns/op     960 B/op  20 allocs/op
Benchmark_M2S_GoccyGoJson_MarshalUnmarshal-32          6661932     517.4 ns/op    80 B/op   3 allocs/op
Benchmark_S2M_GoccyGoJson_MarshalUnmarshal-32          4716292     769.7 ns/op   513 B/op  12 allocs/op
Benchmark_M2S_JsonIteratorGo_MarshalUnmarshal-32       4364426     831.2 ns/op   188 B/op   8 allocs/op
Benchmark_S2M_JsonIteratorGo_MarshalUnmarshal-32       3411016    1072 ns/op     497 B/op  14 allocs/op
Benchmark_M2S_PquernaFfjson_MarshalUnmarshal-32        1537023    2446 ns/op     600 B/op  16 allocs/op
Benchmark_S2M_PquernaFfjson_MarshalUnmarshal-32        1559665    2338 ns/op     689 B/op  18 allocs/op
Benchmark_M2S_UngorjiGoCodec_CborEncodeDecode-32       1000000    3377 ns/op    4340 B/op  23 allocs/op
Benchmark_S2M_UngorjiGoCodec_CborEncodeDecode-32        823958    4117 ns/op    4888 B/op  34 allocs/op
Benchmark_M2S_UngorjiGocodec_BincEncodeDecode-32       1000000    3221 ns/op    4340 B/op  23 allocs/op
Benchmark_S2M_UngorjiGocodec_BincEncodeDecode-32        836186    4188 ns/op    4888 B/op  34 allocs/op
Benchmark_M2S_UngorjiGocodec_JsonEncodeDecode-32        941674    3932 ns/op    4956 B/op  25 allocs/op
Benchmark_S2M_UngorjiGocodec_JsonEncodeDecode-32        731814    4685 ns/op    5504 B/op  36 allocs/op
Benchmark_M2S_UngorjiGocodec_SimpleEncodeDecode-32      984807    3521 ns/op    4340 B/op  23 allocs/op
Benchmark_S2M_UngorjiGocodec_SimpleEncodeDecode-32      771837    4243 ns/op    4888 B/op  34 allocs/op
Benchmark_M2S_VmihailencoMspackV5_MarhsalUnmarshal-32  4572597     741.3 ns/op   188 B/op   5 allocs/op
Benchmark_S2M_VmihailencoMspackV5_MarhsalUnmarshal-32  3700621     951.5 ns/op   606 B/op  12 allocs/op
Benchmark_M2S_ShamatonMsgpackV2_MarshalUnmarshal-32    3544892    1001 ns/op     188 B/op   5 allocs/op
Benchmark_S2M_ShamatonMsgpackV2_MarshalUnmarshal-32    2758134    1265 ns/op     606 B/op  12 allocs/op
Benchmark_M2S_GoccyGoYaml_MarshalUnmarshal-32           194260   19278 ns/op    7815 B/op 214 allocs/op
Benchmark_S2M_GoccyGoYaml_MarshalUnmarshal-32           193806   18297 ns/op    7574 B/op 202 allocs/op
Benchmark_M2S_GopkgInYamlV3_MarshalUnmarshal-32         216579   16130 ns/op   14104 B/op  79 allocs/op
Benchmark_S2M_GopkgInYamlV3_MarshalUnmarshal-32         246944   14790 ns/op   14392 B/op  80 allocs/op
Benchmark_M2S_GhodssYaml_MarshalUnmarshal-32            152730   27637 ns/op   21344 B/op 160 allocs/op
Benchmark_S2M_GhodssYaml_MarshalUnmarshal-32            129429   27043 ns/op   21407 B/op 160 allocs/op
PASS
```

## FAQ: am I allowed to add codegen?

Yes

1. as long as it have `go:generate` and proper `go install` comment
2. as long as it can serialize and deserialize dynamic structure properly (map to struct, struct to map)