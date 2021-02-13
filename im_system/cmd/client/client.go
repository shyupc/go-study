package main

import (
	"flag"
	"fmt"

	"github.com/shyupc/go-study/im_system/client"
)

var (
	serverIp   string
	serverPort int
)

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口(默认是8888)")
}

func main() {
	flag.Parse()

	c := client.NewClient(serverIp, serverPort)
	if c == nil {
		fmt.Println(">>>>>>连接服务器失败.")
		return
	}

	go c.DealResponse()

	fmt.Println(">>>>>>连接服务器成功.")

	c.Run()
}
