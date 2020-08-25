package main

import "fmt"

func addition(x, y int) (int, int, int) {
	return x, y, x + y

}

func main() {
	s := "Seongbin" + " " + "Bernie" + " " + "Cho\n" //문자열 Concate 하는 부분은 Python을 연상시킴
	fmt.Printf("%v", s)                              //문자열의 출력은 C언어를 연상시킴
	x, y, sum := addition(4, 14)                     //Function에 값을 넣고 받아오기
	fmt.Printf("%d + %d = %d", x, y, sum)            // Function 에서 return값을 출력

}
