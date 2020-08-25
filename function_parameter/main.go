package main

import "fmt"

func callFunction(f func()) {
	f()
}

func main() {
	// 함수를 파라미터로 취하는 함수
	callFunction(func() {
		fmt.Println("나는야 Function 입니다")
	})
}
