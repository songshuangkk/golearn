package main

import "fmt"

func pointerCopy() {
	var array1 [3]*string

	array2 := [3]*string{new(string), new(string), new(string)}

	*array2[0] = "Red"
	*array2[1] = "Blue"
	*array2[2] = "Green"

	// 由于使用的是一个指针，复制之后，两个数组的元素指向的都是同一个值
	array1 = array2

	fmt.Println(array1)
}

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

	// 数组只有在相同的类型和长度之间才能进行一个复制操作。
	var arr5 [5]int
	// 这里是不能进行直接复制的，因为这个arr4的长度是需要运行时推算出来的，长度不明确
	//arr5 = arr4
	arr5 = arr1
	fmt.Println(arr5)
}
