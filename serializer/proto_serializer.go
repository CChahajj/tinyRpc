package serializer

import (
	"errors"

	"github.com/golang/protobuf/proto"
)

var Proto = ProtoSerializer{}

type ProtoSerializer struct {
}

var NotImplementProtoSerialsizer = errors.New("param does not implement proto.Message")

func (_ ProtoSerializer) Marshal(m interface{}) ([]byte, error) {
	var body proto.Message
	if m == nil {
		return []byte{}, nil
	}
	var ok bool
	if body, ok = m.(proto.Message); !ok {
		return []byte{}, NotImplementProtoSerialsizer
	}
	return proto.Marshal(body)
}

func (_ ProtoSerializer) Unmarshal(b []byte, m interface{}) error {
	var body proto.Message
	if m == nil {
		return nil
	}
	var ok bool
	if body, ok = m.(proto.Message); !ok {
		return NotImplementProtoSerialsizer
	}
	return proto.Unmarshal(b, body)
}
