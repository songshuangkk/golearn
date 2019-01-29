package main

import "fmt"

// 注意: 如果[]运算符当中指定了一个值，这个时候就是一个数组了，而不是一个切片了

// 切片内部有三个数据结构: 1.指向底层数组的指针. 2.切片的访问元素的个数. 3.切片的长度
func main() {
	slice1 := make([]int, 3, 5)
	fmt.Println(slice1)

	// 创建字符串切片
	slice2 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println(slice2)

	// 创建一个整型的切片，长度和容量都是3个元素
	slice3 := []int{10, 20, 30}
	fmt.Println(slice3)

	// 使用索引声明切片, 第五个元素索引为1
	slice4 := []string{5: "1"}
	fmt.Println(slice4)

	// slice 和newSlices内部指向的底层数组是同一个底层数组
	// newSlice5, 其实就是指向的首地址向后偏移，然后容量变为底层总容量减去偏移量，长度变为初始复制的范围，这里的长度就是2，容量是4
	slice5 := []int{10, 20, 30, 40, 50}
	newSlice5 := slice5[1:3]
	fmt.Println(newSlice5)

	fmt.Println(slice5)
	newSlice5[1] = 35
	fmt.Println(slice5)

	// 如果容量能够满足底层数组，那么对增长的切片进行的一个操作实际会对底层的数组进行操作
	newSlice5 = append(newSlice5, 60, 60)
	fmt.Println(newSlice5)
	fmt.Println(slice5)
	// 如果容量不能满足底层的数组，那么将会创建一个新的底层数组，然后这个扩容后的切片指向这个新的底层数组.
	newSlice5 = append(newSlice5, 60, 60)
	fmt.Println(slice5)
	fmt.Println(newSlice5)
}
