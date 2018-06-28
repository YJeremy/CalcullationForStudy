package main

import "fmt"

func max(a, b int) int { //变量类型省略，返回值省略写法
	if a > b {
		return a
	}
	return b
}

//多个返回值例子
func SumAndProduct(A, B int) (add int, Multiplied int) { //(返回值可以省略，但是为了可读性，要写好)
	add = A + B
	Multiplied = A * B
	return //因为返回值写了名字，所以return就自动把对应名字返回了。
}

//传指针使得多个函数能够操作同一个对象
//传指针比较轻量级（8bytes），只是传内存地址，可以用指针传递体积大的结构体。如果用参数传递的话，每次copy上面就会话费较多的系统开销（内存和时间）
//

//简单的一个函数，实现了参数+1的操作

func add(a *int) int { //返回的是int ,不是指针类型
	*a = *a + 1
	return *a
}

func main() {
	x := 3
	y := 4
	z := 5

	max_xy := max(x, y)
	max_xz := max(x, z)
	xPLUSy, xTIMESy := SumAndProduct(x, y)

	fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
	fmt.Printf("max(%d, %d) = %d\n", x, z, max_xz)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y, z)) //直接调用，函数也是一种“对象的实例”

	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)

	fmt.Printf("x+1 =%d\n", add(&x))
	fmt.Printf("x = %d\n", x)

}
