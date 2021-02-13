package server

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip        string
	Port      int
	OnlineMap map[string]*User // 在线用户列表
	mapLock   sync.RWMutex
	Message   chan string //消息广播的channel
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
}

func (s *Server) Handler(conn net.Conn) {
	u := NewUser(conn, s)

	u.Online()

	isLive := make(chan bool)

	// 接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				u.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}

			// 提取用户消息(去除\n)
			msg := string(buf[:n-1])
			u.DoMessage(msg)

			isLive <- true
		}
	}()

	// 当前handler阻塞
	for {
		select {
		case <-isLive:
		case <-time.After(300 * time.Second):
			u.SendMsg("你被踢了...\n")

			close(u.C)

			_ = conn.Close()

			return
		}
	}
}

//启动服务器接口
func (s *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// close listen socket
	defer listener.Close()

	// 启动监听Message广播消息channel的goroutine
	go s.ListenMessage()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}

		// do handler
		go s.Handler(conn)
	}
}

// 监听Message广播消息channel的goroutine，一旦有消息就发送给所有在线用户
func (s *Server) ListenMessage() {
	for {
		msg := <-s.Message

		// 将msg发送给所有在线用户
		s.mapLock.Lock()
		for _, user := range s.OnlineMap {
			user.C <- msg
		}
		s.mapLock.Unlock()
	}
}

// 广播消息
func (s *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	s.Message <- sendMsg
}
