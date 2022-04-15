package main

import (
	"fmt"
	"math"
)

//匿名函数测试
func anonymousFuncMain() {
	fmt.Println("************************anonymousFuncMain***********************")
	normalAnonyMousFunc()
	arr := []float64{1, 9, 16, 25, 36}
	//求平方根
	visitAnonyMousFunc(arr, func(data float64) {
		sqrt := math.Sqrt(float64(data))
		fmt.Printf("%f 的平方根 = %f\n", data, sqrt)
	})
	//求平方
	visitAnonyMousFunc(arr, func(data float64) {
		pow := math.Pow(data, 2)
		fmt.Printf("%f 的平方 = %f\n", data, pow)
	})
}

//普通的匿名函数,注意这样会直接将数值转成 ascii编码或者unicode编码
func normalAnonyMousFunc() {
	result := func(intStr int32) string {
		return string(intStr)
	}(100)
	fmt.Printf("普通匿名函数结果为: %s\n", result)
}

//用于调用的匿名函数
func visitAnonyMousFunc(slice []float64, f func(data float64)) {
	for _, value := range slice {
		f(value)
	}
}
