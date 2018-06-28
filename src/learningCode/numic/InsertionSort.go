/**
1、了解数组的声明写法
2、了解函数的声明写法
2.1 交换两个数的写法
3、了解for 的写法
4、了解 插入排序 算法
4.1 从第一个元素开始，该元素可以认为已经被排序
4.2 取出下一个元素，在已经排序的元素序列中从后向前扫描
4.3 如果该元素（已排序）大于新元素，将该元素移到下一位置
4.4 重复步骤3,直到找到已排序的元素小于或者等于新元素的位置
4.5 将新元素插入到该位置后
4.6 重复步骤2～5
**/
package main

import (
	"fmt"
)

func insertSort_min(p []int) {
	a := len(p)
	for i := 1; i < a; i++ {
		for j := i - 1; j >= 0 && p[j+1] < p[j]; j-- {
			p[j+1], p[j] = p[j], p[j+1]
		}
	}
}

func insertSort_max(p []int) {
	a := len(p)
	for i := 1; i < a; i++ {
		for j := i - 1; j >= 0 && p[j+1] > p[j]; j-- {
			p[j+1], p[j] = p[j], p[j+1]
		}
	}
}

func main() {
	a := []int{1, 3, 4, 2, 8}
	insertSort_min(a)
	fmt.Println("Min Sort is :", a)
	insertSort_max(a)
	fmt.Println("Max Sort is :", a)

}
