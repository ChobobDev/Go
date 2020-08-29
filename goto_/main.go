package main

import "fmt"

func main() {
	fmt.Println("A")
	goto L
	fmt.Println("B") // 여기에 있는 모든 것들은 출력되지 않는다.
L:
	fmt.Println("This will be Printed")
}
