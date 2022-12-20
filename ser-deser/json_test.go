package main

import (
	"bytes"
	"encoding/json"
	"testing"

	gojson "github.com/goccy/go-json"
	"github.com/kokizzu/json5b/encoding/json5b"
)

// encoding/json

func EncodingJson_EncodeDecode(in, out any) {
	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(in)
	_ = json.NewDecoder(buf).Decode(out)
}

func Benchmark_M2S_EncodingJson_EncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		EncodingJson_EncodeDecode(myMap1, &resultA)
	}
}
func Benchmark_S2M_EncodingJson_EncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		EncodingJson_EncodeDecode(myRow1, &resultA)
	}
}

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
