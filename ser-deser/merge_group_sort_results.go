package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
)

// merge result from multiple runs

var results = `

Benchmark_M2S_GoccyGoJson_MarshalUnmarshal-32          6661932    517.4 ns/op      80 B/op    3 allocs/op
Benchmark_M2S_VmihailencoMspackV5_MarhsalUnmarshal-32  4572597    741.3 ns/op     188 B/op    5 allocs/op
Benchmark_M2S_JsonIteratorGo_MarshalUnmarshal-32       4364426    831.2 ns/op     188 B/op    8 allocs/op
Benchmark_M2S_FxamackerCbor_MarshalUnmarshal-32        3801982    853.6 ns/op     112 B/op    8 allocs/op
Benchmark_M2S_SurrealdbCork_EncodeDecode-32            3105571   1155   ns/op    1217 B/op    6 allocs/op
Benchmark_M2S_GopkgInMgoV2Bson_MarshalUnmarshal-32     3125451   1181   ns/op     232 B/op   13 allocs/op
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
Benchmark_S2M_GopkgInMgoV2Bson_MarshalUnmarshal-32     2769507   1299   ns/op     528 B/op   15 allocs/op
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

Benchmark_M2S_EncodingJson_MarshalUnmarshal-32                   1430934              2619 ns/op             609 B/op         16 allocs/op
Benchmark_S2M_EncodingJson_MarshalUnmarshal-32                   1441177              2462 ns/op             697 B/op         18 allocs/op
Benchmark_M2S_KokizzuJson5b_MarshalUnmarshal-32                  1561827              2230 ns/op             632 B/op         16 allocs/op
Benchmark_S2M_KokizzuJson5b_MarshalUnmarshal-32                  1414522              2460 ns/op             960 B/op         20 allocs/op
Benchmark_M2S_GoccyGoJson_MarshalUnmarshal-32                    5243158               650.2 ns/op            88 B/op          3 allocs/op
Benchmark_S2M_GoccyGoJson_MarshalUnmarshal-32                    3362835              1048 ns/op             522 B/op         12 allocs/op
Benchmark_M2S_JsonIteratorGo_MarshalUnmarshal-32                 3774579               949.4 ns/op           196 B/op          8 allocs/op
Benchmark_S2M_JsonIteratorGo_MarshalUnmarshal-32                 3068481              1106 ns/op             505 B/op         14 allocs/op
Benchmark_M2S_PquernaFfjson_MarshalUnmarshal-32                  1545321              2482 ns/op             609 B/op         16 allocs/op
Benchmark_S2M_PquernaFfjson_MarshalUnmarshal-32                  1461468              2324 ns/op             697 B/op         18 allocs/op
Benchmark_M2S_MongoDriverBson_MarshalUnmarshal-32                1952024              1861 ns/op             414 B/op         14 allocs/op
Benchmark_S2M_MongoDriverBson_MarshalUnmarshal-32                1685284              2215 ns/op             759 B/op         18 allocs/op
Benchmark_M2S_HjsonHjsonGoV4_MarshalUnmarshal-32                  266498             12623 ns/op            3932 B/op         78 allocs/op
Benchmark_S2M_HjsonHjsonGoV4_MarshalUnmarshal-32                  291206             11345 ns/op            3863 B/op         66 allocs/op
Benchmark_M2S_GopkgInMgoV2Bson_MarshalUnmarshal-32               2839399              1318 ns/op             232 B/op         13 allocs/op
Benchmark_S2M_GopkgInMgoV2Bson_MarshalUnmarshal-32               2379944              1548 ns/op             528 B/op         15 allocs/op
Benchmark_M2S_UngorjiGoCodec_CborEncodeDecode-32                 1000000              3224 ns/op            4340 B/op         23 allocs/op
Benchmark_S2M_UngorjiGoCodec_CborEncodeDecode-32                  914655              4246 ns/op            4888 B/op         34 allocs/op
Benchmark_M2S_UngorjiGocodec_BincEncodeDecode-32                 1000000              3406 ns/op            4340 B/op         23 allocs/op
Benchmark_S2M_UngorjiGocodec_BincEncodeDecode-32                 1000000              4046 ns/op            4888 B/op         34 allocs/op
Benchmark_M2S_UngorjiGocodec_JsonEncodeDecode-32                 1000000              3695 ns/op            4956 B/op         25 allocs/op
Benchmark_S2M_UngorjiGocodec_JsonEncodeDecode-32                 1000000              4219 ns/op            5504 B/op         36 allocs/op
Benchmark_M2S_UngorjiGocodec_SimpleEncodeDecode-32               1000000              3319 ns/op            4340 B/op         23 allocs/op
Benchmark_S2M_UngorjiGocodec_SimpleEncodeDecode-32                811300              4310 ns/op            4888 B/op         34 allocs/op
Benchmark_M2S_FxamackerCbor_MarshalUnmarshal-32                  3383008               915.9 ns/op           120 B/op          8 allocs/op
Benchmark_S2M_FxamackerCbor_MarshalUnmarshal-32                  2444036              1434 ns/op             452 B/op         11 allocs/op
Benchmark_M2S_EtNikBinngo_MarshalUnmarshal-32                    1550070              2493 ns/op             425 B/op         39 allocs/op
Benchmark_S2M_EtNikBinngo_MarshalUnmarshal-32                    1401469              2397 ns/op             768 B/op         45 allocs/op
Benchmark_M2S_IchibanTnetstrings_MarshalUnmarshal-32              674178              6222 ns/op            9329 B/op         48 allocs/op
Benchmark_S2M_IchibanTnetstrings_MarshalUnmarshal-32              575835              6277 ns/op            9561 B/op         46 allocs/op
Benchmark_M2S_Vmihailenco_MarhsalUnmarshal-32                    3407040              1038 ns/op             188 B/op          5 allocs/op
Benchmark_S2M_Vmihailenco_MarhsalUnmarshal-32                    2886090              1120 ns/op             607 B/op         12 allocs/op
Benchmark_M2S_ShamatonMsgpackV2_MarshalUnmarshal-32              2211585              1472 ns/op             956 B/op         15 allocs/op
Benchmark_S2M_ShamatonMsgpackV2_MarshalUnmarshal-32              2777444              1344 ns/op             556 B/op         15 allocs/op
Benchmark_M2S_DONUTSLz4Msgpack_MarshalUnmarshal-32                182338             18439 ns/op            1347 B/op         16 allocs/op
Benchmark_S2M_DONUTSLz4Msgpack_MarshalUnmarshal-32                267619             15190 ns/op             763 B/op         16 allocs/op
Benchmark_M2S_SurrealdbCork_EncodeDecode-32                      2511390              1443 ns/op            1217 B/op          6 allocs/op
Benchmark_S2M_SurrealdbCork_EncodeDecode-32                      2304793              1587 ns/op            1611 B/op         12 allocs/op
Benchmark_M2S_MitchellhMapstructure_Decode-32                    1913150              1879 ns/op             720 B/op         18 allocs/op
Benchmark_S2M_MitchellhMapstructure_Decode-32                    4118420               918.9 ns/op           536 B/op         12 allocs/op
Benchmark_M2S_BurntSushiToml_EncodeUnmarshal-32                   300060             10839 ns/op            7958 B/op         71 allocs/op
Benchmark_S2M_BurntSushiToml_EncodeUnmarshal-32                   392454             10341 ns/op            8231 B/op         73 allocs/op
Benchmark_M2S_PelletierGoTomlV2_MarshalUnmarshal-32               885229              3485 ns/op            1600 B/op         27 allocs/op
Benchmark_S2M_PelletierGoTomlV2_MarshalUnmarshal-32               886214              3788 ns/op            1800 B/op         31 allocs/op
Benchmark_M2S_NaoinaToml_MarshalUnmarshal-32                       48402             67834 ns/op          398544 B/op         77 allocs/op
Benchmark_S2M_NaoinaToml_MarshalUnmarshal-32                       52870             69643 ns/op          398112 B/op         80 allocs/op
Benchmark_M2S_GoccyGoYaml_MarshalUnmarshal-32                     208440             20988 ns/op            7821 B/op        214 allocs/op
Benchmark_S2M_GoccyGoYaml_MarshalUnmarshal-32                     189339             17815 ns/op            7580 B/op        202 allocs/op
Benchmark_M2S_GopkgInYamlV3_MarshalUnmarshal-32                   260914             15408 ns/op           14112 B/op         80 allocs/op
Benchmark_S2M_GopkgInYamlV3_MarshalUnmarshal-32                   269636             16052 ns/op           14400 B/op         81 allocs/op
Benchmark_M2S_GhodssYaml_MarshalUnmarshal-32                      125688             29624 ns/op           21377 B/op        161 allocs/op
Benchmark_S2M_GhodssYaml_MarshalUnmarshal-32                      118972             30211 ns/op           21441 B/op        161 allocs/op

Benchmark_M2S_EncodingJson_MarshalUnmarshal-32                   1425469              2545 ns/op             609 B/op         16 allocs/op
Benchmark_S2M_EncodingJson_MarshalUnmarshal-32                   1519557              2509 ns/op             697 B/op         18 allocs/op
Benchmark_M2S_KokizzuJson5b_MarshalUnmarshal-32                  1489426              2439 ns/op             632 B/op         16 allocs/op
Benchmark_S2M_KokizzuJson5b_MarshalUnmarshal-32                  1420410              2560 ns/op             960 B/op         20 allocs/op
Benchmark_M2S_GoccyGoJson_MarshalUnmarshal-32                    4751290               741.3 ns/op            88 B/op          3 allocs/op
Benchmark_S2M_GoccyGoJson_MarshalUnmarshal-32                    3378453              1040 ns/op             522 B/op         12 allocs/op
Benchmark_M2S_JsonIteratorGo_MarshalUnmarshal-32                 3695068               995.9 ns/op           196 B/op          8 allocs/op
Benchmark_S2M_JsonIteratorGo_MarshalUnmarshal-32                 3023834              1158 ns/op             505 B/op         14 allocs/op
Benchmark_M2S_PquernaFfjson_MarshalUnmarshal-32                  1468940              2464 ns/op             609 B/op         16 allocs/op
Benchmark_S2M_PquernaFfjson_MarshalUnmarshal-32                  1461277              2511 ns/op             697 B/op         18 allocs/op
Benchmark_M2S_MongoDriverBson_MarshalUnmarshal-32                1954384              1800 ns/op             414 B/op         14 allocs/op
Benchmark_S2M_MongoDriverBson_MarshalUnmarshal-32                1373968              2641 ns/op             759 B/op         18 allocs/op
Benchmark_M2S_HjsonHjsonGoV4_MarshalUnmarshal-32                  231424             14451 ns/op            3926 B/op         78 allocs/op
Benchmark_S2M_HjsonHjsonGoV4_MarshalUnmarshal-32                  322134             12762 ns/op            3864 B/op         66 allocs/op
Benchmark_M2S_GopkgInMgoV2Bson_MarshalUnmarshal-32               2375430              1495 ns/op             232 B/op         13 allocs/op
Benchmark_S2M_GopkgInMgoV2Bson_MarshalUnmarshal-32               2339065              1619 ns/op             528 B/op         15 allocs/op
Benchmark_M2S_UngorjiGoCodec_CborEncodeDecode-32                 1000000              3531 ns/op            4340 B/op         23 allocs/op
Benchmark_S2M_UngorjiGoCodec_CborEncodeDecode-32                  767800              4566 ns/op            4888 B/op         34 allocs/op
Benchmark_M2S_UngorjiGocodec_BincEncodeDecode-32                  922044              3531 ns/op            4340 B/op         23 allocs/op
Benchmark_S2M_UngorjiGocodec_BincEncodeDecode-32                  736339              4388 ns/op            4888 B/op         34 allocs/op
Benchmark_M2S_UngorjiGocodec_JsonEncodeDecode-32                  827829              4235 ns/op            4956 B/op         25 allocs/op
Benchmark_S2M_UngorjiGocodec_JsonEncodeDecode-32                  673035              5057 ns/op            5504 B/op         36 allocs/op
Benchmark_M2S_UngorjiGocodec_SimpleEncodeDecode-32                948152              3725 ns/op            4340 B/op         23 allocs/op
Benchmark_S2M_UngorjiGocodec_SimpleEncodeDecode-32                776654              4528 ns/op            4888 B/op         34 allocs/op
Benchmark_M2S_FxamackerCbor_MarshalUnmarshal-32                  3151551              1103 ns/op             120 B/op          8 allocs/op
Benchmark_S2M_FxamackerCbor_MarshalUnmarshal-32                  2468354              1416 ns/op             452 B/op         11 allocs/op
Benchmark_M2S_EtNikBinngo_MarshalUnmarshal-32                    1591635              2344 ns/op             425 B/op         39 allocs/op
Benchmark_S2M_EtNikBinngo_MarshalUnmarshal-32                    1387429              2583 ns/op             768 B/op         45 allocs/op
Benchmark_M2S_IchibanTnetstrings_MarshalUnmarshal-32              497610              6951 ns/op            9329 B/op         48 allocs/op
Benchmark_S2M_IchibanTnetstrings_MarshalUnmarshal-32              510625              6235 ns/op            9561 B/op         46 allocs/op
Benchmark_M2S_Vmihailenco_MarhsalUnmarshal-32                    3438524              1035 ns/op             188 B/op          5 allocs/op
Benchmark_S2M_Vmihailenco_MarhsalUnmarshal-32                    3033379              1198 ns/op             606 B/op         12 allocs/op
Benchmark_M2S_ShamatonMsgpackV2_MarshalUnmarshal-32              2251254              1645 ns/op             956 B/op         15 allocs/op
Benchmark_S2M_ShamatonMsgpackV2_MarshalUnmarshal-32              2328894              1545 ns/op             556 B/op         15 allocs/op
Benchmark_M2S_DONUTSLz4Msgpack_MarshalUnmarshal-32                198080             17282 ns/op            1251 B/op         16 allocs/op
Benchmark_S2M_DONUTSLz4Msgpack_MarshalUnmarshal-32                199928             15795 ns/op             763 B/op         16 allocs/op
Benchmark_M2S_SurrealdbCork_EncodeDecode-32                      2377034              1447 ns/op            1217 B/op          6 allocs/op
Benchmark_S2M_SurrealdbCork_EncodeDecode-32                      2190968              1546 ns/op            1611 B/op         12 allocs/op
Benchmark_M2S_MitchellhMapstructure_Decode-32                    1894234              1893 ns/op             720 B/op         18 allocs/op
Benchmark_S2M_MitchellhMapstructure_Decode-32                    3702127               890.4 ns/op           536 B/op         12 allocs/op
Benchmark_M2S_BurntSushiToml_EncodeUnmarshal-32                   336931             10748 ns/op            7958 B/op         71 allocs/op
Benchmark_S2M_BurntSushiToml_EncodeUnmarshal-32                   303885             11374 ns/op            8231 B/op         73 allocs/op
Benchmark_M2S_PelletierGoTomlV2_MarshalUnmarshal-32              1000000              3604 ns/op            1600 B/op         27 allocs/op
Benchmark_S2M_PelletierGoTomlV2_MarshalUnmarshal-32               869896              3841 ns/op            1800 B/op         31 allocs/op
Benchmark_M2S_NaoinaToml_MarshalUnmarshal-32                       43172             70813 ns/op          398544 B/op         77 allocs/op
Benchmark_S2M_NaoinaToml_MarshalUnmarshal-32                       59991             71117 ns/op          398112 B/op         80 allocs/op
Benchmark_M2S_GoccyGoYaml_MarshalUnmarshal-32                     159710             22008 ns/op            7821 B/op        214 allocs/op
Benchmark_S2M_GoccyGoYaml_MarshalUnmarshal-32                     175780             18073 ns/op            7580 B/op        202 allocs/op
Benchmark_M2S_GopkgInYamlV3_MarshalUnmarshal-32                   271116             17226 ns/op           14112 B/op         80 allocs/op
Benchmark_S2M_GopkgInYamlV3_MarshalUnmarshal-32                   193048             16709 ns/op           14400 B/op         81 allocs/op
Benchmark_M2S_GhodssYaml_MarshalUnmarshal-32                      116172             32118 ns/op           21377 B/op        161 allocs/op
Benchmark_S2M_GhodssYaml_MarshalUnmarshal-32                      109956             31529 ns/op           21442 B/op        161 allocs/op

             M2S_GoccyGoJson_MarshalUnmarshal-32  6661932   517     80   3
     M2S_VmihailencoMspackV5_MarhsalUnmarshal-32  4572597   741    188   5
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
        M2S_DONUTSLz4Msgpack_MarshalUnmarshal-32   247033 17272   1271  16
             M2S_GoccyGoYaml_MarshalUnmarshal-32   194260 19278   7815 214
              M2S_GhodssYaml_MarshalUnmarshal-32   152730 27637  21344 160
              M2S_NaoinaToml_MarshalUnmarshal-32    48402 67834 398544  77

             S2M_GoccyGoJson_MarshalUnmarshal-32  4716292   770    513  12
             S2M_MitchellhMapstructure_Decode-32  4160382   879    536  12
     S2M_VmihailencoMspackV5_MarhsalUnmarshal-32  3700621   952    606  12
          S2M_JsonIteratorGo_MarshalUnmarshal-32  3411016  1072    497  14
           S2M_FxamackerCbor_MarshalUnmarshal-32  3251618  1084    444  11
             S2M_Vmihailenco_MarhsalUnmarshal-32  2886090  1120    607  12
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

`

