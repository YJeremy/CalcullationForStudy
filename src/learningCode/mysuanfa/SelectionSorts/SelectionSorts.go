package SelectionSorts

//使用slice,引用操作，无需返回值
//函数为包内公开
func swap(slice []int, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func SelectionSort(s []int) {
	l := len(s)
	m := l - 1
	for i := 0; i < m; i++ {
		k := i
		for j := i + 1; j < l; j++ {
			if s[k] > s[j] {
				k = j
			}
		}

		if k != i {
			swap(s, k, i)
		}
	}

}
