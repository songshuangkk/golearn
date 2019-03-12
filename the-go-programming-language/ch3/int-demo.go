package main

import "fmt"

func main() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	fmt.Println(i, i+1, i*i)

	var x uint8 = 1<<1 | 1<<5 //  0000 0010 | 0010 0000 	 -> 0010 0010
	var y uint8 = 1<<1 | 1<<2 //  0000 0010 | 0000 0100	 -> 0000 0110

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y) // 0010 0010 & 0000 0110  	 -> 0000 0010
	fmt.Printf("%08b\n", x|y) // 0010 0010 | 0000 0110	 -> 0010 0110
	fmt.Printf("%08b\n", x^y) // 0010 0010 ^ 0000 0110	 -> 0010 0100

	// &^ : 位清除(AND NOT), 表达式: z = x &^ y中，若y的某位是1，则z的对应位等于0，否则，就等于x的对应位
	fmt.Printf("%08b\n", x&^y) // 0010 0010 &^ 0000 0110	 -> 0010 0000

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)
}
