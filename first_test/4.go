package first_test

import (
	"fmt"
	"math/rand"
	"time"
)

// 删除切片(slice) 中的一个元素，
func DelSliceARandromElment(s []int) []int {
	len := len(s)
	rand.Seed(time.Now().Unix())
	num := rand.Intn(len)
	result := append(s[:num], s[num+1:]...)
	fmt.Println(result)
	return result
}
