package main
import(
	"fmt"
	//"reflect"
)

type User struct{
	Id int
	Name string
	Age int
}

func (u *User)Read()int{
	u.Id = u.Id + 100
	return 10
}

func main(){
	a := User{1 ,"Jeremy", 29}
	fmt.Println(a)
	b := a.Read()
	fmt.Println(b)
}
