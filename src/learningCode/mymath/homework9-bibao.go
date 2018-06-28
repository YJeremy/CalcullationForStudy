package main
import(
	"fmt"
)

func main(){
	b := func(){
		fmt.Println("Func b")
	}
	b()

	a := closure(10)
	fmt.Println(a(1)) //这时是一个a(10)的 返回 func函数
	fmt.Println(a(2))
}

//匿名函数func（）
// 函数本身作为一种定义 a = func(){... },a 就是函数类型
// 函数定义格式 ： func 函数名 左括号 变量 变量类型 右括号  返回值 返回值类型 {}
// 其中返回值可以是一个函数（该返回函数又带返回值），如下闭包
//闭包调用


func closure(x int) func(int) int{
	fmt.Printf("%p\n", &x)
	return func(y int) int{
		fmt.Printf("%p\n", &x)
		return x + y
	}
}
