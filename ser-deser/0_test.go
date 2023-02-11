package main

import (
	"testing"

	"github.com/kokizzu/gotro/M"
)

type myStruct struct {
	Name string
	Age  int64
}

func (m myStruct) Match(t *testing.T) {
	if m.Name != name || m.Age != age {
		t.Errorf("value mismatched: %v", m)
	}
}

type myStructYaml struct {
	Name string `yaml:"Name"`
	Age  int64  `yaml:"Age"`
}

func (m myStructYaml) Match(t *testing.T) {
	if m.Name != name || m.Age != age {
		t.Errorf("value mismatched: %v", m)
	}
}

type myStructBson struct {
	Name string `bson:"Name"`
	Age  int64  `bson:"Age"`
}

func (m myStructBson) Match(t *testing.T) {
	if m.Name != name || m.Age != age {
		t.Errorf("value mismatched: %v", m)
	}
}

type myStructToml struct {
	Name string `toml:"Name"`
	Age  int64  `toml:"Age"`
}

func (m myStructToml) Match(t *testing.T) {
	if m.Name != name || m.Age != age {
		t.Errorf("value mismatched: %v", m)
	}
}

const name = `Tony` // "1234567890123456789012345678901234567890123456789012345678901234567890"
const age = 123     // math.MaxInt64

var myMap1 = map[string]any{
	"Name": name,
	"Age":  age,
}
var myRow1 = myStruct{
	Name: name,
	Age:  age,
}

var myRowYaml = myStructYaml{
	Name: name,
	Age:  age,
}

var myRowBson = myStructBson{
	Name: name,
	Age:  age,
}

var myRowToml = myStructToml{
	Name: name,
	Age:  age,
}

func mapMatch(t *testing.T, m map[string]any) {
	m2 := M.SX(m)
	if m2.GetStr("Name") != name || m2.GetInt("Age") != age {
		t.Errorf("value mismatched: %v", m)
	}
}

type matcher interface {
	Match(t *testing.T)
}

func testFunc[T matcher](t *testing.T, f func(in, out any)) {
	var resultA T
	f(myMap1, &resultA)
	resultA.Match(t)

	resultB := map[string]any{}
	f(resultA, &resultB)
	mapMatch(t, resultB)

	var resultC T
	f(resultA, &resultC)
	resultC.Match(t)
}

