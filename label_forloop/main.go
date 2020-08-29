package main

import "fmt"

func main() {
L:
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if j > 1 {
				continue L
			}
			fmt.Printf("%d X %d = %d\n", i, j, i*j)
		}
	}

}
