package main

import (
	"math"
	"testing"

	"github.com/kokizzu/gotro/M"
)

type myStruct struct {
	Name string `yaml:"Name" bson:"Name" toml:"Name"`
	Age  int64  `yaml:"Age" bson:"Age" toml:"Age"`
}

const name = "Tony"
const age = math.MaxInt64

var myMap1 = map[string]any{
	"Name": name,
	"Age":  age,
}
var myRow1 = myStruct{
	Name: name,
	Age:  age,
}

func structMatch(t *testing.T, m *myStruct) {
	if m.Name != name || m.Age != age {
		t.Errorf("value mismatched: %v", *m)
	}
}

func mapMatch(t *testing.T, m map[string]any) {
	m2 := M.SX(m)
	if m2.GetStr("Name") != name || m2.GetInt("Age") != age {
		t.Errorf("value mismatched: %v", m)
	}
}

func TestVerify(t *testing.T) {
	var resultA myStruct
	var resultB map[string]any
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
	t.Run("PquernaFfjson_MarshalUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		PquernaFfjson_MarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		PquernaFfjson_MarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("MongoDriverBson_MarshalUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		MongoDriverBson_MarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		MongoDriverBson_MarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("BurntSushiToml_EncodeUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		BurntSushiToml_EncodeUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		BurntSushiToml_EncodeUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("PelletierGoTomlV2_MarshalUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		PelletierGoTomlV2_MarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		PelletierGoTomlV2_MarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("MitchellhMapstructure_Decode", func(t *testing.T) {
		resultA = myStruct{}
		MitchellhMapstructure_Decode(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		MitchellhMapstructure_Decode(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("NaoinaToml_MarshalUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		NaoinaToml_MarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		NaoinaToml_MarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("HjsonHjsonGoV4_MarshalUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		HjsonHjsonGoV4_MarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		HjsonHjsonGoV4_MarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("DONUTSLz4Msgpack_MarshalUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		DONUTSLz4Msgpack_MarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		DONUTSLz4Msgpack_MarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("SurrealdbCork_EncodeDecode", func(t *testing.T) {
		resultA = myStruct{}
		SurrealdbCork_EncodeDecode(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		SurrealdbCork_EncodeDecode(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("EtNikBinngo_MarshalUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		EtNikBinngo_MarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		EtNikBinngo_MarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("IchibanTnetstrings_MarshalUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		IchibanTnetstrings_MarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		IchibanTnetstrings_MarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
	t.Run("GopkgInMgoV2Bson_MarshalUnmarshal", func(t *testing.T) {
		resultA = myStruct{}
		GopkgInMgoV2Bson_MarshalUnmarshal(myMap1, &resultA)
		structMatch(t, &resultA)
		resultB = map[string]any{}
		GopkgInMgoV2Bson_MarshalUnmarshal(resultA, &resultB)
		mapMatch(t, resultB)
	})
}
