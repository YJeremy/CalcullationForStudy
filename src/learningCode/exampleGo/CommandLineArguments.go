package main

import "os"
import "fmt"

func main() {
	argsWithProg := os.Args
	argsWithoutProp := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProp)
	fmt.Println(arg)
}
