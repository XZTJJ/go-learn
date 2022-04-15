package main

import "fmt"

//if 条件语句使用
func ifTestMain() {
	fmt.Println("**********************************ifTestMain*******************************")
	score := 60
	fmt.Printf("分数为 %d ", score)
	//多成嵌套关系
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 70 {
		fmt.Println("中等")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	//特殊写法
	fmt.Print("特殊写法,在if条件中包含初始化声明, 奇偶数判断: ")
	if num := 30; num%2 == 0 {
		fmt.Printf("%d 是偶数\n", num)
	} else {
		fmt.Printf("%d 是奇数\n", num)
	}
}
