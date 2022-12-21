package main

import (
	"testing"

	"github.com/ghodss/yaml"
	goyaml "github.com/goccy/go-yaml"
	yaml3 "gopkg.in/yaml.v3"
)

// github.com/goccy/go-yaml

func GoccyGoYaml_MarshalUnmarshal(in, out any) {
	b, _ := goyaml.Marshal(in)
	_ = goyaml.Unmarshal(b, out)
}

func Benchmark_M2S_GoccyGoYaml_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		GoccyGoYaml_MarshalUnmarshal(myMap1, &resultA)
	}
}
func Benchmark_S2M_GoccyGoYaml_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		GoccyGoYaml_MarshalUnmarshal(myRow1, &resultA)
	}
}

// gopkg.in/yaml.v3

func GopkgInYamlV3_MarshalUnmarshal(in, out any) {
	b, _ := yaml3.Marshal(in)
	_ = yaml3.Unmarshal(b, out)

}

func Benchmark_M2S_GopkgInYamlV3_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		GopkgInYamlV3_MarshalUnmarshal(myMap1, &resultA)
	}
}

func Benchmark_S2M_GopkgInYamlV3_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		GopkgInYamlV3_MarshalUnmarshal(myRow1, &resultA)
	}
}

// github.com/ghodss/yaml

func GhodssYaml_MarshalUnmarshal(in, out any) {
	// the only library that doesn't require yaml tag on the struct
	b, _ := yaml.Marshal(in)
	_ = yaml.Unmarshal(b, out)
}

func Benchmark_M2S_GhodssYaml_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		GhodssYaml_MarshalUnmarshal(myMap1, &resultA)
	}
}

func Benchmark_S2M_GhodssYaml_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		GhodssYaml_MarshalUnmarshal(myRow1, &resultA)
	}
}
