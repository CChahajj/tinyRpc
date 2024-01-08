package header

import (
	"github.com/CChahajj/tinyRpc/compressor"
	"sync"
)

// RequestHeader request header structure looks like:
// +--------------+----------------+----------+------------+----------+
// | CompressType |      Method    |    ID    | RequestLen | Checksum |
// +--------------+----------------+----------+------------+----------+
// |    uint16    | uvarint+string |  uvarint |   uvarint  |  uint32  |
// +--------------+----------------+----------+------------+----------+
type RequestHeader struct {
	sync.RWMutex
	CompressType compressor.CompressType
	Method       string
	ID           uint64
	RequestLen   uint32
	CheckNum     uint32
}

// ResponseHeader request header structure looks like:
// +--------------+----------------+----------+-------------+----------+
// | CompressType |      Error     |    ID    | ResponseLen | Checksum |
// +--------------+----------------+----------+-------------+----------+
// |    uint16    | uvarint+string |  uvarint |   uvarint   |  uint32  |
// +--------------+----------------+----------+-------------+----------+
type ResponseHeader struct {
	sync.RWMutex
	CompressType compressor.CompressType
	ID           uint64
	Error        string
	ResponseLen  uint32
	CheckNum     uint32
}

func (r *RequestHeader) Marshal() error {

}

func (r *RequestHeader) UnMarshal() error {

}

func (r *ResponseHeader) Marshal() error {

}

func (r *ResponseHeader) UnMarshal() error {

}

func (r *RequestHeader) ResetHeader() {

}
