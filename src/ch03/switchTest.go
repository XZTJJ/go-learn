package main

import "fmt"

//switch使用
func switchTestMain() {
	fmt.Println("**********************************switchTestMain*******************************")
	fmt.Println("普通的表达式测试,switch后面没有拼接任何的表达式,默认为 true:")
	source := 60
	switch {
	case source >= 90:
		fmt.Printf("%d 分数登记为 A\n", source)
	case source >= 80:
		fmt.Printf("%d 分数登记为 B\n", source)
	case source >= 70:
		fmt.Printf("%d 分数登记为 C\n", source)
	case source >= 60:
		fmt.Printf("%d 分数登记为 D\n", source)
	default:
		fmt.Printf("%d 分数登记为 E\n", source)
	}

	fmt.Println("现在是 case 后面拼接多个条件测试:")
	year := 2018
	month := 2
	days := 0

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
	case 4, 6, 9, 11:
		days = 30
	case 2:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			days = 29
		} else {
			days = 28
		}
	default:
		days = -1
	}
	fmt.Printf("%d 年 %d 月有 %d 天 \n", year, month, days)

}
