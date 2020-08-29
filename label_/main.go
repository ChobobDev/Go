package main

import "fmt"

func main() {
LOOP:
	for {
		for {
			for {
				fmt.Println("Start")
				break LOOP
			}
			fmt.Println("Not Yet") // 실행되지 않는 부분
		}
		fmt.Println("Not Yet") // 실행되지 않는 부분
	}
	fmt.Println("Finalized")

LOOP_2:
	for {
		for {
			for {
				fmt.Println("Start")
				break
			}
			fmt.Println("Not Yet")
			break LOOP_2
		}
		fmt.Println("Not Yet") // 실행되지 않는 부분
	}
	fmt.Println("Finalized")
}
