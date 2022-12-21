package main

import (
	"bytes"
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/ugorji/go/codec"
)

// github.com/ugorji/go/codec

func UngorjiGoCodec_EncodeDecode(in, out any) {
	h := new(codec.CborHandle)
	b := new(bytes.Buffer)
	enc := codec.NewEncoder(b, h)
	dec := codec.NewDecoder(b, h)
	_ = enc.Encode(in)
	_ = dec.Decode(out)
}

func Benchmark_M2S_UngorjiGoCodec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		UngorjiGoCodec_EncodeDecode(myMap1, &resultA)
	}
}

// github.com/fxamacker/cbor/v2

func Benchmark_S2M_UngorjiGoCodec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		UngorjiGoCodec_EncodeDecode(myRow1, &resultA)
	}
}

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
