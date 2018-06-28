package main
import (
	"fmt"
)

func main(){
	a := [10]int{}
	// 声明一个数组
	
	a[1] = 2
	// 给数组赋值

	fmt.Println(a)
	p := new([10]int)
	p[1] = 2
	// 声明一个数组指针（指向数组的指针），并赋值

	fmt.Println(p)
}
