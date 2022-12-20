package main

import (
	"bytes"
	"testing"

	"github.com/vmihailenco/msgpack/v5"
)

// github.com/vmihailenco/msgpack/v5
func Vmihailenco_EncodeDecode(in, out any) {
	buf := new(bytes.Buffer)
	_ = msgpack.NewEncoder(buf).Encode(in)
	_ = msgpack.NewDecoder(buf).Decode(out)
}

func Benchmark_M2S_Vmihailenco_EncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		Vmihailenco_EncodeDecode(myMap1, &resultA)
	}
}

func Benchmark_S2M_Vmihailenco_EncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		Vmihailenco_EncodeDecode(myRow1, &resultA)
	}
}

func Vmihailenco_MarhsalUnmarshal(in, out any) {
	b, _ := msgpack.Marshal(in)
	_ = msgpack.Unmarshal(b, out)
}
func Benchmark_M2S_Vmihailenco_MarhsalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		Vmihailenco_MarhsalUnmarshal(myMap1, &resultA)
	}
}

func Benchmark_S2M_Vmihailenco_MarhsalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		Vmihailenco_MarhsalUnmarshal(myRow1, &resultA)
	}
}
