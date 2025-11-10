package main

import "fmt"

func map1() {

	var m = map[string]int{"1": 1, "2": 2, "3": 3}
	fmt.Println(m["1"])

	// 遍历 Map
	printMap(m)

	// 删除键值对
	delete(m, "1")

	//4 会添加到 1 的位置
	m["4"] = 4
	m["5"] = 5
	fmt.Println("删除后的map：")

	// 遍历 Map
	printMap(m)

}
