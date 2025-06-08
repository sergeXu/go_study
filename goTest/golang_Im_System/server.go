package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"
	"time"
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
	user := NewUser(conn, this)
	user.Online()
	//活跃度保持chan
	isLive := make(chan bool)
	//处理用户的信息输入
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if err != nil && err != io.EOF {
				fmt.Println("conn.Read err:", err)
				return
			}
			if n <= 0 {
				user.Offline()
				return
			}

			msg := string(buf[:n-1])
			//处理用户输入
			user.DoMessage(msg)

			//激活活跃度
			isLive <- true
		}
	}()
	//当前handler阻塞，处理用户输入处理
	for {
		select {
		case <-isLive:
			//处理活跃度队列，激活select，更新下面定时器
		case <-time.After(time.Second * 10):
			user.SendMessage("你被踢了")
			user.Offline()
			//channel 资源关闭
			close(user.C)
			//关闭连接
			conn.Close()
			return
		}
	}
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
