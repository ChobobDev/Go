package main

import "fmt"

func main() {
	x := 1
	for x < 100 {
		if x%2 == 0 {
			fmt.Printf("%d는 Even 이다 뷁\n", x)
		} else {
			fmt.Printf("%d는 Odd 이다 뷁\n", x)
		}
		x++
	}

}
