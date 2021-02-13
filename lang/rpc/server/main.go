package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	rpcdemo "github.com/shyupc/go-study/lang/rpc"
)

func main() {
	_ = rpc.Register(rpcdemo.DemoService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
