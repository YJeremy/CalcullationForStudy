package main
import(
	"fmt"
)

func main(){
	var fs = [4]func(){} //声明含有4个元素为匿名函数的数组


	for i := 0; i<4 ; i++{
		defer fmt.Println("defer i =", i)
		defer func(){fmt.Println("defer_closure i - ", i)}()
		fs[i] = func() {fmt.Println("closure i = ", i)}
	
	}

	// defer后先进后出，直接记录操作数i，所以3 2 1 0
	// 闭包func调用次数是4次，defer调用最后的i值，所以为4
	//闭包作用，夺取了外层函数的i，即for循环i的引用地址，当第一次for执行完
	//i已经等于4了，然后在第个for里面的f()调用的i就是4了;

	for _,f := range fs{
		f()
	}
	
	//for定义变量，直接取fs数组的元素值，第一个下划线是空意思，指fs下标;
	//定义变量f，轮寻取f为 数组fs的一个元素，循环fs的元素个数次数
	// 调用f()
}
