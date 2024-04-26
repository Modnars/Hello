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
	fmt.Println(<-ch) // 即使 channel 被 close，依然可以从其中读出未读取的值
	fmt.Println(<-ch)
}
