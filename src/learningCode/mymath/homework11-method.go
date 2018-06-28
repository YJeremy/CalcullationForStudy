package main
import(
	"fmt"
)

type A struct{
	Name string
}

type B struct{
	Name string
}

func main(){
	a := A{}
	a.Print()
	fmt.Println(a.Name)

	b := B{}
	b.Print()
	fmt.Println(b.Name)

}
// 定义一个函数（方法），该函数的参数是 A结构体的类型 可以使用
// 通过 以类型的指针定义变量，那么方法里面的变量所使用的改变，都会影响到原来的指针所指向地址的变量的值的
func (kkkk *A) Print(){
	kkkk.Name = "AA"
	fmt.Println("not just a")
	fmt.Println("Ause")
}

func (b B) Print(){
	b.Name = "BB"
	fmt.Println("B")
}
