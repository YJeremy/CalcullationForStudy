package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	//注册函数
	//注册路由 “/”
	// 设置路由;设置访问路由的规则，用http包
	http.HandleFunc("/", sayHello) //第一个参数相当有beego的routepath; 访问什么路径，就用什么handler
	//第二个参数叫做handler 则是路由注册信息
	// http.HandleFunc(ResponseWriter, *Request)
	// 一个自定义“接口”类型，一个结构体指针（因为结构体节省开销用指针）

	err := http.ListenAndServe(":8080", nil) //告诉机器，我注册完成了，可以开始运行了
	//add= "：8080" 是本地地址，handler= "nil" 表示使用默认的handler
	//这个handler 是另一一回事是 servemux,不是注册的handler
	//下一个版本则自定义这个地方

	if err != nil {
		log.Fatal(err)
	}

}

//路由要求什么参数，就对应给什么参数，才会把路由注册成功,所以自定以一个含符合类型的函数
func sayHello(w http.ResponseWriter, r *http.Request) {
	//对参数进行写入，beego 用run string

	io.WriteString(w, "hello ver1")

	//io包下 func : WriteString(w Writer, s string)(n int, err error){}
	/* w参数的 type Writer是个定义的接口类型:
	    type Writer interfac{
	    Write(p []byte)(n int, err error)
	   }

	*/

}
