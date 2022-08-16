package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

func NewServer(ip string, port int) *Server {

	return &Server{
		Ip:   ip,
		Port: port,
	}

}

func (s *Server) Handle(conn net.Conn) {
	// 处理业务逻辑
	fmt.Println("连接建立成功")
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
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("close listener error:", err)
		}
	}(listener)
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
