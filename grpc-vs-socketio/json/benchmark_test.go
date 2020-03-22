package json

import (
	"testing"

	"github.com/golang/protobuf/jsonpb"
)

func BenchmarkUnmarshal(b *testing.B) {
	const data = `
	{
	  "option": 0
	}
	`

	var c Container
	for i := 0; i < b.N; i++ {
		if err := jsonpb.UnmarshalString(data, &c); err != nil {
			b.Fatalf("unmarshal: %v\n", err)
		}
	}
}

func BenchmarkMarshal(b *testing.B) {
	m := &jsonpb.Marshaler{}

	c := Container{
		Option: Options_ONE,
	}

	for i := 0; i < b.N; i++ {
		if _, err := m.MarshalToString(&c); err != nil {
			b.Fatalf("marshal: %v\n", err)
		}
	}
}
