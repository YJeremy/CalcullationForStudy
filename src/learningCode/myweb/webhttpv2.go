package main

import (
	"io"       //handle 写屏幕输出的内容时
	"log"      // 记录出错日记时
	"net/http" // 传建web时
	"os"       //获得路径;因为有函数需要绝对路径；创建静态文件服务器时
)

func main() {
	mux := http.NewServeMux() //创建 NewServeMux

	mux.Handle("/", &myHandler{}) //"/"注册路由

	mux.HandleFunc("/hello", sayHello) //使用函数来handle

	wd, err := os.Getwd() //获取当前目录（以相对路径来获取当前绝对路径）。 创建静态文件路径，静态文件实现
	//若是其他文件不在同一路径下，需要自己另外配添加置文件说明路径哦。

	// wd 也有可能会有错误，所以也做打印错误的判断
	if err != nil {
		log.Fatal(err)
	}

	//实现静态文件服务器
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(wd))))
	//StripPefix(r,h)函数 第一参数，去掉 “static”前缀，以防混淆绝对路径; 第二参数，传入handle转发
	//FileServer（root） 就是一个handler, 它需要传入一个“目录”
	//http.Dir 方法，定位目录；再传入 wd 这个相对路径得到的绝对路径，然后转换为静态服务器的路径

	err = http.ListenAndServe(":8080", mux) //把mux传入原本是nil handle的参数处
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { //取指针，多个函数同时控制

	io.WriteString(w, "URL: "+r.URL.String()) //这就是传入request的原因，有URL

}

//定义handle 内容
func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello ver2")
}

/*直接用函数作为handler传进去  handlefunc -- sayhello
1、有两种handle 一个是mux.HandFunc的，其参数是
2、一种handle是http.ListenAndServe的参数传入


*/
