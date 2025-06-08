package main

import (
	"net"
	"strings"
)

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
	//消息指令处理，查询在线人员
	if msg == "who" {
		for _, v := range this.server.OnlineMap {
			onlineMsg := "[" + v.Addr + "]" + v.Name + ": 在线"
			this.SendMessage(onlineMsg)
		}
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//修改用户名称，前缀过滤判断
		//消息格式：rename|张三
		newName := strings.Split(msg, "|")[1]
		//判断map中是否已有
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMessage("当前用户名已被使用")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()
			this.Name = newName
			this.SendMessage("您已更新用户名：" + this.Name)
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		remoteNmae := strings.Split(msg, "|")[1]
		if remoteNmae == "" {
			this.SendMessage("格式不正确，参考 \"to|张三|你好啊 \"")
			return
		}
		remoteUser, ok := this.server.OnlineMap[remoteNmae]
		if !ok {
			this.SendMessage("目标用户不存在，请检查 ")
		}
		msg := strings.Split(msg, "|")[2]
		if msg == "" {
			this.SendMessage("无发送内容，请检查 ")
		} else {
			remoteUser.SendMessage(this.Name + " 对您说： " + msg)
		}

	} else {
		//数据广播处理
		this.server.BroadCast(this, msg)
	}
}

// 指定向user发送消息
func (this *User) SendMessage(msg string) {
	this.conn.Write([]byte(msg + "\n"))
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
