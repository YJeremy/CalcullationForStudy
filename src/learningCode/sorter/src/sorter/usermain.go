package main

import (
	//"algorithms/SelectionSort"
	"fmt"
)

func main() {
	//a := []int{23, 53, 11, 23, 23, 44, 54, 76, 4}
	//fmt.Println(a)
	//SlectionSort.SlectionSort(a)
	//fmt.Println(a)

	b := []int{91, 60, 96, 13, 35, 65, 46, 65, 10, 30, 20, 31, 77, 81, 22}
	fmt.Println(b)
	HeapAdjust(b)
	//HeapInit(b)
	fmt.Println(b)
	HeapAdjust(b)
	fmt.Println(b)

}

func HeapAdjust(a []int) { //调整大顶堆
	k := 0
	temp := a[0] //取出当前元素
	l := len(a)
	m := l - 1

	//从下往上，右往左传递
	//以i=1 开始，而不是i=0,结果都一样，因为最初的一层a[0] a[1] a[2]，并不只有这三个元素，会被后来加入进来的元素把a[1]往后推，再进行比对
	//
	for i := 1; i < l; i = i*2 + 1 { //从左字节点开始；第二次遍历，就是子节点的子节点
		if i < m && a[i] < a[i+1] { //若左字节点小于右字节点，i指向右字节点
			i++
		}
		if a[i] > temp { //如果字节点大于父节点，将子节点值赋给父节点，不用进行交换
			a[k] = a[i] //父节点位置 由子节点值替代了了
			k = i       //字节点的下标取出，记住位置
		} else {
			break //跳出for
		}
	}
	a[k] = temp // 将temp值放到最终的位置(原字节点下标位置)
}

/*
输出的是一个最大顶堆的子树！
多次调用该（递归），就是对树的扩展，加一个顶点
*/

func HeapInit(a []int) { //构造初始堆
	l := len(a)
	for i := l / 2; i >= 0; i-- { //从第一个非子叶结点从下至上，从左至右调整结构
		//数组一半是最底层
		HeapAdjust(a[i:]) //i到末尾下表的元素
	}
}

/*
1、堆的构造，就是建立完全二叉树
2、把l/2 往前退的数，作为一个个从顶点的位置，从根节点位置加入树里；
3、一维数组就是树扩展是从根部加入元素，然后再通过内部函数的各种排序，由小树变成大树；
3、

*/

func HeapSort(a []int) {
	HeapInit(a)
	//2.调整堆结构+交换堆元素与末尾元素
	for i := len(a) - 1; i > 0; i-- {
		a[i], a[0] = a[0], a[i] //将堆顶元素与末尾元素进行交换
		HeapAdjust(a[:i])       //重新调整堆
	}
	fmt.Println(a)
}

func HeapSort2(s []int) {
	N := len(s) - 1
	for k := N / 2; k >= 1; k-- {
		sink(s, k, N)
	}
	for N > 1 {
		swap(s, 1, N)
		N--
		sink(s, 1, N)
	}
}

func sink(s []int, k, N int) { //K 序标
	for {
		i := 2 * k //左子树
		if i > N { //查询子树的位置超出范围了（没有子树）
			break
		}
		if i < N && s[i+1] > s[i] { //若右节点大，则使用右节点
			i++
		}
		if s[k] >= s[i] { //父节点比较，父节点大则退出
			break
		}
		swap(s, k, i) //交换
		k = i         //子树的序标 与 父节点的序标交换
		//第二轮遍历，由经过交换的那个子节点 的 下面开始比较，兄弟节点及其下面的树不用动，因为兄弟节点之间大小无关；

	}

}

func swap(s []int, i int, j int) {
	s[i], s[j] = s[j], s[i]
}

/*
比较 k 2k 2k+1 三者的大小

*/
