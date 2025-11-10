package main

import (
	"fmt"
)

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}

/**
 * 打印 slice
 */
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

/**
 * 打印 map
 */
func printMap(m map[string]int) {
	for k, v := range m {
		fmt.Printf("key=%s, value=%d\n", k, v)
	}

}

// 递归函数计算阶乘
func factorial(n int) int {
	// 基准条件
	if n == 0 {
		return 1
	}
	// 递归条件
	return n * factorial(n-1)
}

// 自定义类型约束
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}
