package client

import (
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(ip string, port int) *Client {
	client := &Client{
		ServerIp:   ip,
		ServerPort: port,
		flag:       999,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return nil
	}

	client.conn = conn

	return client
}

func (c *Client) menu() bool {
	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.修改用户名")
	fmt.Println("0.退出")

	_, _ = fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		c.flag = flag
		return true
	} else {
		fmt.Println(">>>>>>请输入合法范围内的数字<<<<<<")
		return false
	}
}

func (c *Client) Run() {
	for c.flag != 0 {
		for c.menu() != true {

		}

		switch c.flag {
		case 1:
			c.PublicChat()
		case 2:
			c.PrivateChat()
		case 3:
			c.UpdateName()
		}
	}
}

func (c *Client) DealResponse() {
	_, _ = io.Copy(os.Stdout, c.conn)
}

func (c *Client) SelectUsers() {
	sendMsg := "who\n"

	_, err := c.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
}

func (c *Client) PrivateChat() {
	c.SelectUsers()

	var remoteName string
	var chatMsg string

	fmt.Println(">>>>>>请输入聊天对象的[用户名],exit退出")
	_, _ = fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>>>>请输入聊天内容,exit退出")
		_, _ = fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				_, err := c.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn.Write err:", err)
					break
				}
			}

			chatMsg = ""
			fmt.Println(">>>>>>请输入聊天内容,exit退出")
			_, _ = fmt.Scanln(&chatMsg)
		}

		c.SelectUsers()
		fmt.Println(">>>>>>请输入聊天对象的[用户名],exit退出")
		_, _ = fmt.Scanln(&remoteName)
	}
}

func (c *Client) PublicChat() {
	var chatMsg string

	fmt.Println(">>>>>>请输入聊天内容,exit退出")
	_, _ = fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		if len(chatMsg) != 0 {
			_, err := c.conn.Write([]byte(chatMsg + "\n"))
			if err != nil {
				fmt.Println("conn.Write err:", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println(">>>>>>请输入聊天内容,exit退出")
		_, _ = fmt.Scanln(&chatMsg)
	}
}

func (c *Client) UpdateName() bool {
	fmt.Println(">>>>>>请输入用户名")
	_, _ = fmt.Scanln(&c.Name)

	sendMsg := "rename|" + c.Name + "\n"
	_, err := c.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}

	return true
}
