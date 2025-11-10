package main

import "fmt"

func control() {
	var a = 1
	a++
	// if
	if a > 0 {
		fmt.Printf("a > 0,a = %d\n", a)
	} else {
		fmt.Printf("a <= 0,a = %d\n", a)
	}
	var s = "1"
	// switch
	switch s {
	case "1":
		fmt.Println("s = 1")
	default:
		fmt.Println("s != 1")
	}

	// for
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
		fmt.Println(sum)
	}

	nums := []int{2, 3, 4}

	//range
	// 忽略索引
	for _, num := range nums {
		fmt.Println("value:", num)
	}

	// 忽略值
	for i := range nums {
		fmt.Println("index:", i)
	}
}
