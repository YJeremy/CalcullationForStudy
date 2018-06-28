//通过反射来查看一个结构体的内容
package main
import(
	"fmt"
	"reflect"
)

type User struct{
	ID int
	Name string
	Age int
}

func (u User)Hello(){ //函数书写 （）成为funtion body
	fmt.Println("Hello.")
}


func main(){
	u := User{1, "OK", 12}
	Info(u)
}

func Info(o interface{}){
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")
	for i := 0; i<t.NumField(); i++{
		f:= t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}
	for i := 0;i< t.NumMethod(); i++{
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}
/*
Type: User
Fields:
	ID:int =1
	Name:string = OK
	Age:int = 12
	Hello: func(main.User)
*/
