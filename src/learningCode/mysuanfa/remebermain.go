package main

import "fmt"

func main() {
	a := []int{23, 33, 112, 3, 23, 5, 76, 886, 54}
	fmt.Println(a)
	sort(a)
	fmt.Println(a)

}

func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]

}

func sort(s []int) {
	//创建大小阔值
	l := len(s)
	m := l - 1
	for i := 0; i < m; i++ { //创建序标，与阔值产生迭代1
		k := i                       // 创建定标
		for j := i + 1; j < l; j++ { //创建大序标，与大阔值产生迭代2
			if s[j] > s[k] { // 定标与序标做对比
				k = j
			}

		}
		swap(s, k, i) //定标与序标做交换
	}

}
