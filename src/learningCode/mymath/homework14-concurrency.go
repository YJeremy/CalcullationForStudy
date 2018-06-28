package main
import(
	"fmt"
	//"time"
)

func main(){
	c := make(chan bool)
	go func(){
		fmt.Println("Go Go Go!!!")
		c <- true
		//close(c)
	}()
	//<-c //产生阻塞，主函数等待c的状态有true，然后再执行完成程序，释放你内存
	for v:= range c {
		fmt.Println(v)
	}



}
