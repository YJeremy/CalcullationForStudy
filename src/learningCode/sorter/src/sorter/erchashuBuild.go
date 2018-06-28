package main

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

//
func (node *Node) Print() {
	fmt.Print(node.Value, "")
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("setting value to nil" + " node.Ignored.")
		return
	}
	node.Value = value
}

//
func (node *Node) preOrder() {
	if node == nil {
		return
	}
	node.Print()
	node.Left.preOrder()
	node.Right.preOrder()
}

func (node *Node) middleOrder() {
	if node == nil {
		return
	}
	node.Left.middleOrder()
	node.Print()
	node.Right.middleOrder()
}

//
func (node *Node) postOrder() {
	if node == nil {
		return
	}
	node.Left.postOrder()
	node.Right.postOrder()
	node.Print()
}

func CreateNode(value int) *Node {
	return &Node{Value: value}

}

func main() {
	root := Node{Value: 3}
	root.Left = &Node{} // 设置了零值，是0,而不是nil
	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node)
	root.Right.Left.SetValue(4)
	root.Left.Right = CreateNode(2)

	fmt.Print("前序遍历：")
	root.preOrder()
	fmt.Println()
	fmt.Print("中序遍历：")
	root.middleOrder()
	fmt.Println()
	fmt.Print("后序遍历：")
	root.postOrder()
	fmt.Println()

}
