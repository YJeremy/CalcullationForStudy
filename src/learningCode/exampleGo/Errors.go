package main

import "fmt"
import "errors"

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42") //重新定义
	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string { //因为是传入地址了，理解为结构体的逻辑操作
	return fmt.Sprintf("%d-%s", e.arg, e.prob) //sprintf
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"} //取地址的值，理解为为结构体添加“逻辑操作”的函数
		//把传入型参 arg,作为结构体指针argerror的元素 arg值，“can't”作为 元素prob的值，然后地址指针自动执行函数Error()返回值
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} { // 定义i 在切片元素内取值，下标值不取
		if r, e := f1(i); e != nil { //定义两个变量，来存f1（i） 返回的两个值

			fmt.Println("f1 failed:", e) //输出e ，非空的内容

		} else {
			fmt.Println("f1 worked", r) //输出r 的值
		}

	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e) //error 类型是一串字符串
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)                   //接受f2() 的第二个返回值 42-can't work with it
	if ae, ok := e.(*argError); ok { //
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}
