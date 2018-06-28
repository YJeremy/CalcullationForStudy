package cg

import (
	"fmt"
)

type Player struct {
	Name  string "name"
	Level int    "level"
	Exp   int    "exp"
	Room  int    "room"

	mq chan *Message //等待收取的消息
}

func NewPlayer() *Player {
	m := make(chan *Message, 1024)
	player := &Player{"", 0, 0, 0, m}

	go func(p *Player) {
		for { //单个for,不带表达式，表示无限循环

			msg := <-p.mq
			fmt.Println(p.Name, "received message:", msg.Content)
		}
	}(player) //括号后面直接跟参数列表表示函数调用

	return player
} //由通道来传递 “通讯”，在通道里死循环等候～

//每个人都起了独立的goroutine, 定义在结构体方法里面的goroutine!
