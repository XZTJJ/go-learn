package main

import "fmt"

//闭包功能实现
func closeFuncMain() {
	fmt.Println("***********************closeFuncMain************************")
	counter1 := counter(10)
	fmt.Printf("counter1 type : %T\n", counter1)
	fmt.Printf("closeFuncMain() first call: %d\n", counter1())
	fmt.Printf("closeFuncMain() second call: %d\n", counter1())
	fmt.Printf("closeFuncMain() third call: %d\n", counter1())

	counter2 := counter(3)
	fmt.Printf("counter2 type : %T", counter2)
	fmt.Printf("closeFuncMain() first call: %d\n", counter2())
	fmt.Printf("closeFuncMain() second call: %d\n", counter2())
	fmt.Printf("closeFuncMain() third call: %d\n", counter2())
}

//
func counter(init int32) func() int32 {
	start := init
	return func() int32 {
		fmt.Printf("before add : %d", start)
		start++
		fmt.Printf(" , after add : %d \n", start)
		return start
	}
}
