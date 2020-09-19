package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	/*
		채널 c에 세번째 값을 전송하는 부분을 별도의 고루틴으로 동작시킨다.
		고루틴은 채널 c에 값을 전송할 수 있을때까지 대기한다.
	*/
	go func() { c <- 3 }()
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
