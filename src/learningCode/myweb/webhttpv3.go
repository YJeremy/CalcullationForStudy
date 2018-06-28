package main

import (
	"io"       // handle 打印输出字符串内容时
	"log"      // 错误报错时
	"net/http" // web功能时
	"time"     //程序错误超时的时候
)

//通过map来保存你注册一个handler,然后再进行底层的serverhttp 来进行转发，这样效率是最高的，因为没有经过其他任何的封装了
//这已经是一个最底层的封装了

// 自定义一个名为mux字典，其中键是 string类型，键值是函数func类型，且该func类型有两个型参;
//分别是 http.ResponseWriter 类型（其实就是一个接口类型） 和 指针*http.Request
var mux map[string]func(http.ResponseWriter, *http.Request)

//定义主函数
func main() {
	//定义并初始化一个http.Server类型的变量 server，
	server := http.Server{
		//初始化里面的各个属性，其他的属性为默认
		Addr:    ":8080",
		Handler: &myHandler{}, // 指向一个自定义的结构体
		//且这个handle通过指针地址，与外部同名结构体类型，进行参数变化，实现不同的Handler值
		ReadTimeout: 5 * time.Second, //自定义一个超时时间
	}

	//路由注册

	//初始化字典
	mux = make(map[string]func(http.ResponseWriter, *http.Request))

	//单纯地使用一个访问路径作为一个路由，是最单纯的一种方法，
	mux["/hello"] = sayHello //键值是一个自定义函数，且传入的参数类型一致
	mux["/bye"] = sayBye

	err := server.ListenAndServe()
	// http.Server 类型里面，已初始化了addr,handler;这时调用它的方法ListenAndServe
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct{} //自定义结构体

//由于没有具体传入参数，所以省略传入参数，直接传类型指针，即若多个相同指针类型的函数都可以共用一个参数地址修改相同的值
func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//怎么知道mux？先判断mux
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		//直接转发 w r 去func； 查找 路径的正则是否有匹配的路由统一集中到同一个serveHTTP进行服务
		//找到相应注册的handle了，就进行转发
		//比如我自己在浏览器的地址栏输入地址名称 /hello ,那么就serveHPPT里去匹配查找到对应的handle
		//并传入 http.ResponseWriter 和  *http.Request

		return //返回不往下执行了
	}

	io.WriteString(w, "URL:"+r.URL.String())
	//后面其他的handle没有定义，r *http.Request是空的，但是内容w不空，会变化，根据URL由浏览器传入handle,再提取出来
}

//自定义的函数
func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Jeremy,Today is mother's Day!")
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Jeremy,Bye!")
}
