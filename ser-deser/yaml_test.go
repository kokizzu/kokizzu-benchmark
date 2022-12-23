package main

import (
	"testing"

	ghyaml "github.com/ghodss/yaml"
	goyaml "github.com/goccy/go-yaml"
	yaml3 "gopkg.in/yaml.v3"
)

// github.com/goccy/go-yaml
// require yaml tag on the struct

func GoccyGoYaml_MarshalUnmarshal(in, out any) {
	b, _ := goyaml.Marshal(in)
	_ = goyaml.Unmarshal(b, out)
}

func Benchmark_M2S_GoccyGoYaml_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStructYaml{}
		GoccyGoYaml_MarshalUnmarshal(myMap1, &resultA)
	}
}
func Benchmark_S2M_GoccyGoYaml_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		GoccyGoYaml_MarshalUnmarshal(myRow1, &resultA)
	}
}
func Benchmark_S2S_GoccyGoYaml_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStructYaml{}
		GoccyGoYaml_MarshalUnmarshal(myRow1, &resultA)
	}
}

// gopkg.in/yaml.v3
// require yaml tag on the struct

func GopkgInYamlV3_MarshalUnmarshal(in, out any) {
	b, _ := yaml3.Marshal(in)
	_ = yaml3.Unmarshal(b, out)

}

func Benchmark_M2S_GopkgInYamlV3_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStructYaml{}
		GopkgInYamlV3_MarshalUnmarshal(myMap1, &resultA)
	}
}

func Benchmark_S2M_GopkgInYamlV3_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		GopkgInYamlV3_MarshalUnmarshal(myRow1, &resultA)
	}
}

func Benchmark_S2S_GopkgInYamlV3_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStructYaml{}
		GopkgInYamlV3_MarshalUnmarshal(myRow1, &resultA)
	}
}

// github.com/ghodss/yaml
// the only library that doesn't require yaml tag on the struct

func GhodssYaml_MarshalUnmarshal(in, out any) {
	b, _ := ghyaml.Marshal(in)
	_ = ghyaml.Unmarshal(b, out)
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

func Benchmark_S2S_GhodssYaml_MarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		GhodssYaml_MarshalUnmarshal(myRow1, &resultA)
	}
}
