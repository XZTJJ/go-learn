package main

import "fmt"

const (
	Unknown = 0
	Female  = 1
	Male    = 2
	Third
)

//常量枚举类型
func constEnumMain() {
	fmt.Println("**********************************constEnum操作**********************************")
	fmt.Println(Unknown, Female, Male, Third)
}
