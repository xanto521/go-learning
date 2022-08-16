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
	conn   net.Conn
	server *Server
}

//创建一个用户的api
func NewUser(conn net.Conn, server *Server) *User {

	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	//启动监听
	go user.ListenMessage()
	return user
}

// 用户上线功能
func (u *User) Online() {
	// 用户上线, 加入onlinemap, 然后进行广播
	u.server.mapLock.Lock()
	u.server.OnlineMap[u.Name] = u
	u.server.mapLock.Unlock()
	// 广播当前用户上线消息
	u.server.Broadcast(u, "上线了")
}

// 用户下线功能
func (u *User) Offline() {
	// 用户下线, 将用户从onlinemap中删除, 然后进行广播
	u.server.mapLock.Lock()
	u.server.OnlineMap[u.Name] = u
	delete(u.server.OnlineMap, u.Name)
	u.server.mapLock.Unlock()
	// 广播当前用户上线消息
	u.server.Broadcast(u, "下线了")
}

func (u *User) SendMsg(msg string) {
	u.conn.Write([]byte(msg + "\n"))
	//u.C <- msg
}

// 用户消息处理功能
func (u *User) DoMessage(msg string) {

	if msg == "who" {
		// who命令返回当前在线用户列表
		u.server.mapLock.RLock()
		for _, user := range u.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + "在线中..."
			u.SendMsg(onlineMsg)
			// u.C <- user.Name
		}
		u.server.mapLock.RUnlock()

	} else {
		// 正常处理广播消息
		u.server.Broadcast(u, msg)

	}

}

// 监听channel中的消息发送给客户端
func (u User) ListenMessage() {
	for {
		msg := <-u.C
		u.conn.Write([]byte(msg + "\n"))
	}
}
