package main

import "fmt"

func main() {
	sum, i := 0, 0
	for i < 10 {
		sum += i
		i++
	}
	fmt.Println(sum)
}
