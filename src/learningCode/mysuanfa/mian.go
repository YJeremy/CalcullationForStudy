package main

import (
	"fmt"

	"mysuanfa/SelectionSorts"
)

func main() {
	a := []int{11, 22, 11, 23, 23, 44, 64, 21, 53, 72}
	fmt.Println(a)
	SelectionSorts.SelectionSort(a)
	fmt.Println(a)

}
