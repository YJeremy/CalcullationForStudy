/*************
package main
import(
	"fmt"
	"reflect"
)

type User struct{
	Id int
	Name string
	Age int
}

type Manager struct{
	User
	title string
}

func main(){
	m := Manager{User: User{1, "OK", 3}, title: "123"}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.Field(0))
	//反射匿名或者嵌入字段

	x := 123
	v := reflect.ValueOf(&x)//传入地址，以进行修改
	v.Elem().SetInt(999) //必须传入地址才可以使用，已知123是Int型

	fmt.Println(x)

}


****************************************************************************************/
/***************************************************************************************
package main()
import(
	"fmt"
	"reflect"
)

//对于结构，有很多基本类型，先分解出各种类型

type User struct{
	Id int
	Name string
	Age int
}
func main(){
	u := User{1, "OK", 12}
	Set(&u) //调用自定义函数，取地址，修改；使用接口对象，用反射进行修改
	fmt.Println(u) //输出信息
	
}

func Set(o interface{}){
	v := reflect.ValueOf(o) //取反射参数对象的值
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet(){
		//使用对象方法Kind() 判断是否是Point Interface 指针接口
		//kind正常返回是 reflect.Struct的类型
		//Elem().CanSet() 返回布尔值 方法判断对象是否可以修改,这些工作是正规的工作
		fmt.Println("XX")
		return
	} else{
		v = v.Elem()//修改对象
	}

	f := v.FieldByName("Name")//引用对象里面的”Name“字符命名的属性
	if !f.IsValid(){//是否存在f
		fmt.Println("BAD")
		return
	}

		if; f.Kind() == reflect.String{//类型可以修改
			f.SetString("BYEBYE") //设置修改改属性
		}
}


************************************************************************************88*/

//**************************************************************************************

package main
import(
	"fmt"
	"reflect"
)

type User struct{
	Id int
	Name string
	Age int
}

func (u User)Hello(name string){
	fmt.Println("Hello",name,"my name is",u.Name)
}

func main(){
	u := User{1,"OK",12}
	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")

	args := []reflect.Value{reflect.ValueOf("joe")}

	mv.Call(args)
}


//*************************************************************************************/
