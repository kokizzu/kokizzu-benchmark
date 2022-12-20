package main

import (
	"testing"

	"github.com/fxamacker/cbor/v2"
)

func FxamackerCbor_MarshalUnmarshal(in, out any) {
	b, _ := cbor.Marshal(in)
	_ = cbor.Unmarshal(b, out)
}

func Benchmark_M2S_FxamackerCbor_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		FxamackerCbor_MarshalUnmarshal(myMap1, &resultA)
	}
}
func Benchmark_S2M_FxamackerCbor_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		FxamackerCbor_MarshalUnmarshal(myRow1, &resultA)
	}
}
