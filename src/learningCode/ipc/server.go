package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "method"
	Params string "params"
}

type Response struct {
	Code string "code"
	Body string "body"
}

//定义服务器功能的接口类型
type Server interface {
	Name() string //返回值类型记得也要写上，会区别同名函数
	Handle(method, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer { //把函数的实参的值 传入IpcServer结构体类型的地址指针
	return &IpcServer{server}
}

func (server *IpcServer) Connect() chan string { //IpcServer 方法，返回一个字符串
	session := make(chan string, 0) //通道，0缓存？

	go func(c chan string) { //匿名函数 ，型参是通道类型 ，无返回值

		for {
			request := <-c

			if request == "CLOSE" { //关闭链接
				break
			}
			var req Request                              //Requset结构体型的req
			err := json.Unmarshal([]byte(request), &req) //把request类型 强制转换为byte切片
			if err != nil {
				fmt.Println("Invalid request format:", request)
			}

			resp := server.Handle(req.Method, req.Params)

			b, err := json.Marshal(resp)

			c <- string(b)

		}
		fmt.Println("Sesssion closed.")
	}(session)

	fmt.Println("A new session has been created successfully.")

	return session

}