/*实现嵌套接口的使用，以及说明结构体的范围大于接口类型的范围*/

package main
import (
	"fmt"
)
//定义USB接口类型,接口是一个或多个方法签名的集合
// 只有方法声明，没有实现，没有数据字段
// 只要某个类型拥有该接口的所有方法签名，即算实现该接口，无需显示声明实现了哪个接口，Structural Typing
type USB interface{

	Name() string //接口拥有的方法名，有返回值,注意类型必须写，且与方法定义的时候一致
	Connect() //接口拥有的方法名
	Inputer//嵌套接口
}

// 接口嵌套
type Inputer interface{
		Input() string //接口定义要一致;
}
func (a PhoneConnecter) Input() string{
	//b := "input user"
	fmt.Println("I m input",a.name)
	return a.name
}
//定义空接口，所有类型都是属于它，有点像继承的概念
type Myempty interface{
}

//定义结构体，为了配合上面USB接口的方法，创建一个结构体
type PhoneConnecter struct{
	name string //结构体内部的属性
}

//声明结构提Name方法，是一个自定义PhoneConnecter结构体类型方法且方法名已在USB接口中
func (pc PhoneConnecter) Name() string{
	return pc.name //该方法调用了结构体的属性
}

//声明结构体Connect方法，是一个自定义PhoneConnecter结构体类型方法且方法名已在USB接口中
func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name) //调用了结构体属性
}

//实现主函数
func main(){

	a := PhoneConnecter{"abc"} //由于USB 拥有了方法PhoneConnecter,所以无需声明 var a USB,
	//a已经带有USB接口类型了,abc要注意字符放入name属性中
	var a2 Inputer
	a2 = Inputer(a)	//对象赋值给接口时会发生拷贝，接口内部存储指向的是复制品的指针
	//a2.Connect() //这里是错误，
	//a2定义是Inputer接口类型变量，然后把大范围的接口类型a，强制转换小范围，因此不能使用Connect()
	//大可以转小，小不可以转大,虽然它是结构体a;
	//只是把该方法（接口）给限制了，但是变量还是这个接口类型;
	a.Connect()
	USBcheck(a)
	InterfaceCheck(a)
	InterfaceCheck(a2)
	a.name = "pc"
	a2.Input()//尽管a.name已修改，但是a2里面的.name还是不变;
	//对象赋值给接口时会发生拷贝，接口内部存储指向的是复制品的指针


}

func USBcheck(device USB){
	hi, good := device.(PhoneConnecter) //此行也可以写在if 后面，定义两个变量,类型判断device.(),
	// if hi,good := device.(PhoneConnecter),good{  ，初始化表达式注意写法！多一个点多一个判断
	//这是一个判断是否是PhoneConnecter,有可能是其他的USB类型，可以进入使用USBcheck的函数
	//但是不是PhoneConnecter 这样的结构体类型的，即不同结构体有相同的接口可能！结构体类型范围大
	if good{//判断good 是否ture
		fmt.Println(device.(PhoneConnecter),good) // 输出{abc} ture
	fmt.Println("This device is Phone", hi.name) // hi.name输出 abc
	return
	}

	fmt.Println(good,"This device is not Phone .")

}

//这里做一个接口判断函数，type switch类型判断
func InterfaceCheck(usb interface{}){ //定义是空的接口类型，因此它包括一切“继承”的变量
	switch usb.(type){  //使用type，自己不定义变量，由switch自己来判断;它会判断go的各种类型
		//包括 接口、结构体、基本数据类型等等.因此可以写PhoneConnecter 或USB来判断;
	case Inputer:
		fmt.Println("This is USB Interface:")
		//fallthrough
	/* 错误写法，经过第一次type,不能接下来转换类型了
	case PhoneConnecter:
		fmt.Println("This is Phone",v.name)
	*/
	
	default:
		fmt.Println("Not USB Interface:")
	}
	switch v:=usb.(type){ //v 值可以利用，但是由type 和 PhoneConnecter来决定;
	case PhoneConnecter:
		fmt.Println("This is Phone",v.name)
	default:
		fmt.Println("Not Phone")
	}//说明只是接口关闭，但是还是这个结构体



}



