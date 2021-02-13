package server

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

// 创建一个User的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	// 启动监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}

//监听当前user channel的方法，一旦有消息，就直接发送给对端客户端
func (u *User) ListenMessage() {
	for {
		msg := <-u.C

		_, _ = u.conn.Write([]byte(msg + "\n"))
	}
}

func (u *User) Online() {
	// 用户上线
	u.server.mapLock.Lock()
	u.server.OnlineMap[u.Name] = u
	u.server.mapLock.Unlock()

	// 广播
	u.server.BroadCast(u, "已上线.")
}

func (u *User) Offline() {
	// 用户下线
	u.server.mapLock.Lock()
	delete(u.server.OnlineMap, u.Name)
	u.server.mapLock.Unlock()

	// 广播
	u.server.BroadCast(u, "下线啦.")
}

func (u *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询当前在线用户
		u.server.mapLock.Lock()
		for _, user := range u.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + "在线...\n"
			u.SendMsg(onlineMsg)
		}
		u.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 修改名称: rename|张三
		newName := strings.Split(msg, "|")[1]

		_, ok := u.server.OnlineMap[newName]
		if ok {
			u.SendMsg("当前用户名被使用\n")
		} else {
			u.server.mapLock.Lock()
			delete(u.server.OnlineMap, u.Name)
			u.server.OnlineMap[newName] = u
			u.server.mapLock.Unlock()

			u.Name = newName
			u.SendMsg("您已更新用户名:" + u.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 私聊: to|张三|xxx
		s := strings.Split(msg, "|")
		remoteName := s[1]
		if remoteName == "" {
			u.SendMsg("消息格式不正确,请使用\"to|张三|xxx\"格式.\n")
			return
		}

		remoteUser, ok := u.server.OnlineMap[remoteName]
		if ok {
			u.SendMsg("该用户名不存在\n")
			return
		} else {
			content := s[2]
			if content == "" {
				u.SendMsg("无消息内容,请重发.\n")
				return
			}

			remoteUser.SendMsg(u.Name + "对您说:" + content)
		}

	} else {
		u.server.BroadCast(u, msg)
	}
}

func (u *User) SendMsg(msg string) {
	_, _ = u.conn.Write([]byte(msg))
}
