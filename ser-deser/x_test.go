package main

import (
	"bytes"
	"encoding/json"
	"log"
	"testing"

	json5b "github.com/kokizzu/json5b/encoding/json5b"

	gojson "github.com/goccy/go-json"
)

// how to run:
// go test -bench .

// how to add new test:
// create a function
// add on TestVerify
// create Benchmark[RepoName][MethodName](b *testing.B)
// go test .

type myStruct struct {
	Name string
	Age  int64
}

var myData = map[string]any{
	"Name": "Tony",
	"Age":  23,
}

func mustMatch(m *myStruct) {
	if m.Name != "Tony" || m.Age != 23 {
		log.Fatalf("value mismatched: %v", *m)
	}
}

func TestVerify(t *testing.T) {
	resultA := myStruct{}
	EncodingJsonEncodeDecode(myData, &resultA)
	resultA = myStruct{}
	EncodingJsonMarshalUnmarshal(myData, &resultA)
	resultA = myStruct{}
	KokizzuJson5bMarshalUnmarshal(myData, &resultA)
	resultA = myStruct{}
	GoccyGoJsonMarshalUnmarshal(myData, &resultA)
}

func EncodingJsonEncodeDecode(in, out any) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(in)
	json.NewDecoder(buf).Decode(out)
}

func BenchmarkEncodingJsonEncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		EncodingJsonEncodeDecode(myData, &resultA)
	}
}

func EncodingJsonMarshalUnmarshal(in, out any) {
	b, _ := json.Marshal(in)
	json.Unmarshal(b, out)
}

func BenchmarkEncodingJsonMarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		EncodingJsonMarshalUnmarshal(myData, &resultA)
	}
}

func KokizzuJson5bMarshalUnmarshal(in, out any) {
	b, _ := json5b.Marshal(in)
	json5b.Unmarshal(b, out)
}

func BenchmarkKokizzuJson5bMarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		KokizzuJson5bMarshalUnmarshal(myData, &resultA)
	}
}

func GoccyGoJsonMarshalUnmarshal(in, out any) {
	b, _ := gojson.Marshal(in)
	gojson.Unmarshal(b, out)
}

func BenchmarkGoccyGoJsonMarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		GoccyGoJsonMarshalUnmarshal(myData, &resultA)
	}
}
