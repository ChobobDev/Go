package main

import "fmt"

// Defer는 함수 안에서 자유롭게 등록될 수 있다
// 함수 종료시 나중에 등록된 순으로 실행됨
func runDefer() {
	defer fmt.Println("DEFER")
	fmt.Println("DONE")
}

func main() {
	fmt.Println("이거 다음")
	runDefer()
}
