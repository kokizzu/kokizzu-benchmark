package main

import (
	"testing"

	"github.com/hjson/hjson-go/v4"
	mdbson "go.mongodb.org/mongo-driver/bson"
	mgobson "gopkg.in/mgo.v2/bson"
)

// go.mongodb.org/mongo-driver/bson
// require bson tag on the struct

func MongoDriverBson_MarshalUnmarshal(in, out any) {
	b, _ := mdbson.Marshal(in)
	_ = mdbson.Unmarshal(b, out)
}

func Benchmark_M2S_MongoDriverBson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStructBson{}
		MongoDriverBson_MarshalUnmarshal(myMap1, &resultA)
	}
}

func Benchmark_S2M_MongoDriverBson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		MongoDriverBson_MarshalUnmarshal(myRow1, &resultA)
	}
}

func Benchmark_S2S_MongoDriverBson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStructBson{}
		MongoDriverBson_MarshalUnmarshal(myRow1, &resultA)
	}
}

// github.com/hjson/hjson-go/v4

func HjsonHjsonGoV4_MarshalUnmarshal(in, out any) {
	b, _ := hjson.Marshal(in)
	_ = hjson.Unmarshal(b, out)
}

func Benchmark_M2S_HjsonHjsonGoV4_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		HjsonHjsonGoV4_MarshalUnmarshal(myMap1, &resultA)
	}
}

func Benchmark_S2M_HjsonHjsonGoV4_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		HjsonHjsonGoV4_MarshalUnmarshal(myRow1, &resultA)
	}
}
func Benchmark_S2S_HjsonHjsonGoV4_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		HjsonHjsonGoV4_MarshalUnmarshal(myRow1, &resultA)
	}
}

// gopkg.in/mgo.v2/bson

func GopkgInMgoV2Bson_MarshalUnmarshal(in, out any) {
	b, _ := mgobson.Marshal(in)
	_ = mgobson.Unmarshal(b, out)
}

func Benchmark_M2S_GopkgInMgoV2Bson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStructBson{}
		GopkgInMgoV2Bson_MarshalUnmarshal(myMap1, &resultA)
	}
}

func Benchmark_S2M_GopkgInMgoV2Bson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		GopkgInMgoV2Bson_MarshalUnmarshal(myRow1, &resultA)
	}
}

func Benchmark_S2S_GopkgInMgoV2Bson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStructBson{}
		GopkgInMgoV2Bson_MarshalUnmarshal(myRow1, &resultA)
	}
}
