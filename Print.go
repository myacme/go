package main

import (
	"fmt"
	"time"
)

func print() {

	fmt.Println("Hello World")
	var stockcode = 123
	var enddate = time.Now().Format("2006-01-02 03:04:05")
	var url = "Code=%d&endDate=%s"
	var targetUrl = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(targetUrl)

	m := max(1, 2)
	fmt.Printf("最大值：%d\n", m)

	g, j := swap("Go", "Java")
	fmt.Println(g, j)
}
