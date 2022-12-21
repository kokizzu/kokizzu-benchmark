package main

import (
	"bytes"
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/ugorji/go/codec"
)

// github.com/ugorji/go/codec

func UngorjiGoCodec_CborEncodeDecode(in, out any) {
	h := new(codec.CborHandle)
	b := new(bytes.Buffer)
	enc := codec.NewEncoder(b, h)
	dec := codec.NewDecoder(b, h)
	_ = enc.Encode(in)
	_ = dec.Decode(out)
}

func Benchmark_M2S_UngorjiGoCodec_CborEncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		UngorjiGoCodec_CborEncodeDecode(myMap1, &resultA)
	}
}

func Benchmark_S2M_UngorjiGoCodec_CborEncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		UngorjiGoCodec_CborEncodeDecode(myRow1, &resultA)
	}
}

func UngorjiGocodec_BincEncodeDecode(in, out any) {
	h := new(codec.BincHandle)
	b := new(bytes.Buffer)
	enc := codec.NewEncoder(b, h)
	dec := codec.NewDecoder(b, h)
	_ = enc.Encode(in)
	_ = dec.Decode(out)
}

func Benchmark_M2S_UngorjiGocodec_BincEncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		UngorjiGocodec_BincEncodeDecode(myMap1, &resultA)
	}
}

func Benchmark_S2M_UngorjiGocodec_BincEncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		UngorjiGocodec_BincEncodeDecode(myRow1, &resultA)
	}
}

func UngorjiGocodec_JsonEncodeDecode(in, out any) {
	h := new(codec.JsonHandle)
	b := new(bytes.Buffer)
	enc := codec.NewEncoder(b, h)
	dec := codec.NewDecoder(b, h)
	_ = enc.Encode(in)
	_ = dec.Decode(out)
}

func Benchmark_M2S_UngorjiGocodec_JsonEncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		UngorjiGocodec_JsonEncodeDecode(myMap1, &resultA)
	}
}

func Benchmark_S2M_UngorjiGocodec_JsonEncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		UngorjiGocodec_JsonEncodeDecode(myRow1, &resultA)
	}
}

func UngorjiGocodec_SimpleEncodeDecode(in, out any) {
	h := new(codec.SimpleHandle)
	b := new(bytes.Buffer)
	enc := codec.NewEncoder(b, h)
	dec := codec.NewDecoder(b, h)
	_ = enc.Encode(in)
	_ = dec.Decode(out)
}

func Benchmark_M2S_UngorjiGocodec_SimpleEncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		UngorjiGocodec_SimpleEncodeDecode(myMap1, &resultA)
	}
}

func Benchmark_S2M_UngorjiGocodec_SimpleEncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		UngorjiGocodec_SimpleEncodeDecode(myRow1, &resultA)
	}
}

// github.com/fxamacker/cbor/v2

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
