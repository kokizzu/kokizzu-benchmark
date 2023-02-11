package main

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	expjson "github.com/go-json-experiment/json"
	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	"github.com/kokizzu/json5b/encoding/json5b"
	"github.com/pquerna/ffjson/ffjson"
	segmentio "github.com/segmentio/encoding/json"
)

// encoding/json

func EncodingJson_MarshalUnmarshal(in, out any) {
	b, _ := json.Marshal(in)
	_ = json.Unmarshal(b, out)
}

func Benchmark_M2S_EncodingJson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		EncodingJson_MarshalUnmarshal(myMap1, &resultA)
	}
}
func Benchmark_S2M_EncodingJson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		EncodingJson_MarshalUnmarshal(myRow1, &resultA)
	}
}

func Benchmark_S2S_EncodingJson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		EncodingJson_MarshalUnmarshal(myRow1, &resultA)
	}
}

// github.com/kokizzu/json5b

func KokizzuJson5b_MarshalUnmarshal(in, out any) {
	b, _ := json5b.Marshal(in)
	_ = json5b.Unmarshal(b, out)
}

func Benchmark_M2S_KokizzuJson5b_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		KokizzuJson5b_MarshalUnmarshal(myMap1, &resultA)
	}
}
func Benchmark_S2M_KokizzuJson5b_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		KokizzuJson5b_MarshalUnmarshal(myRow1, &resultA)
	}
}
func Benchmark_S2S_KokizzuJson5b_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		KokizzuJson5b_MarshalUnmarshal(myRow1, &resultA)
	}
}

// github.com/goccy/go-json

func GoccyGoJson_MarshalUnmarshal(in, out any) {
	b, _ := gojson.Marshal(in)
	_ = gojson.Unmarshal(b, out)
}

func Benchmark_M2S_GoccyGoJson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		GoccyGoJson_MarshalUnmarshal(myMap1, &resultA)
	}
}

func Benchmark_S2M_GoccyGoJson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		GoccyGoJson_MarshalUnmarshal(myRow1, &resultA)
	}
}
func Benchmark_S2S_GoccyGoJson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		GoccyGoJson_MarshalUnmarshal(myRow1, &resultA)
	}
}

// github.com/json-iterator/go

func JsonIteratorGo_MarshalUnmarshal(in, out any) {
	b, _ := jsoniter.Marshal(in)
	_ = jsoniter.Unmarshal(b, out)
}

func Benchmark_M2S_JsonIteratorGo_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		JsonIteratorGo_MarshalUnmarshal(myMap1, &resultA)
	}
}

func Benchmark_S2M_JsonIteratorGo_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		JsonIteratorGo_MarshalUnmarshal(myRow1, &resultA)
	}
}

func Benchmark_S2S_JsonIteratorGo_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		JsonIteratorGo_MarshalUnmarshal(myRow1, &resultA)
	}
}

// github.com/pquerna/ffjson/ffjson

func PquernaFfjson_MarshalUnmarshal(in, out any) {
	b, _ := ffjson.Marshal(in)
	_ = ffjson.Unmarshal(b, out)
}

func Benchmark_M2S_PquernaFfjson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		PquernaFfjson_MarshalUnmarshal(myMap1, &resultA)
	}
}

func Benchmark_S2M_PquernaFfjson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		PquernaFfjson_MarshalUnmarshal(myRow1, &resultA)
	}
}

func Benchmark_S2S_PquernaFfjson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		PquernaFfjson_MarshalUnmarshal(myRow1, &resultA)
	}
}

// github.com/bytedance/sonic

func BytedanceSonic_MarshalUnmarshal(in, out any) {
	b, _ := sonic.Marshal(in)
	_ = sonic.Unmarshal(b, out)
}

func Benchmark_M2S_BytedanceSonic_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		BytedanceSonic_MarshalUnmarshal(myMap1, &resultA)
	}
}

func Benchmark_S2M_BytedanceSonic_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		BytedanceSonic_MarshalUnmarshal(myRow1, &resultA)
	}
}

func Benchmark_S2S_BytedanceSonic_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		BytedanceSonic_MarshalUnmarshal(myRow1, &resultA)
	}
}

// github.com/segmentio/encoding/json

func SegmentioEncodingJson_MarshalUnmarshal(in, out any) {
	b, _ := segmentio.Marshal(in)
	_ = segmentio.Unmarshal(b, out)
}

func Benchmark_M2S_SegmentioEncodingJson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		SegmentioEncodingJson_MarshalUnmarshal(myMap1, &resultA)
	}
}

func Benchmark_S2M_SegmentioEncodingJson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		SegmentioEncodingJson_MarshalUnmarshal(myRow1, &resultA)
	}
}

func Benchmark_S2S_SegmentioEncodingJson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		SegmentioEncodingJson_MarshalUnmarshal(myRow1, &resultA)
	}
}

// github.com/go-json-experiment/json

func GoJsonExperimentJson_MarshalUnmarshal(in, out any) {
	b, _ := expjson.Marshal(in)
	_ = expjson.Unmarshal(b, out)
}

func Benchmark_M2S_GoJsonExperimentJson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		GoJsonExperimentJson_MarshalUnmarshal(myMap1, &resultA)
	}
}

func Benchmark_S2M_GoJsonExperimentJson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		GoJsonExperimentJson_MarshalUnmarshal(myRow1, &resultA)
	}
}

func Benchmark_S2S_GoJsonExperimentJson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		GoJsonExperimentJson_MarshalUnmarshal(myRow1, &resultA)
	}
}
