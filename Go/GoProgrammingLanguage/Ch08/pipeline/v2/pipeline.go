/*
 * @Author: modnarshen
 * @Date: 2024.03.04 10:39:32
 * @Note: Copyrights (c) 2024 modnarshen. All rights reserved.
 */
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}
