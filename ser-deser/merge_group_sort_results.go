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

Benchmark_M2S_EncodingJson_MarshalUnmarshal-32                   1408900              2732 ns/op             609 B/op         16 allocs/op
Benchmark_S2M_EncodingJson_MarshalUnmarshal-32                   1440130              2539 ns/op             697 B/op         18 allocs/op
Benchmark_M2S_KokizzuJson5b_MarshalUnmarshal-32                  1446464              2450 ns/op             632 B/op         16 allocs/op
Benchmark_S2M_KokizzuJson5b_MarshalUnmarshal-32                  1313719              2726 ns/op             960 B/op         20 allocs/op
Benchmark_M2S_GoccyGoJson_MarshalUnmarshal-32                    5052811               723.5 ns/op            88 B/op          3 allocs/op
Benchmark_S2M_GoccyGoJson_MarshalUnmarshal-32                    3332824              1069 ns/op             522 B/op         12 allocs/op
Benchmark_M2S_JsonIteratorGo_MarshalUnmarshal-32                 3434266              1058 ns/op             196 B/op          8 allocs/op
Benchmark_S2M_JsonIteratorGo_MarshalUnmarshal-32                 3171970              1149 ns/op             505 B/op         14 allocs/op
Benchmark_M2S_PquernaFfjson_MarshalUnmarshal-32                  1356283              2614 ns/op             609 B/op         16 allocs/op
Benchmark_S2M_PquernaFfjson_MarshalUnmarshal-32                  1396503              2542 ns/op             697 B/op         18 allocs/op
Benchmark_M2S_MongoDriverBson_MarshalUnmarshal-32                1862202              2033 ns/op             414 B/op         14 allocs/op
Benchmark_S2M_MongoDriverBson_MarshalUnmarshal-32                1447026              2679 ns/op             759 B/op         18 allocs/op
Benchmark_M2S_HjsonHjsonGoV4_MarshalUnmarshal-32                  269106             15129 ns/op            3932 B/op         78 allocs/op
Benchmark_S2M_HjsonHjsonGoV4_MarshalUnmarshal-32                  277578             13279 ns/op            3861 B/op         66 allocs/op
Benchmark_M2S_GopkgInMgoV2Bson_MarshalUnmarshal-32               2501360              1453 ns/op             232 B/op         13 allocs/op
Benchmark_S2M_GopkgInMgoV2Bson_MarshalUnmarshal-32               2181133              1622 ns/op             528 B/op         15 allocs/op
Benchmark_M2S_UngorjiGoCodec_CborEncodeDecode-32                 1000000              3449 ns/op            4340 B/op         23 allocs/op
Benchmark_S2M_UngorjiGoCodec_CborEncodeDecode-32                  795828              4633 ns/op            4888 B/op         34 allocs/op
Benchmark_M2S_UngorjiGocodec_BincEncodeDecode-32                  929955              3708 ns/op            4340 B/op         23 allocs/op
Benchmark_S2M_UngorjiGocodec_BincEncodeDecode-32                  791056              4478 ns/op            4888 B/op         34 allocs/op
Benchmark_M2S_UngorjiGocodec_JsonEncodeDecode-32                  818464              4263 ns/op            4956 B/op         25 allocs/op
Benchmark_S2M_UngorjiGocodec_JsonEncodeDecode-32                  686474              5051 ns/op            5504 B/op         36 allocs/op
Benchmark_M2S_UngorjiGocodec_SimpleEncodeDecode-32               1000000              3696 ns/op            4340 B/op         23 allocs/op
Benchmark_S2M_UngorjiGocodec_SimpleEncodeDecode-32                745822              4699 ns/op            4888 B/op         34 allocs/op
Benchmark_M2S_FxamackerCbor_MarshalUnmarshal-32                  3357404              1120 ns/op             120 B/op          8 allocs/op
Benchmark_S2M_FxamackerCbor_MarshalUnmarshal-32                  2523915              1397 ns/op             452 B/op         11 allocs/op
Benchmark_M2S_EtNikBinngo_MarshalUnmarshal-32                    1420594              2611 ns/op             424 B/op         39 allocs/op
Benchmark_S2M_EtNikBinngo_MarshalUnmarshal-32                    1393450              2644 ns/op             768 B/op         45 allocs/op
Benchmark_M2S_IchibanTnetstrings_MarshalUnmarshal-32              551206              6791 ns/op            9329 B/op         48 allocs/op
Benchmark_S2M_IchibanTnetstrings_MarshalUnmarshal-32              644035              6207 ns/op            9561 B/op         46 allocs/op
Benchmark_M2S_VmihailencoMspackV5_MarhsalUnmarshal-32            3586522              1037 ns/op             188 B/op          5 allocs/op
Benchmark_S2M_VmihailencoMspackV5_MarhsalUnmarshal-32            3018156              1241 ns/op             607 B/op         12 allocs/op
Benchmark_M2S_ShamatonMsgpackV2_MarshalUnmarshal-32              2201278              1598 ns/op             956 B/op         15 allocs/op
Benchmark_S2M_ShamatonMsgpackV2_MarshalUnmarshal-32              2404066              1559 ns/op             556 B/op         15 allocs/op
Benchmark_M2S_DONUTSLz4Msgpack_MarshalUnmarshal-32                201740             17736 ns/op            1246 B/op         16 allocs/op
Benchmark_S2M_DONUTSLz4Msgpack_MarshalUnmarshal-32                198936             15700 ns/op             737 B/op         16 allocs/op
Benchmark_M2S_SurrealdbCork_EncodeDecode-32                      2401036              1501 ns/op            1217 B/op          6 allocs/op
Benchmark_S2M_SurrealdbCork_EncodeDecode-32                      2227837              1557 ns/op            1611 B/op         12 allocs/op
Benchmark_M2S_MitchellhMapstructure_Decode-32                    1842063              1959 ns/op             720 B/op         18 allocs/op
Benchmark_S2M_MitchellhMapstructure_Decode-32                    3650191               963.2 ns/op           536 B/op         12 allocs/op
Benchmark_M2S_BurntSushiToml_EncodeUnmarshal-32                   321775             10896 ns/op            7958 B/op         71 allocs/op
Benchmark_S2M_BurntSushiToml_EncodeUnmarshal-32                   317946             11109 ns/op            8231 B/op         73 allocs/op
Benchmark_M2S_PelletierGoTomlV2_MarshalUnmarshal-32              1000000              3691 ns/op            1600 B/op         27 allocs/op
Benchmark_S2M_PelletierGoTomlV2_MarshalUnmarshal-32               877527              3845 ns/op            1800 B/op         31 allocs/op
Benchmark_M2S_NaoinaToml_MarshalUnmarshal-32                       44401             71361 ns/op          398544 B/op         77 allocs/op
Benchmark_S2M_NaoinaToml_MarshalUnmarshal-32                       55065             68090 ns/op          398112 B/op         80 allocs/op
Benchmark_M2S_GoccyGoYaml_MarshalUnmarshal-32                     157342             21894 ns/op            7821 B/op        214 allocs/op
Benchmark_S2M_GoccyGoYaml_MarshalUnmarshal-32                     238287             19752 ns/op            7580 B/op        202 allocs/op
Benchmark_M2S_GopkgInYamlV3_MarshalUnmarshal-32                   204738             17534 ns/op           14112 B/op         80 allocs/op
Benchmark_S2M_GopkgInYamlV3_MarshalUnmarshal-32                   202671             17183 ns/op           14400 B/op         81 allocs/op
Benchmark_M2S_GhodssYaml_MarshalUnmarshal-32                      118881             32157 ns/op           21377 B/op        161 allocs/op
Benchmark_S2M_GhodssYaml_MarshalUnmarshal-32                      111811             30728 ns/op           21442 B/op        161 allocs/op

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
