package main

import (
	"fmt"
)

func main() {
	//声明map,为非nil的map；声明
	var numbers map[string]int
	//可以用这种方式简便 ; 直接创建
	numbers1 := make(map[string]int)
	numbers1["Hello"] = 2
	//make map，变为nil才可以赋值
	numbers = make(map[string]int)
	numbers["ten"] = 10
	numbers["three"] = 3
	fmt.Printf("第三个数字是：%d\n", numbers["three"])
	fmt.Println("空map number:", numbers1)

	//初始化 + 赋值一体化
	rating := map[string]int{"C": 5, "Go": 4, "Python": 4, "C++": 2}
	//map 有两个返回值，key 和 key存在则true,不在false
	rating2, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is", rating2)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}
	delete(rating, "C") //删除Key为C的元素
	fmt.Println(rating)

	//说明map本质是引用 ，如果map引用map
	rating3 := rating
	rating3["C"] = 44
	fmt.Println("m", rating)

}
