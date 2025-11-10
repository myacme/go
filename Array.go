package main

import "fmt"

func array() {

	var a = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	s := []int{0, 1, 2}
	printSlice(s)

	//
	fmt.Println(s[1])
	fmt.Println(s[1:])
	fmt.Println(s[1:3])
	fmt.Println(s[:3])
	fmt.Println(s[:1])

	//
	s = append(s, 3, 4)
	printSlice(s)
	i := make([]int, len(s))
	copy(i, s)
	i = append(i, 5, 6)
	printSlice(i)

}
