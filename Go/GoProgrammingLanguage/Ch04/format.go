/*
 * @Author: modnarshen
 * @Date: 2024.04.08 11:58:23
 * @Note: Copyrights (c) 2024 modnarshen. All rights reserved.
 */
package main

import "fmt"

type Point struct{ x, y int }

func main() {
	p := Point{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)

}
