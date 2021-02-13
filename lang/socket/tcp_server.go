package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError1(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError1(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {

	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	request := make([]byte, 128)
	defer conn.Close()

	for {
		read, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		if read == 0 {
			break
		} else if strings.TrimSpace(string(request[:read])) == "timestamp" {
			dayTime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(dayTime))
		} else {
			dayTime := time.Now().String()
			conn.Write([]byte(dayTime))
		}
		request = make([]byte, 128)
	}

}

func checkError1(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