type sorter struct {
	TestName   string
	RunCount   int64
	Duration   float64
	AllocBytes int
	AllocCount int
}

func printHeader(prefix string) {
	fmt.Printf("\n%48s %8s %5s %6s %3s\n", prefix, "total", "ns/op", "B/op", "allocs/op")
}
func (s sorter) PrintUnique(m map[string]bool) {
	if _, ok := m[s.TestName]; !ok {
		fmt.Printf("%48s %8d %5.0f %6d %3d\n", s.TestName, s.RunCount, s.Duration, s.AllocBytes, s.AllocCount)
		m[s.TestName] = true
	}
}

func main() {
	map2struct := []sorter{}
	struct2map := []sorter{}
	lines := S.Split(results, "\n")
	for _, line := range lines {
		line = S.Trim(line)
		if len(line) == 0 {
			continue
		}
		cells := strings.Fields(line)
		var row sorter
		if len(cells) == 8 {
			if cells[3] != `ns/op` {
				L.Print(cells, `not in ns/op`)
				continue
			}
			if cells[5] != `B/op` {
				L.Print(cells, `not in B/op`)
				continue
			}
			if cells[7] != `allocs/op` {
				L.Print(cells, `not in allocs/op`)
				continue
			}
			row = sorter{
				TestName:   S.RightOf(cells[0], `_`),
				RunCount:   S.ToI(cells[1]),
				Duration:   S.ToF(cells[2]),
				AllocBytes: S.ToInt(cells[4]),
				AllocCount: S.ToInt(cells[6]),
			}
		} else if len(cells) == 5 {
			row = sorter{
				TestName:   cells[0],
				RunCount:   S.ToI(cells[1]),
				Duration:   S.ToF(cells[2]),
				AllocBytes: S.ToInt(cells[3]),
				AllocCount: S.ToInt(cells[4]),
			}
		} else {
			L.Print(cells, `wrong cell count`)
			continue
		}

		if S.StartsWith(row.TestName, "M2S_") {
			map2struct = append(map2struct, row)
		} else {
			struct2map = append(struct2map, row)
		}
	}
	sort.Slice(map2struct, func(i, j int) bool {
		return map2struct[i].Duration < map2struct[j].Duration
	})
	sort.Slice(struct2map, func(i, j int) bool {
		return struct2map[i].Duration < struct2map[j].Duration
	})

	printHeader("map to struct")
	m := map[string]bool{}
	for _, row := range map2struct {
		row.PrintUnique(m)
	}
	printHeader("struct to map")
	m = map[string]bool{}
	for _, row := range struct2map {
		row.PrintUnique(m)
	}
}
