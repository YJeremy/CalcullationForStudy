package SlectionSort

//分步
func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]
}

//
func SlectionSort(s []int) {
	l := len(s)              //创建大阔值
	m := l - 1               //创建小阔值
	for i := 0; i < m; i++ { //创建序标，迭代
		k := i                       //创建定标
		for j := i + 1; j < l; j++ { //创建小序标，迭代
			if s[j] < s[k] { //定标比较
				k = j //定标交换
			}

		}
		if k != i { //定标比较
			swap(s, k, i) //定标值交换
		}

	}
}
