package main

import (
	"fmt"
	"net"
	"strconv"
)

type Server struct {
	Ip   string
	Port int
}

// 创建一个服务的接口
func NewServer(ip string, port int) *Server {
	server := &Server{ip, port}
	return server
}

func (this *Server) Start() {
	//创建一个tcp监听器
	listenser, err := net.Listen("tcp", this.Ip+":"+strconv.Itoa(this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	fmt.Println("net.Listen ok:", this.Ip+":"+strconv.Itoa(this.Port))
	defer listenser.Close()
	for {
		//接收返回一个连接
		conn, err := listenser.Accept()
		if err != nil {
			fmt.Println("listenser.Accept err:", err)
			continue
		}
		// do handler  阻塞循环
		go this.Handler(conn)
	}
	fmt.Println("net.Listen end")
}

// 处理连接的业务部分
func (this *Server) Handler(conn net.Conn) {
	fmt.Println("Handler success")
	//
}
