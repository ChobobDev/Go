package main

import "fmt"

const (
	Running = 1 << iota
	Waiting
	Send
	Receive
)

func main() {
	stat := Running | Send
	display(stat)
}

func display(stat int) {
	if stat&Running == Running {
		fmt.Println("Running")
	}
	if stat&Waiting == Waiting {
		fmt.Println("Waiting")
	}
	if stat&Send == Send {
		fmt.Println("Send")
	}
	if stat&Receive == Receive {
		fmt.Println("Receive")
	}
}
