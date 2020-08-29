package main

import "fmt"

func main() {
	fmt.Println("A")
	goto L
	fmt.Println("B") // 여기에 있는 모든 것들은 출력되지 않는다.
L:
	fmt.Println("This will be Printed")
	goto M //M 으로 선언된 부분으로 이동한다

M:
	fmt.Println("This will be Printed as well")

}
