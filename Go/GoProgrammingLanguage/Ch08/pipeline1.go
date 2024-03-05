/*
 * @Author: modnarshen
 * @Date: 2024.03.04 10:30:39
 * @Note: Copyrights (c) 2024 modnarshen. All rights reserved.
 */
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-squares)
	}
}
