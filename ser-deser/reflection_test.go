package main

import (
	"testing"

	"github.com/mitchellh/mapstructure"
)

// github.com/mitchellh/mapstructure

func MitchellhMapstructure_Decode(in, out any) {
	_ = mapstructure.Decode(in, out)
}

func Benchmark_M2S_MitchellhMapstructure_Decode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		MitchellhMapstructure_Decode(myMap1, &resultA)
	}
}

func Benchmark_S2M_MitchellhMapstructure_Decode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := map[string]any{}
		MitchellhMapstructure_Decode(myRow1, &resultA)
	}
}

func Benchmark_S2S_MitchellhMapstructure_Decode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultA := myStruct{}
		MitchellhMapstructure_Decode(myRow1, &resultA)
	}
}
