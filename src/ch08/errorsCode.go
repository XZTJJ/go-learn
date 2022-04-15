package main

import (
	"errors"
	"fmt"
)

/*
 *go中使用异常的处理基础，是一个error接口
 *并没有对应的try catch finally，error只是
 *一个普通返回结果。这和其他语言非常不同
 *通过error与nil判断
 *注意:panic和recover关键字，不建议使用
 */

//处理异常
func main() {
	fmt.Println("***********************errorsCode**********************")
	//延迟函数执行,可以作为finally使用
	defer normalFunc(1)
	//其他错误处理,正常和异常处理逻辑
	value, ers := hasErrors(10, 5)
	if ers != nil {
		fmt.Printf("hasErrors(10, 5)发生错误，ers类型=%T,值=%v\n", ers, ers)
	} else {
		fmt.Printf("hasErrors(10, 5)正常返回结果，结果类型为=%T,值=%v\n", value, value)
	}
	value, ers = hasErrors(4, 0)
	if ers != nil {
		fmt.Printf("hasErrors(4, 0)发生错误，ers类型=%T,值=%v\n", ers, ers)
	} else {
		fmt.Printf("hasErrors(4, 0)正常返回结果，结果类型为=%T,值=%v", value, value)
	}
	//延迟函数执行,可以作为finally使用
	defer normalFunc(2)
	fmt.Println("main函数执行完了")
}

//普通的函数，处理异常
func hasErrors(dividee float64, divider float64) (float64, error) {
	if divider == 0 {
		return 0, errors.New("除数不能为0")
	} else {
		return dividee / divider, nil
	}
}

//普通函数，注意 defer 的调用顺序，defer 只是修饰在 方法 和 函数 上
func normalFunc(date int32) {
	fmt.Printf("normalFunc的序号为: %v\n", date)
}
