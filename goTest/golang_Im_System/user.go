package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

func NewUser(conn net.Conn) *User {
	user := &User{
		//拿到客户端连接的地址和名称
		Name: conn.RemoteAddr().String(),
		Addr: conn.RemoteAddr().String(),
		C:    make(chan string),
		conn: conn,
	}
	go user.ListenMessage()
	return user
}

// 监听消息发送
func (this *User) ListenMessage() {
	//循环
	for {
		msg := <-this.C
		//网络数据传输
		this.conn.Write([]byte(msg + "\n"))
	}
}
