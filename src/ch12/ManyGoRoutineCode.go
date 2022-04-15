package main

import (
	"fmt"
	"time"
)

/*
 *多个goroutine线程同时执行任务,
 *goroutine只能作用与方法后者函数上
 *如果主goroutine结束，其他的goroutine也自动结束了
 */
func ManyGoRoutineCodeMain() {
	fmt.Println("*********************ManyGoRoutineCodeMain**********************")
	go printNum()
	go printChar()
	time.Sleep(30 * time.Second)
	var intput string
	fmt.Scanln(&intput)
}

//一个简单goroutine函数
func printNum() {
	for i := 1; i < 50; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

//一个简单的goroutine函数
func printChar() {
	for i := 'a'; i < 'z'; i++ {
		time.Sleep(350 * time.Millisecond)
		fmt.Printf("%c", i)
	}
}
