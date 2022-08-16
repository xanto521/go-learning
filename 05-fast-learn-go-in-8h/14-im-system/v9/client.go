package main

import (
	"flag"
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
	flag       int //当前客户端的模式
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}

	// 连接服务端
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("连接服务端失败")
		return nil
	}
	client.conn = conn

	// 返回对象
	return client

}

// 处理server返回的消息
func (client Client) DealResponse() {
	// 当连接收到消息后直接打印到标准输出
	io.Copy(os.Stdout, client.conn)

}

func (client *Client) menu() bool {
	// 获取命令菜单
	fmt.Println("1. 公聊发送消息")
	fmt.Println("2. 私聊发送消息")
	fmt.Println("3. 更新用户名")
	fmt.Println("0. 退出系统")
	fmt.Println("请选择:")
	var choice int
	fmt.Scanln(&choice)
	if choice >= 0 && choice <= 3 {
		client.flag = choice
		return true
	} else {
		fmt.Println(">>>请输入合法的选项<<<")
		return false
	}

}
func (client *Client) SelectAllUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println(">>> client.conn.Write error,查询所有用户失败!", err)
		return
	}
}

// 私聊模式
func (client *Client) PrivateChat() {
	var toUserName, chatMsg string
	client.SelectAllUsers()
	fmt.Println(">>> 请输入私聊用户名(exit退出):")
	fmt.Scanln(&toUserName)
	for toUserName != "exit" {
		// 发送消息逻辑
		fmt.Println(">>> 请输入消息内容(exit退出):")
		fmt.Scanln(&chatMsg)
		for chatMsg != "exit" {
			// 执行发送消息
			if len(chatMsg) != 0 {
				// 发送到服务器
				sendMsg := "to|" + toUserName + "|" + chatMsg + "\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println(">>> client.conn.Write error,发送私聊消息失败!", err)
					break
				}
			}
			chatMsg = ""
			client.SelectAllUsers()
			fmt.Println(">>> 请输入消息内容(exit退出):")
			fmt.Scanln(&chatMsg)
		}
		fmt.Println(">>> 请输入私聊用户名(exit退出):")
		fmt.Scanln(&toUserName)
	}
}

func (client *Client) PublicChat() {
	// 发送公聊消息
	var chatMsg string
	fmt.Println(">>> 请输入公聊消息(exit退出):")
	fmt.Scanln(&chatMsg)
	for chatMsg != "exit" {

		if len(chatMsg) != 0 {
			// 发送到服务器
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println(">>> client.conn.Write error,发送公聊消息失败!", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println(">>> 请输入公聊消息(exit退出):")
		fmt.Scanln(&chatMsg)
	}
}
func (client *Client) UpdateName() bool {
	fmt.Println(">>> 请输入用户名:")
	fmt.Scanln(&client.Name)
	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println(">>> client.conn.Write error,更新用户名失败!", err)
		return false
	}
	return true
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}

		// 根据不同的选项执行不同的业务
		switch client.flag {
		case 1:
			// 公聊模式
			fmt.Println(">>> 公聊模式")
			client.PublicChat()
			break
		case 2:
			// 私聊模式
			fmt.Println(">>> 私聊模式")
			client.PrivateChat()
			break
		case 3:
			// 更新用户名
			fmt.Println(">>> 更改用户名")
			client.UpdateName()
			break
		}
	}

}

var serverIp string
var serverPort int

// 设置命令行参数 -ip -port
// usage: ./client -ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "服务端ip地址")
	flag.IntVar(&serverPort, "port", 8888, "服务端端口号")
}
func main() {
	// 命令行解析变量
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>> 创建客户端失败!")
		return
	}
	fmt.Println(">>> 连接服务端成功...")

	go client.DealResponse()

	// 开始业务流程
	client.Run()

}
