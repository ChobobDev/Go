package main

import "fmt"

func main() {
	/* Simple Slice Ecpression or 간이 슬라이스 식은 Subarray를 구현하는 것이라고 생각하면 이해하기 쉽다*/
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // Slice를 선언해 주는 코드
	b := a[1:5]                                // Slice에서 인덱스 1부터 인덱스 4까지의 수
	fmt.Println(b)

	/*문자열이 들어가 있을때의 경우*/
	c := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"[4:10] //String에서 substring을 빼오는 코드
	fmt.Println(c)

}
