/*
1、冒泡排序
1.1 对比相邻元素。如果第一个比第二个大，就交换它们两个;
1.2 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数;
1.3 针对所有的元素重复以上的步骤，除了最后一个;
1.4 重复步骤1～3,直到排序完成;
*/
package numic

//import "fmt"

func BubbleSort(p []int) {
	a := len(p) - 1
	for i := 0; i < a; i++ {
		for j := 0; j < a-i; j++ {
			if p[j] < p[j+1] {
				p[j], p[j+1] = p[j+1], p[j]
			}
		}
	}
}

/*
func main() {
	a := []int{32, 22, 1, 55, 48, 22}
	BubbleSort(a)
	fmt.Println(a)
}
*/
