package main

import "fmt"

func runDefer() {
	defer fmt.Println("DEFER")
	fmt.Println("DONE")
}

func main() {
	runDefer()

}
