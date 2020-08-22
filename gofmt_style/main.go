package main

import "fmt"

func main() {
	//구조체
	type Rect struct {
		width  int
		height int
	}

	r := Rect{1, 2}
	fmt.Println(r.width*2 + r.height*2)
}
