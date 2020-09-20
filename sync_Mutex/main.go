package main

import (
	"fmt"
	"runtime"
)

type counter struct {
	i int64
}

func (c *counter) increment() {
	c.i += 1
}

func (c *counter) display() {
	fmt.Println(c.i)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := counter{i: 0}
	done := make(chan struct{})

	for i := 0; i < 1000; i++ {
		go func() {
			c.increment()
			done <- struct{}{}
		}()
	}

	for i := 0; i < 1000; i++ {
		<-done
	}

	c.display()
}
