package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("rocket launches!!!")
}

func main() {
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})

	fmt.Println("Commencing countdown. Press enter to abort.")

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	for i := 10; i > 0; i-- {
		fmt.Println(i)
		select {
		case <-tick:
			// noop
		case <-abort:
			fmt.Println("launch aborted!!!")
			return
		}
	}
	launch()
}
