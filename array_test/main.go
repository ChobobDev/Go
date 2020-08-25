package main

import "fmt"

func main() {
	a := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a)
	i := 0
	for {
		if i >= 7 {
			break
		}
		fmt.Printf(" %d 번째 index : %v\n", i+1, a[i])
		i++
	}
}
