package main

import "fmt"

// panic 은 다음가 같이 정의되어 있음 "func panic(v interface{})""

func main() {
	panic("Runtime Error")
	fmt.Println("Hello 世界!") //이 부분은 출력되지 않음

}
