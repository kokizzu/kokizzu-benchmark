package main

import (
	"bytes"
	"testing"

	"github.com/BurntSushi/toml"
	gotomlv2 "github.com/pelletier/go-toml/v2"
)

// github.com/BurntSushi/toml
func BurntSushiToml_MarshalUnmarshal(in, out any) {
	b := new(bytes.Buffer)
	enc := toml.NewEncoder(b)
	enc.Encode(in)
	_ = toml.Unmarshal(b.Bytes(), out)
}
func Benchmark_M2S_BurntSushiToml_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		BurntSushiToml_MarshalUnmarshal(myMap1, &resultA)
	}
}
func Benchmark_S2M_BurntSushiToml_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		BurntSushiToml_MarshalUnmarshal(myRow1, &resultA)
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
