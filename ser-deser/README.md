
# Map/Struct Serialization-Deserialization Benchmark

https://kokizzu.blogspot.com/2022/12/map-to-struct-and-struct-to-map-golang.html

- encoding/json *
- github.com/kokizzu/json5b/encoding/json5b *
- github.com/goccy/go-json *
- github.com/vmihailenco/msgpack/v5
- github.com/fxamacker/cbor/v2
- gopkg.in/yaml.v3
- github.com/ghodss/yaml *
- github.com/goccy/go-yaml
- github.com/ugorji/go/codec
- github.com/json-iterator/go *
- github.com/shamaton/msgpack/v2
- github.com/pquerna/ffjson *
- go.mongodb.org/mongo-driver/bson
- github.com/BurntSushi/toml
- github.com/pelletier/go-toml/v2
- github.com/mitchellh/mapstructure
- github.com/naoina/toml
- github.com/hjson/hjson-go/v4 *
- github.com/d-o-n-u-t-s/lz4msgpack
- github.com/surrealdb/cork
- github.com/et-nik/binngo
- github.com/ichiban/tnetstrings
- github.com/bytedance/sonic *
- github.com/segmentio/encoding/json *
- github.com/go-json-experiment/json *
- github.com/sugawarayuuta/sonnet

* = not safe for >2^53

## TL;DR

