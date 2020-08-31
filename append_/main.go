package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	s = append(s, 6) // s에 append 하는 line
	fmt.Println(s)
}
