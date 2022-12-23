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
          M2S_JsonIteratorGo_MarshalUnmarshal-32  4892611   724    196   8
     M2S_VmihailencoMspackV5_MarhsalUnmarshal-32  4572597   741    188   5
           M2S_FxamackerCbor_MarshalUnmarshal-32  4598559   800    120   8
               M2S_SurrealdbCork_EncodeDecode-32  3080282  1080   1217   6
        M2S_GopkgInMgoV2Bson_MarshalUnmarshal-32  3227905  1092    232  13
       M2S_ShamatonMsgpackV2_MarshalUnmarshal-32  3062677  1161    956  15
             M2S_MitchellhMapstructure_Decode-32  2487428  1395    720  18
         M2S_MongoDriverBson_MarshalUnmarshal-32  2477983  1459    414  14
           M2S_KokizzuJson5b_MarshalUnmarshal-32  1987240  1711    632  16
            M2S_EncodingJson_MarshalUnmarshal-32  2056944  1780    600  16
             M2S_EtNikBinngo_MarshalUnmarshal-32  1985595  1857    425  39
           M2S_PquernaFfjson_MarshalUnmarshal-32  1805090  1987    609  16
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
              M2S_GhodssYaml_MarshalUnmarshal-32   139384 24330  21378 161
              M2S_NaoinaToml_MarshalUnmarshal-32    57607 58331 398544  77

             S2M_MitchellhMapstructure_Decode-32  5055402   716    536  12
             S2M_GoccyGoJson_MarshalUnmarshal-32  4660224   747    522  12
          S2M_JsonIteratorGo_MarshalUnmarshal-32  4283262   835    505  14
     S2M_VmihailencoMspackV5_MarhsalUnmarshal-32  4009863   908    607  12
           S2M_FxamackerCbor_MarshalUnmarshal-32  3562352  1023    452  11
       S2M_ShamatonMsgpackV2_MarshalUnmarshal-32  3180010  1089    556  15
        S2M_GopkgInMgoV2Bson_MarshalUnmarshal-32  3047396  1145    528  15
               S2M_SurrealdbCork_EncodeDecode-32  2976328  1196   1611  12
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
           S2M_GopkgInYamlV3_MarshalUnmarshal-32   266643 13343  14400  81
        S2M_DONUTSLz4Msgpack_MarshalUnmarshal-32   224689 14576    732  16
             S2M_GoccyGoYaml_MarshalUnmarshal-32   230042 14919   7580 202
              S2M_GhodssYaml_MarshalUnmarshal-32   151023 22682  21441 161
              S2M_NaoinaToml_MarshalUnmarshal-32    60916 52047 398112  80

