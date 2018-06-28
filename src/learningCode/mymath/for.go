package main

import "fmt"

func main() {
	sum := 0
	for index := 0; index < 10; index++ { // expression1循环开始用;expression2判断;expression3循环结束用
		sum += index
	}
	fmt.Println("sum is equal to ", sum)

	sum1 := 1
	for sum1 < 100 { //只有expression2
		sum1 += sum1
	}
	fmt.Println("sum1 is equal to ", sum1)

	sum2 := 1
	for sum2 < 1000 { //都省略;号，只有expression2，即while功能
		sum2 += sum2
	}
	fmt.Println("sum2 is equal to ", sum2)

	/********************************************************************/
	for index := 10; index > 0; index-- {
		if index == 5 {
			break //或者continue
		}
		fmt.Println(index)
		// break print 10 9 8 7 6
		//continue 10 9 8 7 6 4 3 2 1
	}
	/*
		for k,v := rang map{
			fmt.Println("map's key:", k)
			fmt.Println("map's val:", v)
		}
	*/

}
