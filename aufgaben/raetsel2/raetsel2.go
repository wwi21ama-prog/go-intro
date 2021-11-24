package main

import (
	"fmt"
)

func main() {
	fmt.Println("Was machen nur diese Funktionen?")
}

func foo1(n int) []int {
	result := make([]int, 0)
	for ; n != 0; n /= 10 {
		result = append(result, n%10)
	}

	return result
}

func foo2(list []int) []int {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return list
}

func foo3(list []int, size int) []int {
	for i := len(list); i < size; i++ {
		list = append([]int{0}, list...)
	}
	return list
}

func foo4(list []int, size int) [][]int {
	for len(list)%3 != 0 {
		list = foo3(list, len(list)+1)
	}

	result := make([][]int, 0)
	for i := 0; i < len(list); i += 3 {
		result = append(result, []int{list[i], list[i+1], list[i+2]})
	}
	return result
}
