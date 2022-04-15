package main

import (
	"fmt"
	"strconv"
)

/*
 *定义接口,同样也是使用的interface关键字,注意实现接口，
 *只需要实现对应的方法，默认就是实现这个接口了
 */
type MyInterface interface {
	call(MyInterface, string) (int64, float64)
	toType()
}

//定义class，注意go中没有class的概念，使用的struct，struct是按照值传递的
type ParentStruct struct {
	name, mark string
	age        int32
}

/*
 *go通过匿名字段实现，实现Java中的继承关系,比如，
 *下面两个 Child1Struct , Child1Struct就继承了上面 ParentStruct 类
 *可以直接调用 父类的字段和方法
 */
type Child1Struct struct {
	ParentStruct
	school string
}
type Child2Struct struct {
	ParentStruct
	work string
}

/*
 *定义方法，在go中方法和函数是非常类似的，增加方法
 *主要是为了模拟面向对象的思想，方法比函数多了一个
 *接受者的角色.（接受者简单来说就是this，也就是调用方法
 *的对象)。方法是可以重名的
 */

//注意这里是按照值传递的,专属于ParentStruct的方法。和Java的方法调用一样
func (p ParentStruct) call(myI MyInterface, str string) (intdata int64, floatdata float64) {
	if instance, ok := myI.(ParentStruct); ok {
		instance.mark = "ParentStruct"
		fmt.Printf("ParentStruct的call, mark = \"ParentStruct\"， 类型= %T, 值为:%+v\n", instance, instance)
	} else if instance, ok := myI.(Child1Struct); ok {
		instance.school = "Child1Struct"
		fmt.Printf("ParentStruct的call, school = \"Child1Struct\"， 类型= %T, 值为:%+v\n", instance, instance)
	} else if instance, ok := myI.(Child2Struct); ok {
		instance.work = "Child2Struct"
		fmt.Printf("ParentStruct的call, work = \"Child2Struct\"， 类型= %T, 值为:%+v\n", instance, instance)
	}
	intdata, _ = strconv.ParseInt(str, 10, 32)
	floatdata, _ = strconv.ParseFloat(str, 32)
	return intdata, floatdata
}

//注意这是传递的是应用，空实现
func (p ParentStruct) toType() {
}

//注意这是传递的是应用，这几修改原值的，这里体现了方法的覆写
func (c *Child2Struct) toRefeType() {
	(*c).age = 80
	fmt.Println("Child2Struct的toType中(*p).age = 80 操作后")
}

//面向对象的编程的思想，注意和Java中区分开来
func main() {
	fmt.Println("**********************oopCode****************************")
	//三种不同的初始化方法
	p1 := ParentStruct{}
	p1.name = "大张"
	p1.age = 30
	p1.mark = "社会人"
	c1 := Child1Struct{ParentStruct: p1, school: "xx大学"}
	c2 := Child2Struct{p1, "xx工厂"}
	fmt.Printf("未修改之前,p1 ：%T, %+v \n", p1, p1)
	fmt.Printf("未修改之前,c1 ：%T, %+v \n", c1, c1)
	fmt.Printf("未修改之前,c2 ：%T, %+v \n", c2, c2)
	//按照值传递的话，应该不会修改才对
	c1.age, c1.mark, c2.age, c2.mark = 20, "大学生", 40, "搬砖工"
	fmt.Println(`c1.age, c1.mark, c2.age, c2.mark = 20, "大学生", 40, "搬砖工" 操作后`)
	fmt.Printf("修改之后,p1 ：%T, %+v \n", p1, p1)
	fmt.Printf("修改之后,c1 ：%T, %+v \n", c1, c1)
	fmt.Printf("修改之后,c2 ：%T, %+v \n", c2, c2)
	fmt.Println()
	var myInt MyInterface
	myInt = c1
	fmt.Printf("将c1赋值给MyInterface变量, 未调用 c1 call之前 ：%T, %+v \n", myInt, myInt)
	intdata, floatdata := myInt.call(myInt, "100")
	fmt.Printf("将c1复制为MyInterface变量, 调用 c1 call之后 ：%T, %+v \n", myInt, myInt)
	fmt.Printf("intdata 返回类型 ：%T, %+v ， floatdata 返回类型 ：%T, %+v \n", intdata, intdata, floatdata, floatdata)
	fmt.Println()
	fmt.Printf("未调用 c2 toRefeType之前：%T, %+v \n", c2, c2)
	c2.toRefeType()
	fmt.Printf("调用 c2 toRefeType之后 ：%T, %+v \n", c2, c2)
}
