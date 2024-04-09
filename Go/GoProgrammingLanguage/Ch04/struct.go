/*
 * @Author: modnarshen
 * @Date: 2024.04.08 10:37:02
 * @Note: Copyrights (c) 2024 modnarshen. All rights reserved.
 */
package main

import "fmt"

type Point struct{ x, y int }

func (p *Point) Info() string {
	return fmt.Sprintf("x: %d, y: %d", p.x, p.y)
}

func (p Point) ObjInfo() string {
	return fmt.Sprintf("x: %d, y: %d", p.x, p.y)
}

func main() {
	pp := &Point{1, 1}
	fmt.Println(pp.Info())
	fmt.Println((&Point{1, 2}).Info())
	fmt.Println(Point{1, 2}.ObjInfo())
}
