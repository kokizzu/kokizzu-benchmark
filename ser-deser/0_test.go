package main

import (
	"testing"

	"github.com/kokizzu/gotro/M"
)

type myStruct struct {
	Name string `yaml:"Name"`
	Age  int64  `yaml:"Age"`
}

var myMap1 = map[string]any{
	"Name": "Tony",
	"Age":  23,
}
var myRow1 = myStruct{
	Name: "Tony",
	Age:  23,
}

func structMatch(t *testing.T, m *myStruct) {
	if m.Name != "Tony" || m.Age != 23 {
		t.Errorf("value mismatched: %v", *m)
	}
}

func mapMatch(t *testing.T, m map[string]any) {
	m2 := M.SX(m)
	if m2.GetStr("Name") != "Tony" || m2.GetInt("Age") != 23 {
		t.Errorf("value mismatched: %v", m)
	}
}

func TestVerify(t *testing.T) {
	var resultA myStruct
	var resultB map[string]any
	//t.Run("EncodingJson_EncodeDecode", func(t *testing.T) {
	//	resultA = myStruct{}
	//	EncodingJson_EncodeDecode(myMap1, &resultA)
	//	structMatch(t, &resultA)
	//	resultB = map[string]any{}
	//	EncodingJson_EncodeDecode(resultA, &resultB)
	//	mapMatch(t, resultB)
	//})
	t.Run("EncodingJson_MarshalUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		EncodingJson_MarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		EncodingJson_MarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("KokizzuJson5b_MarshalUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		KokizzuJson5b_MarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		KokizzuJson5b_MarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("GoccyGoJson_MarshalUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		GoccyGoJson_MarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		GoccyGoJson_MarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
	//t.Run("Vmihailenco_EncodeDecode", func(t *testing.T) {
	//	resultA = myStruct{}
	//	Vmihailenco_EncodeDecode(myMap1, &resultA)
	//	structMatch(t, &resultA)
	//	resultB = map[string]any{}
	//	Vmihailenco_EncodeDecode(resultA, &resultB)
	//	mapMatch(t, resultB)
	//})
	t.Run("VmihailencoMsgpackV5_MarhsalUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		VmihailencoMsgpackV5_MarhsalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		VmihailencoMsgpackV5_MarhsalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("FxamackerCbor_MarshalUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		FxamackerCbor_MarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		FxamackerCbor_MarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("GoccyGoYaml_MarshalUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		GoccyGoYaml_MarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		GoccyGoYaml_MarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("GopkgInYamlV3_MarshalUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		GopkgInYamlV3_MarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		GopkgInYamlV3_MarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("GhodssYaml_MarshalUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		GhodssYaml_MarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		GhodssYaml_MarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("UngorjiGoCodec_CborEncodeDecode", func(t *testing.T) {
		resultA = myStruct{}
		UngorjiGoCodec_CborEncodeDecode(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		UngorjiGoCodec_CborEncodeDecode(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("UngorjiGocodec_BincEncodeDecode", func(t *testing.T) {
		resultA = myStruct{}
		UngorjiGocodec_BincEncodeDecode(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		UngorjiGocodec_BincEncodeDecode(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("UngorjiGocodec_JsonEncodeDecode", func(t *testing.T) {
		resultA = myStruct{}
		UngorjiGocodec_JsonEncodeDecode(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		UngorjiGocodec_JsonEncodeDecode(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("UngorjiGocodec_SimpleEncodeDecode", func(t *testing.T) {
		resultA = myStruct{}
		UngorjiGocodec_SimpleEncodeDecode(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		UngorjiGocodec_SimpleEncodeDecode(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("JsonIteratorGo_MarshalUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		JsonIteratorGo_MarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		JsonIteratorGo_MarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("ShamatonMsgpackV2_MarshalUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		ShamatonMsgpackV2_MarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		ShamatonMsgpackV2_MarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
}
