package main

import "fmt"

func main() {
	/* int타입의 슬라이스 */
	var a []int
	b := make([]int, 10)
	var c [10]int

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
