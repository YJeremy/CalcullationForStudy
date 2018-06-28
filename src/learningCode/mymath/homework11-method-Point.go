//创建一个新的int类型的方法，每次i调用该方法可以使该类型的数自身增加100
package main
import(
	"fmt"
)

type MY int

func main(){
	var a,a2 MY = 0,2  //注意是一个类型重名，不是一个结构体定义
	a.Increase()
	a2.IncreaseN(10)
	fmt.Println(a)
	fmt.Println(a2)
}

// 创建方法注意指针
func (a *MY) Increase(){
	*a = *a + MY(100)
}

// 第二种方式，方法里面有型参，增加型参内的数
func (a *MY) IncreaseN(num int){             //注意型参要写类型
	*a += MY(num)
}
