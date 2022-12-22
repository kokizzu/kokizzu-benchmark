
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
          M2S_JsonIteratorGo_MarshalUnmarshal-32  4364426   831    188   8
           M2S_FxamackerCbor_MarshalUnmarshal-32  3801982   854    112   8
               M2S_SurrealdbCork_EncodeDecode-32  3105571  1155   1217   6
        M2S_GopkgInMgoV2Bson_MarshalUnmarshal-32  3125451  1181    232  13
       M2S_ShamatonMsgpackV2_MarshalUnmarshal-32  2567893  1379    956  15
           M2S_KokizzuJson5b_MarshalUnmarshal-32  1987240  1711    632  16
         M2S_MongoDriverBson_MarshalUnmarshal-32  2083712  1719    413  14
             M2S_MitchellhMapstructure_Decode-32  1988281  1750    720  18
            M2S_EncodingJson_MarshalUnmarshal-32  2056944  1780    600  16
             M2S_EtNikBinngo_MarshalUnmarshal-32  1813059  1965    425  39
           M2S_PquernaFfjson_MarshalUnmarshal-32  1537023  2446    600  16
          M2S_UngorjiGocodec_BincEncodeDecode-32  1000000  3221   4340  23
          M2S_UngorjiGoCodec_CborEncodeDecode-32  1000000  3224   4340  23
        M2S_UngorjiGocodec_SimpleEncodeDecode-32  1000000  3319   4340  23
       M2S_PelletierGoTomlV2_MarshalUnmarshal-32  1000000  3392   1600  27
          M2S_UngorjiGocodec_JsonEncodeDecode-32  1000000  3695   4956  25
      M2S_IchibanTnetstrings_MarshalUnmarshal-32   773697  5747   9328  48
           M2S_BurntSushiToml_EncodeUnmarshal-32   361131  9382   7950  70
          M2S_HjsonHjsonGoV4_MarshalUnmarshal-32   369955 12260   3904  77
           M2S_GopkgInYamlV3_MarshalUnmarshal-32   260914 15408  14112  80
        M2S_DONUTSLz4Msgpack_MarshalUnmarshal-32   216382 16634   1249  16
             M2S_GoccyGoYaml_MarshalUnmarshal-32   194260 19278   7815 214
              M2S_GhodssYaml_MarshalUnmarshal-32   152730 27637  21344 160
              M2S_NaoinaToml_MarshalUnmarshal-32    48402 67834 398544  77

                                   struct to map    total ns/op   B/op allocs/op
             S2M_GoccyGoJson_MarshalUnmarshal-32  4716292   770    513  12
             S2M_MitchellhMapstructure_Decode-32  4160382   879    536  12
     S2M_VmihailencoMspackV5_MarhsalUnmarshal-32  3700621   952    606  12
          S2M_JsonIteratorGo_MarshalUnmarshal-32  3411016  1072    497  14
           S2M_FxamackerCbor_MarshalUnmarshal-32  3251618  1084    444  11
               S2M_SurrealdbCork_EncodeDecode-32  2806642  1289   1611  12
       S2M_ShamatonMsgpackV2_MarshalUnmarshal-32  2669648  1291    556  15
        S2M_GopkgInMgoV2Bson_MarshalUnmarshal-32  2769507  1299    528  15
            S2M_EncodingJson_MarshalUnmarshal-32  1914165  1782    688  18
           S2M_KokizzuJson5b_MarshalUnmarshal-32  1888774  1884    960  20
         S2M_MongoDriverBson_MarshalUnmarshal-32  1685284  2215    759  18
             S2M_EtNikBinngo_MarshalUnmarshal-32  1586864  2260    768  45
           S2M_PquernaFfjson_MarshalUnmarshal-32  1461468  2324    697  18
       S2M_PelletierGoTomlV2_MarshalUnmarshal-32  1000000  3687   1800  31
          S2M_UngorjiGocodec_BincEncodeDecode-32  1000000  4046   4888  34
          S2M_UngorjiGoCodec_CborEncodeDecode-32   823958  4117   4888  34
          S2M_UngorjiGocodec_JsonEncodeDecode-32  1000000  4219   5504  36
        S2M_UngorjiGocodec_SimpleEncodeDecode-32   771837  4243   4888  34
      S2M_IchibanTnetstrings_MarshalUnmarshal-32   612975  5449   9553  46
           S2M_BurntSushiToml_EncodeUnmarshal-32   368166 10156   8222  72
          S2M_HjsonHjsonGoV4_MarshalUnmarshal-32   325735 10531   3842  65
           S2M_GopkgInYamlV3_MarshalUnmarshal-32   246944 14790  14392  80
        S2M_DONUTSLz4Msgpack_MarshalUnmarshal-32   252949 15064    730  16
             S2M_GoccyGoYaml_MarshalUnmarshal-32   189339 17815   7580 202
              S2M_GhodssYaml_MarshalUnmarshal-32   129429 27043  21407 160
              S2M_NaoinaToml_MarshalUnmarshal-32    58723 62006 397936  75

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
