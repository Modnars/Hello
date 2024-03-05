/*
 * @Author: modnarshen
 * @Date: 2023.11.15 11:40:50
 * @Note: Copyrights (c) 2023 modnarshen. All rights reserved.
 */
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
