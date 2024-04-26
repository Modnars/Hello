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
	abort := make(chan struct{})
	countdown := time.After(10 * time.Second)

	fmt.Println("Commencing countdown.  Press return to abort.")

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	select {
	case <-countdown:
		launch()
	case <-abort:
		fmt.Printf("launch aborted!!!\n")
	}
}
