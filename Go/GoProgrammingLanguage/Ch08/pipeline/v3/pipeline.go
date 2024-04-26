package main

import "fmt"

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Printf("%d\n", x)
	}
}

func main() {
	natruals := make(chan int)
	squares := make(chan int)
	go counter(natruals)
	go squarer(squares, natruals)
	printer(squares)
}
