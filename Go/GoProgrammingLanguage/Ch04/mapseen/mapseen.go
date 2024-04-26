/*
 * @Author: modnarshen
 * @Date: 2023.11.21 11:18:39
 * @Note: Copyrights (c) 2023 modnarshen. All rights reserved.
 */
package main

import "fmt"

var seen = make(map[string]struct{}) // set of strings

func test(s string) {
	if _, ok := seen[s]; !ok {
		fmt.Println(len(seen)) // 虽然判断时调用了 seen[]，但其作为右值不改变 seen
		// seen[s] = struct{}{}
		// ...first time seeing s...
	}
}

func main() {
	test("1")
	fmt.Println(len(seen))
}
