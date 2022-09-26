package main

import (
	"first-test/first_test"
	"flag"
	"fmt"
)

func main() {
	fmt.Println("start first test!")

	first_test.DelSliceARandromElment([]int{123, 1, 213, 34, 23, 34, 56})
	first_test.SliceOpterator()
	first_test.GenerateHaskell()
}

func test3() {
	var strs string
	flag.StringVar(&strs, "s", "string", "要获取字符个数的字符串")
	flag.Parse()
	fmt.Println(strs)
	first_test.GetStrShowCount(strs)
}
