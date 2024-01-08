package codec

import (
	"bufio"
	"github.com/CChahajj/tinyRpc/header"
	"github.com/CChahajj/tinyRpc/serializer"
	"io"
	"net/rpc"
	"sync"
)

type ServerCodec struct {
	r io.Reader
	w io.Writer
	c io.Closer

	request header.RequestHeader
	serializer.Serializer
	seq     uint64
	pending map[uint64]uint64
	mutex   sync.Mutex
}

func NewServerCodec(conn io.ReadWriteCloser, s serializer.Serializer) rpc.ServerCodec {
	return &ServerCodec{
		r:          bufio.NewReader(conn),
		w:          bufio.NewWriter(conn),
		c:          conn,
		Serializer: s,
		pending:    make(map[uint64]uint64),
	}
}

func (s *ServerCodec) ReadRequestHeader(r *rpc.Request) error {
	//1，重置请求头
	s.request.ResetHeader()
	//2，从 io 中读取被序列化的请求头
	data, err := recvFrame(s.r)

	//3，将请求头反序列化（此序列化是根据特定格式写死的）

	//4，将编码器（codec）序列号递增（），并给 rpc 内置 Request 赋值
}

func (s *ServerCodec) ReadRequestBody(body interface{}) error {

	//1，校验请求头里的 RequestLen 没被赋值，则空请求体

	//2，根据 RequestLen 长度读取，body 内容赋值字节切片

	//3，根据 crc32 校验 CheckNum，根据请求头中的压缩器类型，校验压缩器 (compress)

	//4，压缩器解压请求体，并将二进制数据反序列化 ( protocol buffer，当然也可以自定义 json ) 实例
}

func (s *ServerCodec) WriteResponse(r *rpc.Response, data interface{}) error {

	//1，判断是否是合法请求，是，响应序列号删除。

	//2，序列化请求 params

	//3，压缩响应头，存到对象池里

	//4，将请求头和请求体分别写入 bufio 刷入 io

}

func (s *ServerCodec) Close() error {
	return s.c.Close()
}
