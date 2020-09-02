package main

import "fmt"

func main() {
	/*cap이란 capacity의 줄인말로 len과는 다르게 slice의 용량(Capacity)의 값을 가진다*/
	/* 크기가10, 용량이 10인 슬라이스*/
	a := make([]int, 10)
	fmt.Println(len(a)) // 출력값 = 10
	fmt.Println(cap(a)) // 출력값 = 10
	/* 크기가10, 용량이 15인 슬라이스*/
	b := make([]int, 10, 15)
	fmt.Println(len(b)) // 출력값 = 10
	fmt.Println(cap(b)) // 출력값 = 15

}
