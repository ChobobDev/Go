package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	s = append(s, 6) // s에 append 하는 line
	fmt.Println(s)

	s = append(s, 7, 8, 9, 10, 11, 12, 13) //s 에 multiple append 하기
	fmt.Println(s)
}
