package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//go的io操作类,读写相关
func main() {
	fmt.Println("***********************ioCode*********************")
	readFile()
	writeFile()
}

//读文件，这里使用的scanner,不是bufio中的reader
func readFile() {
	filepath := "F:/logs/cfg/test.yml"
	//如果仅仅是打开文件的话，直接用Open方法就好了，如果不存在的同时创建这个文件，用OpenFile方法
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("文件打开失败,%v\n", err)
		return
	}
	fmt.Printf("开始读取%v文件\n", filepath)
	//关闭文件
	defer file.Close()
	//创建文件读取,这里使用的Scanner的方式
	fileScanner := bufio.NewScanner(bufio.NewReader(file))
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
}

func writeFile() {
	filepath := "F:/logs/tst/a/test.txt"
	//创建需要的父类文件夹
	allpart := strings.Split(filepath, "/")
	err := os.MkdirAll(strings.Join(allpart[:len(allpart)-1], "/"), os.ModePerm)
	if err != nil {
		fmt.Printf("文件夹创建失败,%v\n", err)
		return
	}
	//如果文件不存在，直接创建文件
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Printf("文件打开失败,%v\n", err)
		return
	}
	//需要关闭文件
	defer file.Close()
	//创建写的流程
	fileWriter := bufio.NewWriter(file)
	_, err = fileWriter.WriteString("发分发分发我的啊发我发的无法发放额的哇发我的发恶法而发违法我的晚饭问的哇\n")
	if err != nil {
		fmt.Printf("内容写入失败,%v\n", err)
	}
	_, err = fileWriter.WriteString("瞎写吧，不知道做什么")
	if err != nil {
		fmt.Printf("内容写入失败,%v\n", err)
	}
	_, err = fileWriter.WriteString("啊哈哈哈哈哈哈哈哈哈好玩哈哈哈哈哈啊哈")
	if err != nil {
		fmt.Printf("内容写入失败,%v\n", err)
	}
	err = fileWriter.Flush()
	if err != nil {
		fmt.Printf("内容写入失败,%v\n", err)
	}
	fmt.Printf("文件写入成功\n")
}
