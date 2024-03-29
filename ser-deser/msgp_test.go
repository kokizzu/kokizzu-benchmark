package main

import (
	"testing"

	lzmsgpack "github.com/d-o-n-u-t-s/lz4msgpack"
	msgpack2 "github.com/shamaton/msgpack/v2"
	"github.com/surrealdb/cork"
	msgpack5 "github.com/vmihailenco/msgpack/v5"
)

// github.com/vmihailenco/msgpack/v5

func VmihailencoMsgpackV5_MarhsalUnmarshal(in, out any) {
	b, _ := msgpack5.Marshal(in)
	_ = msgpack5.Unmarshal(b, out)
}
func Benchmark_M2S_VmihailencoMspackV5_MarhsalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		VmihailencoMsgpackV5_MarhsalUnmarshal(myMap1, &resultA)
	}
}

func Benchmark_S2M_VmihailencoMspackV5_MarhsalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		VmihailencoMsgpackV5_MarhsalUnmarshal(myRow1, &resultA)
	}
}
func Benchmark_S2S_VmihailencoMspackV5_MarhsalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		VmihailencoMsgpackV5_MarhsalUnmarshal(myRow1, &resultA)
	}
}

// github.com/shamaton/msgpack/v2

func ShamatonMsgpackV2_MarshalUnmarshal(in, out any) {
	b, _ := msgpack2.Marshal(in)
	_ = msgpack2.Unmarshal(b, out)
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

func Benchmark_S2S_ShamatonMsgpackV2_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		ShamatonMsgpackV2_MarshalUnmarshal(myRow1, &resultA)
	}
}

// github.com/d-o-n-u-t-s/lz4msgpack

func DONUTSLz4Msgpack_MarshalUnmarshal(in, out any) {
	b, _ := lzmsgpack.Marshal(in)
	_ = lzmsgpack.Unmarshal(b, out)
}

func Benchmark_M2S_DONUTSLz4Msgpack_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		DONUTSLz4Msgpack_MarshalUnmarshal(myMap1, &resultA)
	}
}

func Benchmark_S2M_DONUTSLz4Msgpack_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		DONUTSLz4Msgpack_MarshalUnmarshal(myRow1, &resultA)
	}
}

func Benchmark_S2S_DONUTSLz4Msgpack_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		DONUTSLz4Msgpack_MarshalUnmarshal(myRow1, &resultA)
	}
}

// github.com/surrealdb/cork

func SurrealdbCork_EncodeDecode(in, out any) {
	b := cork.Encode(in)
	cork.DecodeInto(b, out)
}

func Benchmark_M2S_SurrealdbCork_EncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		SurrealdbCork_EncodeDecode(myMap1, &resultA)
	}
}

func Benchmark_S2M_SurrealdbCork_EncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		SurrealdbCork_EncodeDecode(myRow1, &resultA)
	}
}

func Benchmark_S2S_SurrealdbCork_EncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		SurrealdbCork_EncodeDecode(myRow1, &resultA)
	}
}