Benchmark_M2S_EncodingJson_MarshalUnmarshal-32                   1655204              2088 ns/op             609 B/op         16 allocs/op
Benchmark_S2M_EncodingJson_MarshalUnmarshal-32                   1864470              1979 ns/op             697 B/op         18 allocs/op
Benchmark_S2S_EncodingJson_MarshalUnmarshal-32                   2618066              1328 ns/op             304 B/op          9 allocs/op
Benchmark_M2S_KokizzuJson5b_MarshalUnmarshal-32                  1877216              1942 ns/op             632 B/op         16 allocs/op
Benchmark_S2M_KokizzuJson5b_MarshalUnmarshal-32                  1665078              2077 ns/op             960 B/op         20 allocs/op
Benchmark_S2S_KokizzuJson5b_MarshalUnmarshal-32                  2583661              1387 ns/op             504 B/op          9 allocs/op
Benchmark_M2S_GoccyGoJson_MarshalUnmarshal-32                    6535316               563.7 ns/op            88 B/op          3 allocs/op
Benchmark_S2M_GoccyGoJson_MarshalUnmarshal-32                    4644348               804.8 ns/op           522 B/op         12 allocs/op
Benchmark_S2S_GoccyGoJson_MarshalUnmarshal-32                   12046497               317.4 ns/op           112 B/op          4 allocs/op
Benchmark_M2S_JsonIteratorGo_MarshalUnmarshal-32                 4919262               756.0 ns/op           196 B/op          8 allocs/op
Benchmark_S2M_JsonIteratorGo_MarshalUnmarshal-32                 3952065               911.2 ns/op           505 B/op         14 allocs/op
Benchmark_S2S_JsonIteratorGo_MarshalUnmarshal-32                 6927200               512.1 ns/op            92 B/op          6 allocs/op
Benchmark_M2S_PquernaFfjson_MarshalUnmarshal-32                  1771654              2122 ns/op             609 B/op         16 allocs/op
Benchmark_S2M_PquernaFfjson_MarshalUnmarshal-32                  1783028              1998 ns/op             697 B/op         18 allocs/op
Benchmark_S2S_PquernaFfjson_MarshalUnmarshal-32                  2734357              1330 ns/op             304 B/op          9 allocs/op
Benchmark_M2S_MongoDriverBson_MarshalUnmarshal-32                2341857              1559 ns/op             414 B/op         14 allocs/op
Benchmark_S2M_MongoDriverBson_MarshalUnmarshal-32                1786195              2046 ns/op             759 B/op         18 allocs/op
Benchmark_S2S_MongoDriverBson_MarshalUnmarshal-32                3418500              1044 ns/op             321 B/op          8 allocs/op
Benchmark_M2S_HjsonHjsonGoV4_MarshalUnmarshal-32                  299737             11488 ns/op            3931 B/op         78 allocs/op
Benchmark_S2M_HjsonHjsonGoV4_MarshalUnmarshal-32                  350539             10044 ns/op            3863 B/op         66 allocs/op
Benchmark_S2S_HjsonHjsonGoV4_MarshalUnmarshal-32                  330090             11303 ns/op            4582 B/op         79 allocs/op
Benchmark_M2S_GopkgInMgoV2Bson_MarshalUnmarshal-32               3172689              1162 ns/op             232 B/op         13 allocs/op
Benchmark_S2M_GopkgInMgoV2Bson_MarshalUnmarshal-32               2858161              1242 ns/op             528 B/op         15 allocs/op
Benchmark_S2S_GopkgInMgoV2Bson_MarshalUnmarshal-32               5094012               727.7 ns/op           144 B/op          9 allocs/op
Benchmark_M2S_UngorjiGoCodec_CborEncodeDecode-32                 1403346              2735 ns/op            4340 B/op         23 allocs/op
Benchmark_S2M_UngorjiGoCodec_CborEncodeDecode-32                 1000000              3668 ns/op            4888 B/op         34 allocs/op
Benchmark_S2S_UngorjiGoCodec_CborEncodeDecode-32                 1290734              2920 ns/op            4364 B/op         24 allocs/op
Benchmark_M2S_UngorjiGocodec_BincEncodeDecode-32                 1283144              2786 ns/op            4340 B/op         23 allocs/op
Benchmark_S2M_UngorjiGocodec_BincEncodeDecode-32                  834700              3697 ns/op            4888 B/op         34 allocs/op
Benchmark_S2S_UngorjiGocodec_BincEncodeDecode-32                 1000000              3059 ns/op            4364 B/op         24 allocs/op
Benchmark_M2S_UngorjiGocodec_JsonEncodeDecode-32                 1000000              3550 ns/op            4956 B/op         25 allocs/op
Benchmark_S2M_UngorjiGocodec_JsonEncodeDecode-32                  956674              4049 ns/op            5504 B/op         36 allocs/op
Benchmark_S2S_UngorjiGocodec_JsonEncodeDecode-32                 1000000              3500 ns/op            4980 B/op         26 allocs/op
Benchmark_M2S_UngorjiGocodec_SimpleEncodeDecode-32               1243546              2845 ns/op            4340 B/op         23 allocs/op
Benchmark_S2M_UngorjiGocodec_SimpleEncodeDecode-32                925716              3676 ns/op            4888 B/op         34 allocs/op
Benchmark_S2S_UngorjiGocodec_SimpleEncodeDecode-32               1205648              2919 ns/op            4364 B/op         24 allocs/op
Benchmark_M2S_FxamackerCbor_MarshalUnmarshal-32                  4093239               856.7 ns/op           120 B/op          8 allocs/op
Benchmark_S2M_FxamackerCbor_MarshalUnmarshal-32                  3124657              1136 ns/op             452 B/op         11 allocs/op
Benchmark_S2S_FxamackerCbor_MarshalUnmarshal-32                  7038808               510.7 ns/op            80 B/op          5 allocs/op
Benchmark_M2S_EtNikBinngo_MarshalUnmarshal-32                    1820691              1982 ns/op             425 B/op         39 allocs/op
Benchmark_S2M_EtNikBinngo_MarshalUnmarshal-32                    1778528              1991 ns/op             768 B/op         45 allocs/op
Benchmark_S2S_EtNikBinngo_MarshalUnmarshal-32                    1983145              1878 ns/op             400 B/op         41 allocs/op
Benchmark_M2S_IchibanTnetstrings_MarshalUnmarshal-32              751692              5361 ns/op            9329 B/op         48 allocs/op
Benchmark_S2M_IchibanTnetstrings_MarshalUnmarshal-32              744561              4938 ns/op            9561 B/op         46 allocs/op
Benchmark_S2S_IchibanTnetstrings_MarshalUnmarshal-32              722493              4950 ns/op            9289 B/op         47 allocs/op
Benchmark_M2S_VmihailencoMspackV5_MarhsalUnmarshal-32            4271594               830.8 ns/op           188 B/op          5 allocs/op
Benchmark_S2M_VmihailencoMspackV5_MarhsalUnmarshal-32            3834799               942.0 ns/op           607 B/op         12 allocs/op
Benchmark_S2S_VmihailencoMspackV5_MarhsalUnmarshal-32            4549700               817.5 ns/op           213 B/op          6 allocs/op
Benchmark_M2S_ShamatonMsgpackV2_MarshalUnmarshal-32              2932880              1273 ns/op             956 B/op         15 allocs/op
Benchmark_S2M_ShamatonMsgpackV2_MarshalUnmarshal-32              2870930              1229 ns/op             556 B/op         15 allocs/op
Benchmark_S2S_ShamatonMsgpackV2_MarshalUnmarshal-32              7897488               458.5 ns/op           148 B/op          6 allocs/op
Benchmark_M2S_DONUTSLz4Msgpack_MarshalUnmarshal-32                239769             15670 ns/op            1243 B/op         16 allocs/op
Benchmark_S2M_DONUTSLz4Msgpack_MarshalUnmarshal-32                264483             14480 ns/op             724 B/op         16 allocs/op
Benchmark_S2S_DONUTSLz4Msgpack_MarshalUnmarshal-32                272510             12414 ns/op             235 B/op          7 allocs/op
Benchmark_M2S_SurrealdbCork_EncodeDecode-32                      2964792              1212 ns/op            1217 B/op          6 allocs/op
Benchmark_S2M_SurrealdbCork_EncodeDecode-32                      2677389              1359 ns/op            1611 B/op         12 allocs/op
Benchmark_S2S_SurrealdbCork_EncodeDecode-32                      2518413              1407 ns/op            1241 B/op          7 allocs/op
Benchmark_M2S_MitchellhMapstructure_Decode-32                    2314702              1563 ns/op             720 B/op         18 allocs/op
Benchmark_S2M_MitchellhMapstructure_Decode-32                    4631634               782.7 ns/op           536 B/op         12 allocs/op
Benchmark_M2S_BurntSushiToml_EncodeUnmarshal-32                   460612              8801 ns/op            7958 B/op         71 allocs/op
Benchmark_S2M_BurntSushiToml_EncodeUnmarshal-32                   390661              9149 ns/op            8231 B/op         73 allocs/op
Benchmark_S2S_BurntSushiToml_EncodeUnmarshal-32                   398366              8458 ns/op            7918 B/op         72 allocs/op
Benchmark_M2S_PelletierGoTomlV2_MarshalUnmarshal-32              1000000              3076 ns/op            1600 B/op         27 allocs/op
Benchmark_S2M_PelletierGoTomlV2_MarshalUnmarshal-32              1000000              3208 ns/op            1800 B/op         31 allocs/op
Benchmark_S2S_PelletierGoTomlV2_MarshalUnmarshal-32              1460683              2459 ns/op            1440 B/op         23 allocs/op
Benchmark_M2S_NaoinaToml_MarshalUnmarshal-32                       51778             60458 ns/op          398544 B/op         77 allocs/op
Benchmark_S2M_NaoinaToml_MarshalUnmarshal-32                       61222             58972 ns/op          398112 B/op         80 allocs/op
Benchmark_S2S_NaoinaToml_MarshalUnmarshal-32                       56059             62123 ns/op          399064 B/op         83 allocs/op
Benchmark_M2S_GoccyGoYaml_MarshalUnmarshal-32                     217489             17248 ns/op            7821 B/op        214 allocs/op
Benchmark_S2M_GoccyGoYaml_MarshalUnmarshal-32                     249741             15550 ns/op            7581 B/op        202 allocs/op
Benchmark_S2S_GoccyGoYaml_MarshalUnmarshal-32                     211952             15542 ns/op            7982 B/op        208 allocs/op
Benchmark_M2S_GopkgInYamlV3_MarshalUnmarshal-32                   249645             14678 ns/op           14112 B/op         80 allocs/op
Benchmark_S2M_GopkgInYamlV3_MarshalUnmarshal-32                   237858             14348 ns/op           14400 B/op         81 allocs/op
Benchmark_S2S_GopkgInYamlV3_MarshalUnmarshal-32                   277396             13527 ns/op           14016 B/op         76 allocs/op
Benchmark_M2S_GhodssYaml_MarshalUnmarshal-32                      142334             26180 ns/op           21377 B/op        161 allocs/op
Benchmark_S2M_GhodssYaml_MarshalUnmarshal-32                      124528             26024 ns/op           21442 B/op        161 allocs/op
Benchmark_S2S_GhodssYaml_MarshalUnmarshal-32                      135193             25603 ns/op           21073 B/op        154 allocs/op



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
	struct2struct := []sorter{}
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
		} else if S.StartsWith(row.TestName, `S2M_`) {
			struct2map = append(struct2map, row)
		} else {
			struct2struct = append(struct2struct, row)
		}
	}
	sort.Slice(map2struct, func(i, j int) bool {
		return map2struct[i].Duration < map2struct[j].Duration
	})
	sort.Slice(struct2map, func(i, j int) bool {
		return struct2map[i].Duration < struct2map[j].Duration
	})
	sort.Slice(struct2struct, func(i, j int) bool {
		return struct2struct[i].Duration < struct2struct[j].Duration
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
	printHeader("struct to struct")
	m = map[string]bool{}
	for _, row := range struct2struct {
		row.PrintUnique(m)
	}
}
