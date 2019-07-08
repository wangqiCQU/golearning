package main

import "fmt"

func main(){
	const PI float32  = 3.1415926
	fmt.Println(PI)


	//const beef,two,c = "eat", 3, "veg"    常量支持并行赋值
	var a int8 = 12
	var b int8 = a
	fmt.Println(&a)
	fmt.Println(&b)
}
