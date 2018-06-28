package main

import (
	"fmt"
)

func main() {
	num := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 3, 46, 4, 19, 50, 48}
	fmt.Println(num)
	sort(num)
	fmt.Println(num)

}

/*
[3 44 38 5 47 15 36 26 27 3 46 4 19 50 48]
[3 3 4 5 15 19 26 27 36 38 44 46 47 48 50]
*/

//交换的步骤，耗时不随n改变
//使用切片,直接通过数据类型自身做引用操作，无需返回值
//传入切片，以及下标的形式
//小写开头的函数，只在包内调用；协助包内其他函数；
//把步骤提炼出来；函数式编程；分步；
func swap(slice []int, i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func sort(a []int) {
	l := len(a)              //以免每次for遍历都要算有一次len;给出排序结束的条件；
	m := len(a) - 1          //总长度少1个首个元素
	for i := 0; i < m; i++ { //该层遍历是作为对整个序列的修改

		//该层遍历是作为对比查找，不做修改
		k := i                       //值传递，用作对比的基准
		for j := i + 1; j < l; j++ { //大于排序i的所有元素进行遍历搜索;j比i大1,所以用l判断
			if a[j] < a[k] {
				k = j //更改对比的基准
			}
		}

		if k != i {
			swap(a, k, i)
		}

	}

}

/*

 */
