package first_test

import (
	"fmt"
)

// generate 从2开始 把值传给chan通道
func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// 过滤
func filter(in chan int, out chan int, prime int) {
	// 阻塞主线程从通道取值
	for {
		// 取值
		num := <-in
		// 把满足条件的值传给另一个通道
		if num%prime != 0 {
			out <- num
		}
	}
}

func main() {
	// 创建int 类型 chan
	ch := make(chan int)
	// 用携程执行generate函数
	go generate(ch)
	for i := 0; i < 6; i++ {
		prime := <-ch
		fmt.Printf("prime:%d\n", prime)
		out := make(chan int)
		go filter(ch, out, prime)
		// 这行作用不懂
		ch = out
	}
}
