package main

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

// go.mongodb.org/mongo-driver/bson
// require bson tag on the struct

func MongoDriverBson_MarshalUnmarshal(in, out any) {
	b, _ := bson.Marshal(in)
	_ = bson.Unmarshal(b, out)
}

func Benchmark_M2S_MongoDriverBson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		MongoDriverBson_MarshalUnmarshal(myMap1, &resultA)
	}
}

func Benchmark_S2M_MongoDriverBson_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		MongoDriverBson_MarshalUnmarshal(myRow1, &resultA)
	}
}
