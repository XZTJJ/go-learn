package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
 *在go中使用channel进行 gorutine 之间的数据交互。
 *channel中的数据只能被消费一次,
 *channel中的数据时候类型的,只能读写指定的类型
 *channel分为及时channel和缓冲区channel。他们都会阻塞线程
 *及时channel:
 *    1.A Goroutine往channel中写数据，没有其他Goroutine读数据，A Goroutine会阻塞
 *    2.A Goroutine从channel中读数据, 没有其他Goroutine写数据，A Goroutine会阻塞
 *缓冲区channel：
 *    1.A Goroutine往channel中写数据，直到channel的缓冲区满了，A Goroutine才会阻塞
 *    2.A Goroutine从channel中读数据, 直到channel的缓冲区空了，A Goroutine才会阻塞
 */
func channelCodeMain() {
	fmt.Println("******************channelCodeMain*********************")
	instantChannel()
	time.Sleep(3 * time.Second)
	fmt.Println()
	bufChannel()
}

//生产者
func producer(ch chan string) {
	//关闭通道
	defer close(ch)
	fmt.Printf("生产者开始生产数据...\n")
	for i := 0; i < 10; i++ {
		fmt.Printf("生产者生产了糕点:%d\n", i)
		ch <- fmt.Sprintf("糕点%d", i)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

//消费者,多消费者，mark用于每个消费者是否消费完标识
func consumer(num int, ch chan string, mark chan bool) {
	//标识管道是消费者自己主动关闭才行
	defer close(mark)
	for data := range ch {
		fmt.Printf("----- %d 消费者，消费了 %s\n", num, data)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
	mark <- true
}

//及时channel,
func instantChannel() {
	fmt.Println("及时channel开始运行了")
	//创建一个生产者共享数据管道，多个消费者标识符
	var ch chan string = make(chan string)
	var cu1 chan bool = make(chan bool)
	var cu2 chan bool = make(chan bool)
	var cu3 chan bool = make(chan bool)
	//创建生产者
	go producer(ch)
	//创建消费者
	go consumer(1, ch, cu1)
	go consumer(2, ch, cu2)
	go consumer(3, ch, cu3)
	//主线程等待
	<-cu1
	<-cu2
	<-cu3
}

//缓冲区channel,
func bufChannel() {
	fmt.Println("缓冲区channel开始运行了")
	//创建一个生产者共享数据管道，多个消费者标识符
	var ch chan string = make(chan string, 20)
	var cu1 chan bool = make(chan bool)
	var cu2 chan bool = make(chan bool)
	var cu3 chan bool = make(chan bool)
	//创建生产者
	go producer(ch)
	//创建消费者
	go consumer(1, ch, cu1)
	go consumer(2, ch, cu2)
	go consumer(3, ch, cu3)
	//主线程等待
	<-cu1
	<-cu2
	<-cu3
}
