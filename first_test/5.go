package first_test

// 创建切片(slice) 使其元素为数字1-50，从切⽚删掉为3的倍数的数，(3, 6, 9, ...)在末尾再增加⼀个数666，输出切⽚
func SliceOpterator() []int {
	result := make([]int, 0, 50)

	for i := 1; i <= 50; i++ {
		result = append(result, i)
	}

	for i, v := range result {
		if v%3 == 0 {
			result = append(result[:i], result[i+1:]...)
		}
	}
	result = append(result, 666)
	return result
}
