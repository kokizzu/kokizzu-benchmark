package json

import (
	proto "github.com/golang/protobuf/proto"
)

type Options int32

const (
	Options_ONE Options = 0
	Options_TWO Options = 1
)

var Options_name = map[int32]string{
	0: "ONE",
	1: "TWO",
}
var Options_value = map[string]int32{
	"ONE": 0,
	"TWO": 1,
}

func (x Options) String() string {
	return proto.EnumName(Options_name, int32(x))
}

type Container struct {
	Option Options `protobuf:"varint,1,opt,name=option,enum=main.Options" json:"option,omitempty"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
