package main

import "fmt"

func main() {
	var x interface{} = false
	switch x.(type) {
	case bool:
		fmt.Println("Boolean") // x의 Type이 Boolean일때 출력이됨
	case int, uint:
		fmt.Println("Integer or Unsigned Integer") // x의 Type이 int 이거나 uint일때 출력이됨
	case string:
		fmt.Println("String") // x의 Type이 String일때 출력이됨
	default:
		fmt.Println("Not Satisfied") // x의 Type을 모를때 출력이됨
	}
}
