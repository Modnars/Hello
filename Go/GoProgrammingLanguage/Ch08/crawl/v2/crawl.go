package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // 获取一个 token
	list, err := links.Extract(url)
	<-tokens // 释放一个 token
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int // 用于标识当前有多少个 list 需要进行爬虫，用于退出程序，否则 worklist 通道返回默认值会一直持续循环下去

	n++ // 首先处理命令行 list
	go func() { worklist <- os.Args[1:] }()

	// 并发爬取以下网页
	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
