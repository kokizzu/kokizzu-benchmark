package main

import (
	"testing"

	"github.com/vmihailenco/msgpack/v5"
)

// github.com/vmihailenco/msgpack/v5

func VmihailencoMsgpackV5_MarhsalUnmarshal(in, out any) {
	b, _ := msgpack.Marshal(in)
	_ = msgpack.Unmarshal(b, out)
}
func Benchmark_M2S_Vmihailenco_MarhsalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		VmihailencoMsgpackV5_MarhsalUnmarshal(myMap1, &resultA)
	}
}

func Benchmark_S2M_Vmihailenco_MarhsalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		VmihailencoMsgpackV5_MarhsalUnmarshal(myRow1, &resultA)
	}
}

// github.com/shamaton/msgpack/v2

func ShamatonMsgpackV2_MarshalUnmarshal(in, out any) {
	b, _ := msgpack.Marshal(in)
	_ = msgpack.Unmarshal(b, out)
}

func Benchmark_M2S_ShamatonMsgpackV2_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		ShamatonMsgpackV2_MarshalUnmarshal(myMap1, &resultA)
	}
}

func Benchmark_S2M_ShamatonMsgpackV2_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		ShamatonMsgpackV2_MarshalUnmarshal(myRow1, &resultA)
	}
}
