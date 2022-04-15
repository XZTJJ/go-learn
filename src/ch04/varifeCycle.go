package main

import "fmt"

var a1 int = 1
var b1 int = 2

//变量的生命周期 以及 基础类型是否是按值传递还是引用传递.(值传递)
func varifeCycleMain() {
	fmt.Println("*******************************varifeCycleMain**************************************")
	a1, b1, c1 := 10, 20, 0
	fmt.Printf("main()函数中 a1 = %d \n", a1)
	fmt.Printf("main()函数中 b1 = %d \n", b1)
	fmt.Printf("main()函数中 c1 = %d \n", c1)
	c1 = sum(a1, b1)
	fmt.Printf("main()函数中 c1 = %d \n", c1)
	fmt.Printf("main()函数中 a1 = %d \n", a1)
}

func sum(a1, b1 int) (c1 int) {
	a1++
	b1 += 2
	c1 = a1 + b1
	fmt.Printf("sum()函数中 a1 = %d \n", a1)
	fmt.Printf("sum()函数中 b1 = %d \n", b1)
	fmt.Printf("sum()函数中 c1 = %d \n", c1)
	return c1
}
