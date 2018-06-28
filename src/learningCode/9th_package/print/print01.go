package main

import "fmt"

func main() {
	var i = 42
	var s = "Answer"
	fmt.Print(s, "is", i, 3.14, '\n', "\n")
	fmt.Println("=====================================")
	fmt.Println(s, "Print line", i, "\n", 3.14, '\n')
	fmt.Println("=====================================")
	fmt.Printf("Printf:\n", " %s is %d\n %f%c %x\n", s, i, 3.14, '\n', '\n', "\n")
	fmt.Println("=====================================")
	fmt.Printf("Printf 2:\n  %s is %d\n %f%c %x\n", s, i, 3.14, '\n', '\n')
	fmt.Println("=====================================")
	fmt.Printf("精度输出 %6d and %0.4f", 42, 3.14159)
	fmt.Println("\n=====================================")
	fmt.Printf("Say %6.4s!", "hello", "\n")
	fmt.Printf("Say %*.4s!", 6, "hello", "\n")

}
