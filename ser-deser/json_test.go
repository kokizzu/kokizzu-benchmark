package main_test

import (
	"bytes"
	"encoding/json"
	"testing"

	gojson "github.com/goccy/go-json"
	"github.com/kokizzu/json5b/encoding/json5b"
)

// encoding/json

func EncodingJsonEncodeDecode(in, out any) {
	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(in)
	_ = json.NewDecoder(buf).Decode(out)
}

func Benchmark_M2S_EncodingJson_EncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		EncodingJsonEncodeDecode(myMap1, &resultA)
	}
}
func Benchmark_S2M_EncodingJson_EncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		EncodingJsonEncodeDecode(myRow1, &resultA)
	}
}

func EncodingJsonMarshalUnmarshal(in, out any) {
	b, _ := json.Marshal(in)
	json.Unmarshal(b, out)
}

func Benchmark_M2S_EncodingJson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		EncodingJsonMarshalUnmarshal(myMap1, &resultA)
	}
}
func Benchmark_S2M_EncodingJson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		EncodingJsonMarshalUnmarshal(myRow1, &resultA)
	}
}

// github.com/kokizzu/json5b

func KokizzuJson5bMarshalUnmarshal(in, out any) {
	b, _ := json5b.Marshal(in)
	_ = json5b.Unmarshal(b, out)
}

func Benchmark_M2S_KokizzuJson5b_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		KokizzuJson5bMarshalUnmarshal(myMap1, &resultA)
	}
}
func Benchmark_S2M_KokizzuJson5b_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		KokizzuJson5bMarshalUnmarshal(myRow1, &resultA)
	}
}

// github.com/goccy/go-json

func GoccyGoJsonMarshalUnmarshal(in, out any) {
	b, _ := gojson.Marshal(in)
	gojson.Unmarshal(b, out)
}

func Benchmark_M2S_GoccyGoJson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		GoccyGoJsonMarshalUnmarshal(myMap1, &resultA)
	}
}


func Benchmark_S2M_GoccyGoJson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		GoccyGoJsonMarshalUnmarshal(myRow1, &resultA)
	}
}
