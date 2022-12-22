package main

import (
	"bytes"
	"testing"

	burntoml "github.com/BurntSushi/toml"
	notoml "github.com/naoina/toml"
	gotomlv2 "github.com/pelletier/go-toml/v2"
)

// github.com/BurntSushi/toml
func BurntSushiToml_EncodeUnmarshal(in, out any) {
	b := new(bytes.Buffer)
	_ = burntoml.NewEncoder(b).Encode(in)
	_ = burntoml.Unmarshal(b.Bytes(), out)
}
func Benchmark_M2S_BurntSushiToml_EncodeUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		BurntSushiToml_EncodeUnmarshal(myMap1, &resultA)
	}
}
func Benchmark_S2M_BurntSushiToml_EncodeUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		BurntSushiToml_EncodeUnmarshal(myRow1, &resultA)
	}
}

// github.com/pelletier/go-toml/v2

func PelletierGoTomlV2_MarshalUnmarshal(in, out any) {
	b, _ := gotomlv2.Marshal(in)
	_ = gotomlv2.Unmarshal(b, out)
}

func Benchmark_M2S_PelletierGoTomlV2_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		PelletierGoTomlV2_MarshalUnmarshal(myMap1, &resultA)
	}
}

func Benchmark_S2M_PelletierGoTomlV2_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		PelletierGoTomlV2_MarshalUnmarshal(myRow1, &resultA)
	}
}

// github.com/naoina/toml
// requires toml tag on struct fields

func NaoinaToml_MarshalUnmarshal(in, out any) {
	b, _ := notoml.Marshal(in)
	_ = notoml.Unmarshal(b, out)
}

func Benchmark_M2S_NaoinaToml_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStructToml{}
		NaoinaToml_MarshalUnmarshal(myMap1, &resultA)
	}
}

func Benchmark_S2M_NaoinaToml_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		NaoinaToml_MarshalUnmarshal(myRow1, &resultA)
	}
}
