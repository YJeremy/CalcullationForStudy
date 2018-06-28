package main
import(
	"fmt"
)
func main(){
	var m0 map[string]int = make(map[string]int)
	m1 := map[int]string{1:"o", 2:"t", 3:"t", 4:"f"}
	fmt.Println(m1)
	for k,v := range m1{
		m0[v] = k
	}
	fmt.Println(m0)
}
 
