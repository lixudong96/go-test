package first_test

// 在正整数范围内，大于1并且只有1和自身两个约数的数
func PrimeNumber(number int) []int {
	a := 1
	result := make([]int, 0, 50)
W:
	for a < number {
		a++
		for c := 2; c < a; c++ {
			if a%c == 0 {
				goto W
			}
		}
		result = append(result, a)
	}
	return result
}
