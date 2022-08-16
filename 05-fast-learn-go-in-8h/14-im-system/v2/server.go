package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int
	//在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	// 服务端的消息广播channel
	Message chan string
}

func NewServer(ip string, port int) *Server {

	return &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

}

// 监听广播消息channel的goroutine, 一旦有消息就发送给所有在线用户
func (s *Server) ListenMessage() {
	for {
		msg := <-s.Message
		s.mapLock.RLock()
		for _, user := range s.OnlineMap {
			user.C <- msg
		}
		s.mapLock.RUnlock()
	}
}

// 广播消息
func (s *Server) Broadcast(user *User, msg string) {
	sendMsg := "[" + user.Name + "]" + msg
	s.Message <- sendMsg
}

func (s *Server) Handle(conn net.Conn) {
	// 处理业务逻辑
	fmt.Println("连接建立成功")
	// 创建一个用户
	user := NewUser(conn)
	// 用户上线, 加入onlinemap, 然后进行广播
	s.mapLock.Lock()
	s.OnlineMap[user.Name] = user
	s.mapLock.Unlock()
	// 广播当前用户上线消息
	s.Broadcast(user, "上线了")

	// 阻塞防止用户进程退出
	select {}
}

//启动服务
func (s *Server) Start() {
	// stock listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	// close listener
	defer listener.Close()

	// 启动监听
	go s.ListenMessage()

	for {
		//accept

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		//handle
		go s.Handle(conn)

	}

}
