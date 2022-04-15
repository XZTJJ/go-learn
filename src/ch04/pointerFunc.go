package main

import "fmt"

//指针变量使用
func pointerFuncMain() {
	fmt.Println("***********************pointerFuncMain**************************")
	var num int32 = 100
	var pnum *int32 = &num
	var ppnum **int32 = &pnum
	fmt.Printf("num的类型为 %T,其值为: %v \n", num, num)             //int32  100
	fmt.Printf("pnum的类型为 %T,其值为: %v \n", pnum, pnum)          //*int32  num的内存地址
	fmt.Printf("*pnum的类型为 %T,其值为: %v \n", *pnum, *pnum)       //int  100
	fmt.Printf("ppnum的类型为 %T,其值为: %v \n", ppnum, ppnum)       //**int  pnum的变量地址
	fmt.Printf("*ppnum的类型为 %T,其值为: %v \n", *ppnum, *ppnum)    // *int32 num的内存地址
	fmt.Printf("**ppnum的类型为 %T,其值为: %v \n", **ppnum, **ppnum) //int  100
}
