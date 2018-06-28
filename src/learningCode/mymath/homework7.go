package main
import(
	"fmt"
)
func main(){
	s1 := make([]int, 3, 10)
	fmt.Println("%p", s1)
	fmt.Printf("%p", s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1)
	s2 := s1[:]
	fmt.Println("s2=", s2)
}
