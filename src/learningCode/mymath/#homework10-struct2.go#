
package main
import(
	"fmt"
)
type person struct{
	Name string
	Age int
	Contact struct{
		Phone, City string
	}
}


func main(){
	a := person{Name: "Jeremy", Age: 29,} //对于结构体的字段可以直接声明，但是字段是结构体里面的内容，则声明不;只能外层的字段声明；
	a.Contact.Phone = "1231231321"
	a.Contact.City = "Beijing"
	fmt.Println(a)
	
}
