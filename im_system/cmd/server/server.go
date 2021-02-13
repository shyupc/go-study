package main

import "github.com/shyupc/go-study/im_system/server"

func main() {
	s := server.NewServer("127.0.0.1", 8888)
	s.Start()
}
