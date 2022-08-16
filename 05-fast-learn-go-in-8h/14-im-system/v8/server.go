package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
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
	sendMsg := "[" + user.Addr + "]" + user.Name + "> " + msg
	s.Message <- sendMsg
}

func (s *Server) Handle(conn net.Conn) {
	// 处理业务逻辑
	fmt.Println("连接建立成功")
	defer fmt.Println("连接已关闭")
	// 创建一个用户
	user := NewUser(conn, s)

	user.Online()

	// 监听用户是否活跃的channel
	isOnline := make(chan bool)

	// 创建一个goroutine接收用户传来的消息
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			// 如果异常不为空 且没有正常读取 正常读取完毕是EOF
			if err != nil && err != io.EOF {
				fmt.Println("read error:", err)
				return
			}
			// 读取用户消息 去除'\n'
			msg := string(buf[:n-1])
			// 发送给所有在线用户
			user.DoMessage(msg)

			// 用户的任意消息都是活跃的
			isOnline <- true
		}
	}()

	// 阻塞防止用户进程退出
	for {
		select {
		case <-isOnline:
			// 用户活跃
			// 不做任何操作 为了select执行下面的定时器
		case <-time.After(300 * time.Second):
			// 用户不活跃
			// 关闭用户
			user.SendMsg("你已经10秒没有活跃了, 请重新登录")
			// 销毁用户资源
			close(isOnline)
			close(user.C)
			conn.Close()
			return
			//runtime.Goexit()
		}
	}
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
