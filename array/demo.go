package main

import "fmt"

func main() {
	var arr1 [5]int
	fmt.Printf("arr1 size = %d, %d", len(arr1), arr1)
	fmt.Println()

	// 初始化并指定长度
	arr2 := [2]int{1, 2}
	fmt.Printf("arr2 size = %d, %d", len(arr2), arr2)
	fmt.Println()

	// 使用...自动计算数组长度
	arr3 := [...]int{1, 2, 3}
	fmt.Printf("arr3 size = %d, %d", len(arr3), arr3)
	fmt.Println()

	// 初始化指定长度
	arr4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("arr4 size = %d, %d", len(arr4), arr4)
	fmt.Println()
}
