package main

import (
	"testing"

	"github.com/kokizzu/gotro/M"
)

type myStruct struct {
	Name string
	Age  int64
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
	t.Run("EncodingJson_EncodeDecode", func(t *testing.T) {
		resultA = myStruct{}
		EncodingJson_EncodeDecode(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		EncodingJson_EncodeDecode(resultA, &resultB)
		mapMatch(t, resultB)
	})
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
	t.Run("Vmihailenco_EncodeDecode", func(t *testing.T) {
		resultA = myStruct{}
		Vmihailenco_EncodeDecode(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		Vmihailenco_EncodeDecode(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("Vmihailenco_MarhsalUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		Vmihailenco_MarhsalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		Vmihailenco_MarhsalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
}
