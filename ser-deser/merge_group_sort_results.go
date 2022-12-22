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
Benchmark_M2S_EncodingJson_MarshalUnmarshal-32                   1666660              1945 ns/op             609 B/op         16 allocs/op
Benchmark_S2M_EncodingJson_MarshalUnmarshal-32                   1991836              1895 ns/op             697 B/op         18 allocs/op
Benchmark_M2S_KokizzuJson5b_MarshalUnmarshal-32                  1991420              1875 ns/op             632 B/op         16 allocs/op
Benchmark_S2M_KokizzuJson5b_MarshalUnmarshal-32                  1762370              2025 ns/op             960 B/op         20 allocs/op
Benchmark_M2S_GoccyGoJson_MarshalUnmarshal-32                    6385722               532.8 ns/op            88 B/op          3 allocs/op
Benchmark_S2M_GoccyGoJson_MarshalUnmarshal-32                    4482018               777.5 ns/op           522 B/op         12 allocs/op
Benchmark_M2S_JsonIteratorGo_MarshalUnmarshal-32                 4953243               760.6 ns/op           196 B/op          8 allocs/op
Benchmark_S2M_JsonIteratorGo_MarshalUnmarshal-32                 3984788               908.4 ns/op           505 B/op         14 allocs/op
Benchmark_M2S_PquernaFfjson_MarshalUnmarshal-32                  1785348              2057 ns/op             609 B/op         16 allocs/op
Benchmark_S2M_PquernaFfjson_MarshalUnmarshal-32                  1789437              1940 ns/op             697 B/op         18 allocs/op
Benchmark_M2S_MongoDriverBson_MarshalUnmarshal-32                2477983              1459 ns/op             414 B/op         14 allocs/op
Benchmark_S2M_MongoDriverBson_MarshalUnmarshal-32                1782861              2076 ns/op             759 B/op         18 allocs/op
Benchmark_M2S_HjsonHjsonGoV4_MarshalUnmarshal-32                  325690             11219 ns/op            3923 B/op         78 allocs/op
Benchmark_S2M_HjsonHjsonGoV4_MarshalUnmarshal-32                  347148              9758 ns/op            3864 B/op         66 allocs/op
Benchmark_M2S_GopkgInMgoV2Bson_MarshalUnmarshal-32               3328545              1093 ns/op             232 B/op         13 allocs/op
Benchmark_S2M_GopkgInMgoV2Bson_MarshalUnmarshal-32               3043587              1210 ns/op             528 B/op         15 allocs/op
Benchmark_M2S_UngorjiGoCodec_CborEncodeDecode-32                 1309371              2728 ns/op            4340 B/op         23 allocs/op
Benchmark_S2M_UngorjiGoCodec_CborEncodeDecode-32                 1043264              3398 ns/op            4888 B/op         34 allocs/op
Benchmark_M2S_UngorjiGocodec_BincEncodeDecode-32                 1293718              2823 ns/op            4340 B/op         23 allocs/op
Benchmark_S2M_UngorjiGocodec_BincEncodeDecode-32                 1000000              3309 ns/op            4888 B/op         34 allocs/op
Benchmark_M2S_UngorjiGocodec_JsonEncodeDecode-32                 1000000              3356 ns/op            4956 B/op         25 allocs/op
Benchmark_S2M_UngorjiGocodec_JsonEncodeDecode-32                  924037              4083 ns/op            5504 B/op         36 allocs/op
Benchmark_M2S_UngorjiGocodec_SimpleEncodeDecode-32               1215002              2871 ns/op            4340 B/op         23 allocs/op
Benchmark_S2M_UngorjiGocodec_SimpleEncodeDecode-32                865138              3683 ns/op            4888 B/op         34 allocs/op
Benchmark_M2S_FxamackerCbor_MarshalUnmarshal-32                  4411034               831.9 ns/op           120 B/op          8 allocs/op
Benchmark_S2M_FxamackerCbor_MarshalUnmarshal-32                  3030465              1222 ns/op             452 B/op         11 allocs/op
Benchmark_M2S_EtNikBinngo_MarshalUnmarshal-32                    1733204              2148 ns/op             425 B/op         39 allocs/op
Benchmark_S2M_EtNikBinngo_MarshalUnmarshal-32                    1611408              2214 ns/op             768 B/op         45 allocs/op
Benchmark_M2S_IchibanTnetstrings_MarshalUnmarshal-32              599776              5071 ns/op            9329 B/op         48 allocs/op
Benchmark_S2M_IchibanTnetstrings_MarshalUnmarshal-32              752912              5081 ns/op            9561 B/op         46 allocs/op
Benchmark_M2S_VmihailencoMspackV5_MarhsalUnmarshal-32            4005979               873.4 ns/op           188 B/op          5 allocs/op
Benchmark_S2M_VmihailencoMspackV5_MarhsalUnmarshal-32            3676345              1032 ns/op             607 B/op         12 allocs/op
Benchmark_M2S_ShamatonMsgpackV2_MarshalUnmarshal-32              2633756              1416 ns/op             956 B/op         15 allocs/op
Benchmark_S2M_ShamatonMsgpackV2_MarshalUnmarshal-32              2834528              1315 ns/op             556 B/op         15 allocs/op
Benchmark_M2S_DONUTSLz4Msgpack_MarshalUnmarshal-32                209420             16371 ns/op            1268 B/op         16 allocs/op
Benchmark_S2M_DONUTSLz4Msgpack_MarshalUnmarshal-32                217772             14595 ns/op             731 B/op         16 allocs/op
Benchmark_M2S_SurrealdbCork_EncodeDecode-32                      2911629              1250 ns/op            1217 B/op          6 allocs/op
Benchmark_S2M_SurrealdbCork_EncodeDecode-32                      2674888              1427 ns/op            1611 B/op         12 allocs/op
Benchmark_M2S_MitchellhMapstructure_Decode-32                    2164524              1628 ns/op             720 B/op         18 allocs/op
Benchmark_S2M_MitchellhMapstructure_Decode-32                    4512402               798.7 ns/op           536 B/op         12 allocs/op
Benchmark_M2S_BurntSushiToml_EncodeUnmarshal-32                   369919              9201 ns/op            7958 B/op         71 allocs/op
Benchmark_S2M_BurntSushiToml_EncodeUnmarshal-32                   372052              8783 ns/op            8231 B/op         73 allocs/op
Benchmark_M2S_PelletierGoTomlV2_MarshalUnmarshal-32              1000000              3146 ns/op            1600 B/op         27 allocs/op
Benchmark_S2M_PelletierGoTomlV2_MarshalUnmarshal-32               924240              3342 ns/op            1800 B/op         31 allocs/op
Benchmark_M2S_NaoinaToml_MarshalUnmarshal-32                       57607             58331 ns/op          398544 B/op         77 allocs/op
Benchmark_S2M_NaoinaToml_MarshalUnmarshal-32                       60916             52047 ns/op          398112 B/op         80 allocs/op
Benchmark_M2S_GoccyGoYaml_MarshalUnmarshal-32                     214776             16192 ns/op            7821 B/op        214 allocs/op
Benchmark_S2M_GoccyGoYaml_MarshalUnmarshal-32                     239161             14976 ns/op            7581 B/op        202 allocs/op
Benchmark_M2S_GopkgInYamlV3_MarshalUnmarshal-32                   271190             13524 ns/op           14112 B/op         80 allocs/op
Benchmark_S2M_GopkgInYamlV3_MarshalUnmarshal-32                   266643             13343 ns/op           14400 B/op         81 allocs/op
Benchmark_M2S_GhodssYaml_MarshalUnmarshal-32                      139384             24330 ns/op           21378 B/op        161 allocs/op
Benchmark_S2M_GhodssYaml_MarshalUnmarshal-32                      138720             24937 ns/op           21442 B/op        161 allocs/op



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
