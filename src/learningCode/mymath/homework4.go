
//export GOPATH=/home/jeremy/mygo/development
package main

/*
 6: 0110
11: 1011
-----------
&   0010 = 2
|   1111 = 15
^   1101 = 13
&^  0100 = 4 (第二位是1,则第一位数强制为0)
*/

import "fmt"
const(
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
)
func main(){

	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println("GB is =",GB)

}

