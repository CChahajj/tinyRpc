package tinyRpc

import (
	"github.com/CChahajj/tinyRpc/codec"
	"github.com/CChahajj/tinyRpc/serializer"
	"net"
	"net/rpc"
)

type Server struct {
	*rpc.Server
	serializer.Serializer
}

func NewServer(opts ...Option) *Server {
    options := Options{
		Serializer:serializer.Proto,
	}

	for _,option := range opts{
        option(&options)
	}

	return &Server{&rpc.Server{},options.Serializer}
}

func (server *Server) Register(rcvr interface{}) error {
	return server.Server.Register(rcvr)
}

func (server *Server) RegisterName(name string, rcvr interface{}) error {
	return server.Server.RegisterName(name, rcvr)
}

func (server *Server) Serve(lis net.Listener){
	for {
		conn,err := lis.Accept()
		if err != nil{
			continue
		}
		go server.Server.ServeCodec(codec.NewServerCodec(conn,server.Serializer))
	}
}
