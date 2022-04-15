package main

import "fmt"

//go的contanier使用，主要是array, slice ,map
func main() {
	fmt.Println("*****************************goContainer*******************************")
	array1 := [4]int32{1, 2, 3, 4}
	//注意这个是按照值传递的
	fmt.Printf("\n数组操作:按照值传递的\n")
	array2 := array1
	fmt.Printf("array1的值为 %v \n", array1)
	fmt.Printf("array2的值为 %v \n", array2)
	array1[0] = 10
	fmt.Println("array1[0] = 10 操作后")
	fmt.Printf("array1的值为 %v \n", array1)
	fmt.Printf("array2的值为 %v \n", array2)

	fmt.Printf("\nslice操作:按照引用传递，除了copy后的操作\n")
	//切片,切片是按照应用传递的
	fmt.Println("append操作-----------")
	slice1 := []int32{1, 2, 3, 4, 5, 6, 7}
	slice2 := slice1[1:4]
	fmt.Printf("slice1的值为 %v\n", slice1)
	fmt.Printf("slice2的值为 %v\n", slice2)
	slice2[1] = 10
	fmt.Println("slice2[1] = 10 操作后")
	fmt.Printf("slice1的值为 %v\n", slice1)
	fmt.Printf("slice2的值为 %v\n", slice2)
	//开始拼接
	slice3 := append(slice1, 100, 200)
	fmt.Println("append(slice1, 100, 200)后")
	fmt.Printf("slice1的值为 %v\n", slice1)
	fmt.Printf("slice2的值为 %v\n", slice2)
	fmt.Printf("slice3的值为 %v\n", slice3)

	fmt.Println("copy(slice2_copy,slice2)操作-----------")
	slice2_copy := make([]int32, len(slice2))
	count := copy(slice2_copy, slice2)
	fmt.Printf("slice2_copy的值为 %v\n", slice2_copy)
	fmt.Printf("slice2的值为 %v\n", slice2)
	slice2_copy[0] = 200
	fmt.Println("slice2_copy[0] = 200")
	fmt.Printf("slice2_copy的值为 %v\n", slice2_copy)
	fmt.Printf("slice2的值为 %v\n", slice2)
	fmt.Printf("复制的元素个数为: %d \n", count)

	fmt.Printf("\nmap操作:按照引用传递，\n")
	souceMap := make(map[string]string)
	souceMap["1"] = "1"
	souceMap["2"] = "2"
	souceMap["3"] = "3"
	sourceMap2 := souceMap
	fmt.Printf("souceMap的值为: %v\n", souceMap)
	fmt.Printf("sourceMap2的值为: %v\n", sourceMap2)
	if _, ok := souceMap["3"]; ok {
		delete(sourceMap2, "3")
	}
	fmt.Println(`if _, ok := souceMap["3"]; ok {	delete(sourceMap2, "3")}操作后`)
	fmt.Printf("souceMap的值为: %v\n", souceMap)
	fmt.Printf("sourceMap2的值为: %v\n", sourceMap2)
}
