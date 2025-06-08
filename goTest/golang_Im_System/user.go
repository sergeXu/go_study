package main

import "net"

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	user := &User{
		//拿到客户端连接的地址和名称
		Name:   conn.RemoteAddr().String(),
		Addr:   conn.RemoteAddr().String(),
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	go user.ListenMessage()
	return user
}

// 用户上线
func (this *User) Online() {
	//用户上线
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()
	//广播用户上线消息
	this.server.BroadCast(this, "已上线")
}

func (this *User) Offline() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()
	//广播用户上线消息
	this.server.BroadCast(this, "已下线")
}

// 处理用户端发送消息
func (this *User) DoMessage(msg string) {
	//数据广播处理
	this.server.BroadCast(this, msg)
}

// 监听消息进行发送
func (this *User) ListenMessage() {
	//循环
	for {
		msg := <-this.C
		//网络数据传输
		this.conn.Write([]byte(msg + "\n"))
	}
}
