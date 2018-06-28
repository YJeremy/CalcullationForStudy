package main

import "fmt"

func main() {
	var arr [10]int
	arr[0] = 44
	arr[1] = 13
	fmt.Printf("The [10]int first element arr[0] is %d\n", arr[0])
	fmt.Printf("The [10]int last element arr[9] is %d\n", arr[9])
	//slice和声明array一样，只是少了长度;它是一个切片数组，要基于某个数组基础上切片;它是一个空间里引用;
	//但当 append 元素，使其长度超出原本引用数组的空间 cap()>len(),那么将分配新的空间给切片，且不再影响原来引用;
	var fslice []byte
	mm := [4]byte{'a', 'd', 'g', 'k'}

	//如此先定义一个基础数组，然后在上面取值，而不是直接当作数组用
	fslice = mm[:] // 切片全部数组，不要只写a
	fislic2 := []byte{'k', 'b'}
	var a, b []byte
	a = fslice[1:3]
	b = fslice[0:2]
	fmt.Printf("mm= %d\n", mm)
	fmt.Printf("fslice=%d\n,fislic2=%d\n", fslice, fislic2) //打印出了的是地址
	fmt.Printf("cap a = %d\n a=%d", cap(a), a)
	fmt.Printf("b = ", b, "\n")

}
