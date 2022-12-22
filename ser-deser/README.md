
# Map/Struct Serialization-Deserialization Benchmark

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
- go.mongodb.org/mongo-driver/bson
- github.com/BurntSushi/toml
- github.com/pelletier/go-toml/v2
- github.com/mitchellh/mapstructure
- github.com/naoina/toml
- github.com/hjson/hjson-go/v4
- github.com/d-o-n-u-t-s/lz4msgpack
- github.com/surrealdb/cork
- github.com/et-nik/binngo
- github.com/ichiban/tnetstrings

## TL;DR

- `goccy/go-json` the fastest

## How to run?

```bash
go test -bench=.
go test -bench=. -benchmem
go test -bench=. -benchtime 3s
go test -benchmem -benchtime=3s -bench=SpecificName
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
go test -benchmem -benchtime=3s -bench=.
goos: linux
goarch: amd64
pkg: ser_deser 
 
Benchmark_M2S_GoccyGoJson_MarshalUnmarshal-32          6661932    517.4 ns/op      80 B/op    3 allocs/op
Benchmark_M2S_VmihailencoMspackV5_MarhsalUnmarshal-32  4572597    741.3 ns/op     188 B/op    5 allocs/op
Benchmark_M2S_JsonIteratorGo_MarshalUnmarshal-32       4364426    831.2 ns/op     188 B/op    8 allocs/op
Benchmark_M2S_FxamackerCbor_MarshalUnmarshal-32        3801982    853.6 ns/op     112 B/op    8 allocs/op
Benchmark_M2S_SurrealdbCork_EncodeDecode-32            3105571   1155   ns/op    1217 B/op    6 allocs/op
Benchmark_M2S_ShamatonMsgpackV2_MarshalUnmarshal-32    2567893   1379   ns/op     956 B/op   15 allocs/op
Benchmark_M2S_KokizzuJson5b_MarshalUnmarshal-32        1987240   1711   ns/op     632 B/op   16 allocs/op
Benchmark_M2S_MongoDriverBson_MarshalUnmarshal-32      2083712   1719   ns/op     413 B/op   14 allocs/op
Benchmark_M2S_MitchellhMapstructure_Decode-32          1988281   1750   ns/op     720 B/op   18 allocs/op
Benchmark_M2S_EncodingJson_MarshalUnmarshal-32         2056944   1780   ns/op     600 B/op   16 allocs/op
Benchmark_M2S_EtNikBinngo_MarshalUnmarshal-32          1813059   1965   ns/op     425 B/op   39 allocs/op
Benchmark_M2S_PquernaFfjson_MarshalUnmarshal-32        1537023   2446   ns/op     600 B/op   16 allocs/op
Benchmark_M2S_UngorjiGocodec_BincEncodeDecode-32       1000000   3221   ns/op    4340 B/op   23 allocs/op
Benchmark_M2S_UngorjiGoCodec_CborEncodeDecode-32       1000000   3377   ns/op    4340 B/op   23 allocs/op
Benchmark_M2S_PelletierGoTomlV2_MarshalUnmarshal-32    1000000   3392   ns/op    1600 B/op   27 allocs/op
Benchmark_M2S_UngorjiGocodec_SimpleEncodeDecode-32      984807   3521   ns/op    4340 B/op   23 allocs/op
Benchmark_M2S_UngorjiGocodec_JsonEncodeDecode-32        941674   3932   ns/op    4956 B/op   25 allocs/op
Benchmark_M2S_IchibanTnetstrings_MarshalUnmarshal-32    773697   5747   ns/op    9328 B/op   48 allocs/op
Benchmark_M2S_BurntSushiToml_EncodeUnmarshal-32         361131   9382   ns/op    7950 B/op   70 allocs/op
Benchmark_M2S_HjsonHjsonGoV4_MarshalUnmarshal-32        369955  12260   ns/op    3904 B/op   77 allocs/op
Benchmark_M2S_GopkgInYamlV3_MarshalUnmarshal-32         216579  16130   ns/op   14104 B/op   79 allocs/op
Benchmark_M2S_DONUTSLz4Msgpack_MarshalUnmarshal-32      247033  17272   ns/op    1271 B/op   16 allocs/op
Benchmark_M2S_GoccyGoYaml_MarshalUnmarshal-32           194260  19278   ns/op    7815 B/op  214 allocs/op
Benchmark_M2S_GhodssYaml_MarshalUnmarshal-32            152730  27637   ns/op   21344 B/op  160 allocs/op
Benchmark_M2S_NaoinaToml_MarshalUnmarshal-32             42057  88518   ns/op  398393 B/op   77 allocs/op


Benchmark_S2M_GoccyGoJson_MarshalUnmarshal-32          4716292    769.7 ns/op     513 B/op   12 allocs/op
Benchmark_S2M_MitchellhMapstructure_Decode-32          4160382    878.7 ns/op     536 B/op   12 allocs/op
Benchmark_S2M_VmihailencoMspackV5_MarhsalUnmarshal-32  3700621    951.5 ns/op     606 B/op   12 allocs/op
Benchmark_S2M_JsonIteratorGo_MarshalUnmarshal-32       3411016   1072   ns/op     497 B/op   14 allocs/op
Benchmark_S2M_FxamackerCbor_MarshalUnmarshal-32        3251618   1084   ns/op     444 B/op   11 allocs/op
Benchmark_S2M_SurrealdbCork_EncodeDecode-32            2806642   1289   ns/op    1611 B/op   12 allocs/op
Benchmark_S2M_ShamatonMsgpackV2_MarshalUnmarshal-32    2669648   1291   ns/op     556 B/op   15 allocs/op
Benchmark_S2M_EncodingJson_MarshalUnmarshal-32         1914165   1782   ns/op     688 B/op   18 allocs/op
Benchmark_S2M_KokizzuJson5b_MarshalUnmarshal-32        1888774   1884   ns/op     960 B/op   20 allocs/op
Benchmark_S2M_EtNikBinngo_MarshalUnmarshal-32          1586864   2260   ns/op     768 B/op   45 allocs/op
Benchmark_S2M_PquernaFfjson_MarshalUnmarshal-32        1559665   2338   ns/op     689 B/op   18 allocs/op
Benchmark_S2M_MongoDriverBson_MarshalUnmarshal-32      1476170   2668   ns/op     759 B/op   18 allocs/op
Benchmark_S2M_PelletierGoTomlV2_MarshalUnmarshal-32    1000000   3687   ns/op    1800 B/op   31 allocs/op
Benchmark_S2M_UngorjiGoCodec_CborEncodeDecode-32        823958   4117   ns/op    4888 B/op   34 allocs/op
Benchmark_S2M_UngorjiGocodec_BincEncodeDecode-32        836186   4188   ns/op    4888 B/op   34 allocs/op
Benchmark_S2M_UngorjiGocodec_SimpleEncodeDecode-32      771837   4243   ns/op    4888 B/op   34 allocs/op
Benchmark_S2M_UngorjiGocodec_JsonEncodeDecode-32        731814   4685   ns/op    5504 B/op   36 allocs/op
Benchmark_S2M_IchibanTnetstrings_MarshalUnmarshal-32    612975   5449   ns/op    9553 B/op   46 allocs/op
Benchmark_S2M_BurntSushiToml_EncodeUnmarshal-32         368166  10156   ns/op    8222 B/op   72 allocs/op
Benchmark_S2M_HjsonHjsonGoV4_MarshalUnmarshal-32        325735  10531   ns/op    3842 B/op   65 allocs/op
Benchmark_S2M_GopkgInYamlV3_MarshalUnmarshal-32         246944  14790   ns/op   14392 B/op   80 allocs/op
Benchmark_S2M_DONUTSLz4Msgpack_MarshalUnmarshal-32      252949  15064   ns/op     730 B/op   16 allocs/op
Benchmark_S2M_GoccyGoYaml_MarshalUnmarshal-32           193806  18297   ns/op    7574 B/op  202 allocs/op
Benchmark_S2M_GhodssYaml_MarshalUnmarshal-32            129429  27043   ns/op   21407 B/op  160 allocs/op
Benchmark_S2M_NaoinaToml_MarshalUnmarshal-32             58723  62006   ns/op  397936 B/op   75 allocs/op

PASS
```