func TestVerify(t *testing.T) {
	t.Run("EncodingJson_MarshalUnmarshal", func(t *testing.T) {
		testFunc[myStruct](t, EncodingJson_MarshalUnmarshal)
	})
	t.Run("KokizzuJson5b_MarshalUnmarshal", func(t *testing.T) {
		testFunc[myStruct](t, KokizzuJson5b_MarshalUnmarshal)
	})
	t.Run("GoccyGoJson_MarshalUnmarshal", func(t *testing.T) {
		testFunc[myStruct](t, GoccyGoJson_MarshalUnmarshal)
	})
	t.Run("VmihailencoMsgpackV5_MarhsalUnmarshal", func(t *testing.T) {
		testFunc[myStruct](t, VmihailencoMsgpackV5_MarhsalUnmarshal)
	})
	t.Run("FxamackerCbor_MarshalUnmarshal", func(t *testing.T) {
		testFunc[myStruct](t, FxamackerCbor_MarshalUnmarshal)
	})
	t.Run("GoccyGoYaml_MarshalUnmarshal", func(t *testing.T) {
		testFunc[myStructYaml](t, GoccyGoYaml_MarshalUnmarshal)
	})
	t.Run("GopkgInYamlV3_MarshalUnmarshal", func(t *testing.T) {
		testFunc[myStructYaml](t, GopkgInYamlV3_MarshalUnmarshal)
	})
	t.Run("GhodssYaml_MarshalUnmarshal", func(t *testing.T) {
		testFunc[myStruct](t, GhodssYaml_MarshalUnmarshal)
	})
	t.Run("UngorjiGoCodec_CborEncodeDecode", func(t *testing.T) {
		testFunc[myStruct](t, UngorjiGoCodec_CborEncodeDecode)
	})
	t.Run("UngorjiGocodec_BincEncodeDecode", func(t *testing.T) {
		testFunc[myStruct](t, UngorjiGocodec_BincEncodeDecode)
	})
	t.Run("UngorjiGocodec_JsonEncodeDecode", func(t *testing.T) {
		testFunc[myStruct](t, UngorjiGocodec_JsonEncodeDecode)
	})
	t.Run("UngorjiGocodec_SimpleEncodeDecode", func(t *testing.T) {
		testFunc[myStruct](t, UngorjiGocodec_SimpleEncodeDecode)
	})
	t.Run("JsonIteratorGo_MarshalUnmarshal", func(t *testing.T) {
		testFunc[myStruct](t, JsonIteratorGo_MarshalUnmarshal)
	})
	t.Run("ShamatonMsgpackV2_MarshalUnmarshal", func(t *testing.T) {
		testFunc[myStruct](t, ShamatonMsgpackV2_MarshalUnmarshal)
	})
	t.Run("PquernaFfjson_MarshalUnmarshal", func(t *testing.T) {
		testFunc[myStruct](t, PquernaFfjson_MarshalUnmarshal)
	})
	t.Run("MongoDriverBson_MarshalUnmarshal", func(t *testing.T) {
		testFunc[myStructBson](t, MongoDriverBson_MarshalUnmarshal)
	})
	t.Run("BurntSushiToml_EncodeUnmarshal", func(t *testing.T) {
		testFunc[myStruct](t, BurntSushiToml_EncodeUnmarshal)
	})
	t.Run("PelletierGoTomlV2_MarshalUnmarshal", func(t *testing.T) {
		testFunc[myStruct](t, PelletierGoTomlV2_MarshalUnmarshal)
	})
	t.Run("MitchellhMapstructure_Decode", func(t *testing.T) {
		testFunc[myStruct](t, MitchellhMapstructure_Decode)
	})
	t.Run("NaoinaToml_MarshalUnmarshal", func(t *testing.T) {
		testFunc[myStructToml](t, NaoinaToml_MarshalUnmarshal)
	})
	t.Run("HjsonHjsonGoV4_MarshalUnmarshal", func(t *testing.T) {
		testFunc[myStruct](t, HjsonHjsonGoV4_MarshalUnmarshal)
	})
	t.Run("DONUTSLz4Msgpack_MarshalUnmarshal", func(t *testing.T) {
		testFunc[myStruct](t, DONUTSLz4Msgpack_MarshalUnmarshal)
	})
	t.Run("SurrealdbCork_EncodeDecode", func(t *testing.T) {
		testFunc[myStruct](t, SurrealdbCork_EncodeDecode)
	})
	t.Run("EtNikBinngo_MarshalUnmarshal", func(t *testing.T) {
		testFunc[myStruct](t, EtNikBinngo_MarshalUnmarshal)
	})
	t.Run("IchibanTnetstrings_MarshalUnmarshal", func(t *testing.T) {
		testFunc[myStruct](t, IchibanTnetstrings_MarshalUnmarshal)
	})
	t.Run("GopkgInMgoV2Bson_MarshalUnmarshal", func(t *testing.T) {
		testFunc[myStructBson](t, GopkgInMgoV2Bson_MarshalUnmarshal)
	})
	t.Run("BytedanceSonic_MarshalUnmarshal", func(t *testing.T) {
		testFunc[myStruct](t, BytedanceSonic_MarshalUnmarshal)
	})
	t.Run("SegmentioEncodingJson_MarshalUnmarshal", func(t *testing.T) {
		testFunc[myStruct](t, SegmentioEncodingJson_MarshalUnmarshal)
	})
}
