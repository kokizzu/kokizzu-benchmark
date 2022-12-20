package main_test

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
	{
		resultA = myStruct{}
		EncodingJsonEncodeDecode(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		EncodingJsonEncodeDecode(resultA, &resultB)
		mapMatch(t, resultB)
	}
	{
		resultA = myStruct{}
		EncodingJsonMarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		EncodingJsonMarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	}
	{
		resultA = myStruct{}
		KokizzuJson5bMarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		KokizzuJson5bMarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	}
	{
		resultA = myStruct{}
		GoccyGoJsonMarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		GoccyGoJsonMarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	}
}
