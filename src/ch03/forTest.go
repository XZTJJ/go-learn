package main

import "fmt"

//for循环使用
func forTestMain() {
	fmt.Println("*******************************forTestMain********************************")
	for i := 0; i < 20; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	source := 20
	for source > 0 {
		fmt.Printf("%d ", source)
		source--
		if source < 5 {
			break
		}
	}
	fmt.Println()

	const str string = "12345678"
	for index, value := range str {
		fmt.Printf("第 %d 位的ASCII的值=%d, 字符是 %c \n", index, value, value)
	}
	fmt.Println()

	for i := 0; i < 2; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
