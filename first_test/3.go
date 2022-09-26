package first_test

import "fmt"

// 输入一个字符串 统计字符出现次数统计
func GetStrShowCount(str string) map[string]int {
	result := make(map[string]int)
	for _, r := range str {
		key := string(r)
		if _, ok := result[key]; !ok {
			result[key] = 1
		} else {
			result[key]++
		}
	}
	printResult(result)
	return result
}

func printResult(result map[string]int) {
	for k, v := range result {
		fmt.Printf("%s 出现次数 %d\n", k, v)
	}
}
