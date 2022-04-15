package main

import (
	"fmt"
	"strings"
)

//函数变量，函数作为值传递，以及非常类似接口的 type  函数变量使用
func funcVarMain() {
	fmt.Println("*******************************funcVarMain**************************************")
	strRes := stringToCus("ABCDEFGHIJK", oddCharToLow)
	fmt.Println(strRes)
	strRes = stringToCus2("ABCDEFGHIJK", oddCharToLow)
	fmt.Println(strRes)
}

//定义一个函数，将索引为奇数的变成小写, 因为只有一个返回类型，所以省略返回类型的括号
func oddCharToLow(str string) string {
	result := ""
	for index, value := range str {
		if index%2 == 0 {
			result += strings.ToUpper(string(value))
		} else {
			result += strings.ToLower(string(value))
		}
	}
	return result
}

//定义一个调用函数的方法
func stringToCus(str string, f func(string) string) string {
	fmt.Printf("调用函数的类型为= %T \n", f)
	return f(str)
}

//type 声明一个函数类型，非常类似于接口，但是比接口更加简介
type cusFunc func(string) string

func stringToCus2(str string, f cusFunc) string {
	fmt.Printf("调用函数的类型为= %T \n", f)
	return f(str)
}
