package main

import "net"

type User struct {
	// 用户名
	Name string
	// 用户的地址
	Addr string
	// 用户的通道, 负责接收服务端发送的消息
	C chan string
	// 用户的连接对象
	conn net.Conn
}

//创建一个用户的api
func NewUser(conn net.Conn) *User {

	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}
	//启动监听
	go user.ListenMessage()
	return user
}

// 监听channel中的消息发送给客户端
func (u User) ListenMessage() {
	for {
		msg := <-u.C
		u.conn.Write([]byte(msg + "\n"))
	}
}