## FAQ: is it ok to add codegen format?

Yes

1. as long as it have `go:generate THE_CODEGEN_COMMAND` and proper `go:generate go install` comment
2. as long as it can serialize and deserialize dynamic structure properly (map to struct, struct to map)

## FAQ: failed attempts?

- github.com/hprose/hprose-golang/v2
- github.com/ikkerens/ikeapack
- github.com/glycerine/zebrapack
- github.com/niubaoshu/gotiny
- github.com/jinzhu/copier
- github.com/mashingan/smapping
- github.com/davecgh/go-xdr/xdr2
- github.com/tinylib/msgp
- github.com/mailru/easyjson
- github.com/theodesp/binpack
- github.com/kelindar/binary
- github.com/polydawn/refmt
- github.com/renproject/surge
- github.com/vipally/binary
- github.com/the729/lcs
- github.com/cristalhq/bencode
- github.com/mprot/msgpack-go
- github.com/arloliu/jsonpack
- github.com/Itay2805/rawbin
- github.com/superp00t/etc
- github.com/wwalexander/go-bencode
- github.com/fengyoulin/schema
- github.com/near/borsh-go
- github.com/go-restruct/restruct

## TODO

- for libraries that default doesn't have buffer pool, clone benchmark function and add a private `sync.Pool`
