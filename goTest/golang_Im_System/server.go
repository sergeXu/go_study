package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	//消息广播channel
	Message chan string
}

// 创建一个服务的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
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
	//启动监听Message go程
	go this.ListenMessager()
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
	//用户上线
	this.mapLock.Lock()
	user := NewUser(conn)
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()
	//广播用户上线消息
	this.BroadCast(user, "已上线")
	//当前handler阻塞，处理用户输入处理
	select {}
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

// 监听Message广播消息，消息发送给全部User
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message
		this.mapLock.Lock()
		for _, v := range this.OnlineMap {
			v.C <- msg
		}
		this.mapLock.Unlock()
	}
}
