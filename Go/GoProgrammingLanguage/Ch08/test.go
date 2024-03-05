/*
 * @Author: modnarshen
 * @Date: 2024.03.05 11:09:32
 * @Note: Copyrights (c) 2024 modnarshen. All rights reserved.
 */
package main

import "fmt"

func main() {
	ch := make(chan rune, 3)

	ch <- 'A'
	ch <- 'B'
	fmt.Println(<-ch)
	close(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
