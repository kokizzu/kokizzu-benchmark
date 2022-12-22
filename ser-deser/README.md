
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

- `goccy/go-json` the fastest, but if you need to store integer more than the JSON standard (2^53), then  `vmihailenco/msgpack/v5` is the fastest.

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

                                   map to struct    total ns/op   B/op allocs/op
             M2S_GoccyGoJson_MarshalUnmarshal-32  6661932   517     80   3
     M2S_VmihailencoMspackV5_MarhsalUnmarshal-32  4572597   741    188   5
          M2S_JsonIteratorGo_MarshalUnmarshal-32  5025380   759    196   8
           M2S_FxamackerCbor_MarshalUnmarshal-32  4545553   826    120   8
        M2S_GopkgInMgoV2Bson_MarshalUnmarshal-32  3227905  1092    232  13
               M2S_SurrealdbCork_EncodeDecode-32  3105571  1155   1217   6
       M2S_ShamatonMsgpackV2_MarshalUnmarshal-32  3062677  1161    956  15
         M2S_MongoDriverBson_MarshalUnmarshal-32  2375272  1525    414  14
           M2S_KokizzuJson5b_MarshalUnmarshal-32  1987240  1711    632  16
            M2S_EncodingJson_MarshalUnmarshal-32  2056944  1780    600  16
             M2S_EtNikBinngo_MarshalUnmarshal-32  1867093  1918    425  39
           M2S_PquernaFfjson_MarshalUnmarshal-32  1830831  2052    609  16
          M2S_UngorjiGocodec_BincEncodeDecode-32  1401453  2582   4340  23
          M2S_UngorjiGoCodec_CborEncodeDecode-32  1304828  2636   4340  23
       M2S_PelletierGoTomlV2_MarshalUnmarshal-32  1220426  2874   1600  27
        M2S_UngorjiGocodec_SimpleEncodeDecode-32  1291774  2918   4340  23
          M2S_UngorjiGocodec_JsonEncodeDecode-32  1000000  3028   4956  25
      M2S_IchibanTnetstrings_MarshalUnmarshal-32   656162  5263   9329  48
           M2S_BurntSushiToml_EncodeUnmarshal-32   415000  8763   7958  71
          M2S_HjsonHjsonGoV4_MarshalUnmarshal-32   306690 11155   3931  78
           M2S_GopkgInYamlV3_MarshalUnmarshal-32   243672 14200  14112  80
        M2S_DONUTSLz4Msgpack_MarshalUnmarshal-32   236578 15933   1287  16
             M2S_GoccyGoYaml_MarshalUnmarshal-32   208022 17474   7821 214
              M2S_GhodssYaml_MarshalUnmarshal-32   142357 25266  21377 161
              M2S_NaoinaToml_MarshalUnmarshal-32    52935 63322 398544  77

                                   struct to map    total ns/op   B/op allocs/op
             S2M_GoccyGoJson_MarshalUnmarshal-32  4782322   752    522  12
             S2M_MitchellhMapstructure_Decode-32  4821429   765    536  12
          S2M_JsonIteratorGo_MarshalUnmarshal-32  3857022   904    505  14
     S2M_VmihailencoMspackV5_MarhsalUnmarshal-32  3748212   939    607  12
           S2M_FxamackerCbor_MarshalUnmarshal-32  3317041  1079    452  11
       S2M_ShamatonMsgpackV2_MarshalUnmarshal-32  3261236  1188    556  15
        S2M_GopkgInMgoV2Bson_MarshalUnmarshal-32  2824772  1215    528  15
               S2M_SurrealdbCork_EncodeDecode-32  2806642  1289   1611  12
            S2M_EncodingJson_MarshalUnmarshal-32  1914165  1782    688  18
           S2M_KokizzuJson5b_MarshalUnmarshal-32  1888774  1884    960  20
             S2M_EtNikBinngo_MarshalUnmarshal-32  1806198  1934    768  45
           S2M_PquernaFfjson_MarshalUnmarshal-32  1854583  1951    697  18
         S2M_MongoDriverBson_MarshalUnmarshal-32  1742617  2058    760  18
       S2M_PelletierGoTomlV2_MarshalUnmarshal-32  1000000  3072   1800  31
          S2M_UngorjiGocodec_BincEncodeDecode-32  1000000  3234   4888  34
          S2M_UngorjiGoCodec_CborEncodeDecode-32  1000000  3448   4888  34
        S2M_UngorjiGocodec_SimpleEncodeDecode-32  1000000  3597   4888  34
          S2M_UngorjiGocodec_JsonEncodeDecode-32   986792  3884   5504  36
      S2M_IchibanTnetstrings_MarshalUnmarshal-32   696022  4916   9561  46
           S2M_BurntSushiToml_EncodeUnmarshal-32   399859  9023   8231  73
          S2M_HjsonHjsonGoV4_MarshalUnmarshal-32   353668  9813   3871  66
           S2M_GopkgInYamlV3_MarshalUnmarshal-32   272811 13918  14400  81
        S2M_DONUTSLz4Msgpack_MarshalUnmarshal-32   224689 14576    732  16
             S2M_GoccyGoYaml_MarshalUnmarshal-32   235486 15645   7580 202
              S2M_GhodssYaml_MarshalUnmarshal-32   147253 25277  21442 161
              S2M_NaoinaToml_MarshalUnmarshal-32    59755 60212 398112  80

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
- github.com/dedis/protobuf
- github.com/linkedin/goavro/v2
- encoding/asn1
- encoding/gob
- encoding/xml

## TODO

- for libraries that default doesn't have buffer pool, clone benchmark function and add a private `sync.Pool`
