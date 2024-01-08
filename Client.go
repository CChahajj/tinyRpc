package tinyRpc

import (
	"github.com/CChahajj/tinyRpc/compressor"
	"github.com/CChahajj/tinyRpc/serializer"
)

type Option func(o *Options)

type Options struct {
	compressor.CompressType
	serializer.Serializer
}

func WithSerializer(s serializer.Serializer) Option {
	return func(o *Options) {
		o.Serializer = s
	}
}

func WithCompressType(c compressor.CompressType) Option {
	return func(o *Options) {
		o.CompressType = c
	}
}