- `goccy/go-json` is the average fastest on all use case (rank 1-2-2), but if you need to store integer more than the JSON standard number/double (2^53), then  `vmihailenco/msgpack/v5` on average is the fastest and best for map to struct use case (rank 1-2-5 excluding all json). `mitchellh/mapstructure` the best for struct to struct/map use case (rank 5-1-1 excluding all json).
- beware that some library not having same behavior as stdlib (`encoding/json`) even if they are claim to be (`goccy/go-json` is one that I found compatible, while `jsoniter` and easyjson` [didn't](//github.com/kokizzu/gotro/tree/master/W2)).
- These are the list of encoding that can serialize-deserialize everything properly (see `verifyCorrectness` for checking which edge case/data type/value they will fail):
  - vmihailenco/msgpack/v5
  - fxamacker/cbor/v2
  - goccyy/go-yaml
  - gopkg.in/yaml.v3
  - ungorji/go/codec/cbor
  - ungorji/go/codec/binc
  - ungorji/go/codec/simple
  - shamanon/msgpack/v2
  - mitchellh/mapstructure
  - DONUTS/lz4msgpack
  - surreald/cork
- only mapstructure can serialize-deserialize enum and anonymous struct properly

## How to run?

```bash
go test -bench=.
go test -bench=. -benchmem
go test -bench=. -benchtime 3s
go test -benchmem -benchtime=3s -bench=SpecificName

# normal flow
go test -benchmem -benchtime=3s -bench=. | tee last.txt && go run merge_group_sort_results.go
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

## Last Result? Go 1.20

best of N runs
```bash
                                   map to struct    total ns/op   B/op allocs/op
             M2S_GoccyGoJson_MarshalUnmarshal-32  7190296   500     88   3
          M2S_BytedanceSonic_MarshalUnmarshal-32  6610046   534    183   6
          M2S_JsonIteratorGo_MarshalUnmarshal-32  5754914   630    196   8
     M2S_SugawarayuutaSonnet_MarshalUnmarshal-32  5668898   661    309   9
   M2S_SegmentioEncodingJson_MarshalUnmarshal-32  5152272   679     60   3
     M2S_VmihailencoMspackV5_MarhsalUnmarshal-32  5474974   680    189   5
           M2S_FxamackerCbor_MarshalUnmarshal-32  4393375   754    120   8
       M2S_ShamatonMsgpackV2_MarshalUnmarshal-32  3660501  1050    956  15
        M2S_GopkgInMgoV2Bson_MarshalUnmarshal-32  3631568  1065    232  13
               M2S_SurrealdbCork_EncodeDecode-32  3276034  1127   1217   6
             M2S_MitchellhMapstructure_Decode-32  2743123  1366    720  18
    M2S_GoJsonExperimentJson_MarshalUnmarshal-32  2545573  1407    128   8
         M2S_MongoDriverBson_MarshalUnmarshal-32  2631243  1410    414  14
             M2S_EtNikBinngo_MarshalUnmarshal-32  2177028  1629    425  39
            M2S_EncodingJson_MarshalUnmarshal-32  2220351  1736    609  16
           M2S_KokizzuJson5b_MarshalUnmarshal-32  2088553  1738    632  16
           M2S_PquernaFfjson_MarshalUnmarshal-32  1809913  1833    609  16
          M2S_UngorjiGocodec_BincEncodeDecode-32  1414644  2554   4340  23
       M2S_PelletierGoTomlV2_MarshalUnmarshal-32  1357699  2603   1600  26
          M2S_UngorjiGoCodec_CborEncodeDecode-32  1331397  2635   4340  23
        M2S_UngorjiGocodec_SimpleEncodeDecode-32  1302295  2638   4340  23
          M2S_UngorjiGocodec_JsonEncodeDecode-32  1201784  3013   4956  25
      M2S_IchibanTnetstrings_MarshalUnmarshal-32   841498  4737   9329  47
           M2S_BurntSushiToml_EncodeUnmarshal-32   466965  8014   7960  71
          M2S_HjsonHjsonGoV4_MarshalUnmarshal-32   395845  9945   3940  78
           M2S_GopkgInYamlV3_MarshalUnmarshal-32   270820 13223  14112  80
        M2S_DONUTSLz4Msgpack_MarshalUnmarshal-32   237379 15144   1192  16
             M2S_GoccyGoYaml_MarshalUnmarshal-32   201076 15172   7825 214
              M2S_GhodssYaml_MarshalUnmarshal-32   141508 24973  21383 161
              M2S_NaoinaToml_MarshalUnmarshal-32    74859 54616 398546  77

                                   struct to map    total ns/op   B/op allocs/op
          S2M_BytedanceSonic_MarshalUnmarshal-32  5552894   653    645  11
     S2M_SugawarayuutaSonnet_MarshalUnmarshal-32  5073622   694    597  10
             S2M_MitchellhMapstructure_Decode-32  5045251   704    536  12
             S2M_GoccyGoJson_MarshalUnmarshal-32  5175942   720    522  12
   S2M_SegmentioEncodingJson_MarshalUnmarshal-32  4651759   799    442  10
          S2M_JsonIteratorGo_MarshalUnmarshal-32  4423248   814    506  14
     S2M_VmihailencoMspackV5_MarhsalUnmarshal-32  4112706   899    607  12
       S2M_ShamatonMsgpackV2_MarshalUnmarshal-32  3560764  1011    556  15
           S2M_FxamackerCbor_MarshalUnmarshal-32  3695653  1021    452  11
        S2M_GopkgInMgoV2Bson_MarshalUnmarshal-32  3230205  1102    528  15
               S2M_SurrealdbCork_EncodeDecode-32  3000410  1198   1611  12
    S2M_GoJsonExperimentJson_MarshalUnmarshal-32  2447996  1415    482  10
            S2M_EncodingJson_MarshalUnmarshal-32  1992739  1693    697  18
           S2M_PquernaFfjson_MarshalUnmarshal-32  2175736  1701    697  18
             S2M_EtNikBinngo_MarshalUnmarshal-32  2143041  1724    768  45
           S2M_KokizzuJson5b_MarshalUnmarshal-32  1947457  1800    960  20
         S2M_MongoDriverBson_MarshalUnmarshal-32  1950174  1818    760  18
       S2M_PelletierGoTomlV2_MarshalUnmarshal-32  1357994  2767   1800  30
          S2M_UngorjiGoCodec_CborEncodeDecode-32  1000000  3197   4888  34
          S2M_UngorjiGocodec_BincEncodeDecode-32   963999  3221   4888  34
        S2M_UngorjiGocodec_SimpleEncodeDecode-32  1000000  3265   4888  34
          S2M_UngorjiGocodec_JsonEncodeDecode-32   979472  3673   5504  36
      S2M_IchibanTnetstrings_MarshalUnmarshal-32   915262  4870   9553  45
           S2M_BurntSushiToml_EncodeUnmarshal-32   394602  8424   8231  73
          S2M_HjsonHjsonGoV4_MarshalUnmarshal-32   392012  8862   3877  66
           S2M_GopkgInYamlV3_MarshalUnmarshal-32   300262 12883  14400  81
             S2M_GoccyGoYaml_MarshalUnmarshal-32   260446 14117   7584 202
        S2M_DONUTSLz4Msgpack_MarshalUnmarshal-32   239043 15550    733  16
              S2M_GhodssYaml_MarshalUnmarshal-32   147769 23811  21447 161
              S2M_NaoinaToml_MarshalUnmarshal-32    63631 56368 398113  80

                                struct to struct    total ns/op   B/op allocs/op
             S2S_MitchellhMapstructure_Decode-32 23032750   167    136   4
             S2S_GoccyGoJson_MarshalUnmarshal-32 12658130   281    112   4
          S2S_BytedanceSonic_MarshalUnmarshal-32  8614616   396    213   7
       S2S_ShamatonMsgpackV2_MarshalUnmarshal-32  9214420   397    148   6
     S2S_SugawarayuutaSonnet_MarshalUnmarshal-32  8810007   400    245   6
          S2S_JsonIteratorGo_MarshalUnmarshal-32  7616412   476     92   6
           S2S_FxamackerCbor_MarshalUnmarshal-32  7092066   498     80   5
   S2S_SegmentioEncodingJson_MarshalUnmarshal-32  6367107   564     84   4
     S2S_VmihailencoMspackV5_MarhsalUnmarshal-32  4708036   686    213   6
        S2S_GopkgInMgoV2Bson_MarshalUnmarshal-32  5171426   723    144   9
    S2S_GoJsonExperimentJson_MarshalUnmarshal-32  4504526   824    104   4
         S2S_MongoDriverBson_MarshalUnmarshal-32  3810334   964    321   8
            S2S_EncodingJson_MarshalUnmarshal-32  3190252  1103    304   9
           S2S_PquernaFfjson_MarshalUnmarshal-32  3009652  1163    304   9
           S2S_KokizzuJson5b_MarshalUnmarshal-32  2886364  1244    504   9
               S2S_SurrealdbCork_EncodeDecode-32  2714652  1356   1241   7
             S2S_EtNikBinngo_MarshalUnmarshal-32  2343198  1580    400  41
       S2S_PelletierGoTomlV2_MarshalUnmarshal-32  1612135  2207   1440  22
          S2S_UngorjiGoCodec_CborEncodeDecode-32  1465347  2416   4364  24
        S2S_UngorjiGocodec_SimpleEncodeDecode-32  1372903  2580   4364  24
          S2S_UngorjiGocodec_BincEncodeDecode-32  1363026  2583   4364  24
          S2S_UngorjiGocodec_JsonEncodeDecode-32  1000000  3092   4980  26
      S2S_IchibanTnetstrings_MarshalUnmarshal-32   802863  4774   9286  46
           S2S_BurntSushiToml_EncodeUnmarshal-32   458978  8448   7919  72
          S2S_HjsonHjsonGoV4_MarshalUnmarshal-32   362943  9530   4588  79
           S2S_GopkgInYamlV3_MarshalUnmarshal-32   296682 12103  14016  76
        S2S_DONUTSLz4Msgpack_MarshalUnmarshal-32   237526 14740    231   7
             S2S_GoccyGoYaml_MarshalUnmarshal-32   259622 14772   7986 208
              S2S_GhodssYaml_MarshalUnmarshal-32   136220 23678  21079 154
              S2S_NaoinaToml_MarshalUnmarshal-32    60412 60310 399006  83
```

## Last Result? Go 1.19

best of N runs

```bash
                                   map to struct    total ns/op   B/op allocs/op
             M2S_GoccyGoJson_MarshalUnmarshal-32  6661932   517     80   3
          M2S_BytedanceSonic_MarshalUnmarshal-32  4952025   640    193   6
          M2S_JsonIteratorGo_MarshalUnmarshal-32  4892611   724    196   8
     M2S_VmihailencoMspackV5_MarhsalUnmarshal-32  4572597   741    188   5
           M2S_FxamackerCbor_MarshalUnmarshal-32  4418558   799    120   8
   M2S_SegmentioEncodingJson_MarshalUnmarshal-32  3559434   858     60   3
        M2S_GopkgInMgoV2Bson_MarshalUnmarshal-32  3288061   971    232  13
               M2S_SurrealdbCork_EncodeDecode-32  3080282  1080   1217   6
       M2S_ShamatonMsgpackV2_MarshalUnmarshal-32  3062677  1161    956  15
             M2S_MitchellhMapstructure_Decode-32  2487428  1395    720  18
         M2S_MongoDriverBson_MarshalUnmarshal-32  2477983  1459    414  14
    M2S_GoJsonExperimentJson_MarshalUnmarshal-32  2041227  1637    128   8
           M2S_KokizzuJson5b_MarshalUnmarshal-32  1987240  1711    632  16
            M2S_EncodingJson_MarshalUnmarshal-32  2056944  1780    600  16
             M2S_EtNikBinngo_MarshalUnmarshal-32  1985595  1857    425  39
           M2S_PquernaFfjson_MarshalUnmarshal-32  1739968  1986    609  16
          M2S_UngorjiGocodec_BincEncodeDecode-32  1401453  2582   4340  23
          M2S_UngorjiGoCodec_CborEncodeDecode-32  1304828  2636   4340  23
       M2S_PelletierGoTomlV2_MarshalUnmarshal-32  1284037  2787   1600  27
        M2S_UngorjiGocodec_SimpleEncodeDecode-32  1295926  2810   4340  23
          M2S_UngorjiGocodec_JsonEncodeDecode-32  1000000  3028   4956  25
      M2S_IchibanTnetstrings_MarshalUnmarshal-32   749947  5056   9329  48
           M2S_BurntSushiToml_EncodeUnmarshal-32   425335  8065   7958  71
          M2S_HjsonHjsonGoV4_MarshalUnmarshal-32   355784 10870   3936  78
           M2S_GopkgInYamlV3_MarshalUnmarshal-32   271190 13524  14112  80
        M2S_DONUTSLz4Msgpack_MarshalUnmarshal-32   240619 15498   1264  16
             M2S_GoccyGoYaml_MarshalUnmarshal-32   214776 16192   7821 214
              M2S_GhodssYaml_MarshalUnmarshal-32   156412 23347  21378 161
              M2S_NaoinaToml_MarshalUnmarshal-32    57607 58331 398544  77

                                   struct to map    total ns/op   B/op allocs/op
             S2M_MitchellhMapstructure_Decode-32  5055402   716    536  12
             S2M_GoccyGoJson_MarshalUnmarshal-32  4660224   747    522  12
          S2M_BytedanceSonic_MarshalUnmarshal-32  4376922   821    669  11
          S2M_JsonIteratorGo_MarshalUnmarshal-32  4283262   835    505  14
     S2M_VmihailencoMspackV5_MarhsalUnmarshal-32  4009863   908    607  12
           S2M_FxamackerCbor_MarshalUnmarshal-32  3562352  1023    452  11
   S2M_SegmentioEncodingJson_MarshalUnmarshal-32  3752232  1077    442  10
       S2M_ShamatonMsgpackV2_MarshalUnmarshal-32  3180010  1089    556  15
        S2M_GopkgInMgoV2Bson_MarshalUnmarshal-32  3047396  1145    528  15
               S2M_SurrealdbCork_EncodeDecode-32  2976328  1196   1611  12
    S2M_GoJsonExperimentJson_MarshalUnmarshal-32  2049212  1602    482  10
            S2M_EncodingJson_MarshalUnmarshal-32  1914165  1782    688  18
           S2M_PquernaFfjson_MarshalUnmarshal-32  1911950  1845    697  18
             S2M_EtNikBinngo_MarshalUnmarshal-32  1948802  1859    768  45
           S2M_KokizzuJson5b_MarshalUnmarshal-32  1888774  1884    960  20
         S2M_MongoDriverBson_MarshalUnmarshal-32  1857649  1995    759  18
       S2M_PelletierGoTomlV2_MarshalUnmarshal-32  1244012  2864   1800  31
          S2M_UngorjiGocodec_BincEncodeDecode-32  1000000  3234   4888  34
          S2M_UngorjiGoCodec_CborEncodeDecode-32   989671  3358   4888  34
        S2M_UngorjiGocodec_SimpleEncodeDecode-32  1000000  3400   4888  34
          S2M_UngorjiGocodec_JsonEncodeDecode-32   912512  3639   5504  36
      S2M_IchibanTnetstrings_MarshalUnmarshal-32   776796  4744   9561  46
           S2M_BurntSushiToml_EncodeUnmarshal-32   447216  8538   8231  73
          S2M_HjsonHjsonGoV4_MarshalUnmarshal-32   389476  9416   3868  66
           S2M_GopkgInYamlV3_MarshalUnmarshal-32   315939 13338  14400  81
        S2M_DONUTSLz4Msgpack_MarshalUnmarshal-32   220008 14281    715  16
             S2M_GoccyGoYaml_MarshalUnmarshal-32   230042 14919   7580 202
              S2M_GhodssYaml_MarshalUnmarshal-32   151023 22682  21441 161
              S2M_NaoinaToml_MarshalUnmarshal-32    60916 52047 398112  80

                                struct to struct    total ns/op   B/op allocs/op
             S2S_MitchellhMapstructure_Decode-32 17987224   181    136   4
             S2S_GoccyGoJson_MarshalUnmarshal-32 12046497   317    112   4
       S2S_ShamatonMsgpackV2_MarshalUnmarshal-32  7897488   458    148   6
          S2S_BytedanceSonic_MarshalUnmarshal-32  7702740   469    218   7
          S2S_JsonIteratorGo_MarshalUnmarshal-32  7853592   494     92   6
           S2S_FxamackerCbor_MarshalUnmarshal-32  7038808   511     80   5
        S2S_GopkgInMgoV2Bson_MarshalUnmarshal-32  5105343   715    144   9
   S2S_SegmentioEncodingJson_MarshalUnmarshal-32  4891458   763     84   4
     S2S_VmihailencoMspackV5_MarhsalUnmarshal-32  4549700   818    213   6
         S2S_MongoDriverBson_MarshalUnmarshal-32  3560946  1019    321   8
    S2S_GoJsonExperimentJson_MarshalUnmarshal-32  3451814  1026    104   4
            S2S_EncodingJson_MarshalUnmarshal-32  2679778  1127    304   9
           S2S_PquernaFfjson_MarshalUnmarshal-32  2734357  1330    304   9
           S2S_KokizzuJson5b_MarshalUnmarshal-32  2594728  1343    504   9
               S2S_SurrealdbCork_EncodeDecode-32  2555745  1397   1241   7
             S2S_EtNikBinngo_MarshalUnmarshal-32  1995468  1840    400  41
       S2S_PelletierGoTomlV2_MarshalUnmarshal-32  1460683  2459   1440  23
        S2S_UngorjiGocodec_SimpleEncodeDecode-32  1205648  2919   4364  24
          S2S_UngorjiGoCodec_CborEncodeDecode-32  1290734  2920   4364  24
          S2S_UngorjiGocodec_BincEncodeDecode-32  1207327  3007   4364  24
          S2S_UngorjiGocodec_JsonEncodeDecode-32  1000000  3223   4980  26
      S2S_IchibanTnetstrings_MarshalUnmarshal-32   722493  4950   9289  47
           S2S_BurntSushiToml_EncodeUnmarshal-32   398366  8458   7918  72
          S2S_HjsonHjsonGoV4_MarshalUnmarshal-32   304369 11189   4578  79
        S2S_DONUTSLz4Msgpack_MarshalUnmarshal-32   282505 12215    237   7
           S2S_GopkgInYamlV3_MarshalUnmarshal-32   279332 12541  14016  76
             S2S_GoccyGoYaml_MarshalUnmarshal-32   211952 15542   7982 208
              S2S_GhodssYaml_MarshalUnmarshal-32   160660 23148  21073 154
              S2S_NaoinaToml_MarshalUnmarshal-32    64468 59672 399065  83

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
- muzzammil.xyz/jsonc (only decoder)
- github.com/wI2L/jettison (only encoder)
- github.com/minio/simdjson-go (need manual convert on each structure)

## TODO

- for libraries that default doesn't have buffer pool, clone benchmark function and add a private `sync.Pool`
- github.com/kanosaki/flexbuffers
- github.com/mojura/enkodo
