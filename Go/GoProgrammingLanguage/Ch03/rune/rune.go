package main

import "fmt"

func main() {
	fmt.Printf("%c\n", '\u2318')
	fmt.Printf("%c\n", '\U0001F602')
	fmt.Println(len("世界"))
	fmt.Println(len("\u4e16\u754c"))
	fmt.Printf("%c\n", '\uFFFD')
}
