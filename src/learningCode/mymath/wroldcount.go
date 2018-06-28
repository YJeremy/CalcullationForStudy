package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "I love my work and I"
	res := WordCount(s)
	fmt.Println(res)
}

func WordCount(s string) map[string]int {
	s_arr := strings.Fields(s)    //分割字符串为字符数组
	s_map := make(map[string]int) //建立map
	//对s_arr中的每个字符进行循环
	for i := 0; i < len(s_arr); i++ {
		if s_map[s_arr[i]] == 0 { //当还没有统计过该字符时，赋值为1
			s_map[s_arr[i]] = 1
		} else { //当统计过该字符时，更新计数值+1
			s_map[s_arr[i]] = s_map[s_arr[i]] + 1
		}
	}
	return s_map
}
