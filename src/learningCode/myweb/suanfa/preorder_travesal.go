package main

import (
	"errors"
	"fmt"
)

//展示栈的使用，非递归遍历树用
const stackSize int = 22 //栈的容量

func main() {
	var a [5]int
	fmt.Println("空数组展示:\n", a)
	fmt.Println()

	b := [5]int{1, 2, 3, 4, 5} //中括号
	fmt.Println("数组展示 b: \n", b)
	fmt.Println()

	c := []int{10, 2, 3, 3, 1, 8, 2, 23, 4, 45, 2, 3, 9, 54, 4}
	fmt.Println("切片展示 c:\n", c)
	fmt.Println()

	//显示一下二叉树的表示
	t := creatTree(c)
	fmt.Println("二叉树形式展示 c:\n", t)
	fmt.Println()

	//使用下面的用例定义
	//空指针引用，需要用new来创建返回是一个指针；否则会报错
	//创建一个结构体指针但是没有使用，不会报错not used？？？
	//有操作，把它的值从nil改变了都算！
	var root *Student = new(Student)

	root.Name = "stu01"
	root.Age = 18
	root.Score = 100

	var left1 *Student = new(Student)
	left1.Name = "stu02"
	root.Age = 20
	root.Score = 99

	root.left = left1

	var right1 *Student = new(Student)
	right1.Name = "stu04"
	right1.Age = 29
	right1.Score = 101

	root.right = right1

	var left2 *Student = new(Student)
	left2.Name = "stu03"
	left2.Age = 30
	left2.Score = 100

	left1.left = left2

	trans(root)
	fmt.Println()

}

//树的结构体表示，有节点，以及指向它的左右子树节点的地址
type tree struct {
	data  int
	left  *tree
	right *tree
}

//栈结构
type Stack struct {
	Size   int
	Values []*tree
}

//一个二叉树的类型应用
type Student struct {
	Name  string
	Age   int
	Score float32
	left  *Student
	right *Student
}

//函数创建实例;展示二叉树的递归来替代迭代；递归概念展示
func trans(root *Student) {
	if root == nil { //基线条件
		return
	}
	fmt.Println("前序遍历：", root) //前序遍历

	trans(root.left)
	//fmt.Println("中序遍历：", root) //中序遍历

	trans(root.right)
	//fmt.Println("后序遍历: ",root)//后序遍历
}

//把切片转为二叉树类型
func creatTree(arr []int) []tree {
	//先遍历切片，把所有元素都转换为可含有左右节点的结构体，放进新的切片
	d := make([]tree, 0)
	for i, ar := range arr {
		d = append(d, tree{})
		d[i].data = ar
	}

	//为左右节点添加内容，即指向的下级字节点的地址（并不是值）
	for i := 0; i < len(arr)/2; i++ {
		d[i].left = &d[i*2+1]
		if i*2+2 < len(d) {
			d[i].right = &d[i*2+2]
		}
	}

	//fmt.Println(d)
	return d

}

//创建栈
func creatStack() Stack {
	s := Stack{}
	s.Size = stackSize
	return s
}

//栈满; 方法展示
func (s *Stack) IsFull() bool {
	return len(s.Values) == s.Size
}

//栈空
func (s *Stack) IsEmpty() bool {
	return len(s.Values) == 0
}

//入栈;
func (s *Stack) Push(a *tree) error {
	if s.IsFull() {
		return errors.New("Stack is full,push failed la~~")
	}
	s.Values = append(s.Values, a)
	return nil
}

//出栈
func (s *Stack) Pop() (*tree, error) { // 出栈，并返回其值
	if s.IsEmpty() {
		return nil, errors.New("Out of index,len(stakck) is 0,空栈")
	}
	res := s.Values[len(s.Values)-1]
	s.Values = s.Values[:len(s.Values)-1]
	return res, nil
}

//查看栈顶元素
func (s *Stack) Peek() (*tree, error) {
	if s.IsEmpty() {
		return nil, errors.New("Out of index,len(stack) is 0,栈顶是空")
	}
	return s.Values[len(s.Values)-1], nil
}

//前序遍历
func preorder(root tree) {
	fmt.Print(root.data, " ")
	if root.left != nil {
		preorder(*root.left)
	}
	if root.right != nil {
		preorder(*root.right)
	}

}
